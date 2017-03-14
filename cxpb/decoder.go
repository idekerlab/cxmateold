package cxpb

import (
	"encoding/json"
	"io"
  "strconv"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

//Decoder decodes a JSON stream into a stream of protocol buffer elements, and hands them off to it's internal streamer for further processing.
type Decoder struct {
	r        io.Reader
	dec      *json.Decoder
	streamer func(*Element) error
	options  *DecoderOptions
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
func NewDecoder(r io.Reader, o *DecoderOptions, streamer func(*Element) error) *Decoder {
	d := &Decoder{
		r:        r,
		dec:      json.NewDecoder(r),
		streamer: streamer,
		options:  o,
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
		dec.delim('[', "opening collection bracket [") //Remove collection opening [
		var n int64 = 1
		for n <= dec.options.NumNetworks {
			err := dec.decodeNetwork(n) //Decode network n
			n++
			if err != nil {
				dec.streamer(&Element{NetworkId: 0, Value: &Element_Error{Error: err}})
				return
			}
		}
		dec.delim(']', "closing collection bracket ]") //Remove collection closing ]
	} else {
		err := dec.decodeNetwork(0) //There's only one network
		if err != nil {
			dec.streamer(&Element{NetworkId: 0, Value: &Element_Error{Error: err}})
			return
		}

	}
	stripEOF(d) //Should find EOF at end of stream
}

//decodeNetwork handles the decoding of elements from a single network, creating elements that are tagged with a network number, networkNum.
//The decoder will then call the streamer that was passed into NewDecoder with the element as the sole parameter. If an fragment is not recognized
//or required, it is discared.
func (dec *Decoder) decodeNetwork(networkNum int64) *Error {
	d := dec.dec
	num := strconv.FormatInt(networkNum, 10)
	err := dec.delim('[', "opening bracket [ of unparsed network " + num) //Opeing network [
	if err != nil {
		return err
	}
	for d.More() { //Iterate over every fragment in the network
		err = dec.delim('{', "opening brace { of an unparsed aspect fragment in network " + num) //Opening fragment {
		if err != nil {
			return err
		}
		fragmentName, err := dec.fragmentName() //Fragment name
		if err != nil {
			return err
		}
		err = dec.delim('[', "opening bracket [ of an unparsed element list in netowrk" + num + ".") //Opening element [
		if err != nil {
			return err
		}
		element, known := KnownAspects[fragmentName]
		required := isAspectRequired(fragmentName, dec.options.RequiredAspects)
		for d.More() { //Iterate over every element in the fragment
			if known && required { //Convert known aspects to protobuf
				jsonpb.UnmarshalNext(d, element)
				dec.streamer(wrapElement(networkNum, element))
			} else { //Consume and do nothing with unknown aspects
				var v map[string]interface{}
				d.Decode(&v)
			}
		}
		err = dec.delim(']', "closing bracket ] of a parsed fragment "+ fragmentName + " element list in network " + num + ".") //Closing element ]
		if err != nil {
			return err
		}
		err = dec.delim('}', "closing brace } of a parsed fragment " + fragmentName + " in network " + num + ".") //Closing fragment }
		if err != nil {
			return err
		}
	}
	err = dec.delim(']', "closing bracket ] of a parsed network " + num) //Closing network ]
	if err != nil {
		return err
	}
	return nil
}

func (d *Decoder) delim(expectedToken rune, desc string) *Error {
	token, err := d.dec.Token()
	if err != nil {
		return &Error{
			Code:    "cy:/cxmate/555",
			Message: "Error decoding token from stream, " + err.Error(),
		}
	}
	delim, ok := token.(json.Delim)
	if !ok {
		return &Error{
			Code:    "cy/cxmate/444",
			Message: "Error in CX stream, could not decode delimiter where " + desc + "should be.",
		}
	}
	if rune(delim) != expectedToken {
		return &Error{
			Code:    "cy://cxmate/666",
			Message: "Error in CX stream, could not find " + desc + ".",
		}
	}
	return nil
}

func (d *Decoder) fragmentName() (string, *Error) {
	token, err := d.dec.Token()
	if err != nil {
		return "", &Error{
			Code:    "cy:/cxmate/555",
			Message: "Error decoding token from stream, " + err.Error(),
		}
	}
	name, ok := token.(string)
	if !ok {
		return "", &Error{
			Code:    "cy/cxmate/777",
			Message: "Error in CX stream, failed to find fragment name in beginning of fragment",
		}
	}
	return name, nil
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
