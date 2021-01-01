// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_message.proto

package userproto

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

type RequestMetaSortsDirection int32

const (
	RequestMetaSorts_Descending RequestMetaSortsDirection = 0
	RequestMetaSorts_Ascending  RequestMetaSortsDirection = 1
)

var RequestMetaSortsDirection_name = map[int32]string{
	0: "Descending",
	1: "Ascending",
}

var RequestMetaSortsDirection_value = map[string]int32{
	"Descending": 0,
	"Ascending":  1,
}

func (x RequestMetaSortsDirection) String() string {
	return proto.EnumName(RequestMetaSortsDirection_name, int32(x))
}

func (RequestMetaSortsDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{3, 0}
}

//============ The request message
type RequestMessage struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name"`
	Type                 string   `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type"`
	Time                 string   `protobuf:"bytes,3,opt,name=time,proto3" json:"time"`
	Header               *ReqMeta `protobuf:"bytes,4,opt,name=Header,proto3" json:"Header"`
	Body                 *Any     `protobuf:"bytes,5,opt,name=Body,proto3" json:"Body"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMessage) Reset()         { *m = RequestMessage{} }
func (m *RequestMessage) String() string { return proto.CompactTextString(m) }
func (*RequestMessage) ProtoMessage()    {}
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{0}
}

func (m *RequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMessage.Unmarshal(m, b)
}
func (m *RequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMessage.Marshal(b, m, deterministic)
}
func (m *RequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMessage.Merge(m, src)
}
func (m *RequestMessage) XXX_Size() int {
	return xxx_messageInfo_RequestMessage.Size(m)
}
func (m *RequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMessage proto.InternalMessageInfo

func (m *RequestMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RequestMessage) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *RequestMessage) GetHeader() *ReqMeta {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RequestMessage) GetBody() *Any {
	if m != nil {
		return m.Body
	}
	return nil
}

//============ The response message
type ResponseMessage struct {
	Entity               string            `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Meta                 *ResponseMetadata `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta"`
	Data                 *Any              `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResponseMessage) Reset()         { *m = ResponseMessage{} }
func (m *ResponseMessage) String() string { return proto.CompactTextString(m) }
func (*ResponseMessage) ProtoMessage()    {}
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{1}
}

func (m *ResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMessage.Unmarshal(m, b)
}
func (m *ResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMessage.Marshal(b, m, deterministic)
}
func (m *ResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMessage.Merge(m, src)
}
func (m *ResponseMessage) XXX_Size() int {
	return xxx_messageInfo_ResponseMessage.Size(m)
}
func (m *ResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMessage proto.InternalMessageInfo

func (m *ResponseMessage) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ResponseMessage) GetMeta() *ResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ResponseMessage) GetData() *Any {
	if m != nil {
		return m.Data
	}
	return nil
}

//==================== Meta
type ReqMeta struct {
	UTP                  string             `protobuf:"bytes,1,opt,name=UTP,proto3" json:"UTP"`
	UID                  string             `protobuf:"bytes,2,opt,name=UID,proto3" json:"UID"`
	FID                  string             `protobuf:"bytes,3,opt,name=FID,proto3" json:"FID"`
	Page                 uint32             `protobuf:"varint,4,opt,name=page,proto3" json:"page"`
	PerPage              uint32             `protobuf:"varint,5,opt,name=perPage,proto3" json:"perPage"`
	IpAddress            string             `protobuf:"bytes,6,opt,name=ipAddress,proto3" json:"ipAddress"`
	StartAt              string             `protobuf:"bytes,7,opt,name=startAt,proto3" json:"startAt"`
	EndAt                string             `protobuf:"bytes,8,opt,name=endAt,proto3" json:"endAt"`
	Sorts                *RequestMetaSorts  `protobuf:"bytes,9,opt,name=sorts,proto3" json:"sorts"`
	Filters              *RequestMetaFilter `protobuf:"bytes,10,opt,name=filters,proto3" json:"filters"`
	Action               *RequestMetaAction `protobuf:"bytes,11,opt,name=action,proto3" json:"action"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReqMeta) Reset()         { *m = ReqMeta{} }
func (m *ReqMeta) String() string { return proto.CompactTextString(m) }
func (*ReqMeta) ProtoMessage()    {}
func (*ReqMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{2}
}

