// Code generated by protoc-gen-go.
// source: fuzz3.proto
// DO NOT EDIT!

package fuzztests

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type TheTestEnum3 int32

const (
	TheTestEnum3_D TheTestEnum3 = 0
	TheTestEnum3_E TheTestEnum3 = 1
	TheTestEnum3_F TheTestEnum3 = 2
)

var TheTestEnum3_name = map[int32]string{
	0: "D",
	1: "E",
	2: "F",
}
var TheTestEnum3_value = map[string]int32{
	"D": 0,
	"E": 1,
	"F": 2,
}

func (x TheTestEnum3) String() string {
	return proto.EnumName(TheTestEnum3_name, int32(x))
}

type NestedDefinition3_NestedEnum3 int32

const (
	NestedDefinition3_TYPE_NESTED3 NestedDefinition3_NestedEnum3 = 0
	NestedDefinition3_TYPE_NESTE3  NestedDefinition3_NestedEnum3 = 1
)

var NestedDefinition3_NestedEnum3_name = map[int32]string{
	0: "TYPE_NESTED3",
	1: "TYPE_NESTE3",
}
var NestedDefinition3_NestedEnum3_value = map[string]int32{
	"TYPE_NESTED3": 0,
	"TYPE_NESTE3":  1,
}

func (x NestedDefinition3_NestedEnum3) String() string {
	return proto.EnumName(NestedDefinition3_NestedEnum3_name, int32(x))
}

type NinOptNative3 struct {
	Field1  float64 `protobuf:"fixed64,1,opt" json:"Field1,omitempty"`
	Field2  float32 `protobuf:"fixed32,2,opt" json:"Field2,omitempty"`
	Field3  int32   `protobuf:"varint,3,opt" json:"Field3,omitempty"`
	Field4  int64   `protobuf:"varint,4,opt" json:"Field4,omitempty"`
	Field5  uint32  `protobuf:"varint,5,opt" json:"Field5,omitempty"`
	Field6  uint64  `protobuf:"varint,6,opt" json:"Field6,omitempty"`
	Field7  int32   `protobuf:"zigzag32,7,opt" json:"Field7,omitempty"`
	Field8  int64   `protobuf:"zigzag64,8,opt" json:"Field8,omitempty"`
	Field9  uint32  `protobuf:"fixed32,9,opt" json:"Field9,omitempty"`
	Field10 int32   `protobuf:"fixed32,10,opt" json:"Field10,omitempty"`
	Field11 uint64  `protobuf:"fixed64,11,opt" json:"Field11,omitempty"`
	Field12 int64   `protobuf:"fixed64,12,opt" json:"Field12,omitempty"`
	Field13 bool    `protobuf:"varint,13,opt" json:"Field13,omitempty"`
	Field14 string  `protobuf:"bytes,14,opt" json:"Field14,omitempty"`
	Field15 []byte  `protobuf:"bytes,15,opt,proto3" json:"Field15,omitempty"`
}

func (m *NinOptNative3) Reset()         { *m = NinOptNative3{} }
func (m *NinOptNative3) String() string { return proto.CompactTextString(m) }
func (*NinOptNative3) ProtoMessage()    {}

