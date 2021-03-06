// Code generated by protoc-gen-go.
// source: cx.proto
// DO NOT EDIT!

/*
Package cxpb is a generated protocol buffer package.

It is generated from these files:
	cx.proto

It has these top-level messages:
	Element
	NumberVerification
	MetaData
	Property
	Parameter
	Error
	Node
	Edge
	NodeAttribute
	EdgeAttribute
	NetworkAttribute
	CartesianLayout
	AnonymousAspect
*/
package cxpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Element struct {
	NetworkId int64 `protobuf:"varint,1,opt,name=networkId" json:"networkId,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Element_Parameter
	//	*Element_Error
	//	*Element_NumberVerification
	//	*Element_Metadata
	//	*Element_Aspect
	//	*Element_Node
	//	*Element_Edge
	//	*Element_NodeAttribute
	//	*Element_EdgeAttribute
	//	*Element_NetworkAttribute
	//	*Element_CartesianLayout
	Value isElement_Value `protobuf_oneof:"value"`
}

func (m *Element) Reset()                    { *m = Element{} }
func (m *Element) String() string            { return proto.CompactTextString(m) }
func (*Element) ProtoMessage()               {}
func (*Element) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isElement_Value interface {
	isElement_Value()
}

type Element_Parameter struct {
	Parameter *Parameter `protobuf:"bytes,2,opt,name=parameter,oneof"`
}
type Element_Error struct {
	Error *Error `protobuf:"bytes,3,opt,name=error,oneof"`
}
type Element_NumberVerification struct {
	NumberVerification *NumberVerification `protobuf:"bytes,4,opt,name=numberVerification,oneof"`
}
type Element_Metadata struct {
	Metadata *MetaData `protobuf:"bytes,5,opt,name=metadata,oneof"`
}
type Element_Aspect struct {
	Aspect *AnonymousAspect `protobuf:"bytes,6,opt,name=aspect,oneof"`
}
type Element_Node struct {
	Node *Node `protobuf:"bytes,7,opt,name=node,oneof"`
}
type Element_Edge struct {
	Edge *Edge `protobuf:"bytes,8,opt,name=edge,oneof"`
}
type Element_NodeAttribute struct {
	NodeAttribute *NodeAttribute `protobuf:"bytes,9,opt,name=nodeAttribute,oneof"`
}
type Element_EdgeAttribute struct {
	EdgeAttribute *EdgeAttribute `protobuf:"bytes,10,opt,name=edgeAttribute,oneof"`
}
type Element_NetworkAttribute struct {
	NetworkAttribute *NetworkAttribute `protobuf:"bytes,11,opt,name=networkAttribute,oneof"`
}
type Element_CartesianLayout struct {
	CartesianLayout *CartesianLayout `protobuf:"bytes,12,opt,name=cartesianLayout,oneof"`
}

func (*Element_Parameter) isElement_Value()          {}
func (*Element_Error) isElement_Value()              {}
func (*Element_NumberVerification) isElement_Value() {}
func (*Element_Metadata) isElement_Value()           {}
func (*Element_Aspect) isElement_Value()             {}
func (*Element_Node) isElement_Value()               {}
func (*Element_Edge) isElement_Value()               {}
func (*Element_NodeAttribute) isElement_Value()      {}
func (*Element_EdgeAttribute) isElement_Value()      {}
func (*Element_NetworkAttribute) isElement_Value()   {}
func (*Element_CartesianLayout) isElement_Value()    {}

func (m *Element) GetValue() isElement_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Element) GetNetworkId() int64 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

func (m *Element) GetParameter() *Parameter {
	if x, ok := m.GetValue().(*Element_Parameter); ok {
		return x.Parameter
	}
	return nil
}

func (m *Element) GetError() *Error {
	if x, ok := m.GetValue().(*Element_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Element) GetNumberVerification() *NumberVerification {
	if x, ok := m.GetValue().(*Element_NumberVerification); ok {
		return x.NumberVerification
	}
	return nil
}

func (m *Element) GetMetadata() *MetaData {
	if x, ok := m.GetValue().(*Element_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Element) GetAspect() *AnonymousAspect {
	if x, ok := m.GetValue().(*Element_Aspect); ok {
		return x.Aspect
	}
	return nil
}

func (m *Element) GetNode() *Node {
	if x, ok := m.GetValue().(*Element_Node); ok {
		return x.Node
	}
	return nil
}

func (m *Element) GetEdge() *Edge {
	if x, ok := m.GetValue().(*Element_Edge); ok {
		return x.Edge
	}
	return nil
}

func (m *Element) GetNodeAttribute() *NodeAttribute {
	if x, ok := m.GetValue().(*Element_NodeAttribute); ok {
		return x.NodeAttribute
	}
	return nil
}

func (m *Element) GetEdgeAttribute() *EdgeAttribute {
	if x, ok := m.GetValue().(*Element_EdgeAttribute); ok {
		return x.EdgeAttribute
	}
	return nil
}

func (m *Element) GetNetworkAttribute() *NetworkAttribute {
	if x, ok := m.GetValue().(*Element_NetworkAttribute); ok {
		return x.NetworkAttribute
	}
	return nil
}

func (m *Element) GetCartesianLayout() *CartesianLayout {
	if x, ok := m.GetValue().(*Element_CartesianLayout); ok {
		return x.CartesianLayout
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Element) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Element_OneofMarshaler, _Element_OneofUnmarshaler, _Element_OneofSizer, []interface{}{
		(*Element_Parameter)(nil),
		(*Element_Error)(nil),
		(*Element_NumberVerification)(nil),
		(*Element_Metadata)(nil),
		(*Element_Aspect)(nil),
		(*Element_Node)(nil),
		(*Element_Edge)(nil),
		(*Element_NodeAttribute)(nil),
		(*Element_EdgeAttribute)(nil),
		(*Element_NetworkAttribute)(nil),
		(*Element_CartesianLayout)(nil),
	}
}

func _Element_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Element)
	// value
	switch x := m.Value.(type) {
	case *Element_Parameter:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Parameter); err != nil {
			return err
		}
	case *Element_Error:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *Element_NumberVerification:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NumberVerification); err != nil {
			return err
		}
	case *Element_Metadata:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Metadata); err != nil {
			return err
		}
	case *Element_Aspect:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aspect); err != nil {
			return err
		}
	case *Element_Node:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Node); err != nil {
			return err
		}
	case *Element_Edge:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Edge); err != nil {
			return err
		}
	case *Element_NodeAttribute:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodeAttribute); err != nil {
			return err
		}
	case *Element_EdgeAttribute:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EdgeAttribute); err != nil {
			return err
		}
	case *Element_NetworkAttribute:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NetworkAttribute); err != nil {
			return err
		}
	case *Element_CartesianLayout:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CartesianLayout); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Element.Value has unexpected type %T", x)
	}
	return nil
}

func _Element_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Element)
	switch tag {
	case 2: // value.parameter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Parameter)
		err := b.DecodeMessage(msg)
		m.Value = &Element_Parameter{msg}
		return true, err
	case 3: // value.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Error)
		err := b.DecodeMessage(msg)
		m.Value = &Element_Error{msg}
		return true, err
	case 4: // value.numberVerification
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NumberVerification)
		err := b.DecodeMessage(msg)
		m.Value = &Element_NumberVerification{msg}
		return true, err
	case 5: // value.metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MetaData)
		err := b.DecodeMessage(msg)
		m.Value = &Element_Metadata{msg}
		return true, err
	case 6: // value.aspect
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AnonymousAspect)
		err := b.DecodeMessage(msg)
		m.Value = &Element_Aspect{msg}
		return true, err
	case 7: // value.node
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Node)
		err := b.DecodeMessage(msg)
		m.Value = &Element_Node{msg}
		return true, err
	case 8: // value.edge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Edge)
		err := b.DecodeMessage(msg)
		m.Value = &Element_Edge{msg}
		return true, err
	case 9: // value.nodeAttribute
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NodeAttribute)
		err := b.DecodeMessage(msg)
		m.Value = &Element_NodeAttribute{msg}
		return true, err
	case 10: // value.edgeAttribute
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EdgeAttribute)
		err := b.DecodeMessage(msg)
		m.Value = &Element_EdgeAttribute{msg}
		return true, err
	case 11: // value.networkAttribute
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NetworkAttribute)
		err := b.DecodeMessage(msg)
		m.Value = &Element_NetworkAttribute{msg}
		return true, err
	case 12: // value.cartesianLayout
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CartesianLayout)
		err := b.DecodeMessage(msg)
		m.Value = &Element_CartesianLayout{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Element_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Element)
	// value
	switch x := m.Value.(type) {
	case *Element_Parameter:
		s := proto.Size(x.Parameter)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_NumberVerification:
		s := proto.Size(x.NumberVerification)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_Metadata:
		s := proto.Size(x.Metadata)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_Aspect:
		s := proto.Size(x.Aspect)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_Node:
		s := proto.Size(x.Node)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_Edge:
		s := proto.Size(x.Edge)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_NodeAttribute:
		s := proto.Size(x.NodeAttribute)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_EdgeAttribute:
		s := proto.Size(x.EdgeAttribute)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_NetworkAttribute:
		s := proto.Size(x.NetworkAttribute)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Element_CartesianLayout:
		s := proto.Size(x.CartesianLayout)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NumberVerification struct {
	LongNumber int64 `protobuf:"varint,1,opt,name=longNumber" json:"longNumber,omitempty"`
}

func (m *NumberVerification) Reset()                    { *m = NumberVerification{} }
func (m *NumberVerification) String() string            { return proto.CompactTextString(m) }
func (*NumberVerification) ProtoMessage()               {}
func (*NumberVerification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NumberVerification) GetLongNumber() int64 {
	if m != nil {
		return m.LongNumber
	}
	return 0
}

type MetaData struct {
	Name             string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version          string      `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	IdCounter        int64       `protobuf:"varint,3,opt,name=idCounter" json:"idCounter,omitempty"`
	ElementCount     int64       `protobuf:"varint,4,opt,name=elementCount" json:"elementCount,omitempty"`
	ConsistencyGroup int64       `protobuf:"varint,5,opt,name=consistencyGroup" json:"consistencyGroup,omitempty"`
	Checksum         int64       `protobuf:"varint,6,opt,name=checksum" json:"checksum,omitempty"`
	Properties       []*Property `protobuf:"bytes,7,rep,name=properties" json:"properties,omitempty"`
}

