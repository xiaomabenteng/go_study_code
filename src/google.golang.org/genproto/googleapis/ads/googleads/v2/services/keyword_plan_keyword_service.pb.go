// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/keyword_plan_keyword_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [KeywordPlanKeywordService.GetKeywordPlanKeyword][google.ads.googleads.v2.services.KeywordPlanKeywordService.GetKeywordPlanKeyword].
type GetKeywordPlanKeywordRequest struct {
	// The resource name of the ad group keyword to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordPlanKeywordRequest) Reset()         { *m = GetKeywordPlanKeywordRequest{} }
func (m *GetKeywordPlanKeywordRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordPlanKeywordRequest) ProtoMessage()    {}
func (*GetKeywordPlanKeywordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec3e8b0b40b90dff, []int{0}
}

func (m *GetKeywordPlanKeywordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordPlanKeywordRequest.Unmarshal(m, b)
}
func (m *GetKeywordPlanKeywordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordPlanKeywordRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordPlanKeywordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordPlanKeywordRequest.Merge(m, src)
}
func (m *GetKeywordPlanKeywordRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordPlanKeywordRequest.Size(m)
}
func (m *GetKeywordPlanKeywordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordPlanKeywordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordPlanKeywordRequest proto.InternalMessageInfo

func (m *GetKeywordPlanKeywordRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [KeywordPlanKeywordService.MutateKeywordPlanKeywords][google.ads.googleads.v2.services.KeywordPlanKeywordService.MutateKeywordPlanKeywords].
type MutateKeywordPlanKeywordsRequest struct {
	// The ID of the customer whose Keyword Plan keywords are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual Keyword Plan keywords.
	Operations []*KeywordPlanKeywordOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanKeywordsRequest) Reset()         { *m = MutateKeywordPlanKeywordsRequest{} }
func (m *MutateKeywordPlanKeywordsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanKeywordsRequest) ProtoMessage()    {}
func (*MutateKeywordPlanKeywordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec3e8b0b40b90dff, []int{1}
}

func (m *MutateKeywordPlanKeywordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanKeywordsRequest.Unmarshal(m, b)
}
func (m *MutateKeywordPlanKeywordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanKeywordsRequest.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanKeywordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanKeywordsRequest.Merge(m, src)
}
func (m *MutateKeywordPlanKeywordsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanKeywordsRequest.Size(m)
}
func (m *MutateKeywordPlanKeywordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanKeywordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanKeywordsRequest proto.InternalMessageInfo

func (m *MutateKeywordPlanKeywordsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateKeywordPlanKeywordsRequest) GetOperations() []*KeywordPlanKeywordOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateKeywordPlanKeywordsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateKeywordPlanKeywordsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a Keyword Plan keyword.
type KeywordPlanKeywordOperation struct {
	// The FieldMask that determines which resource fields are modified in an
	// update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*KeywordPlanKeywordOperation_Create
	//	*KeywordPlanKeywordOperation_Update
	//	*KeywordPlanKeywordOperation_Remove
	Operation            isKeywordPlanKeywordOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *KeywordPlanKeywordOperation) Reset()         { *m = KeywordPlanKeywordOperation{} }
func (m *KeywordPlanKeywordOperation) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanKeywordOperation) ProtoMessage()    {}
func (*KeywordPlanKeywordOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec3e8b0b40b90dff, []int{2}
}

func (m *KeywordPlanKeywordOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanKeywordOperation.Unmarshal(m, b)
}
func (m *KeywordPlanKeywordOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanKeywordOperation.Marshal(b, m, deterministic)
}
func (m *KeywordPlanKeywordOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanKeywordOperation.Merge(m, src)
}
func (m *KeywordPlanKeywordOperation) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanKeywordOperation.Size(m)
}
func (m *KeywordPlanKeywordOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanKeywordOperation.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanKeywordOperation proto.InternalMessageInfo

func (m *KeywordPlanKeywordOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isKeywordPlanKeywordOperation_Operation interface {
	isKeywordPlanKeywordOperation_Operation()
}

type KeywordPlanKeywordOperation_Create struct {
	Create *resources.KeywordPlanKeyword `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type KeywordPlanKeywordOperation_Update struct {
	Update *resources.KeywordPlanKeyword `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type KeywordPlanKeywordOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*KeywordPlanKeywordOperation_Create) isKeywordPlanKeywordOperation_Operation() {}

func (*KeywordPlanKeywordOperation_Update) isKeywordPlanKeywordOperation_Operation() {}

func (*KeywordPlanKeywordOperation_Remove) isKeywordPlanKeywordOperation_Operation() {}

func (m *KeywordPlanKeywordOperation) GetOperation() isKeywordPlanKeywordOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *KeywordPlanKeywordOperation) GetCreate() *resources.KeywordPlanKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanKeywordOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *KeywordPlanKeywordOperation) GetUpdate() *resources.KeywordPlanKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanKeywordOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *KeywordPlanKeywordOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*KeywordPlanKeywordOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeywordPlanKeywordOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeywordPlanKeywordOperation_Create)(nil),
		(*KeywordPlanKeywordOperation_Update)(nil),
		(*KeywordPlanKeywordOperation_Remove)(nil),
	}
}

