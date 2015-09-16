// Code generated by protoc-gen-gogo.
// source: fuzz2.proto
// DO NOT EDIT!

/*
Package fuzztests is a generated protocol buffer package.

It is generated from these files:
	fuzz2.proto
	fuzz3.proto

It has these top-level messages:
	NinOptNative
	NinRepNative
	NinOptStruct
	NinRepStruct
	NinNestedStruct
	Nil
	NestedDefinition
	NestedScope
	NinOptNativeDefault
	NinOptNative3
	NinRepNative3
	NinOptStruct3
	NinRepStruct3
	NinNestedStruct3
	Nil3
	NestedDefinition3
	NestedScope3
*/
package fuzztests

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NinOptNative struct {
	Field1           *float64 `protobuf:"fixed64,1,opt,name=Field1" json:"Field1,omitempty"`
	Field2           *float32 `protobuf:"fixed32,2,opt,name=Field2" json:"Field2,omitempty"`
	Field3           *int32   `protobuf:"varint,3,opt,name=Field3" json:"Field3,omitempty"`
	Field4           *int64   `protobuf:"varint,4,opt,name=Field4" json:"Field4,omitempty"`
	Field5           *uint32  `protobuf:"varint,5,opt,name=Field5" json:"Field5,omitempty"`
	Field6           *uint64  `protobuf:"varint,6,opt,name=Field6" json:"Field6,omitempty"`
	Field7           *int32   `protobuf:"zigzag32,7,opt,name=Field7" json:"Field7,omitempty"`
	Field8           *int64   `protobuf:"zigzag64,8,opt,name=Field8" json:"Field8,omitempty"`
	Field9           *uint32  `protobuf:"fixed32,9,opt,name=Field9" json:"Field9,omitempty"`
	Field10          *int32   `protobuf:"fixed32,10,opt,name=Field10" json:"Field10,omitempty"`
	Field11          *uint64  `protobuf:"fixed64,11,opt,name=Field11" json:"Field11,omitempty"`
	Field12          *int64   `protobuf:"fixed64,12,opt,name=Field12" json:"Field12,omitempty"`
	Field13          *bool    `protobuf:"varint,13,opt,name=Field13" json:"Field13,omitempty"`
	Field14          *string  `protobuf:"bytes,14,opt,name=Field14" json:"Field14,omitempty"`
	Field15          []byte   `protobuf:"bytes,15,opt,name=Field15" json:"Field15,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NinOptNative) Reset()         { *m = NinOptNative{} }
func (m *NinOptNative) String() string { return proto.CompactTextString(m) }
func (*NinOptNative) ProtoMessage()    {}

func (m *NinOptNative) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

func (m *NinOptNative) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

func (m *NinOptNative) GetField3() int32 {
	if m != nil && m.Field3 != nil {
		return *m.Field3
	}
	return 0
}

func (m *NinOptNative) GetField4() int64 {
	if m != nil && m.Field4 != nil {
		return *m.Field4
	}
	return 0
}

func (m *NinOptNative) GetField5() uint32 {
	if m != nil && m.Field5 != nil {
		return *m.Field5
	}
	return 0
}

func (m *NinOptNative) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return 0
}

func (m *NinOptNative) GetField7() int32 {
	if m != nil && m.Field7 != nil {
		return *m.Field7
	}
	return 0
}

func (m *NinOptNative) GetField8() int64 {
	if m != nil && m.Field8 != nil {
		return *m.Field8
	}
	return 0
}

func (m *NinOptNative) GetField9() uint32 {
	if m != nil && m.Field9 != nil {
		return *m.Field9
	}
	return 0
}

func (m *NinOptNative) GetField10() int32 {
	if m != nil && m.Field10 != nil {
		return *m.Field10
	}
	return 0
}

func (m *NinOptNative) GetField11() uint64 {
	if m != nil && m.Field11 != nil {
		return *m.Field11
	}
	return 0
}

func (m *NinOptNative) GetField12() int64 {
	if m != nil && m.Field12 != nil {
		return *m.Field12
	}
	return 0
}

func (m *NinOptNative) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return false
}

func (m *NinOptNative) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *NinOptNative) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NinRepNative struct {
	Field1           []float64 `protobuf:"fixed64,1,rep,name=Field1" json:"Field1,omitempty"`
	Field2           []float32 `protobuf:"fixed32,2,rep,name=Field2" json:"Field2,omitempty"`
	Field3           []int32   `protobuf:"varint,3,rep,name=Field3" json:"Field3,omitempty"`
	Field4           []int64   `protobuf:"varint,4,rep,name=Field4" json:"Field4,omitempty"`
	Field5           []uint32  `protobuf:"varint,5,rep,name=Field5" json:"Field5,omitempty"`
	Field6           []uint64  `protobuf:"varint,6,rep,name=Field6" json:"Field6,omitempty"`
	Field7           []int32   `protobuf:"zigzag32,7,rep,name=Field7" json:"Field7,omitempty"`
	Field8           []int64   `protobuf:"zigzag64,8,rep,name=Field8" json:"Field8,omitempty"`
	Field9           []uint32  `protobuf:"fixed32,9,rep,name=Field9" json:"Field9,omitempty"`
	Field10          []int32   `protobuf:"fixed32,10,rep,name=Field10" json:"Field10,omitempty"`
	Field11          []uint64  `protobuf:"fixed64,11,rep,name=Field11" json:"Field11,omitempty"`
	Field12          []int64   `protobuf:"fixed64,12,rep,name=Field12" json:"Field12,omitempty"`
	Field13          []bool    `protobuf:"varint,13,rep,name=Field13" json:"Field13,omitempty"`
	Field14          []string  `protobuf:"bytes,14,rep,name=Field14" json:"Field14,omitempty"`
	Field15          [][]byte  `protobuf:"bytes,15,rep,name=Field15" json:"Field15,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *NinRepNative) Reset()         { *m = NinRepNative{} }
func (m *NinRepNative) String() string { return proto.CompactTextString(m) }
func (*NinRepNative) ProtoMessage()    {}

func (m *NinRepNative) GetField1() []float64 {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NinRepNative) GetField2() []float32 {
	if m != nil {
		return m.Field2
	}
	return nil
}

func (m *NinRepNative) GetField3() []int32 {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinRepNative) GetField4() []int64 {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinRepNative) GetField5() []uint32 {
	if m != nil {
		return m.Field5
	}
	return nil
}

func (m *NinRepNative) GetField6() []uint64 {
	if m != nil {
		return m.Field6
	}
	return nil
}

func (m *NinRepNative) GetField7() []int32 {
	if m != nil {
		return m.Field7
	}
	return nil
}

func (m *NinRepNative) GetField8() []int64 {
	if m != nil {
		return m.Field8
	}
	return nil
}

func (m *NinRepNative) GetField9() []uint32 {
	if m != nil {
		return m.Field9
	}
	return nil
}

func (m *NinRepNative) GetField10() []int32 {
	if m != nil {
		return m.Field10
	}
	return nil
}

func (m *NinRepNative) GetField11() []uint64 {
	if m != nil {
		return m.Field11
	}
	return nil
}

func (m *NinRepNative) GetField12() []int64 {
	if m != nil {
		return m.Field12
	}
	return nil
}

func (m *NinRepNative) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

func (m *NinRepNative) GetField14() []string {
	if m != nil {
		return m.Field14
	}
	return nil
}

func (m *NinRepNative) GetField15() [][]byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NinOptStruct struct {
	Field1           *float64      `protobuf:"fixed64,1,opt,name=Field1" json:"Field1,omitempty"`
	Field2           *float32      `protobuf:"fixed32,2,opt,name=Field2" json:"Field2,omitempty"`
	Field3           *NinOptNative `protobuf:"bytes,3,opt,name=Field3" json:"Field3,omitempty"`
	Field4           *NinOptNative `protobuf:"bytes,4,opt,name=Field4" json:"Field4,omitempty"`
	Field6           *uint64       `protobuf:"varint,6,opt,name=Field6" json:"Field6,omitempty"`
	Field7           *int32        `protobuf:"zigzag32,7,opt,name=Field7" json:"Field7,omitempty"`
	Field8           *NinOptNative `protobuf:"bytes,8,opt,name=Field8" json:"Field8,omitempty"`
	Field13          *bool         `protobuf:"varint,13,opt,name=Field13" json:"Field13,omitempty"`
	Field14          *string       `protobuf:"bytes,14,opt,name=Field14" json:"Field14,omitempty"`
	Field15          []byte        `protobuf:"bytes,15,opt,name=Field15" json:"Field15,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NinOptStruct) Reset()         { *m = NinOptStruct{} }
func (m *NinOptStruct) String() string { return proto.CompactTextString(m) }
func (*NinOptStruct) ProtoMessage()    {}

func (m *NinOptStruct) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

func (m *NinOptStruct) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

func (m *NinOptStruct) GetField3() *NinOptNative {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinOptStruct) GetField4() *NinOptNative {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinOptStruct) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return 0
}

func (m *NinOptStruct) GetField7() int32 {
	if m != nil && m.Field7 != nil {
		return *m.Field7
	}
	return 0
}

func (m *NinOptStruct) GetField8() *NinOptNative {
	if m != nil {
		return m.Field8
	}
	return nil
}

func (m *NinOptStruct) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return false
}

func (m *NinOptStruct) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *NinOptStruct) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NinRepStruct struct {
	Field1           []float64       `protobuf:"fixed64,1,rep,name=Field1" json:"Field1,omitempty"`
	Field2           []float32       `protobuf:"fixed32,2,rep,name=Field2" json:"Field2,omitempty"`
	Field3           []*NinOptNative `protobuf:"bytes,3,rep,name=Field3" json:"Field3,omitempty"`
	Field4           []*NinOptNative `protobuf:"bytes,4,rep,name=Field4" json:"Field4,omitempty"`
	Field6           []uint64        `protobuf:"varint,6,rep,name=Field6" json:"Field6,omitempty"`
	Field7           []int32         `protobuf:"zigzag32,7,rep,name=Field7" json:"Field7,omitempty"`
	Field8           []*NinOptNative `protobuf:"bytes,8,rep,name=Field8" json:"Field8,omitempty"`
	Field13          []bool          `protobuf:"varint,13,rep,name=Field13" json:"Field13,omitempty"`
	Field14          []string        `protobuf:"bytes,14,rep,name=Field14" json:"Field14,omitempty"`
	Field15          [][]byte        `protobuf:"bytes,15,rep,name=Field15" json:"Field15,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NinRepStruct) Reset()         { *m = NinRepStruct{} }
func (m *NinRepStruct) String() string { return proto.CompactTextString(m) }
func (*NinRepStruct) ProtoMessage()    {}

func (m *NinRepStruct) GetField1() []float64 {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NinRepStruct) GetField2() []float32 {
	if m != nil {
		return m.Field2
	}
	return nil
}

func (m *NinRepStruct) GetField3() []*NinOptNative {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinRepStruct) GetField4() []*NinOptNative {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinRepStruct) GetField6() []uint64 {
	if m != nil {
		return m.Field6
	}
	return nil
}

func (m *NinRepStruct) GetField7() []int32 {
	if m != nil {
		return m.Field7
	}
	return nil
}

func (m *NinRepStruct) GetField8() []*NinOptNative {
	if m != nil {
		return m.Field8
	}
	return nil
}

func (m *NinRepStruct) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

func (m *NinRepStruct) GetField14() []string {
	if m != nil {
		return m.Field14
	}
	return nil
}

func (m *NinRepStruct) GetField15() [][]byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NinNestedStruct struct {
	Field1           *NinOptStruct   `protobuf:"bytes,1,opt,name=Field1" json:"Field1,omitempty"`
	Field2           []*NinRepStruct `protobuf:"bytes,2,rep,name=Field2" json:"Field2,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NinNestedStruct) Reset()         { *m = NinNestedStruct{} }
func (m *NinNestedStruct) String() string { return proto.CompactTextString(m) }
func (*NinNestedStruct) ProtoMessage()    {}

func (m *NinNestedStruct) GetField1() *NinOptStruct {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NinNestedStruct) GetField2() []*NinRepStruct {
	if m != nil {
		return m.Field2
	}
	return nil
}

type Nil struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Nil) Reset()         { *m = Nil{} }
func (m *Nil) String() string { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()    {}

type NestedDefinition struct {
	Field1           *int64                                          `protobuf:"varint,1,opt,name=Field1" json:"Field1,omitempty"`
	NNM              *NestedDefinition_NestedMessage_NestedNestedMsg `protobuf:"bytes,3,opt,name=NNM" json:"NNM,omitempty"`
	NM               *NestedDefinition_NestedMessage                 `protobuf:"bytes,4,opt,name=NM" json:"NM,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *NestedDefinition) Reset()         { *m = NestedDefinition{} }
func (m *NestedDefinition) String() string { return proto.CompactTextString(m) }
func (*NestedDefinition) ProtoMessage()    {}

func (m *NestedDefinition) GetField1() int64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

func (m *NestedDefinition) GetNNM() *NestedDefinition_NestedMessage_NestedNestedMsg {
	if m != nil {
		return m.NNM
	}
	return nil
}

func (m *NestedDefinition) GetNM() *NestedDefinition_NestedMessage {
	if m != nil {
		return m.NM
	}
	return nil
}

type NestedDefinition_NestedMessage struct {
	NestedField1     *uint64                                         `protobuf:"fixed64,1,opt,name=NestedField1" json:"NestedField1,omitempty"`
	NNM              *NestedDefinition_NestedMessage_NestedNestedMsg `protobuf:"bytes,2,opt,name=NNM" json:"NNM,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *NestedDefinition_NestedMessage) Reset()         { *m = NestedDefinition_NestedMessage{} }
func (m *NestedDefinition_NestedMessage) String() string { return proto.CompactTextString(m) }
func (*NestedDefinition_NestedMessage) ProtoMessage()    {}

func (m *NestedDefinition_NestedMessage) GetNestedField1() uint64 {
	if m != nil && m.NestedField1 != nil {
		return *m.NestedField1
	}
	return 0
}

func (m *NestedDefinition_NestedMessage) GetNNM() *NestedDefinition_NestedMessage_NestedNestedMsg {
	if m != nil {
		return m.NNM
	}
	return nil
}

type NestedDefinition_NestedMessage_NestedNestedMsg struct {
	NestedNestedField1 *string `protobuf:"bytes,10,opt,name=NestedNestedField1" json:"NestedNestedField1,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *NestedDefinition_NestedMessage_NestedNestedMsg) Reset() {
	*m = NestedDefinition_NestedMessage_NestedNestedMsg{}
}
func (m *NestedDefinition_NestedMessage_NestedNestedMsg) String() string {
	return proto.CompactTextString(m)
}
func (*NestedDefinition_NestedMessage_NestedNestedMsg) ProtoMessage() {}

func (m *NestedDefinition_NestedMessage_NestedNestedMsg) GetNestedNestedField1() string {
	if m != nil && m.NestedNestedField1 != nil {
		return *m.NestedNestedField1
	}
	return ""
}

type NestedScope struct {
	A                *NestedDefinition_NestedMessage_NestedNestedMsg `protobuf:"bytes,1,opt,name=A" json:"A,omitempty"`
	C                *NestedDefinition_NestedMessage                 `protobuf:"bytes,3,opt,name=C" json:"C,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *NestedScope) Reset()         { *m = NestedScope{} }
func (m *NestedScope) String() string { return proto.CompactTextString(m) }
func (*NestedScope) ProtoMessage()    {}

func (m *NestedScope) GetA() *NestedDefinition_NestedMessage_NestedNestedMsg {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *NestedScope) GetC() *NestedDefinition_NestedMessage {
	if m != nil {
		return m.C
	}
	return nil
}

type NinOptNativeDefault struct {
	Field1           *float64 `protobuf:"fixed64,1,opt,name=Field1,def=1234.1234" json:"Field1,omitempty"`
	Field2           *float32 `protobuf:"fixed32,2,opt,name=Field2,def=1234.1234" json:"Field2,omitempty"`
	Field3           *int32   `protobuf:"varint,3,opt,name=Field3,def=1234" json:"Field3,omitempty"`
	Field4           *int64   `protobuf:"varint,4,opt,name=Field4,def=1234" json:"Field4,omitempty"`
	Field5           *uint32  `protobuf:"varint,5,opt,name=Field5,def=1234" json:"Field5,omitempty"`
	Field6           *uint64  `protobuf:"varint,6,opt,name=Field6,def=1234" json:"Field6,omitempty"`
	Field7           *int32   `protobuf:"zigzag32,7,opt,name=Field7,def=1234" json:"Field7,omitempty"`
	Field8           *int64   `protobuf:"zigzag64,8,opt,name=Field8,def=1234" json:"Field8,omitempty"`
	Field9           *uint32  `protobuf:"fixed32,9,opt,name=Field9,def=1234" json:"Field9,omitempty"`
	Field10          *int32   `protobuf:"fixed32,10,opt,name=Field10,def=1234" json:"Field10,omitempty"`
	Field11          *uint64  `protobuf:"fixed64,11,opt,name=Field11,def=1234" json:"Field11,omitempty"`
	Field12          *int64   `protobuf:"fixed64,12,opt,name=Field12,def=1234" json:"Field12,omitempty"`
	Field13          *bool    `protobuf:"varint,13,opt,name=Field13,def=1" json:"Field13,omitempty"`
	Field14          *string  `protobuf:"bytes,14,opt,name=Field14,def=1234" json:"Field14,omitempty"`
	Field15          []byte   `protobuf:"bytes,15,opt,name=Field15" json:"Field15,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NinOptNativeDefault) Reset()         { *m = NinOptNativeDefault{} }
func (m *NinOptNativeDefault) String() string { return proto.CompactTextString(m) }
func (*NinOptNativeDefault) ProtoMessage()    {}

const Default_NinOptNativeDefault_Field1 float64 = 1234.1234
const Default_NinOptNativeDefault_Field2 float32 = 1234.1234
const Default_NinOptNativeDefault_Field3 int32 = 1234
const Default_NinOptNativeDefault_Field4 int64 = 1234
const Default_NinOptNativeDefault_Field5 uint32 = 1234
const Default_NinOptNativeDefault_Field6 uint64 = 1234
const Default_NinOptNativeDefault_Field7 int32 = 1234
const Default_NinOptNativeDefault_Field8 int64 = 1234
const Default_NinOptNativeDefault_Field9 uint32 = 1234
const Default_NinOptNativeDefault_Field10 int32 = 1234
const Default_NinOptNativeDefault_Field11 uint64 = 1234
const Default_NinOptNativeDefault_Field12 int64 = 1234
const Default_NinOptNativeDefault_Field13 bool = true
const Default_NinOptNativeDefault_Field14 string = "1234"

func (m *NinOptNativeDefault) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return Default_NinOptNativeDefault_Field1
}

func (m *NinOptNativeDefault) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return Default_NinOptNativeDefault_Field2
}

func (m *NinOptNativeDefault) GetField3() int32 {
	if m != nil && m.Field3 != nil {
		return *m.Field3
	}
	return Default_NinOptNativeDefault_Field3
}

func (m *NinOptNativeDefault) GetField4() int64 {
	if m != nil && m.Field4 != nil {
		return *m.Field4
	}
	return Default_NinOptNativeDefault_Field4
}

func (m *NinOptNativeDefault) GetField5() uint32 {
	if m != nil && m.Field5 != nil {
		return *m.Field5
	}
	return Default_NinOptNativeDefault_Field5
}

func (m *NinOptNativeDefault) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return Default_NinOptNativeDefault_Field6
}

func (m *NinOptNativeDefault) GetField7() int32 {
	if m != nil && m.Field7 != nil {
		return *m.Field7
	}
	return Default_NinOptNativeDefault_Field7
}

func (m *NinOptNativeDefault) GetField8() int64 {
	if m != nil && m.Field8 != nil {
		return *m.Field8
	}
	return Default_NinOptNativeDefault_Field8
}

func (m *NinOptNativeDefault) GetField9() uint32 {
	if m != nil && m.Field9 != nil {
		return *m.Field9
	}
	return Default_NinOptNativeDefault_Field9
}

func (m *NinOptNativeDefault) GetField10() int32 {
	if m != nil && m.Field10 != nil {
		return *m.Field10
	}
	return Default_NinOptNativeDefault_Field10
}

func (m *NinOptNativeDefault) GetField11() uint64 {
	if m != nil && m.Field11 != nil {
		return *m.Field11
	}
	return Default_NinOptNativeDefault_Field11
}

func (m *NinOptNativeDefault) GetField12() int64 {
	if m != nil && m.Field12 != nil {
		return *m.Field12
	}
	return Default_NinOptNativeDefault_Field12
}

func (m *NinOptNativeDefault) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return Default_NinOptNativeDefault_Field13
}

func (m *NinOptNativeDefault) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return Default_NinOptNativeDefault_Field14
}

func (m *NinOptNativeDefault) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}
