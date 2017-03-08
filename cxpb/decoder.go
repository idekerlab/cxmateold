package cxpb

import (
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io"
)

//Decoder decodes a JSON stream into a stream of protocol buffer elements, and hands them off to it's internal streamer for further processing.
type Decoder struct {
	r           io.ReadCloser
	dec         *json.Decoder
	streamer    func(*Element) error
	options     *DecoderOptions
}

//DecoderOptions provides a simple interfaces to configure a cxpb JSON Decoder. The IsCollection flag specifies whether the Decoder should expect
//an array of 0 or more networks, or a singleton network. If IsCollection is set, the Decoder will decode the number of networks specified in NumNetworks.
//Any network Decoded, singleton or nested, should contain at least the Aspects defined in RequiredAspects. RequiredAspects must, in turn, be a subset of KnownAspects.
//The Decoder will only parse and send Aspects that are in both RequiredAspects and KnownAspects to the streamer.
type DecoderOptions struct {
	RequiredAspects []string
	IsCollection    bool
	NumNetworks     int64
}

//NewDecoder initializes and returns a JSON Decoder that decodes to a stream of protobuf elements defined in this package. NewDecoder accepts
//a stream to read from, r, an options struct, o, and a calback function, s, that accepts a single element s that can stream the element elewhere
//or handle it in some way. Before calling NewDecoder, a DecoderOptions struct should be made detailing the expected input from the stream.
func NewDecoder(r io.ReadCloser, o *DecoderOptions, streamer func(*Element) error) *Decoder {
	d := &Decoder{
		r:           r,
		dec:         json.NewDecoder(r),
		streamer:    streamer,
		options:     o,
	}
	return d
}

//Decode parses a json stream as specified by the DecoderOptions in protobuf elements. It can parse a singleton network or a list of one or more networks.
//Every decoded element has a networkId specifing which network it belongs to. A singleton network's elements will all have networkId
//set to 0. A list of networks will have elements with networkId 1 through DecoderOptions.NumNetworks inclusive. For every element in a
//network, Decode will call the streamer passed into NewDecoder with the element as the sole argument. The streamer may then decide what to
//do with the element.
func (dec *Decoder) Decode() {
	d := dec.dec
	if dec.options.IsCollection { //Check if a collection is expected
    d.Token() //Remove collection opening [
		var n int64 = 1
		for n <= dec.options.NumNetworks {
			dec.decodeNetwork(n) //Decode network n
			n++
		}
		d.Token() //Remove collection closing ]
	} else {
			dec.decodeNetwork(0) //There's only one network
  }
	stripEOF(d) //Should find EOF at end of stream
}

//decodeNetwork handles the decoding of elements from a single network, creating elements that are tagged with a network number, networkNum.
//The decoder will then call the streamer that was passed into NewDecoder with the element as the sole parameter. If an fragment is not recognized
//or required, it is discared.
func (dec *Decoder) decodeNetwork(networkNum int64) {
	d := dec.dec
	d.Token()      //Opeing network [
	for d.More() { //Iterate over every fragment in the network
		d.Token()            //Opening fragment {
		name, _ := d.Token() //Fragment name
		d.Token() //Opening element [
		element, known := KnownAspects[name.(string)]
		required := isAspectRequired(name.(string), dec.options.RequiredAspects)
		for d.More() { //Iterate over every element in the fragment
			if known && required { //Convert known aspects to protobuf
				jsonpb.UnmarshalNext(d, element)
				dec.streamer(wrapElement(networkNum, element))
			} else { //Consume and do nothing with known aspects
				var v map[string]interface{}
				d.Decode(&v)
			}
		}
		d.Token() //Closing element ]
		d.Token() //Closing fragment }
	}
	d.Token() //Closing network ]
}

//KnownAspects specifies the full set of aspects the Decoder can understand. It provides a map between
//an aspect name and the Protocol Buffers struct representing the type. The Decoder uses this map to convert
//JSON elements to their protocol buffers representation.
var KnownAspects = map[string]proto.Message{
		"metaData":          &MetaData{},
		"nodes":             &Node{},
		"edges":             &Edge{},
		"nodeAttributes":    &NodeAttribute{},
		"edgeAttributes":    &EdgeAttribute{},
		"networkAttributes": &NetworkAttribute{},
	}

//wrapElement wraps a protocol buffers aspect element in an Element wrapper and tags the wrapper with a network number.
//Elements must be wrapped so that the streamer has a uniform type to stream.
func wrapElement(networkNum int64, element proto.Message) *Element {
	switch e := element.(type) {
	case *MetaData:
		return &Element{networkNum, &Element_Metadata{e}}
	case *Node:
		return &Element{networkNum, &Element_Node{e}}
	case *Edge:
		return &Element{networkNum, &Element_Edge{e}}
	case *NodeAttribute:
		return &Element{networkNum, &Element_NodeAttribute{e}}
	case *EdgeAttribute:
		return &Element{networkNum, &Element_EdgeAttribute{e}}
	case *NetworkAttribute:
		return &Element{networkNum, &Element_NetworkAttribute{e}}
	default:
		return &Element{networkNum, &Element_Aspect{}}
	}
}

//isAspectRequired checks to see if an aspect name encountered during streaming is required by the service, if not, the Decoder
//may simply discard the aspect fragment it is processing.
func isAspectRequired(aspectName string, required []string) bool {
    for _, name := range required {
        if aspectName == name {
            return true
        }
    }
    return false
}

//stripEOF looks for the EOF err, asserting that the JSON document has been fully parsed.
func stripEOF(dec *json.Decoder) {
	_, err := dec.Token()
	if err != io.EOF {
		panic("Should be EOF")
	}
}
