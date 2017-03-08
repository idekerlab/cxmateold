package cxpb

import (
  "github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
  "io"
)

type Encoder struct {
  w io.Writer
  m *jsonpb.Marshaler
  stream func() (*Element, error)
  options *EncoderOptions
  curr *current
}

type EncoderOptions struct {
  RequiredAspects []string
	IsCollection    bool
	NumNetworks     int64
}

type current struct {
  networkId int
  aspect string
  message proto.Message
  err error
}

func NewEncoder(w io.Writer, o *EncoderOptions, stream func() (*Element, error)) *Encoder {
  return &Encoder{
    w: w,
    m: &jsonpb.Marshaler{},
    stream: stream,
    options: o,
    curr: &current{},
  }
}

func (e *Encoder) Encode() {
  e.fetchNext()
  if e.options.IsCollection {
    e.emit("[")
    for i := 0; int64(i) < e.options.NumNetworks; i++ {
      if i != 0 {
        e.emit(",")
      }
      e.encodeNetwork()
    }
    e.emit("]")
  } else {
    e.encodeNetwork()
  }
}

func (e *Encoder) encodeNetwork() {
  e.emit("[")
  e.emitNumberVerification()
  e.emit(",")
  e.emitMetaData()
  openFragment := ""
  currNetwork := e.curr.networkId
  for (e.curr.err != io.EOF) && (e.curr.networkId == currNetwork) {
    if e.curr.aspect != openFragment { //A new fragment is needed
      if openFragment != "" { //This is not first fragment
         e.emit("]},")
      } else {
        e.emit(",")
      }
      e.emitOpenFragment()
      openFragment = e.curr.aspect //Set the curren
    } else {
      e.emit(",")
      e.emitElement()
    }
    e.fetchNext()
  }
  e.emit("]}]")
}


//emitNumberVerification emits a static stanza indicating the max value of long integer on this system so that the caller
//can verify that all the number sent in the response can be represented by the callers runtime.
func (e *Encoder) emitNumberVerification() {
  e.emit("{\"numberVerification\":[{\"longNumber\":\"281474976710655\"}]}")
}

//emitMetaData emits a static stanza which contains the metadata for the emitted network. The encoder always emits a NEW network
//which allows emitMetaData to create simple default metadata objects for the newly minted aspects being sent.
func (e *Encoder) emitMetaData() {
  e.emit("{\"metaData\":[")
  for index, name := range e.options.RequiredAspects {
    if index != 0 {
      e.emit(",")
    }
    e.emit("{\"consistencyGroup\":1,\"name\":\"" + name + "\",\"properties\":[],\"version\":\"1.0\"}")
  }
  e.emit("]}")
}

func (e *Encoder) emitOpenFragment() {
  io.WriteString(e.w, "{\"" + e.curr.aspect + "\":[")
  e.emitElement()
}

func (e *Encoder) emitAppendElement() {
  io.WriteString(e.w, ",")
  e.emitElement()
}

func (e *Encoder) emitElement() {
  if err := e.m.Marshal(e.w, e.curr.message); err != nil {
    panic("Could not encode element")
  }
}

func (e *Encoder) emit(token string) {
  io.WriteString(e.w, token)
}

func (e *Encoder) fetchNext() {
  in, err := e.stream()
  if err != nil {
    e.curr.err = err
    return
  }
  name, message := unwrapElement(in)
  if name == "metadata" {
    e.fetchNext()
  } else {
    e.curr.networkId = int(in.NetworkId)
    e.curr.aspect = name
    e.curr.message = message
  }
}

func unwrapElement(ele *Element) (string, proto.Message) {
	switch ele.Value.(type) {
  case *Element_Metadata:
    return "metadata", ele.GetMetadata()
	case *Element_Node:
		return "nodes", ele.GetNode()
	case *Element_Edge:
		return "edges", ele.GetEdge()
	case *Element_NodeAttribute:
		return "nodeAttributes", ele.GetNodeAttribute()
	case *Element_EdgeAttribute:
		return "edgeAttributes", ele.GetEdgeAttribute()
	case *Element_NetworkAttribute:
		return "networkAttributes", ele.GetNetworkAttribute()
	default:
		return "aspect", ele.GetAspect()
	}
}