func (m *ReqMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMeta.Unmarshal(m, b)
}
func (m *ReqMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMeta.Marshal(b, m, deterministic)
}
func (m *ReqMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMeta.Merge(m, src)
}
func (m *ReqMeta) XXX_Size() int {
	return xxx_messageInfo_ReqMeta.Size(m)
}
func (m *ReqMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMeta proto.InternalMessageInfo

func (m *ReqMeta) GetUTP() string {
	if m != nil {
		return m.UTP
	}
	return ""
}

func (m *ReqMeta) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *ReqMeta) GetFID() string {
	if m != nil {
		return m.FID
	}
	return ""
}

func (m *ReqMeta) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReqMeta) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *ReqMeta) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ReqMeta) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *ReqMeta) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

func (m *ReqMeta) GetSorts() *RequestMetaSorts {
	if m != nil {
		return m.Sorts
	}
	return nil
}

func (m *ReqMeta) GetFilters() *RequestMetaFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *ReqMeta) GetAction() *RequestMetaAction {
	if m != nil {
		return m.Action
	}
	return nil
}

type RequestMetaSorts struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Dir                  uint32   `protobuf:"varint,2,opt,name=dir,proto3" json:"dir"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMetaSorts) Reset()         { *m = RequestMetaSorts{} }
func (m *RequestMetaSorts) String() string { return proto.CompactTextString(m) }
func (*RequestMetaSorts) ProtoMessage()    {}
func (*RequestMetaSorts) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{3}
}

func (m *RequestMetaSorts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMetaSorts.Unmarshal(m, b)
}
func (m *RequestMetaSorts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMetaSorts.Marshal(b, m, deterministic)
}
func (m *RequestMetaSorts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMetaSorts.Merge(m, src)
}
func (m *RequestMetaSorts) XXX_Size() int {
	return xxx_messageInfo_RequestMetaSorts.Size(m)
}
func (m *RequestMetaSorts) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMetaSorts.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMetaSorts proto.InternalMessageInfo

func (m *RequestMetaSorts) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestMetaSorts) GetDir() uint32 {
	if m != nil {
		return m.Dir
	}
	return 0
}

type RequestMetaAction struct {
	ActionType           string   `protobuf:"bytes,1,opt,name=actionType,proto3" json:"actionType"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMetaAction) Reset()         { *m = RequestMetaAction{} }
func (m *RequestMetaAction) String() string { return proto.CompactTextString(m) }
func (*RequestMetaAction) ProtoMessage()    {}
func (*RequestMetaAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{4}
}

func (m *RequestMetaAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMetaAction.Unmarshal(m, b)
}
func (m *RequestMetaAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMetaAction.Marshal(b, m, deterministic)
}
func (m *RequestMetaAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMetaAction.Merge(m, src)
}
func (m *RequestMetaAction) XXX_Size() int {
	return xxx_messageInfo_RequestMetaAction.Size(m)
}
func (m *RequestMetaAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMetaAction.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMetaAction proto.InternalMessageInfo

func (m *RequestMetaAction) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

type RequestMetaFilter struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	Opt                  string   `protobuf:"bytes,2,opt,name=opt,proto3" json:"opt"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMetaFilter) Reset()         { *m = RequestMetaFilter{} }
func (m *RequestMetaFilter) String() string { return proto.CompactTextString(m) }
func (*RequestMetaFilter) ProtoMessage()    {}
func (*RequestMetaFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{5}
}

func (m *RequestMetaFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMetaFilter.Unmarshal(m, b)
}
func (m *RequestMetaFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMetaFilter.Marshal(b, m, deterministic)
}
func (m *RequestMetaFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMetaFilter.Merge(m, src)
}
func (m *RequestMetaFilter) XXX_Size() int {
	return xxx_messageInfo_RequestMetaFilter.Size(m)
}
func (m *RequestMetaFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMetaFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMetaFilter proto.InternalMessageInfo

func (m *RequestMetaFilter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RequestMetaFilter) GetOpt() string {
	if m != nil {
		return m.Opt
	}
	return ""
}

func (m *RequestMetaFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ResponseMetadata struct {
	Total                uint32   `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Page                 uint32   `protobuf:"varint,2,opt,name=page,proto3" json:"page"`
	PerPage              uint32   `protobuf:"varint,3,opt,name=perPage,proto3" json:"perPage"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMetadata) Reset()         { *m = ResponseMetadata{} }
func (m *ResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*ResponseMetadata) ProtoMessage()    {}
func (*ResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{6}
}

func (m *ResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMetadata.Unmarshal(m, b)
}
func (m *ResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMetadata.Marshal(b, m, deterministic)
}
func (m *ResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMetadata.Merge(m, src)
}
func (m *ResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_ResponseMetadata.Size(m)
}
func (m *ResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMetadata proto.InternalMessageInfo

func (m *ResponseMetadata) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ResponseMetadata) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ResponseMetadata) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

type Any struct {
	TypeUrl              string   `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Any) Reset()         { *m = Any{} }
