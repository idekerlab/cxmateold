package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/ericsage/cxmate/cxpb"
	"google.golang.org/grpc"
)

var (
	requiredAspects = []string{
		"metaData",
		"edges",
		"nodes",
		"nodeAttributes",
		"edgeAttributes",
		"networkAttributes",
	}
)

var (
	listeningAddress = getenv("LISTENING_ADDRESS", "0.0.0.0")
	listeningPort    = getenv("LISTENING_PORT", "80")
	serverAddress    = getenv("SERVICE_ADDRESS", "127.0.0.1")
	serverPort       = getenv("SERVICE_PORT", "8080")
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	handler := CXHandler
	http.HandleFunc("/", handler)
	address := listeningAddress + ":" + listeningPort
	log.Fatal(http.ListenAndServe(address, nil))
}

func CXHandler(res http.ResponseWriter, req *http.Request) {
	streamNetwork(req.Body, res)
}

func streamNetwork(network io.ReadCloser, out http.ResponseWriter) {
	address := serverAddress + ":" + serverPort
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic("Could not establish connection")
	}
	client := cxpb.NewCyServiceClient(conn)
	stream, err := client.StreamElements(context.Background())
	if err != nil {
		panic("Could not open stream")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		decOpt := &cxpb.DecoderOptions{
			RequiredAspects: requiredAspects,
			IsCollection:    true,
			NumNetworks:     2,
		}
		decoder := cxpb.NewDecoder(network, decOpt, stream.Send)
		decoder.Decode()
		stream.CloseSend()
	}()

	go func() {
		defer wg.Done()
		encOpt := &cxpb.EncoderOptions{
			RequiredAspects: requiredAspects,
			IsCollection:    true,
			NumNetworks:     2,
		}
		encoder := cxpb.NewEncoder(out, encOpt, stream.Recv)
		encoder.Encode()
	}()

	wg.Wait()
}
