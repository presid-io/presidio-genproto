// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// StreamRequest represents the request to the stream service via GRPC
type StreamRequest struct {
	// The scanner input kind
	Kind string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	// The stream configuration for scanning streams such as Azure EventHub, Kafka and Kinesis
	StreamConfig *StreamConfig `protobuf:"bytes,2,opt,name=streamConfig" json:"streamConfig,omitempty"`
	// The minProbability will filter results which has lower certainty than the provided value.
	MinProbability string `protobuf:"bytes,3,opt,name=minProbability" json:"minProbability,omitempty"`
	// The analyzer template configures the fields to analyze
	AnalyzeTemplate *AnalyzeTemplate `protobuf:"bytes,4,opt,name=analyzeTemplate" json:"analyzeTemplate,omitempty"`
	// The anonymizer template configures how to anonymize the sensitive data [optional]
	AnonymizeTemplate *AnonymizeTemplate `protobuf:"bytes,5,opt,name=anonymizeTemplate" json:"anonymizeTemplate,omitempty"`
	// The datasinkTemplate configures the output destination of the analyzer/anonymizer results
	DatasinkTemplate     *DatasinkTemplate `protobuf:"bytes,6,opt,name=datasinkTemplate" json:"datasinkTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_618904af0445e701, []int{0}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *StreamRequest) GetStreamConfig() *StreamConfig {
	if m != nil {
		return m.StreamConfig
	}
	return nil
}

func (m *StreamRequest) GetMinProbability() string {
	if m != nil {
		return m.MinProbability
	}
	return ""
}

func (m *StreamRequest) GetAnalyzeTemplate() *AnalyzeTemplate {
	if m != nil {
		return m.AnalyzeTemplate
	}
	return nil
}

func (m *StreamRequest) GetAnonymizeTemplate() *AnonymizeTemplate {
	if m != nil {
		return m.AnonymizeTemplate
	}
	return nil
}

func (m *StreamRequest) GetDatasinkTemplate() *DatasinkTemplate {
	if m != nil {
		return m.DatasinkTemplate
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "types.StreamRequest")
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_stream_618904af0445e701) }

var fileDescriptor_stream_618904af0445e701 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x96,
	0xe2, 0x2b, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0x85, 0x08, 0x2b, 0x5d, 0x66, 0xe2, 0xe2,
	0x0d, 0x06, 0xab, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0xce,
	0xcc, 0x4b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xcc, 0x61, 0x86, 0x39,
	0xe7, 0xe7, 0xa5, 0x65, 0xa6, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xeb, 0x81, 0xcd,
	0xd4, 0x0b, 0x46, 0x92, 0x0a, 0x42, 0x51, 0x28, 0xa4, 0xc6, 0xc5, 0x97, 0x9b, 0x99, 0x17, 0x50,
	0x94, 0x9f, 0x94, 0x98, 0x94, 0x99, 0x93, 0x59, 0x52, 0x29, 0xc1, 0x0c, 0x36, 0x16, 0x4d, 0x54,
	0xc8, 0x81, 0x8b, 0x3f, 0x31, 0x2f, 0x31, 0xa7, 0xb2, 0x2a, 0x35, 0x04, 0xea, 0x3e, 0x09, 0x16,
	0xb0, 0x1d, 0x62, 0x50, 0x3b, 0x1c, 0x51, 0x65, 0x83, 0xd0, 0x95, 0x0b, 0xb9, 0x71, 0x09, 0x26,
	0xe6, 0xe5, 0xe7, 0x55, 0xe6, 0x66, 0x22, 0x99, 0xc1, 0x0a, 0x36, 0x43, 0x02, 0x6e, 0x06, 0x9a,
	0x7c, 0x10, 0xa6, 0x16, 0x21, 0x67, 0x2e, 0x81, 0x94, 0xc4, 0x92, 0xc4, 0xe2, 0xcc, 0xbc, 0x6c,
	0xb8, 0x31, 0x6c, 0x60, 0x63, 0xc4, 0xa1, 0xc6, 0xb8, 0xa0, 0x49, 0x07, 0x61, 0x68, 0x48, 0x62,
	0x03, 0x07, 0xae, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xcc, 0x39, 0x1f, 0x83, 0x01, 0x00,
	0x00,
}