func (m *Any) String() string { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()    {}
func (*Any) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{7}
}

func (m *Any) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Any.Unmarshal(m, b)
}
func (m *Any) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Any.Marshal(b, m, deterministic)
}
func (m *Any) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Any.Merge(m, src)
}
func (m *Any) XXX_Size() int {
	return xxx_messageInfo_Any.Size(m)
}
func (m *Any) XXX_DiscardUnknown() {
	xxx_messageInfo_Any.DiscardUnknown(m)
}

var xxx_messageInfo_Any proto.InternalMessageInfo

func (m *Any) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *Any) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("userproto.RequestMetaSortsDirection", RequestMetaSortsDirection_name, RequestMetaSortsDirection_value)
	proto.RegisterType((*RequestMessage)(nil), "userproto.RequestMessage")
	proto.RegisterType((*ResponseMessage)(nil), "userproto.ResponseMessage")
	proto.RegisterType((*ReqMeta)(nil), "userproto.ReqMeta")
	proto.RegisterType((*RequestMetaSorts)(nil), "userproto.RequestMetaSorts")
	proto.RegisterType((*RequestMetaAction)(nil), "userproto.RequestMetaAction")
	proto.RegisterType((*RequestMetaFilter)(nil), "userproto.RequestMetaFilter")
	proto.RegisterType((*ResponseMetadata)(nil), "userproto.ResponseMetadata")
	proto.RegisterType((*Any)(nil), "userproto.Any")
}

func init() {
	proto.RegisterFile("user_message.proto", fileDescriptor_7965fb6944d08275)
}