type NinRepNative3 struct {
	Field1  []float64 `protobuf:"fixed64,1,rep" json:"Field1,omitempty"`
	Field2  []float32 `protobuf:"fixed32,2,rep" json:"Field2,omitempty"`
	Field3  []int32   `protobuf:"varint,3,rep" json:"Field3,omitempty"`
	Field4  []int64   `protobuf:"varint,4,rep" json:"Field4,omitempty"`
	Field5  []uint32  `protobuf:"varint,5,rep" json:"Field5,omitempty"`
	Field6  []uint64  `protobuf:"varint,6,rep" json:"Field6,omitempty"`
	Field7  []int32   `protobuf:"zigzag32,7,rep" json:"Field7,omitempty"`
	Field8  []int64   `protobuf:"zigzag64,8,rep" json:"Field8,omitempty"`
	Field9  []uint32  `protobuf:"fixed32,9,rep" json:"Field9,omitempty"`
	Field10 []int32   `protobuf:"fixed32,10,rep" json:"Field10,omitempty"`
	Field11 []uint64  `protobuf:"fixed64,11,rep" json:"Field11,omitempty"`
	Field12 []int64   `protobuf:"fixed64,12,rep" json:"Field12,omitempty"`
	Field13 []bool    `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	Field14 []string  `protobuf:"bytes,14,rep" json:"Field14,omitempty"`
	Field15 [][]byte  `protobuf:"bytes,15,rep,proto3" json:"Field15,omitempty"`
}

func (m *NinRepNative3) Reset()         { *m = NinRepNative3{} }
func (m *NinRepNative3) String() string { return proto.CompactTextString(m) }
func (*NinRepNative3) ProtoMessage()    {}

type NinRepPackedNative3 struct {
	Field1  []float64 `protobuf:"fixed64,1,rep,packed" json:"Field1,omitempty"`
	Field2  []float32 `protobuf:"fixed32,2,rep,packed" json:"Field2,omitempty"`
	Field3  []int32   `protobuf:"varint,3,rep,packed" json:"Field3,omitempty"`
	Field4  []int64   `protobuf:"varint,4,rep,packed" json:"Field4,omitempty"`
	Field5  []uint32  `protobuf:"varint,5,rep,packed" json:"Field5,omitempty"`
	Field6  []uint64  `protobuf:"varint,6,rep,packed" json:"Field6,omitempty"`
	Field7  []int32   `protobuf:"zigzag32,7,rep,packed" json:"Field7,omitempty"`
	Field8  []int64   `protobuf:"zigzag64,8,rep,packed" json:"Field8,omitempty"`
	Field9  []uint32  `protobuf:"fixed32,9,rep,packed" json:"Field9,omitempty"`
	Field10 []int32   `protobuf:"fixed32,10,rep,packed" json:"Field10,omitempty"`
	Field11 []uint64  `protobuf:"fixed64,11,rep,packed" json:"Field11,omitempty"`
	Field12 []int64   `protobuf:"fixed64,12,rep,packed" json:"Field12,omitempty"`
	Field13 []bool    `protobuf:"varint,13,rep,packed" json:"Field13,omitempty"`
}

func (m *NinRepPackedNative3) Reset()         { *m = NinRepPackedNative3{} }
func (m *NinRepPackedNative3) String() string { return proto.CompactTextString(m) }
func (*NinRepPackedNative3) ProtoMessage()    {}

type NinOptStruct3 struct {
	Field1  float64        `protobuf:"fixed64,1,opt" json:"Field1,omitempty"`
	Field2  float32        `protobuf:"fixed32,2,opt" json:"Field2,omitempty"`
	Field3  *NinOptNative3 `protobuf:"bytes,3,opt" json:"Field3,omitempty"`
	Field4  *NinOptNative3 `protobuf:"bytes,4,opt" json:"Field4,omitempty"`
	Field6  uint64         `protobuf:"varint,6,opt" json:"Field6,omitempty"`
	Field7  int32          `protobuf:"zigzag32,7,opt" json:"Field7,omitempty"`
	Field8  *NinOptNative3 `protobuf:"bytes,8,opt" json:"Field8,omitempty"`
	Field13 bool           `protobuf:"varint,13,opt" json:"Field13,omitempty"`
	Field14 string         `protobuf:"bytes,14,opt" json:"Field14,omitempty"`
	Field15 []byte         `protobuf:"bytes,15,opt,proto3" json:"Field15,omitempty"`
}

func (m *NinOptStruct3) Reset()         { *m = NinOptStruct3{} }
func (m *NinOptStruct3) String() string { return proto.CompactTextString(m) }
func (*NinOptStruct3) ProtoMessage()    {}

func (m *NinOptStruct3) GetField3() *NinOptNative3 {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinOptStruct3) GetField4() *NinOptNative3 {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinOptStruct3) GetField8() *NinOptNative3 {
	if m != nil {
		return m.Field8
	}
	return nil
}

type NinRepStruct3 struct {
	Field1  []float64        `protobuf:"fixed64,1,rep" json:"Field1,omitempty"`
	Field2  []float32        `protobuf:"fixed32,2,rep" json:"Field2,omitempty"`
	Field3  []*NinOptNative3 `protobuf:"bytes,3,rep" json:"Field3,omitempty"`
	Field4  []*NinOptNative3 `protobuf:"bytes,4,rep" json:"Field4,omitempty"`
	Field6  []uint64         `protobuf:"varint,6,rep" json:"Field6,omitempty"`
	Field7  []int32          `protobuf:"zigzag32,7,rep" json:"Field7,omitempty"`
	Field8  []*NinOptNative3 `protobuf:"bytes,8,rep" json:"Field8,omitempty"`
	Field13 []bool           `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	Field14 []string         `protobuf:"bytes,14,rep" json:"Field14,omitempty"`
	Field15 [][]byte         `protobuf:"bytes,15,rep,proto3" json:"Field15,omitempty"`
}

func (m *NinRepStruct3) Reset()         { *m = NinRepStruct3{} }
func (m *NinRepStruct3) String() string { return proto.CompactTextString(m) }
func (*NinRepStruct3) ProtoMessage()    {}

func (m *NinRepStruct3) GetField3() []*NinOptNative3 {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinRepStruct3) GetField4() []*NinOptNative3 {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinRepStruct3) GetField8() []*NinOptNative3 {
	if m != nil {
		return m.Field8
	}
	return nil
}

type NinNestedStruct3 struct {
	Field1 *NinOptStruct3   `protobuf:"bytes,1,opt" json:"Field1,omitempty"`
	Field2 []*NinRepStruct3 `protobuf:"bytes,2,rep" json:"Field2,omitempty"`
}

func (m *NinNestedStruct3) Reset()         { *m = NinNestedStruct3{} }
func (m *NinNestedStruct3) String() string { return proto.CompactTextString(m) }
func (*NinNestedStruct3) ProtoMessage()    {}

func (m *NinNestedStruct3) GetField1() *NinOptStruct3 {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NinNestedStruct3) GetField2() []*NinRepStruct3 {
	if m != nil {
		return m.Field2
	}
	return nil
}