func (m *MetaData) Reset()                    { *m = MetaData{} }
func (m *MetaData) String() string            { return proto.CompactTextString(m) }
func (*MetaData) ProtoMessage()               {}
func (*MetaData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MetaData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetaData) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *MetaData) GetIdCounter() int64 {
	if m != nil {
		return m.IdCounter
	}
	return 0
}

func (m *MetaData) GetElementCount() int64 {
	if m != nil {
		return m.ElementCount
	}
	return 0
}

func (m *MetaData) GetConsistencyGroup() int64 {
	if m != nil {
		return m.ConsistencyGroup
	}
	return 0
}

func (m *MetaData) GetChecksum() int64 {
	if m != nil {
		return m.Checksum
	}
	return 0
}

func (m *MetaData) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Property struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Property) Reset()                    { *m = Property{} }
func (m *Property) String() string            { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()               {}
func (*Property) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Property) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Property) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Parameter struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Parameter) Reset()                    { *m = Parameter{} }
func (m *Parameter) String() string            { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()               {}
func (*Parameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Error struct {
	Status  int64  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Link    string `protobuf:"bytes,4,opt,name=link" json:"link,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Error) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type Node struct {
	Id         int64  `protobuf:"varint,1,opt,name=id,json=@id" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,json=n" json:"name,omitempty"`
	Represents string `protobuf:"bytes,3,opt,name=represents,json=r" json:"represents,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Node) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetRepresents() string {
	if m != nil {
		return m.Represents
	}
	return ""
}

type Edge struct {
	Id          int64  `protobuf:"varint,1,opt,name=id,json=@id" json:"id,omitempty"`
	SourceId    int64  `protobuf:"varint,2,opt,name=sourceId,json=s" json:"sourceId,omitempty"`
	TargetId    int64  `protobuf:"varint,3,opt,name=targetId,json=t" json:"targetId,omitempty"`
	Interaction string `protobuf:"bytes,4,opt,name=interaction,json=i" json:"interaction,omitempty"`
}

func (m *Edge) Reset()                    { *m = Edge{} }
func (m *Edge) String() string            { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()               {}
func (*Edge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Edge) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Edge) GetSourceId() int64 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *Edge) GetTargetId() int64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Edge) GetInteraction() string {
	if m != nil {
		return m.Interaction
	}
	return ""
}

type NodeAttribute struct {
	NodeId   int64  `protobuf:"varint,1,opt,name=nodeId,json=po" json:"nodeId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,json=n" json:"name,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,json=v" json:"value,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,json=d" json:"type,omitempty"`
	SubnetId int64  `protobuf:"varint,5,opt,name=subnetId,json=s" json:"subnetId,omitempty"`
}

func (m *NodeAttribute) Reset()                    { *m = NodeAttribute{} }
func (m *NodeAttribute) String() string            { return proto.CompactTextString(m) }
func (*NodeAttribute) ProtoMessage()               {}
func (*NodeAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *NodeAttribute) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *NodeAttribute) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NodeAttribute) GetSubnetId() int64 {
	if m != nil {
		return m.SubnetId
	}
	return 0
}

type EdgeAttribute struct {
	EdgeId   int64  `protobuf:"varint,1,opt,name=edgeId,json=po" json:"edgeId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,json=n" json:"name,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,json=v" json:"value,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,json=d" json:"type,omitempty"`
	SubnetId int64  `protobuf:"varint,5,opt,name=subnetId,json=s" json:"subnetId,omitempty"`
}

func (m *EdgeAttribute) Reset()                    { *m = EdgeAttribute{} }
func (m *EdgeAttribute) String() string            { return proto.CompactTextString(m) }
func (*EdgeAttribute) ProtoMessage()               {}
func (*EdgeAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *EdgeAttribute) GetEdgeId() int64 {
	if m != nil {
		return m.EdgeId
	}
	return 0
}

func (m *EdgeAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EdgeAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *EdgeAttribute) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EdgeAttribute) GetSubnetId() int64 {
	if m != nil {
		return m.SubnetId
	}
	return 0
}

type NetworkAttribute struct {
	EdgeId   int64  `protobuf:"varint,1,opt,name=edgeId,json=po" json:"edgeId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,json=n" json:"name,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,json=v" json:"value,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,json=d" json:"type,omitempty"`
	SubnetId int64  `protobuf:"varint,5,opt,name=subnetId,json=s" json:"subnetId,omitempty"`
}

func (m *NetworkAttribute) Reset()                    { *m = NetworkAttribute{} }
func (m *NetworkAttribute) String() string            { return proto.CompactTextString(m) }
func (*NetworkAttribute) ProtoMessage()               {}
func (*NetworkAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *NetworkAttribute) GetEdgeId() int64 {
	if m != nil {
		return m.EdgeId
	}
	return 0
}

func (m *NetworkAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *NetworkAttribute) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkAttribute) GetSubnetId() int64 {
	if m != nil {
		return m.SubnetId
	}
	return 0
}

type CartesianLayout struct {
	Nodeid int64   `protobuf:"varint,1,opt,name=nodeid" json:"nodeid,omitempty"`
	X      float64 `protobuf:"fixed64,2,opt,name=x" json:"x,omitempty"`
	Y      float64 `protobuf:"fixed64,3,opt,name=y" json:"y,omitempty"`
	Viewid int64   `protobuf:"varint,4,opt,name=viewid" json:"viewid,omitempty"`
}

func (m *CartesianLayout) Reset()                    { *m = CartesianLayout{} }
func (m *CartesianLayout) String() string            { return proto.CompactTextString(m) }
func (*CartesianLayout) ProtoMessage()               {}
func (*CartesianLayout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CartesianLayout) GetNodeid() int64 {
	if m != nil {
		return m.Nodeid
	}
	return 0
}

func (m *CartesianLayout) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *CartesianLayout) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *CartesianLayout) GetViewid() int64 {
	if m != nil {
		return m.Viewid
	}
	return 0
}

type AnonymousAspect struct {
	Type    string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Element []byte `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`
}

func (m *AnonymousAspect) Reset()                    { *m = AnonymousAspect{} }
func (m *AnonymousAspect) String() string            { return proto.CompactTextString(m) }
func (*AnonymousAspect) ProtoMessage()               {}
func (*AnonymousAspect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AnonymousAspect) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AnonymousAspect) GetElement() []byte {
	if m != nil {
		return m.Element
	}
	return nil
}

func init() {
	proto.RegisterType((*Element)(nil), "cxpb.Element")
	proto.RegisterType((*NumberVerification)(nil), "cxpb.NumberVerification")
	proto.RegisterType((*MetaData)(nil), "cxpb.MetaData")
	proto.RegisterType((*Property)(nil), "cxpb.Property")
	proto.RegisterType((*Parameter)(nil), "cxpb.Parameter")
	proto.RegisterType((*Error)(nil), "cxpb.Error")
	proto.RegisterType((*Node)(nil), "cxpb.Node")
	proto.RegisterType((*Edge)(nil), "cxpb.Edge")
	proto.RegisterType((*NodeAttribute)(nil), "cxpb.NodeAttribute")
	proto.RegisterType((*EdgeAttribute)(nil), "cxpb.EdgeAttribute")
	proto.RegisterType((*NetworkAttribute)(nil), "cxpb.NetworkAttribute")
	proto.RegisterType((*CartesianLayout)(nil), "cxpb.CartesianLayout")
	proto.RegisterType((*AnonymousAspect)(nil), "cxpb.AnonymousAspect")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CyService service

type CyServiceClient interface {
	StreamElements(ctx context.Context, opts ...grpc.CallOption) (CyService_StreamElementsClient, error)
}

type cyServiceClient struct {
	cc *grpc.ClientConn
}

func NewCyServiceClient(cc *grpc.ClientConn) CyServiceClient {
	return &cyServiceClient{cc}
}

func (c *cyServiceClient) StreamElements(ctx context.Context, opts ...grpc.CallOption) (CyService_StreamElementsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CyService_serviceDesc.Streams[0], c.cc, "/cxpb.CyService/StreamElements", opts...)
	if err != nil {
		return nil, err
	}
	x := &cyServiceStreamElementsClient{stream}
	return x, nil
}

type CyService_StreamElementsClient interface {
	Send(*Element) error
	Recv() (*Element, error)
	grpc.ClientStream
}

type cyServiceStreamElementsClient struct {
	grpc.ClientStream
}

func (x *cyServiceStreamElementsClient) Send(m *Element) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cyServiceStreamElementsClient) Recv() (*Element, error) {
	m := new(Element)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for CyService service

type CyServiceServer interface {
	StreamElements(CyService_StreamElementsServer) error
}

func RegisterCyServiceServer(s *grpc.Server, srv CyServiceServer) {
	s.RegisterService(&_CyService_serviceDesc, srv)
}

func _CyService_StreamElements_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CyServiceServer).StreamElements(&cyServiceStreamElementsServer{stream})
}

type CyService_StreamElementsServer interface {
	Send(*Element) error
	Recv() (*Element, error)
	grpc.ServerStream
}

type cyServiceStreamElementsServer struct {
	grpc.ServerStream
}

func (x *cyServiceStreamElementsServer) Send(m *Element) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cyServiceStreamElementsServer) Recv() (*Element, error) {
	m := new(Element)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cxpb.CyService",
	HandlerType: (*CyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamElements",
			Handler:       _CyService_StreamElements_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cx.proto",
}

func init() { proto.RegisterFile("cx.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xdc, 0x36,
	0x10, 0x15, 0xad, 0xfd, 0xd2, 0xd8, 0xce, 0x1a, 0x4c, 0x63, 0x08, 0x41, 0x51, 0x18, 0xea, 0xc5,
	0x28, 0x0a, 0xa7, 0x48, 0xdd, 0x53, 0x0f, 0xe9, 0xd6, 0x31, 0x2a, 0x17, 0x6d, 0x11, 0x30, 0x40,
	0x81, 0x1e, 0x7a, 0xe0, 0x4a, 0xd3, 0x2d, 0xe1, 0x15, 0x29, 0x90, 0xd4, 0xc6, 0xfa, 0xdb, 0x3d,
	0xf5, 0x58, 0x90, 0xa2, 0xb4, 0x5f, 0xb9, 0xf4, 0x90, 0x1b, 0x67, 0xe6, 0xbd, 0x99, 0xd1, 0xf0,
	0x0d, 0x05, 0xb3, 0xe2, 0xe9, 0xa6, 0xd6, 0xca, 0x2a, 0x3a, 0x2a, 0x9e, 0xea, 0x65, 0xf6, 0xef,
	0x08, 0xa6, 0xf7, 0x6b, 0xac, 0x50, 0x5a, 0xfa, 0x39, 0x24, 0x12, 0xed, 0x07, 0xa5, 0x1f, 0x1f,
	0xca, 0x94, 0x5c, 0x91, 0xeb, 0x98, 0x6d, 0x1d, 0xf4, 0x15, 0x24, 0x35, 0xd7, 0xbc, 0x42, 0x8b,
	0x3a, 0x3d, 0xb9, 0x22, 0xd7, 0xa7, 0xaf, 0xe7, 0x37, 0x2e, 0xc7, 0xcd, 0xbb, 0xde, 0x9d, 0x47,
	0x6c, 0x8b, 0xa1, 0x5f, 0xc2, 0x18, 0xb5, 0x56, 0x3a, 0x8d, 0x3d, 0xf8, 0xb4, 0x03, 0xdf, 0x3b,
	0x57, 0x1e, 0xb1, 0x2e, 0x46, 0x7f, 0x06, 0x2a, 0x9b, 0x6a, 0x89, 0xfa, 0x77, 0xd4, 0xe2, 0x2f,
	0x51, 0x70, 0x2b, 0x94, 0x4c, 0x47, 0x9e, 0x91, 0x76, 0x8c, 0xdf, 0x8e, 0xe2, 0x79, 0xc4, 0x3e,
	0xc2, 0xa2, 0x5f, 0xc3, 0xac, 0x42, 0xcb, 0x4b, 0x6e, 0x79, 0x3a, 0xf6, 0x19, 0x9e, 0x75, 0x19,
	0x7e, 0x45, 0xcb, 0xdf, 0x72, 0xcb, 0xf3, 0x88, 0x0d, 0x08, 0xfa, 0x0a, 0x26, 0xdc, 0xd4, 0x58,
	0xd8, 0x74, 0xe2, 0xb1, 0x2f, 0x3a, 0xec, 0x42, 0x2a, 0xd9, 0x56, 0xaa, 0x31, 0x0b, 0x1f, 0xcc,
	0x23, 0x16, 0x60, 0xf4, 0x0a, 0x46, 0x52, 0x95, 0x98, 0x4e, 0x3d, 0x1c, 0x42, 0x73, 0xaa, 0xc4,
	0x3c, 0x62, 0x3e, 0xe2, 0x10, 0x58, 0xae, 0x30, 0x9d, 0xed, 0x22, 0xee, 0xcb, 0x95, 0x47, 0xb8,
	0x08, 0xfd, 0x1e, 0xce, 0x1d, 0x72, 0x61, 0xad, 0x16, 0xcb, 0xc6, 0x62, 0x9a, 0x78, 0xe8, 0xf3,
	0x6d, 0xb2, 0x21, 0x94, 0x47, 0x6c, 0x1f, 0xeb, 0xc8, 0x2e, 0xc9, 0x96, 0x0c, 0xbb, 0xe4, 0xfb,
	0xdd, 0x90, 0x23, 0xef, 0x61, 0xe9, 0x5b, 0xb8, 0x08, 0x77, 0xb9, 0xe5, 0x9f, 0x7a, 0xfe, 0x65,
	0x28, 0x7e, 0x10, 0xcd, 0x23, 0x76, 0xc4, 0xa0, 0x0b, 0x98, 0x17, 0x5c, 0x5b, 0x34, 0x82, 0xcb,
	0x5f, 0x78, 0xab, 0x1a, 0x9b, 0x9e, 0xed, 0x4e, 0xef, 0x6e, 0x3f, 0x98, 0x47, 0xec, 0x10, 0xff,
	0xe3, 0x14, 0xc6, 0x1b, 0xbe, 0x6e, 0x30, 0xbb, 0x05, 0x7a, 0x7c, 0xb5, 0xf4, 0x0b, 0x80, 0xb5,
	0x92, 0xab, 0x2e, 0x12, 0x54, 0xb8, 0xe3, 0xc9, 0xfe, 0x21, 0x30, 0xeb, 0xef, 0x93, 0x52, 0x18,
	0x49, 0x5e, 0xa1, 0x87, 0x25, 0xcc, 0x9f, 0x69, 0x0a, 0xd3, 0x0d, 0x6a, 0xe3, 0x64, 0x74, 0xe2,
	0xdd, 0xbd, 0xe9, 0xf4, 0x2d, 0xca, 0x3b, 0xd5, 0x48, 0xa7, 0xe0, 0xb8, 0xd3, 0xf7, 0xe0, 0xa0,
	0x19, 0x9c, 0x61, 0xb7, 0x08, 0xde, 0xe3, 0x35, 0x18, 0xb3, 0x3d, 0x1f, 0xfd, 0x0a, 0x2e, 0x0a,
	0x25, 0x8d, 0x30, 0x16, 0x65, 0xd1, 0xfe, 0xa4, 0x55, 0x53, 0x7b, 0xa5, 0xc5, 0xec, 0xc8, 0x4f,
	0x5f, 0xc2, 0xac, 0xf8, 0x1b, 0x8b, 0x47, 0xd3, 0x54, 0x5e, 0x61, 0x31, 0x1b, 0x6c, 0x7a, 0x03,
	0x50, 0x6b, 0x55, 0xa3, 0xb6, 0x02, 0x4d, 0x3a, 0xbd, 0x8a, 0xb7, 0x5a, 0x7d, 0xd7, 0xf9, 0x5b,
	0xb6, 0x83, 0xc8, 0x6e, 0x61, 0xd6, 0xfb, 0x3f, 0xfa, 0xcd, 0x9f, 0x85, 0x99, 0x86, 0x2f, 0x0e,
	0x03, 0xfe, 0x0e, 0x92, 0x61, 0x35, 0xff, 0x07, 0x8d, 0xc3, 0xd8, 0x2f, 0x29, 0xbd, 0x84, 0x89,
	0xb1, 0xdc, 0x36, 0x26, 0x5c, 0x43, 0xb0, 0x5c, 0xaa, 0xc2, 0x2d, 0x42, 0xc7, 0xf2, 0x67, 0x37,
	0xf5, 0x0a, 0x8d, 0xe1, 0x2b, 0xf4, 0x93, 0x4d, 0x58, 0x6f, 0x3a, 0xf4, 0x5a, 0xc8, 0x47, 0x3f,
	0xcf, 0x84, 0xf9, 0x73, 0xf6, 0x06, 0x46, 0x4e, 0xeb, 0x74, 0x0e, 0x27, 0xa2, 0x7f, 0x6a, 0xe2,
	0x1f, 0x44, 0x49, 0xe7, 0xa1, 0xcb, 0x2e, 0x35, 0x91, 0xf4, 0x05, 0x80, 0xc6, 0x5a, 0xa3, 0x41,
	0x69, 0x4d, 0x48, 0x4d, 0x74, 0xf6, 0x07, 0x8c, 0x9c, 0xde, 0x8f, 0x13, 0x3c, 0x87, 0x99, 0x51,
	0x8d, 0x2e, 0xf0, 0xa1, 0xf4, 0x49, 0x62, 0x46, 0x8c, 0x73, 0x5a, 0xae, 0x57, 0x68, 0x1f, 0xca,
	0x70, 0xef, 0xc4, 0xd2, 0x4b, 0x38, 0x15, 0xee, 0xe2, 0x79, 0x31, 0x3c, 0x39, 0x09, 0x23, 0x22,
	0x5b, 0xc3, 0xf9, 0xde, 0x1e, 0x52, 0x0a, 0x13, 0xb7, 0x87, 0xc3, 0x9b, 0x78, 0x52, 0xab, 0xe3,
	0x3e, 0x2f, 0xfa, 0x51, 0x86, 0x16, 0x37, 0x0e, 0x62, 0xdb, 0x1a, 0xfb, 0xc4, 0x5d, 0x6b, 0xcd,
	0x52, 0xfa, 0x2e, 0xc6, 0xa1, 0x35, 0x57, 0x6d, 0x6f, 0x71, 0x5d, 0x35, 0xb7, 0xb8, 0x9f, 0xa6,
	0x9a, 0x82, 0x8b, 0xc3, 0x35, 0xff, 0xb4, 0x05, 0xff, 0x84, 0xf9, 0xc1, 0x93, 0xe0, 0x54, 0xe5,
	0xc6, 0x39, 0x5c, 0x5b, 0xb0, 0xe8, 0x19, 0x90, 0x27, 0x5f, 0x90, 0x30, 0xf2, 0xe4, 0xac, 0xd6,
	0x17, 0x23, 0x8c, 0xb4, 0x8e, 0xb3, 0x11, 0xf8, 0x41, 0x94, 0x61, 0x2b, 0x83, 0x95, 0xbd, 0x81,
	0xf9, 0xc1, 0x7b, 0xed, 0xe4, 0xe6, 0xfb, 0x0a, 0x3a, 0x77, 0x67, 0x27, 0xce, 0xb0, 0xc6, 0xbe,
	0xc0, 0x19, 0xeb, 0xcd, 0xd7, 0x0b, 0x48, 0xee, 0xda, 0xf7, 0xa8, 0x37, 0xa2, 0x40, 0x7a, 0x0b,
	0xcf, 0xde, 0x5b, 0x8d, 0xbc, 0x0a, 0x3f, 0x44, 0x43, 0xcf, 0xc3, 0xd3, 0xda, 0xd9, 0x2f, 0xf7,
	0xcd, 0x2c, 0xba, 0x26, 0xdf, 0x90, 0xe5, 0xc4, 0xff, 0x4e, 0xbf, 0xfd, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x89, 0xab, 0x5f, 0x5a, 0x07, 0x00, 0x00,
}
