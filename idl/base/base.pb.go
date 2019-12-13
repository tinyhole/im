// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/base.proto

package base

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 50 以内为保留产品
type Product int32

const (
	Product_ProductUnknown  Product = 0
	Product_ProductPlatform Product = 1
	Product_ProductDayan    Product = 51
)

var Product_name = map[int32]string{
	0:  "ProductUnknown",
	1:  "ProductPlatform",
	51: "ProductDayan",
}

var Product_value = map[string]int32{
	"ProductUnknown":  0,
	"ProductPlatform": 1,
	"ProductDayan":    51,
}

func (x Product) String() string {
	return proto.EnumName(Product_name, int32(x))
}

func (Product) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{0}
}

// 产品线 8位 子系统 8位 模块8位 8
type CODE int32

const (
	CODE_SUCCEED           CODE = 0
	CODE_FAILED            CODE = 1
	CODE_ENCODE_FAILED     CODE = 2
	CODE_DECODE_FAILED     CODE = 3
	CODE_SERVICE_EXCEPTION CODE = 4
)

var CODE_name = map[int32]string{
	0: "SUCCEED",
	1: "FAILED",
	2: "ENCODE_FAILED",
	3: "DECODE_FAILED",
	4: "SERVICE_EXCEPTION",
}

var CODE_value = map[string]int32{
	"SUCCEED":           0,
	"FAILED":            1,
	"ENCODE_FAILED":     2,
	"DECODE_FAILED":     3,
	"SERVICE_EXCEPTION": 4,
}

func (x CODE) String() string {
	return proto.EnumName(CODE_name, int32(x))
}

func (CODE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{1}
}

type Gender int32

const (
	Gender_GenderUnknown     Gender = 0
	Gender_GenderProductMale Gender = 1
	Gender_GenderFemale      Gender = 2
)

var Gender_name = map[int32]string{
	0: "GenderUnknown",
	1: "GenderProductMale",
	2: "GenderFemale",
}

var Gender_value = map[string]int32{
	"GenderUnknown":     0,
	"GenderProductMale": 1,
	"GenderFemale":      2,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{2}
}

type ActiveType int32

const (
	ActiveType_ActiveTypeEmail   ActiveType = 0
	ActiveType_ActiveTypeSMSCode ActiveType = 1
)

var ActiveType_name = map[int32]string{
	0: "ActiveTypeEmail",
	1: "ActiveTypeSMSCode",
}

var ActiveType_value = map[string]int32{
	"ActiveTypeEmail":   0,
	"ActiveTypeSMSCode": 1,
}

func (x ActiveType) String() string {
	return proto.EnumName(ActiveType_name, int32(x))
}

func (ActiveType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{3}
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{0}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("base.Product", Product_name, Product_value)
	proto.RegisterEnum("base.CODE", CODE_name, CODE_value)
	proto.RegisterEnum("base.Gender", Gender_name, Gender_value)
	proto.RegisterEnum("base.ActiveType", ActiveType_name, ActiveType_value)
	proto.RegisterType((*Void)(nil), "base.Void")
}

func init() { proto.RegisterFile("base/base.proto", fileDescriptor_d66ec2e140567106) }

var fileDescriptor_d66ec2e140567106 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x06, 0xe0, 0x75, 0x5f, 0xc9, 0xe0, 0x7c, 0xea, 0xd2, 0x23, 0xfb, 0x03, 0x5e, 0xf6, 0x42,
	0x2f, 0x76, 0xe3, 0xed, 0x4c, 0xce, 0xa4, 0xe0, 0xb6, 0x62, 0xb7, 0x21, 0x22, 0x8c, 0x6c, 0x8d,
	0x50, 0x6c, 0x93, 0x51, 0xab, 0xb2, 0x7f, 0x2f, 0x59, 0x02, 0xf5, 0x26, 0xbc, 0x79, 0x48, 0xc2,
	0x9b, 0x03, 0xe3, 0xbd, 0xfa, 0xd4, 0x77, 0x6e, 0xb9, 0x3d, 0xb6, 0xb6, 0xb3, 0x18, 0xbb, 0x7c,
	0xc3, 0x20, 0xde, 0xda, 0xaa, 0x4c, 0x25, 0x8c, 0xf2, 0xd6, 0x96, 0x5f, 0x87, 0x0e, 0x11, 0xae,
	0x42, 0xdc, 0x98, 0x0f, 0x63, 0x7f, 0x0c, 0x1f, 0xe0, 0x35, 0x8c, 0x83, 0xe5, 0xb5, 0xea, 0xde,
	0x6d, 0xdb, 0xf0, 0x08, 0x39, 0x5c, 0x04, 0x94, 0xea, 0xa4, 0x0c, 0x9f, 0xa6, 0x6f, 0x10, 0x8b,
	0x95, 0x24, 0xfc, 0x0f, 0xa3, 0x62, 0x23, 0x04, 0x91, 0xe4, 0x03, 0x04, 0x60, 0xf3, 0x59, 0xf6,
	0x44, 0x92, 0x47, 0x98, 0xc0, 0x25, 0x2d, 0xdd, 0x91, 0x5d, 0xa0, 0xa1, 0x23, 0x49, 0x7f, 0xe9,
	0x1f, 0x4e, 0x20, 0x29, 0xe8, 0x79, 0x9b, 0x09, 0xda, 0xd1, 0x8b, 0xa0, 0x7c, 0x9d, 0xad, 0x96,
	0x3c, 0x4e, 0x25, 0xb0, 0x47, 0x6d, 0x4a, 0xdd, 0xba, 0x3b, 0x3e, 0xf5, 0x0d, 0x27, 0x90, 0x78,
	0x0a, 0x95, 0x16, 0xaa, 0xd6, 0xbe, 0xa3, 0xe7, 0xb9, 0x6e, 0x9c, 0x0c, 0xd3, 0x7b, 0x80, 0xd9,
	0xa1, 0xab, 0xbe, 0xf5, 0xfa, 0x74, 0xd4, 0xee, 0x63, 0xfd, 0x8e, 0x1a, 0x55, 0xd5, 0xfe, 0xad,
	0x1e, 0x8b, 0x45, 0x21, 0x6c, 0xa9, 0x79, 0xf4, 0xc0, 0x5e, 0xcf, 0x33, 0xdb, 0xb3, 0xf3, 0x00,
	0xa7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xc7, 0xa1, 0x24, 0x53, 0x01, 0x00, 0x00,
}