// Response message for a Keyword Plan keyword mutate.
type MutateKeywordPlanKeywordsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateKeywordPlanKeywordResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *MutateKeywordPlanKeywordsResponse) Reset()         { *m = MutateKeywordPlanKeywordsResponse{} }
func (m *MutateKeywordPlanKeywordsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanKeywordsResponse) ProtoMessage()    {}
func (*MutateKeywordPlanKeywordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec3e8b0b40b90dff, []int{3}
}

func (m *MutateKeywordPlanKeywordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanKeywordsResponse.Unmarshal(m, b)
}
func (m *MutateKeywordPlanKeywordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanKeywordsResponse.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanKeywordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanKeywordsResponse.Merge(m, src)
}
func (m *MutateKeywordPlanKeywordsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanKeywordsResponse.Size(m)
}
func (m *MutateKeywordPlanKeywordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanKeywordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanKeywordsResponse proto.InternalMessageInfo

func (m *MutateKeywordPlanKeywordsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateKeywordPlanKeywordsResponse) GetResults() []*MutateKeywordPlanKeywordResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the Keyword Plan keyword mutate.
type MutateKeywordPlanKeywordResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanKeywordResult) Reset()         { *m = MutateKeywordPlanKeywordResult{} }
func (m *MutateKeywordPlanKeywordResult) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanKeywordResult) ProtoMessage()    {}
func (*MutateKeywordPlanKeywordResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec3e8b0b40b90dff, []int{4}
}

func (m *MutateKeywordPlanKeywordResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanKeywordResult.Unmarshal(m, b)
}
func (m *MutateKeywordPlanKeywordResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanKeywordResult.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanKeywordResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanKeywordResult.Merge(m, src)
}
func (m *MutateKeywordPlanKeywordResult) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanKeywordResult.Size(m)
}
func (m *MutateKeywordPlanKeywordResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanKeywordResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanKeywordResult proto.InternalMessageInfo

func (m *MutateKeywordPlanKeywordResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordPlanKeywordRequest)(nil), "google.ads.googleads.v2.services.GetKeywordPlanKeywordRequest")
	proto.RegisterType((*MutateKeywordPlanKeywordsRequest)(nil), "google.ads.googleads.v2.services.MutateKeywordPlanKeywordsRequest")
	proto.RegisterType((*KeywordPlanKeywordOperation)(nil), "google.ads.googleads.v2.services.KeywordPlanKeywordOperation")
	proto.RegisterType((*MutateKeywordPlanKeywordsResponse)(nil), "google.ads.googleads.v2.services.MutateKeywordPlanKeywordsResponse")
	proto.RegisterType((*MutateKeywordPlanKeywordResult)(nil), "google.ads.googleads.v2.services.MutateKeywordPlanKeywordResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/keyword_plan_keyword_service.proto", fileDescriptor_ec3e8b0b40b90dff)
}

