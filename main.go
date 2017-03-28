package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/ericsage/cxmate/cxpb"
	"github.com/ericsage/cxmate/metrics"
	"github.com/golang/protobuf/jsonpb"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

var (
	requestsCounter = metrics.NewCounter(
		"cxmate_requests_total",
		"Number of total requests processed by cxmate",
		[]string{},
	)
	errorCounter = metrics.NewCounter(
		"cxmate_response_with_errors_total",
		"Number of responses that contained at least one error",
		[]string{},
	)
	errorsCounter = metrics.NewCounter(
		"cxmate_sent_errors_total",
		"Total number of errors sent by cxmate to clients",
		[]string{},
	)
	sentParametersCounter = metrics.NewCounter(
		"cxmate_recieved_parameter_total",
		"Number of times each parameter key/value recieved by cxmate",
		[]string{"name", "value"},
	)
)

var (
	listeningAddress       = getenv("LISTENING_ADDRESS", "0.0.0.0")
	listeningPort          = getenv("LISTENING_PORT", "80")
	serverAddress          = getenv("SERVICE_ADDRESS", "127.0.0.1")
	serverPort             = getenv("SERVICE_PORT", "8080")
	requiresCollection, _  = strconv.ParseBool(getenv("RECEIVES_COLLECTION", "false"))
	requiredNumNetworks, _ = strconv.ParseInt(getenv("EXPECTED_NUM_NETWORKS", "1"), 10, 64)
	requiredAspects        = strings.Split(getenv("RECEIVES_ASPECTS", "edges, nodes, nodeAttributes, edgeAttributes, networkAttributes"), ",")
	sendsCollection, _     = strconv.ParseBool(getenv("SENDS_COLLECTION", "false"))
	sendNumNetworks, _     = strconv.ParseInt(getenv("SENDS_NUM_NETWORKS", "1"), 10, 64)
	sendingAspects         = strings.Split(getenv("SENDS_ASPECTS", "edges, nodes, nodeAttributes, edgeAttributes, networkAttributes"), ",")
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
	address := listeningAddress + ":" + listeningPort
	s := &http.Server{
		Addr:    address,
		Handler: http.HandlerFunc(handler),
	}
	fmt.Println("Config:")
	fmt.Println("Listening at address:", listeningAddress)
	fmt.Println("Listening on port:", listeningPort)
	fmt.Println("Proxying for service at address", serverAddress)
	fmt.Println("Proxying for service on port", serverPort)
	fmt.Println("Receiving a collection:", requiresCollection)
	fmt.Println("Receives", requiredNumNetworks, "network(s)")
	fmt.Println("Receiving aspects:", requiredAspects)
	fmt.Println("Sending a collection:", sendsCollection)
	fmt.Println("Sends", sendNumNetworks, "network(s)")
	fmt.Println("Sending aspects:", sendingAspects)
	fmt.Println("Now listening...")
	go metrics.Serve()
	log.Fatal(s.ListenAndServe())
}

//CXHandler handles CX
func CXHandler(res http.ResponseWriter, req *http.Request) {
	requestsCounter.With(prometheus.Labels{}).Inc()
	params := req.URL.Query()
	streamNetwork(req.Body, res, params)
}

func streamNetwork(in io.Reader, out io.Writer, params map[string][]string) {
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
			IsCollection:    requiresCollection,
			NumNetworks:     requiredNumNetworks,
		}
		sendParams(stream.Send, params)
		decoder := cxpb.NewDecoder(in, decOpt, stream.Send)
		decoder.Decode()
		stream.CloseSend()
	}()

	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in encoder, panic:", r)
			}
		}()
		encOpt := &cxpb.EncoderOptions{
			RequiredAspects: sendingAspects,
			IsCollection:    sendsCollection,
			NumNetworks:     sendNumNetworks,
		}
		encoder := cxpb.NewEncoder(out, encOpt, stream.Recv)
		io.WriteString(out, "{\"data\":")
		errors := encoder.Encode()
		io.WriteString(out, ",\"errors\":[")
		m := jsonpb.Marshaler{}
		if len(errors) != 0 {
			errorCounter.With(prometheus.Labels{}).Inc()
		}
		for index, error := range errors {
			if index != 0 {
				io.WriteString(out, ",")
			}
			m.Marshal(out, error)
			errorsCounter.With(prometheus.Labels{}).Inc()
		}
		io.WriteString(out, "]}")
	}()

	wg.Wait()
}

func sendParams(streamer func(*cxpb.Element) error, params map[string][]string) {
	for name, values := range params {
		for _, value := range values {
			p := &cxpb.Element{
				NetworkId: 0,
				Value: &cxpb.Element_Parameter{
					Parameter: &cxpb.Parameter{
						Name:  name,
						Value: value,
					},
				},
			}
			err := streamer(p)
			if err != nil {
				panic("Could not send parameter")
			}
			sentParametersCounter.With(prometheus.Labels{
				"name":  name,
				"value": value,
			}).Inc()
		}
	}
}
