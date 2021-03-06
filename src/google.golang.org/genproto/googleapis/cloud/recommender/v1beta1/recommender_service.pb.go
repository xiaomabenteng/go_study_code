// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recommender/v1beta1/recommender_service.proto

package recommender

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request for the `ListRecommendations` method.
type ListRecommendationsRequest struct {
	// Required. The container resource on which to execute the request.
	// Acceptable formats:
	//
	// 1.
	// "projects/[PROJECT_NUMBER]/locations/[LOCATION]/recommenders/[RECOMMENDER_ID]",
	//
	// LOCATION here refers to GCP Locations:
	// https://cloud.google.com/about/locations/
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored. If not specified, the server will
	// determine the number of results to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. If present, retrieves the next batch of results from the
	// preceding call to this method. `page_token` must be the value of
	// `next_page_token` from the previous response. The values of other method
	// parameters must be identical to those in the previous call.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Filter expression to restrict the recommendations returned. Supported
	// filter fields: state_info.state
	// Eg: `state_info.state:"DISMISSED" or state_info.state:"FAILED"
	Filter               string   `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRecommendationsRequest) Reset()         { *m = ListRecommendationsRequest{} }
func (m *ListRecommendationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListRecommendationsRequest) ProtoMessage()    {}
func (*ListRecommendationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca265e4ecfbc6d8, []int{0}
}

func (m *ListRecommendationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRecommendationsRequest.Unmarshal(m, b)
}
func (m *ListRecommendationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRecommendationsRequest.Marshal(b, m, deterministic)
}
func (m *ListRecommendationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRecommendationsRequest.Merge(m, src)
}
func (m *ListRecommendationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListRecommendationsRequest.Size(m)
}
func (m *ListRecommendationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRecommendationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRecommendationsRequest proto.InternalMessageInfo

func (m *ListRecommendationsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListRecommendationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListRecommendationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListRecommendationsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// Response to the `ListRecommendations` method.
type ListRecommendationsResponse struct {
	// The set of recommendations for the `parent` resource.
	Recommendations []*Recommendation `protobuf:"bytes,1,rep,name=recommendations,proto3" json:"recommendations,omitempty"`
	// A token that can be used to request the next page of results. This field is
	// empty if there are no additional results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRecommendationsResponse) Reset()         { *m = ListRecommendationsResponse{} }
func (m *ListRecommendationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRecommendationsResponse) ProtoMessage()    {}
func (*ListRecommendationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca265e4ecfbc6d8, []int{1}
}

func (m *ListRecommendationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRecommendationsResponse.Unmarshal(m, b)
}
func (m *ListRecommendationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRecommendationsResponse.Marshal(b, m, deterministic)
}
func (m *ListRecommendationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRecommendationsResponse.Merge(m, src)
}
func (m *ListRecommendationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListRecommendationsResponse.Size(m)
}
func (m *ListRecommendationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRecommendationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRecommendationsResponse proto.InternalMessageInfo

func (m *ListRecommendationsResponse) GetRecommendations() []*Recommendation {
	if m != nil {
		return m.Recommendations
	}
	return nil
}

func (m *ListRecommendationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request to the `GetRecommendation` method.
type GetRecommendationRequest struct {
	// Name of the recommendation.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecommendationRequest) Reset()         { *m = GetRecommendationRequest{} }
func (m *GetRecommendationRequest) String() string { return proto.CompactTextString(m) }
func (*GetRecommendationRequest) ProtoMessage()    {}
func (*GetRecommendationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca265e4ecfbc6d8, []int{2}
}

func (m *GetRecommendationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecommendationRequest.Unmarshal(m, b)
}
func (m *GetRecommendationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecommendationRequest.Marshal(b, m, deterministic)
}
func (m *GetRecommendationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecommendationRequest.Merge(m, src)
}
func (m *GetRecommendationRequest) XXX_Size() int {
	return xxx_messageInfo_GetRecommendationRequest.Size(m)
}
func (m *GetRecommendationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecommendationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecommendationRequest proto.InternalMessageInfo

func (m *GetRecommendationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request for the `MarkRecommendationClaimed` Method.
type MarkRecommendationClaimedRequest struct {
	// Name of the recommendation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// State properties to include with this state. Overwrites any existing
	// `state_metadata`.
	StateMetadata map[string]string `protobuf:"bytes,2,rep,name=state_metadata,json=stateMetadata,proto3" json:"state_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Fingerprint of the Recommendation. Provides optimistic locking.
	Etag                 string   `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkRecommendationClaimedRequest) Reset()         { *m = MarkRecommendationClaimedRequest{} }
func (m *MarkRecommendationClaimedRequest) String() string { return proto.CompactTextString(m) }
func (*MarkRecommendationClaimedRequest) ProtoMessage()    {}
func (*MarkRecommendationClaimedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca265e4ecfbc6d8, []int{3}
}

func (m *MarkRecommendationClaimedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkRecommendationClaimedRequest.Unmarshal(m, b)
}
func (m *MarkRecommendationClaimedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkRecommendationClaimedRequest.Marshal(b, m, deterministic)
}
func (m *MarkRecommendationClaimedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkRecommendationClaimedRequest.Merge(m, src)
}
func (m *MarkRecommendationClaimedRequest) XXX_Size() int {
	return xxx_messageInfo_MarkRecommendationClaimedRequest.Size(m)
}
func (m *MarkRecommendationClaimedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkRecommendationClaimedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarkRecommendationClaimedRequest proto.InternalMessageInfo

func (m *MarkRecommendationClaimedRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MarkRecommendationClaimedRequest) GetStateMetadata() map[string]string {
	if m != nil {
		return m.StateMetadata
	}
	return nil
}

func (m *MarkRecommendationClaimedRequest) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

// Request for the `MarkRecommendationSucceeded` Method.
type MarkRecommendationSucceededRequest struct {
	// Name of the recommendation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// State properties to include with this state. Overwrites any existing
	// `state_metadata`.
	StateMetadata map[string]string `protobuf:"bytes,2,rep,name=state_metadata,json=stateMetadata,proto3" json:"state_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Fingerprint of the Recommendation. Provides optimistic locking.
	Etag                 string   `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkRecommendationSucceededRequest) Reset()         { *m = MarkRecommendationSucceededRequest{} }
func (m *MarkRecommendationSucceededRequest) String() string { return proto.CompactTextString(m) }
func (*MarkRecommendationSucceededRequest) ProtoMessage()    {}
func (*MarkRecommendationSucceededRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca265e4ecfbc6d8, []int{4}
}

func (m *MarkRecommendationSucceededRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkRecommendationSucceededRequest.Unmarshal(m, b)
}
func (m *MarkRecommendationSucceededRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkRecommendationSucceededRequest.Marshal(b, m, deterministic)
}
func (m *MarkRecommendationSucceededRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkRecommendationSucceededRequest.Merge(m, src)
}
func (m *MarkRecommendationSucceededRequest) XXX_Size() int {
	return xxx_messageInfo_MarkRecommendationSucceededRequest.Size(m)
}
func (m *MarkRecommendationSucceededRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkRecommendationSucceededRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarkRecommendationSucceededRequest proto.InternalMessageInfo

func (m *MarkRecommendationSucceededRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MarkRecommendationSucceededRequest) GetStateMetadata() map[string]string {
	if m != nil {
		return m.StateMetadata
	}
	return nil
}

func (m *MarkRecommendationSucceededRequest) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

// Request for the `MarkRecommendationFailed` Method.
type MarkRecommendationFailedRequest struct {
	// Name of the recommendation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// State properties to include with this state. Overwrites any existing
	// `state_metadata`.
	StateMetadata map[string]string `protobuf:"bytes,2,rep,name=state_metadata,json=stateMetadata,proto3" json:"state_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Fingerprint of the Recommendation. Provides optimistic locking.
	Etag                 string   `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkRecommendationFailedRequest) Reset()         { *m = MarkRecommendationFailedRequest{} }
func (m *MarkRecommendationFailedRequest) String() string { return proto.CompactTextString(m) }
func (*MarkRecommendationFailedRequest) ProtoMessage()    {}
func (*MarkRecommendationFailedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca265e4ecfbc6d8, []int{5}
}

func (m *MarkRecommendationFailedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkRecommendationFailedRequest.Unmarshal(m, b)
}
func (m *MarkRecommendationFailedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkRecommendationFailedRequest.Marshal(b, m, deterministic)
}
func (m *MarkRecommendationFailedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkRecommendationFailedRequest.Merge(m, src)
}
func (m *MarkRecommendationFailedRequest) XXX_Size() int {
	return xxx_messageInfo_MarkRecommendationFailedRequest.Size(m)
}
func (m *MarkRecommendationFailedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkRecommendationFailedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarkRecommendationFailedRequest proto.InternalMessageInfo

func (m *MarkRecommendationFailedRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MarkRecommendationFailedRequest) GetStateMetadata() map[string]string {
	if m != nil {
		return m.StateMetadata
	}
	return nil
}

func (m *MarkRecommendationFailedRequest) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func init() {
	proto.RegisterType((*ListRecommendationsRequest)(nil), "google.cloud.recommender.v1beta1.ListRecommendationsRequest")
	proto.RegisterType((*ListRecommendationsResponse)(nil), "google.cloud.recommender.v1beta1.ListRecommendationsResponse")
	proto.RegisterType((*GetRecommendationRequest)(nil), "google.cloud.recommender.v1beta1.GetRecommendationRequest")
	proto.RegisterType((*MarkRecommendationClaimedRequest)(nil), "google.cloud.recommender.v1beta1.MarkRecommendationClaimedRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.recommender.v1beta1.MarkRecommendationClaimedRequest.StateMetadataEntry")
	proto.RegisterType((*MarkRecommendationSucceededRequest)(nil), "google.cloud.recommender.v1beta1.MarkRecommendationSucceededRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.recommender.v1beta1.MarkRecommendationSucceededRequest.StateMetadataEntry")
	proto.RegisterType((*MarkRecommendationFailedRequest)(nil), "google.cloud.recommender.v1beta1.MarkRecommendationFailedRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.recommender.v1beta1.MarkRecommendationFailedRequest.StateMetadataEntry")
}

func init() {
	proto.RegisterFile("google/cloud/recommender/v1beta1/recommender_service.proto", fileDescriptor_0ca265e4ecfbc6d8)
}

var fileDescriptor_0ca265e4ecfbc6d8 = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcf, 0x6f, 0xd3, 0x48,
	0x18, 0xd5, 0xb8, 0x3f, 0xb4, 0x9d, 0xaa, 0xdb, 0xee, 0xec, 0x6a, 0xd7, 0xeb, 0xee, 0xaa, 0x59,
	0x6b, 0x85, 0xaa, 0x48, 0xd8, 0xb4, 0x08, 0x09, 0x05, 0x2a, 0xd1, 0x84, 0x52, 0x21, 0x28, 0x54,
	0x4e, 0x4b, 0x51, 0x25, 0x08, 0x53, 0xe7, 0xab, 0x6b, 0x62, 0x7b, 0x8c, 0x3d, 0x49, 0x69, 0x4b,
	0x39, 0x70, 0xaa, 0x38, 0x70, 0xe1, 0x2f, 0x80, 0x23, 0x47, 0xfe, 0x0c, 0x8e, 0xe5, 0x4f, 0x80,
	0x33, 0x12, 0xa7, 0x5e, 0x91, 0xc7, 0x6e, 0xeb, 0x24, 0x4d, 0x52, 0xd2, 0x48, 0xbd, 0xcd, 0x7c,
	0x33, 0xef, 0xfb, 0xde, 0x7b, 0xb6, 0x9e, 0x06, 0xe7, 0x2c, 0xc6, 0x2c, 0x07, 0x74, 0xd3, 0x61,
	0xd5, 0xb2, 0x1e, 0x80, 0xc9, 0x5c, 0x17, 0xbc, 0x32, 0x04, 0x7a, 0x6d, 0x6a, 0x0d, 0x38, 0x9d,
	0x4a, 0xd7, 0x4a, 0x21, 0x04, 0x35, 0xdb, 0x04, 0xcd, 0x0f, 0x18, 0x67, 0x24, 0x13, 0x63, 0x35,
	0x81, 0xd5, 0x52, 0xf7, 0xb4, 0x04, 0xab, 0xfc, 0x93, 0x74, 0xa7, 0xbe, 0xad, 0x53, 0xcf, 0x63,
	0x9c, 0x72, 0x9b, 0x79, 0x61, 0x8c, 0x57, 0x26, 0x52, 0xa7, 0xeb, 0x36, 0x38, 0xe5, 0xd2, 0x1a,
	0x6c, 0xd0, 0x9a, 0xcd, 0x82, 0xe4, 0xc2, 0x95, 0xd3, 0x93, 0x13, 0x8d, 0x13, 0xd8, 0x5f, 0xa9,
	0xbe, 0xa6, 0x63, 0x83, 0xc7, 0xe3, 0x03, 0x75, 0x0f, 0x61, 0xe5, 0xae, 0x1d, 0x72, 0xa3, 0x0e,
	0x15, 0x1a, 0xf0, 0xac, 0x0a, 0x21, 0x27, 0x7f, 0xe2, 0x41, 0x9f, 0x06, 0xe0, 0x71, 0x19, 0x65,
	0xd0, 0xe4, 0x90, 0x91, 0xec, 0xc8, 0x38, 0x1e, 0xf2, 0xa9, 0x05, 0xa5, 0xd0, 0xde, 0x06, 0x59,
	0xca, 0xa0, 0xc9, 0x01, 0xe3, 0x97, 0xa8, 0x50, 0xb4, 0xb7, 0x81, 0xfc, 0x8b, 0xb1, 0x38, 0xe4,
	0xac, 0x02, 0x9e, 0xdc, 0x27, 0x80, 0xe2, 0xfa, 0x52, 0x54, 0x88, 0x7a, 0xae, 0xdb, 0x0e, 0x87,
	0x40, 0x1e, 0x88, 0x7b, 0xc6, 0x3b, 0xf5, 0x1d, 0xc2, 0xe3, 0x27, 0x52, 0x09, 0x7d, 0xe6, 0x85,
	0x40, 0x56, 0xf1, 0x68, 0xbd, 0xb6, 0x50, 0x46, 0x99, 0xbe, 0xc9, 0xe1, 0xe9, 0x4b, 0x5a, 0x27,
	0xd7, 0xb5, 0xfa, 0x9e, 0x46, 0x63, 0x23, 0x72, 0x01, 0x8f, 0x7a, 0xf0, 0x9c, 0x97, 0x52, 0xbc,
	0x25, 0x41, 0x6e, 0x24, 0x2a, 0x2f, 0x1e, 0x72, 0x57, 0x35, 0x2c, 0xcf, 0x43, 0x03, 0xc3, 0x43,
	0xaf, 0x08, 0xee, 0xf7, 0xa8, 0x0b, 0x89, 0x53, 0x62, 0xad, 0xbe, 0x96, 0x70, 0x66, 0x81, 0x06,
	0x95, 0x7a, 0x44, 0xc1, 0xa1, 0xb6, 0x0b, 0xe5, 0x36, 0x40, 0xf2, 0x02, 0xff, 0x1a, 0x72, 0xca,
	0xa1, 0xe4, 0x02, 0xa7, 0x65, 0xca, 0xa9, 0x2c, 0x09, 0xad, 0xcb, 0x9d, 0xb5, 0x76, 0x9a, 0xa7,
	0x15, 0xa3, 0xc6, 0x0b, 0x49, 0xdf, 0x39, 0x8f, 0x07, 0x5b, 0xc6, 0x48, 0x98, 0xae, 0x45, 0x8c,
	0x80, 0x53, 0x2b, 0xf9, 0x76, 0x62, 0xad, 0xdc, 0xc0, 0xa4, 0x19, 0x48, 0xc6, 0x70, 0x5f, 0x05,
	0xb6, 0x12, 0xea, 0xd1, 0x92, 0xfc, 0x81, 0x07, 0x6a, 0xd4, 0xa9, 0x42, 0x62, 0x60, 0xbc, 0xc9,
	0x49, 0x57, 0x91, 0xfa, 0x46, 0xc2, 0x6a, 0x33, 0xb9, 0x62, 0xd5, 0x34, 0x01, 0xca, 0xed, 0xed,
	0x78, 0xd9, 0xc2, 0x8e, 0x95, 0x6e, 0xec, 0x68, 0x9c, 0x78, 0x6e, 0x86, 0xec, 0x49, 0x78, 0xa2,
	0x99, 0xde, 0x2d, 0x6a, 0x3b, 0xed, 0xdd, 0xd8, 0x69, 0xe1, 0xc6, 0x52, 0x37, 0x6e, 0xd4, 0x8d,
	0x3b, 0x2f, 0x2b, 0xa6, 0x0f, 0x86, 0xf0, 0xb0, 0x71, 0xcc, 0x97, 0x7c, 0x41, 0xf8, 0xf7, 0x13,
	0xc2, 0x80, 0x5c, 0xef, 0x2c, 0xb1, 0x75, 0x9c, 0x29, 0x33, 0x5d, 0xa2, 0xe3, 0x04, 0x52, 0xef,
	0xbf, 0xfa, 0xfc, 0xf5, 0xad, 0x74, 0x9b, 0xcc, 0x1f, 0x85, 0xed, 0x4e, 0x9c, 0x87, 0x33, 0x7e,
	0xc0, 0x9e, 0x82, 0xc9, 0x43, 0x3d, 0xab, 0x3b, 0xcc, 0x8c, 0x61, 0x7a, 0x36, 0x9d, 0xcf, 0xa1,
	0x9e, 0xdd, 0xd5, 0x1b, 0x63, 0x67, 0x1f, 0xe1, 0xdf, 0x9a, 0xf2, 0x84, 0xe4, 0x3a, 0xb3, 0x6c,
	0x15, 0x42, 0xca, 0x4f, 0x67, 0xe1, 0x49, 0xa2, 0xa2, 0x9f, 0xec, 0x74, 0x92, 0x1a, 0x15, 0xe9,
	0xd9, 0x5d, 0xf2, 0x1d, 0xe1, 0xbf, 0x5b, 0x66, 0x10, 0xc9, 0x9f, 0x3d, 0xc0, 0xba, 0x10, 0xf9,
	0x58, 0x88, 0x7c, 0xa8, 0x16, 0x7b, 0x24, 0x32, 0xe7, 0xd2, 0xa0, 0x92, 0xb0, 0xca, 0xa1, 0x2c,
	0x39, 0x40, 0x78, 0xbc, 0x4d, 0xd0, 0x90, 0x9b, 0xbd, 0xc8, 0xa9, 0x2e, 0x74, 0x3f, 0x11, 0xba,
	0x57, 0xd5, 0xe5, 0x5e, 0xea, 0x3e, 0xe2, 0x15, 0x29, 0xff, 0x86, 0xb0, 0xdc, 0x2a, 0x54, 0xc8,
	0xec, 0x99, 0x03, 0xa9, 0x0b, 0xcd, 0x8f, 0x84, 0xe6, 0x15, 0xd5, 0xe8, 0xa5, 0xe6, 0x98, 0x54,
	0x0e, 0x65, 0x95, 0x7b, 0x9f, 0x66, 0x95, 0x34, 0x8b, 0x98, 0x1e, 0xf5, 0xed, 0x50, 0x33, 0x99,
	0xbb, 0x3f, 0xab, 0x6d, 0x70, 0xee, 0x87, 0x39, 0x5d, 0xdf, 0xdc, 0xdc, 0x6c, 0x38, 0xd4, 0x69,
	0x95, 0x6f, 0xc4, 0x4f, 0xb7, 0x8b, 0xbe, 0x43, 0xf9, 0x3a, 0x0b, 0xdc, 0xfc, 0x47, 0x84, 0xff,
	0x37, 0x99, 0xdb, 0x51, 0x66, 0x7e, 0x2c, 0x95, 0x8f, 0x8b, 0xd1, 0xe3, 0x6d, 0x11, 0xad, 0xde,
	0x49, 0x50, 0x16, 0x73, 0xa8, 0x67, 0x69, 0x2c, 0xb0, 0x74, 0x0b, 0x3c, 0xf1, 0xb4, 0xd3, 0x8f,
	0x67, 0xb7, 0x7e, 0x2d, 0x5e, 0x4b, 0xd5, 0xde, 0x4b, 0xfd, 0x05, 0x63, 0xae, 0xf0, 0x41, 0xfa,
	0x6f, 0x3e, 0x6e, 0x5a, 0x10, 0x54, 0xe2, 0xa9, 0x09, 0x97, 0x07, 0x53, 0xf9, 0x08, 0xb6, 0x36,
	0x28, 0x46, 0x5c, 0xfe, 0x11, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xbe, 0x7c, 0x2c, 0x2c, 0x0b, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecommenderClient is the client API for Recommender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecommenderClient interface {
	// Lists recommendations for a Cloud project. Requires the recommender.*.list
	// IAM permission for the specified recommender.
	ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, opts ...grpc.CallOption) (*ListRecommendationsResponse, error)
	// Gets the requested recommendation. Requires the recommender.*.get
	// IAM permission for the specified recommender.
	GetRecommendation(ctx context.Context, in *GetRecommendationRequest, opts ...grpc.CallOption) (*Recommendation, error)
	// Marks the Recommendation State as Claimed. Users can use this method to
	// indicate to the Recommender API that they are starting to apply the
	// recommendation themselves. This stops the recommendation content from being
	// updated. Associated insights are frozen and placed in the ACCEPTED state.
	//
	// MarkRecommendationClaimed can be applied to recommendations in CLAIMED or
	// ACTIVE state.
	//
	// Requires the recommender.*.update IAM permission for the specified
	// recommender.
	MarkRecommendationClaimed(ctx context.Context, in *MarkRecommendationClaimedRequest, opts ...grpc.CallOption) (*Recommendation, error)
	// Marks the Recommendation State as Succeeded. Users can use this method to
	// indicate to the Recommender API that they have applied the recommendation
	// themselves, and the operation was successful. This stops the recommendation
	// content from being updated. Associated insights are frozen and placed in
	// the ACCEPTED state.
	//
	// MarkRecommendationSucceeded can be applied to recommendations in ACTIVE,
	// CLAIMED, SUCCEEDED, or FAILED state.
	//
	// Requires the recommender.*.update IAM permission for the specified
	// recommender.
	MarkRecommendationSucceeded(ctx context.Context, in *MarkRecommendationSucceededRequest, opts ...grpc.CallOption) (*Recommendation, error)
	// Marks the Recommendation State as Failed. Users can use this method to
	// indicate to the Recommender API that they have applied the recommendation
	// themselves, and the operation failed. This stops the recommendation content
	// from being updated. Associated insights are frozen and placed in the
	// ACCEPTED state.
	//
	// MarkRecommendationFailed can be applied to recommendations in ACTIVE,
	// CLAIMED, SUCCEEDED, or FAILED state.
	//
	// Requires the recommender.*.update IAM permission for the specified
	// recommender.
	MarkRecommendationFailed(ctx context.Context, in *MarkRecommendationFailedRequest, opts ...grpc.CallOption) (*Recommendation, error)
}

type recommenderClient struct {
	cc *grpc.ClientConn
}

func NewRecommenderClient(cc *grpc.ClientConn) RecommenderClient {
	return &recommenderClient{cc}
}

func (c *recommenderClient) ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, opts ...grpc.CallOption) (*ListRecommendationsResponse, error) {
	out := new(ListRecommendationsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.recommender.v1beta1.Recommender/ListRecommendations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommenderClient) GetRecommendation(ctx context.Context, in *GetRecommendationRequest, opts ...grpc.CallOption) (*Recommendation, error) {
	out := new(Recommendation)
	err := c.cc.Invoke(ctx, "/google.cloud.recommender.v1beta1.Recommender/GetRecommendation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommenderClient) MarkRecommendationClaimed(ctx context.Context, in *MarkRecommendationClaimedRequest, opts ...grpc.CallOption) (*Recommendation, error) {
	out := new(Recommendation)
	err := c.cc.Invoke(ctx, "/google.cloud.recommender.v1beta1.Recommender/MarkRecommendationClaimed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommenderClient) MarkRecommendationSucceeded(ctx context.Context, in *MarkRecommendationSucceededRequest, opts ...grpc.CallOption) (*Recommendation, error) {
	out := new(Recommendation)
	err := c.cc.Invoke(ctx, "/google.cloud.recommender.v1beta1.Recommender/MarkRecommendationSucceeded", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommenderClient) MarkRecommendationFailed(ctx context.Context, in *MarkRecommendationFailedRequest, opts ...grpc.CallOption) (*Recommendation, error) {
	out := new(Recommendation)
	err := c.cc.Invoke(ctx, "/google.cloud.recommender.v1beta1.Recommender/MarkRecommendationFailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommenderServer is the server API for Recommender service.
type RecommenderServer interface {
	// Lists recommendations for a Cloud project. Requires the recommender.*.list
	// IAM permission for the specified recommender.
	ListRecommendations(context.Context, *ListRecommendationsRequest) (*ListRecommendationsResponse, error)
	// Gets the requested recommendation. Requires the recommender.*.get
	// IAM permission for the specified recommender.
	GetRecommendation(context.Context, *GetRecommendationRequest) (*Recommendation, error)
	// Marks the Recommendation State as Claimed. Users can use this method to
	// indicate to the Recommender API that they are starting to apply the
	// recommendation themselves. This stops the recommendation content from being
	// updated. Associated insights are frozen and placed in the ACCEPTED state.
	//
	// MarkRecommendationClaimed can be applied to recommendations in CLAIMED or
	// ACTIVE state.
	//
	// Requires the recommender.*.update IAM permission for the specified
	// recommender.
	MarkRecommendationClaimed(context.Context, *MarkRecommendationClaimedRequest) (*Recommendation, error)
	// Marks the Recommendation State as Succeeded. Users can use this method to
	// indicate to the Recommender API that they have applied the recommendation
	// themselves, and the operation was successful. This stops the recommendation
	// content from being updated. Associated insights are frozen and placed in
	// the ACCEPTED state.
	//
	// MarkRecommendationSucceeded can be applied to recommendations in ACTIVE,
	// CLAIMED, SUCCEEDED, or FAILED state.
	//
	// Requires the recommender.*.update IAM permission for the specified
	// recommender.
	MarkRecommendationSucceeded(context.Context, *MarkRecommendationSucceededRequest) (*Recommendation, error)
	// Marks the Recommendation State as Failed. Users can use this method to
	// indicate to the Recommender API that they have applied the recommendation
	// themselves, and the operation failed. This stops the recommendation content
	// from being updated. Associated insights are frozen and placed in the
	// ACCEPTED state.
	//
	// MarkRecommendationFailed can be applied to recommendations in ACTIVE,
	// CLAIMED, SUCCEEDED, or FAILED state.
	//
	// Requires the recommender.*.update IAM permission for the specified
	// recommender.
	MarkRecommendationFailed(context.Context, *MarkRecommendationFailedRequest) (*Recommendation, error)
}

// UnimplementedRecommenderServer can be embedded to have forward compatible implementations.
type UnimplementedRecommenderServer struct {
}

func (*UnimplementedRecommenderServer) ListRecommendations(ctx context.Context, req *ListRecommendationsRequest) (*ListRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecommendations not implemented")
}
func (*UnimplementedRecommenderServer) GetRecommendation(ctx context.Context, req *GetRecommendationRequest) (*Recommendation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendation not implemented")
}
func (*UnimplementedRecommenderServer) MarkRecommendationClaimed(ctx context.Context, req *MarkRecommendationClaimedRequest) (*Recommendation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkRecommendationClaimed not implemented")
}
func (*UnimplementedRecommenderServer) MarkRecommendationSucceeded(ctx context.Context, req *MarkRecommendationSucceededRequest) (*Recommendation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkRecommendationSucceeded not implemented")
}
func (*UnimplementedRecommenderServer) MarkRecommendationFailed(ctx context.Context, req *MarkRecommendationFailedRequest) (*Recommendation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkRecommendationFailed not implemented")
}

func RegisterRecommenderServer(s *grpc.Server, srv RecommenderServer) {
	s.RegisterService(&_Recommender_serviceDesc, srv)
}

func _Recommender_ListRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommenderServer).ListRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommender.v1beta1.Recommender/ListRecommendations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommenderServer).ListRecommendations(ctx, req.(*ListRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recommender_GetRecommendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommenderServer).GetRecommendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommender.v1beta1.Recommender/GetRecommendation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommenderServer).GetRecommendation(ctx, req.(*GetRecommendationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recommender_MarkRecommendationClaimed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkRecommendationClaimedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommenderServer).MarkRecommendationClaimed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommender.v1beta1.Recommender/MarkRecommendationClaimed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommenderServer).MarkRecommendationClaimed(ctx, req.(*MarkRecommendationClaimedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recommender_MarkRecommendationSucceeded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkRecommendationSucceededRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommenderServer).MarkRecommendationSucceeded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommender.v1beta1.Recommender/MarkRecommendationSucceeded",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommenderServer).MarkRecommendationSucceeded(ctx, req.(*MarkRecommendationSucceededRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recommender_MarkRecommendationFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkRecommendationFailedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommenderServer).MarkRecommendationFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommender.v1beta1.Recommender/MarkRecommendationFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommenderServer).MarkRecommendationFailed(ctx, req.(*MarkRecommendationFailedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recommender_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.recommender.v1beta1.Recommender",
	HandlerType: (*RecommenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRecommendations",
			Handler:    _Recommender_ListRecommendations_Handler,
		},
		{
			MethodName: "GetRecommendation",
			Handler:    _Recommender_GetRecommendation_Handler,
		},
		{
			MethodName: "MarkRecommendationClaimed",
			Handler:    _Recommender_MarkRecommendationClaimed_Handler,
		},
		{
			MethodName: "MarkRecommendationSucceeded",
			Handler:    _Recommender_MarkRecommendationSucceeded_Handler,
		},
		{
			MethodName: "MarkRecommendationFailed",
			Handler:    _Recommender_MarkRecommendationFailed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/recommender/v1beta1/recommender_service.proto",
}