var fileDescriptor_7965fb6944d08275 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x25, 0xcd, 0x57, 0x3d, 0x21, 0x21, 0xac, 0x10, 0x5a, 0x44, 0x85, 0x2a, 0x9f, 0xaa, 0x1e,
	0x82, 0x68, 0x51, 0xef, 0x46, 0x51, 0x45, 0x0f, 0x85, 0x6a, 0x69, 0xb9, 0x56, 0x4b, 0x3d, 0x44,
	0x96, 0x12, 0xdb, 0xdd, 0x9d, 0x20, 0xe5, 0xc2, 0x2f, 0xe1, 0xa7, 0xf1, 0x63, 0xd0, 0xcc, 0xae,
	0xf3, 0x01, 0x11, 0xb7, 0x37, 0x6f, 0xdf, 0xce, 0x3c, 0xcf, 0x3e, 0x83, 0x5a, 0x7a, 0x74, 0xf7,
	0x0b, 0xf4, 0xde, 0xce, 0x70, 0x52, 0xbb, 0x8a, 0x2a, 0x95, 0x30, 0x27, 0x30, 0xfd, 0xd5, 0x82,
	0x91, 0xc1, 0xc7, 0x25, 0x7a, 0xba, 0x0e, 0x1a, 0xa5, 0xa0, 0xf3, 0xc9, 0x2e, 0x50, 0xb7, 0x8e,
	0x5b, 0x27, 0x89, 0x11, 0xcc, 0xdc, 0xed, 0xaa, 0x46, 0x7d, 0x10, 0x38, 0xc6, 0xcc, 0x51, 0xb1,
	0x40, 0xdd, 0x0e, 0x1c, 0x63, 0x75, 0x0a, 0xbd, 0x8f, 0x68, 0x73, 0x74, 0xba, 0x73, 0xdc, 0x3a,
	0x19, 0x9c, 0xa9, 0xc9, 0x7a, 0xd4, 0xc4, 0xe0, 0xe3, 0x35, 0x92, 0x35, 0x51, 0xa1, 0x52, 0xe8,
	0x7c, 0xa8, 0xf2, 0x95, 0xee, 0x8a, 0x72, 0xb4, 0xa5, 0xcc, 0xca, 0x95, 0x91, 0xb3, 0xf4, 0x27,
	0x3c, 0x33, 0xe8, 0xeb, 0xaa, 0xf4, 0xd8, 0xd8, 0x7b, 0x09, 0x3d, 0x2c, 0xa9, 0xa0, 0x55, 0x34,
	0x18, 0x2b, 0xf5, 0x16, 0x3a, 0x0b, 0x24, 0x2b, 0x16, 0x07, 0x67, 0xaf, 0x77, 0x06, 0x37, 0x1d,
	0xc8, 0xe6, 0x96, 0xac, 0x11, 0x21, 0xcf, 0x9f, 0x5a, 0xb2, 0xe2, 0x7f, 0xcf, 0x7c, 0x3e, 0x4b,
	0x7f, 0x1f, 0x40, 0x3f, 0xfa, 0x56, 0x63, 0x68, 0xdf, 0xdd, 0xde, 0xc4, 0xa9, 0x0c, 0x85, 0xb9,
	0x9a, 0xc6, 0xa5, 0x30, 0x64, 0xe6, 0xf2, 0x6a, 0x1a, 0x57, 0xc2, 0x90, 0xb7, 0x54, 0xdb, 0x19,
	0xca, 0x3e, 0x86, 0x46, 0xb0, 0xd2, 0xd0, 0xaf, 0xd1, 0xdd, 0x30, 0xdd, 0x15, 0xba, 0x29, 0xd5,
	0x11, 0x24, 0x45, 0x9d, 0xe5, 0xb9, 0x43, 0xef, 0x75, 0x4f, 0xba, 0x6c, 0x08, 0xbe, 0xe7, 0xc9,
	0x3a, 0xca, 0x48, 0xf7, 0xe5, 0xac, 0x29, 0xd5, 0x0b, 0xe8, 0x62, 0x99, 0x67, 0xa4, 0x0f, 0x85,
	0x0f, 0x85, 0x7a, 0x07, 0x5d, 0x5f, 0x39, 0xf2, 0x3a, 0xd9, 0xb3, 0x93, 0xf8, 0xe6, 0x64, 0xbf,
	0xb0, 0xc4, 0x04, 0xa5, 0xba, 0x80, 0xfe, 0xf7, 0x62, 0x4e, 0xe8, 0xbc, 0x06, 0xb9, 0x74, 0xb4,
	0xff, 0xd2, 0xa5, 0x88, 0x4c, 0x23, 0x56, 0xef, 0xa1, 0x67, 0x1f, 0xa8, 0xa8, 0x4a, 0x3d, 0xf8,
	0xdf, 0xb5, 0x4c, 0x34, 0x26, 0x6a, 0xd3, 0x1c, 0xc6, 0x7f, 0x1b, 0xe1, 0x85, 0x95, 0x5b, 0xf1,
	0x63, 0xcc, 0x6b, 0xcd, 0x0b, 0x27, 0x8b, 0x1e, 0x1a, 0x86, 0xe9, 0x29, 0x24, 0x79, 0xe1, 0x50,
	0xda, 0xa8, 0x11, 0xc0, 0x14, 0xfd, 0x03, 0x96, 0x79, 0x51, 0xce, 0xc6, 0x4f, 0xd4, 0x10, 0x92,
	0x6c, 0x5d, 0xb6, 0xd2, 0x73, 0x78, 0xfe, 0x8f, 0x05, 0xf5, 0x06, 0x20, 0x98, 0x90, 0x5c, 0x87,
	0x61, 0x5b, 0x4c, 0xfa, 0x79, 0xe7, 0x52, 0xf8, 0x5c, 0x89, 0xfc, 0x46, 0x2e, 0x98, 0xbd, 0x55,
	0x35, 0x35, 0x21, 0xa8, 0x6a, 0x79, 0x8c, 0x1f, 0x76, 0xbe, 0x6c, 0xfe, 0x8c, 0x50, 0xa4, 0x5f,
	0xf9, 0x5b, 0x77, 0x83, 0xc8, 0x4a, 0xaa, 0xc8, 0xce, 0xa5, 0xe1, 0xd0, 0x84, 0x62, 0x1d, 0x99,
	0x83, 0xfd, 0x91, 0x69, 0xef, 0x44, 0x26, 0xbd, 0x80, 0x76, 0x56, 0xae, 0xd4, 0x2b, 0x38, 0x64,
	0x3b, 0xf7, 0x4b, 0x37, 0x8f, 0xf6, 0xfa, 0x5c, 0xdf, 0xb9, 0xf9, 0xc6, 0x0f, 0x37, 0x7c, 0x1a,
	0xfd, 0x7c, 0xeb, 0xc9, 0xe3, 0x9c, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xdd, 0x97, 0x48,
	0x21, 0x04, 0x00, 0x00,
}