type Nil3 struct {
}

func (m *Nil3) Reset()         { *m = Nil3{} }
func (m *Nil3) String() string { return proto.CompactTextString(m) }
func (*Nil3) ProtoMessage()    {}

type NinOptEnum3 struct {
	Field1 TheTestEnum3 `protobuf:"varint,1,opt,enum=fuzztests.TheTestEnum3" json:"Field1,omitempty"`
}

func (m *NinOptEnum3) Reset()         { *m = NinOptEnum3{} }
func (m *NinOptEnum3) String() string { return proto.CompactTextString(m) }
func (*NinOptEnum3) ProtoMessage()    {}

type NinRepEnum3 struct {
	Field1 []TheTestEnum3 `protobuf:"varint,1,rep,enum=fuzztests.TheTestEnum3" json:"Field1,omitempty"`
}

func (m *NinRepEnum3) Reset()         { *m = NinRepEnum3{} }
func (m *NinRepEnum3) String() string { return proto.CompactTextString(m) }
func (*NinRepEnum3) ProtoMessage()    {}

type NestedDefinition3 struct {
	Field1    int64                                              `protobuf:"varint,1,opt" json:"Field1,omitempty"`
	EnumField NestedDefinition3_NestedEnum3                      `protobuf:"varint,2,opt,enum=fuzztests.NestedDefinition3_NestedEnum3" json:"EnumField,omitempty"`
	NNM       *NestedDefinition3_NestedMessage3_NestedNestedMsg3 `protobuf:"bytes,3,opt" json:"NNM,omitempty"`
	NM        *NestedDefinition3_NestedMessage3                  `protobuf:"bytes,4,opt" json:"NM,omitempty"`
}

func (m *NestedDefinition3) Reset()         { *m = NestedDefinition3{} }
func (m *NestedDefinition3) String() string { return proto.CompactTextString(m) }
func (*NestedDefinition3) ProtoMessage()    {}

func (m *NestedDefinition3) GetNNM() *NestedDefinition3_NestedMessage3_NestedNestedMsg3 {
	if m != nil {
		return m.NNM
	}
	return nil
}

func (m *NestedDefinition3) GetNM() *NestedDefinition3_NestedMessage3 {
	if m != nil {
		return m.NM
	}
	return nil
}

type NestedDefinition3_NestedMessage3 struct {
	NestedField1 uint64                                             `protobuf:"fixed64,1,opt" json:"NestedField1,omitempty"`
	NNM          *NestedDefinition3_NestedMessage3_NestedNestedMsg3 `protobuf:"bytes,2,opt" json:"NNM,omitempty"`
}

func (m *NestedDefinition3_NestedMessage3) Reset()         { *m = NestedDefinition3_NestedMessage3{} }
func (m *NestedDefinition3_NestedMessage3) String() string { return proto.CompactTextString(m) }
func (*NestedDefinition3_NestedMessage3) ProtoMessage()    {}

func (m *NestedDefinition3_NestedMessage3) GetNNM() *NestedDefinition3_NestedMessage3_NestedNestedMsg3 {
	if m != nil {
		return m.NNM
	}
	return nil
}

type NestedDefinition3_NestedMessage3_NestedNestedMsg3 struct {
	NestedNestedField1 string `protobuf:"bytes,10,opt" json:"NestedNestedField1,omitempty"`
}

func (m *NestedDefinition3_NestedMessage3_NestedNestedMsg3) Reset() {
	*m = NestedDefinition3_NestedMessage3_NestedNestedMsg3{}
}
func (m *NestedDefinition3_NestedMessage3_NestedNestedMsg3) String() string {
	return proto.CompactTextString(m)
}
func (*NestedDefinition3_NestedMessage3_NestedNestedMsg3) ProtoMessage() {}

type NestedScope3 struct {
	A *NestedDefinition3_NestedMessage3_NestedNestedMsg3 `protobuf:"bytes,1,opt" json:"A,omitempty"`
	B NestedDefinition3_NestedEnum3                      `protobuf:"varint,2,opt,enum=fuzztests.NestedDefinition3_NestedEnum3" json:"B,omitempty"`
	C *NestedDefinition3_NestedMessage3                  `protobuf:"bytes,3,opt" json:"C,omitempty"`
}

func (m *NestedScope3) Reset()         { *m = NestedScope3{} }
func (m *NestedScope3) String() string { return proto.CompactTextString(m) }
func (*NestedScope3) ProtoMessage()    {}

func (m *NestedScope3) GetA() *NestedDefinition3_NestedMessage3_NestedNestedMsg3 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *NestedScope3) GetC() *NestedDefinition3_NestedMessage3 {
	if m != nil {
		return m.C
	}
	return nil
}

func init() {
	proto.RegisterEnum("fuzztests.TheTestEnum3", TheTestEnum3_name, TheTestEnum3_value)
	proto.RegisterEnum("fuzztests.NestedDefinition3_NestedEnum3", NestedDefinition3_NestedEnum3_name, NestedDefinition3_NestedEnum3_value)
}