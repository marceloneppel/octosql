// Code generated by protoc-gen-go. DO NOT EDIT.
// source: execution/process.proto

package execution

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type QueueElement struct {
	// Types that are valid to be assigned to Type:
	//	*QueueElement_Record
	//	*QueueElement_Watermark
	//	*QueueElement_EndOfStream
	//	*QueueElement_Error
	Type                 isQueueElement_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *QueueElement) Reset()         { *m = QueueElement{} }
func (m *QueueElement) String() string { return proto.CompactTextString(m) }
func (*QueueElement) ProtoMessage()    {}
func (*QueueElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9889a525e12bbaa, []int{0}
}

func (m *QueueElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueElement.Unmarshal(m, b)
}
func (m *QueueElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueElement.Marshal(b, m, deterministic)
}
func (m *QueueElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueElement.Merge(m, src)
}
func (m *QueueElement) XXX_Size() int {
	return xxx_messageInfo_QueueElement.Size(m)
}
func (m *QueueElement) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueElement.DiscardUnknown(m)
}

var xxx_messageInfo_QueueElement proto.InternalMessageInfo

type isQueueElement_Type interface {
	isQueueElement_Type()
}

type QueueElement_Record struct {
	Record *Record `protobuf:"bytes,1,opt,name=record,proto3,oneof"`
}

type QueueElement_Watermark struct {
	Watermark *timestamp.Timestamp `protobuf:"bytes,2,opt,name=watermark,proto3,oneof"`
}

type QueueElement_EndOfStream struct {
	EndOfStream bool `protobuf:"varint,3,opt,name=endOfStream,proto3,oneof"`
}

type QueueElement_Error struct {
	Error string `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*QueueElement_Record) isQueueElement_Type() {}

func (*QueueElement_Watermark) isQueueElement_Type() {}

func (*QueueElement_EndOfStream) isQueueElement_Type() {}

func (*QueueElement_Error) isQueueElement_Type() {}

func (m *QueueElement) GetType() isQueueElement_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *QueueElement) GetRecord() *Record {
	if x, ok := m.GetType().(*QueueElement_Record); ok {
		return x.Record
	}
	return nil
}

func (m *QueueElement) GetWatermark() *timestamp.Timestamp {
	if x, ok := m.GetType().(*QueueElement_Watermark); ok {
		return x.Watermark
	}
	return nil
}

func (m *QueueElement) GetEndOfStream() bool {
	if x, ok := m.GetType().(*QueueElement_EndOfStream); ok {
		return x.EndOfStream
	}
	return false
}

func (m *QueueElement) GetError() string {
	if x, ok := m.GetType().(*QueueElement_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QueueElement) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QueueElement_Record)(nil),
		(*QueueElement_Watermark)(nil),
		(*QueueElement_EndOfStream)(nil),
		(*QueueElement_Error)(nil),
	}
}

func init() {
	proto.RegisterType((*QueueElement)(nil), "execution.QueueElement")
}

func init() { proto.RegisterFile("execution/process.proto", fileDescriptor_c9889a525e12bbaa) }

var fileDescriptor_c9889a525e12bbaa = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0x41, 0x2b, 0x91, 0xad, 0x17, 0x39, 0x54, 0xc2, 0x45, 0xd2, 0xc4, 0x48, 0x62, 0xb2,
	0x9b, 0xe0, 0xcd, 0x63, 0x13, 0x13, 0x6e, 0x46, 0xf4, 0xe4, 0x0d, 0xb6, 0x53, 0x24, 0xb2, 0x0c,
	0x0e, 0xb3, 0x51, 0x7f, 0x9b, 0x7f, 0xce, 0x84, 0xb5, 0xb4, 0xc7, 0x99, 0xf7, 0xe5, 0x7d, 0x4f,
	0x5c, 0xc1, 0x37, 0x68, 0xcb, 0x2d, 0xf6, 0x6a, 0x20, 0xd4, 0x30, 0x8e, 0x72, 0x20, 0x64, 0x8c,
	0xc2, 0x39, 0x48, 0x56, 0x07, 0x86, 0x40, 0x23, 0x6d, 0x1d, 0x92, 0x5c, 0x37, 0x88, 0x4d, 0x07,
	0x6a, 0xba, 0x6a, 0xbb, 0x53, 0xdc, 0x1a, 0x18, 0xb9, 0x32, 0x83, 0x03, 0xd6, 0xbf, 0xbe, 0xb8,
	0x78, 0xb6, 0x60, 0xe1, 0xb1, 0x03, 0x03, 0x3d, 0x47, 0x77, 0x22, 0x70, 0x0d, 0xb1, 0x9f, 0xfa,
	0xd9, 0x32, 0xbf, 0x94, 0x73, 0xb5, 0x2c, 0xa7, 0xa0, 0xf0, 0xca, 0x7f, 0x24, 0x7a, 0x10, 0xe1,
	0x57, 0xc5, 0x40, 0xa6, 0xa2, 0x8f, 0xf8, 0x64, 0xe2, 0x13, 0xe9, 0x94, 0x72, 0xaf, 0x94, 0xaf,
	0x7b, 0x65, 0xe1, 0x95, 0x07, 0x3c, 0x5a, 0x8b, 0x25, 0xf4, 0xdb, 0xa7, 0xdd, 0x0b, 0x13, 0x54,
	0x26, 0x3e, 0x4d, 0xfd, 0xec, 0xbc, 0xf0, 0xca, 0xe3, 0x67, 0xb4, 0x12, 0x67, 0x40, 0x84, 0x14,
	0x2f, 0x52, 0x3f, 0x0b, 0x0b, 0xaf, 0x74, 0xe7, 0x26, 0x10, 0x0b, 0xfe, 0x19, 0x60, 0x73, 0xfb,
	0x76, 0xd3, 0xb4, 0xfc, 0x6e, 0x6b, 0xa9, 0xd1, 0x28, 0x6d, 0x6b, 0xc8, 0xf3, 0x3c, 0x57, 0xa8,
	0x19, 0xc7, 0xcf, 0x4e, 0xcd, 0xcb, 0xeb, 0x60, 0x5a, 0x73, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xf3, 0x8f, 0x59, 0xaa, 0x4c, 0x01, 0x00, 0x00,
}