var fileDescriptor_ec3e8b0b40b90dff = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdf, 0x6a, 0xd4, 0x4a,
	0x18, 0xc0, 0x4f, 0xb2, 0x87, 0x9e, 0xd3, 0x49, 0xcf, 0x11, 0x46, 0x8a, 0xe9, 0xb6, 0xd4, 0x35,
	0x16, 0x2c, 0x7b, 0x91, 0x40, 0xa4, 0x28, 0xa9, 0x2b, 0xee, 0x2e, 0xfd, 0x23, 0x52, 0x5b, 0x52,
	0xe8, 0x45, 0x59, 0x09, 0xd3, 0x64, 0xba, 0x84, 0x26, 0x99, 0x38, 0x33, 0x59, 0x29, 0xa5, 0x37,
	0x82, 0x4f, 0xe0, 0x1b, 0xe8, 0x9d, 0x2f, 0x22, 0x14, 0xbc, 0xf2, 0x15, 0xbc, 0xd1, 0x2b, 0x2f,
	0x7c, 0x00, 0x49, 0x26, 0xb3, 0xfd, 0x9b, 0x5d, 0x69, 0xef, 0xbe, 0xf9, 0xe6, 0xcb, 0xef, 0xfb,
	0x3b, 0x5f, 0x40, 0xb7, 0x4f, 0x48, 0x3f, 0xc2, 0x16, 0x0a, 0x98, 0x25, 0xc4, 0x5c, 0x1a, 0xd8,
	0x16, 0xc3, 0x74, 0x10, 0xfa, 0x98, 0x59, 0x07, 0xf8, 0xf0, 0x0d, 0xa1, 0x81, 0x97, 0x46, 0x28,
	0xf1, 0xe4, 0xa1, 0xbc, 0x35, 0x53, 0x4a, 0x38, 0x81, 0x0d, 0xf1, 0xa5, 0x89, 0x02, 0x66, 0x0e,
	0x21, 0xe6, 0xc0, 0x36, 0x25, 0xa4, 0xfe, 0xa4, 0xca, 0x0d, 0xc5, 0x8c, 0x64, 0xb4, 0xca, 0x8f,
	0xe0, 0xd7, 0xe7, 0xe4, 0xd7, 0x69, 0x68, 0xa1, 0x24, 0x21, 0x1c, 0xf1, 0x90, 0x24, 0xac, 0xbc,
	0x2d, 0xbd, 0x5b, 0xc5, 0x69, 0x2f, 0xdb, 0xb7, 0xf6, 0x43, 0x1c, 0x05, 0x5e, 0x8c, 0xd8, 0x41,
	0x69, 0x71, 0xa7, 0xb4, 0xa0, 0xa9, 0x6f, 0x31, 0x8e, 0x78, 0xc6, 0x2e, 0x5c, 0xe4, 0x60, 0x3f,
	0x0a, 0x71, 0xc2, 0xc5, 0x85, 0xd1, 0x05, 0x73, 0x6b, 0x98, 0xbf, 0x10, 0x51, 0x6c, 0x45, 0x28,
	0x29, 0x45, 0x17, 0xbf, 0xce, 0x30, 0xe3, 0xf0, 0x3e, 0xf8, 0x4f, 0x46, 0xee, 0x25, 0x28, 0xc6,
	0xba, 0xd2, 0x50, 0x16, 0x27, 0xdd, 0x29, 0xa9, 0x7c, 0x89, 0x62, 0x6c, 0xfc, 0x54, 0x40, 0x63,
	0x23, 0xe3, 0x88, 0xe3, 0xcb, 0x20, 0x26, 0x49, 0x77, 0x81, 0xe6, 0x67, 0x8c, 0x93, 0x18, 0x53,
	0x2f, 0x0c, 0x4a, 0x0e, 0x90, 0xaa, 0xe7, 0x01, 0x7c, 0x05, 0x00, 0x49, 0x31, 0x15, 0x29, 0xeb,
	0x6a, 0xa3, 0xb6, 0xa8, 0xd9, 0x2d, 0x73, 0x5c, 0xc5, 0xcd, 0xcb, 0x2e, 0x37, 0x25, 0xc5, 0x3d,
	0x03, 0x84, 0x0f, 0xc0, 0xad, 0x14, 0x51, 0x1e, 0xa2, 0xc8, 0xdb, 0x47, 0x61, 0x94, 0x51, 0xac,
	0xd7, 0x1a, 0xca, 0xe2, 0xbf, 0xee, 0xff, 0xa5, 0x7a, 0x55, 0x68, 0xf3, 0x94, 0x07, 0x28, 0x0a,
	0x03, 0xc4, 0xb1, 0x47, 0x92, 0xe8, 0x50, 0xff, 0xbb, 0x30, 0x9b, 0x92, 0xca, 0xcd, 0x24, 0x3a,
	0x34, 0x3e, 0xaa, 0x60, 0x76, 0x84, 0x67, 0xb8, 0x0c, 0xb4, 0x2c, 0x2d, 0x10, 0x79, 0x7b, 0x0a,
	0x84, 0x66, 0xd7, 0x65, 0x36, 0xb2, 0x83, 0xe6, 0x6a, 0xde, 0xc1, 0x0d, 0xc4, 0x0e, 0x5c, 0x20,
	0xcc, 0x73, 0x19, 0x6e, 0x82, 0x09, 0x9f, 0x62, 0xc4, 0x45, 0xb5, 0x35, 0x7b, 0xa9, 0xb2, 0x0a,
	0xc3, 0xa9, 0xba, 0xa2, 0x0c, 0xeb, 0x7f, 0xb9, 0x25, 0x26, 0x07, 0x0a, 0xbc, 0xae, 0xde, 0x10,
	0x28, 0x30, 0x50, 0x07, 0x13, 0x14, 0xc7, 0x64, 0x20, 0x6a, 0x38, 0x99, 0xdf, 0x88, 0x73, 0x47,
	0x03, 0x93, 0xc3, 0xa2, 0x1b, 0x9f, 0x15, 0x70, 0x6f, 0xc4, 0x60, 0xb0, 0x94, 0x24, 0x0c, 0xc3,
	0x55, 0x30, 0x7d, 0xa1, 0x33, 0x1e, 0xa6, 0x94, 0xd0, 0x82, 0xad, 0xd9, 0x50, 0x06, 0x4b, 0x53,
	0xdf, 0xdc, 0x2e, 0xa6, 0xda, 0xbd, 0x7d, 0xbe, 0x67, 0x2b, 0xb9, 0x39, 0xdc, 0x05, 0xff, 0x50,
	0xcc, 0xb2, 0x88, 0xcb, 0xe9, 0x79, 0x36, 0x7e, 0x7a, 0xaa, 0xa2, 0x73, 0x0b, 0x90, 0x2b, 0x81,
	0xc6, 0x0a, 0x98, 0x1f, 0x6d, 0xfa, 0x47, 0x2f, 0xc5, 0xfe, 0x55, 0x03, 0x33, 0x97, 0x09, 0xdb,
	0x22, 0x1a, 0xf8, 0x45, 0x01, 0xd3, 0x57, 0xbe, 0x46, 0xf8, 0x74, 0x7c, 0x26, 0xa3, 0x9e, 0x71,
	0xfd, 0x7a, 0x0d, 0x37, 0x5a, 0x6f, 0xbf, 0x7e, 0x7b, 0xaf, 0x3e, 0x82, 0x4b, 0xf9, 0x06, 0x3b,
	0x3a, 0x97, 0x5e, 0x4b, 0xbe, 0x5c, 0x66, 0x35, 0xe5, 0x4a, 0x3b, 0xdb, 0x5d, 0xab, 0x79, 0x0c,
	0xbf, 0x2b, 0x60, 0xa6, 0xb2, 0xfd, 0xb0, 0x73, 0xfd, 0xee, 0xc8, 0xa5, 0x52, 0xef, 0xde, 0x88,
	0x21, 0xe6, 0xcf, 0xe8, 0x16, 0x59, 0xb6, 0x8c, 0xc7, 0x79, 0x96, 0xa7, 0x69, 0x1d, 0x9d, 0x59,
	0x57, 0xad, 0xe6, 0xf1, 0x55, 0x49, 0x3a, 0x71, 0x01, 0x77, 0x94, 0x66, 0x7d, 0xf6, 0xa4, 0xad,
	0x9f, 0x06, 0x50, 0x4a, 0x69, 0xc8, 0x4c, 0x9f, 0xc4, 0x9d, 0x77, 0x2a, 0x58, 0xf0, 0x49, 0x3c,
	0x36, 0xd8, 0xce, 0x7c, 0xe5, 0x70, 0x6c, 0xe5, 0x2b, 0x63, 0x4b, 0xd9, 0x5d, 0x2f, 0x19, 0x7d,
	0x12, 0xa1, 0xa4, 0x6f, 0x12, 0xda, 0xb7, 0xfa, 0x38, 0x29, 0x16, 0x8a, 0x75, 0xea, 0xb5, 0xfa,
	0x37, 0xb7, 0x2c, 0x85, 0x0f, 0x6a, 0x6d, 0xad, 0xdd, 0xfe, 0xa4, 0x36, 0xd6, 0x04, 0xb0, 0x1d,
	0x30, 0x53, 0x88, 0xb9, 0xb4, 0x63, 0x9b, 0xa5, 0x63, 0x76, 0x22, 0x4d, 0x7a, 0xed, 0x80, 0xf5,
	0x86, 0x26, 0xbd, 0x1d, 0xbb, 0x27, 0x4d, 0x7e, 0xa8, 0x0b, 0x42, 0xef, 0x38, 0xed, 0x80, 0x39,
	0xce, 0xd0, 0xc8, 0x71, 0x76, 0x6c, 0xc7, 0x91, 0x66, 0x7b, 0x13, 0x45, 0x9c, 0x0f, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x75, 0x3d, 0xaf, 0xb0, 0x8d, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordPlanKeywordServiceClient is the client API for KeywordPlanKeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanKeywordServiceClient interface {
	// Returns the requested Keyword Plan keyword in full detail.
	GetKeywordPlanKeyword(ctx context.Context, in *GetKeywordPlanKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanKeyword, error)
	// Creates, updates, or removes Keyword Plan keywords. Operation statuses are
	// returned.
	MutateKeywordPlanKeywords(ctx context.Context, in *MutateKeywordPlanKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanKeywordsResponse, error)
}

type keywordPlanKeywordServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordPlanKeywordServiceClient(cc *grpc.ClientConn) KeywordPlanKeywordServiceClient {
	return &keywordPlanKeywordServiceClient{cc}
}

func (c *keywordPlanKeywordServiceClient) GetKeywordPlanKeyword(ctx context.Context, in *GetKeywordPlanKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanKeyword, error) {
	out := new(resources.KeywordPlanKeyword)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordPlanKeywordService/GetKeywordPlanKeyword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanKeywordServiceClient) MutateKeywordPlanKeywords(ctx context.Context, in *MutateKeywordPlanKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanKeywordsResponse, error) {
	out := new(MutateKeywordPlanKeywordsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordPlanKeywordService/MutateKeywordPlanKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanKeywordServiceServer is the server API for KeywordPlanKeywordService service.
type KeywordPlanKeywordServiceServer interface {
	// Returns the requested Keyword Plan keyword in full detail.
	GetKeywordPlanKeyword(context.Context, *GetKeywordPlanKeywordRequest) (*resources.KeywordPlanKeyword, error)
	// Creates, updates, or removes Keyword Plan keywords. Operation statuses are
	// returned.
	MutateKeywordPlanKeywords(context.Context, *MutateKeywordPlanKeywordsRequest) (*MutateKeywordPlanKeywordsResponse, error)
}

// UnimplementedKeywordPlanKeywordServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanKeywordServiceServer struct {
}

func (*UnimplementedKeywordPlanKeywordServiceServer) GetKeywordPlanKeyword(ctx context.Context, req *GetKeywordPlanKeywordRequest) (*resources.KeywordPlanKeyword, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetKeywordPlanKeyword not implemented")
}
func (*UnimplementedKeywordPlanKeywordServiceServer) MutateKeywordPlanKeywords(ctx context.Context, req *MutateKeywordPlanKeywordsRequest) (*MutateKeywordPlanKeywordsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateKeywordPlanKeywords not implemented")
}

func RegisterKeywordPlanKeywordServiceServer(s *grpc.Server, srv KeywordPlanKeywordServiceServer) {
	s.RegisterService(&_KeywordPlanKeywordService_serviceDesc, srv)
}

func _KeywordPlanKeywordService_GetKeywordPlanKeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanKeywordServiceServer).GetKeywordPlanKeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordPlanKeywordService/GetKeywordPlanKeyword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanKeywordServiceServer).GetKeywordPlanKeyword(ctx, req.(*GetKeywordPlanKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanKeywordService_MutateKeywordPlanKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanKeywordServiceServer).MutateKeywordPlanKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordPlanKeywordService/MutateKeywordPlanKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanKeywordServiceServer).MutateKeywordPlanKeywords(ctx, req.(*MutateKeywordPlanKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanKeywordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.KeywordPlanKeywordService",
	HandlerType: (*KeywordPlanKeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlanKeyword",
			Handler:    _KeywordPlanKeywordService_GetKeywordPlanKeyword_Handler,
		},
		{
			MethodName: "MutateKeywordPlanKeywords",
			Handler:    _KeywordPlanKeywordService_MutateKeywordPlanKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/keyword_plan_keyword_service.proto",
}
