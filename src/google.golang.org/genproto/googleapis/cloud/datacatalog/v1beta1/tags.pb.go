// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/tags.proto

package datacatalog

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FieldType_PrimitiveType int32

const (
	// This is the default invalid value for a type.
	FieldType_PRIMITIVE_TYPE_UNSPECIFIED FieldType_PrimitiveType = 0
	// A double precision number.
	FieldType_DOUBLE FieldType_PrimitiveType = 1
	// An UTF-8 string.
	FieldType_STRING FieldType_PrimitiveType = 2
	// A boolean value.
	FieldType_BOOL FieldType_PrimitiveType = 3
	// A timestamp.
	FieldType_TIMESTAMP FieldType_PrimitiveType = 4
)

var FieldType_PrimitiveType_name = map[int32]string{
	0: "PRIMITIVE_TYPE_UNSPECIFIED",
	1: "DOUBLE",
	2: "STRING",
	3: "BOOL",
	4: "TIMESTAMP",
}

var FieldType_PrimitiveType_value = map[string]int32{
	"PRIMITIVE_TYPE_UNSPECIFIED": 0,
	"DOUBLE":                     1,
	"STRING":                     2,
	"BOOL":                       3,
	"TIMESTAMP":                  4,
}

func (x FieldType_PrimitiveType) String() string {
	return proto.EnumName(FieldType_PrimitiveType_name, int32(x))
}

func (FieldType_PrimitiveType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{4, 0}
}

// Tags are used to attach custom metadata to Data Catalog resources. Tags
// conform to the specifications within their tag template.
//
// See [Data Catalog IAM](/data-catalog/docs/concepts/iam) for information on
// the permissions needed to create or view tags.
type Tag struct {
	// The resource name of the tag in URL format. Example:
	//
	// * projects/{project_id}/locations/{location}/entrygroups/{entry_group_id}/entries/{entry_id}/tags/{tag_id}
	//
	// where `tag_id` is a system-generated identifier.
	// Note that this Tag may not actually be stored in the location in this name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The resource name of the tag template that this tag uses. Example:
	//
	// * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id}
	//
	// This field cannot be modified after creation.
	Template string `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Output only. The display name of the tag template.
	TemplateDisplayName string `protobuf:"bytes,5,opt,name=template_display_name,json=templateDisplayName,proto3" json:"template_display_name,omitempty"`
	// The scope within the parent resource that this tag is attached to. If not
	// provided, the tag is attached to the parent resource itself.
	// Deleting the scope from the parent resource will delete all tags attached
	// to that scope. These fields cannot be updated after creation.
	//
	// Types that are valid to be assigned to Scope:
	//	*Tag_Column
	Scope isTag_Scope `protobuf_oneof:"scope"`
	// Required. This maps the ID of a tag field to the value of and additional information
	// about that field. Valid field IDs are defined by the tag's template. A tag
	// must have at least 1 field and at most 500 fields.
	Fields               map[string]*TagField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{0}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *Tag) GetTemplateDisplayName() string {
	if m != nil {
		return m.TemplateDisplayName
	}
	return ""
}

type isTag_Scope interface {
	isTag_Scope()
}

type Tag_Column struct {
	Column string `protobuf:"bytes,4,opt,name=column,proto3,oneof"`
}

func (*Tag_Column) isTag_Scope() {}

func (m *Tag) GetScope() isTag_Scope {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *Tag) GetColumn() string {
	if x, ok := m.GetScope().(*Tag_Column); ok {
		return x.Column
	}
	return ""
}

func (m *Tag) GetFields() map[string]*TagField {
	if m != nil {
		return m.Fields
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Tag) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Tag_Column)(nil),
	}
}

// Contains the value and supporting information for a field within
// a [Tag][google.cloud.datacatalog.v1beta1.Tag].
type TagField struct {
	// Output only. The display name of this field.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The value of this field.
	//
	// Types that are valid to be assigned to Kind:
	//	*TagField_DoubleValue
	//	*TagField_StringValue
	//	*TagField_BoolValue
	//	*TagField_TimestampValue
	//	*TagField_EnumValue_
	Kind                 isTagField_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TagField) Reset()         { *m = TagField{} }
func (m *TagField) String() string { return proto.CompactTextString(m) }
func (*TagField) ProtoMessage()    {}
func (*TagField) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{1}
}

func (m *TagField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagField.Unmarshal(m, b)
}
func (m *TagField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagField.Marshal(b, m, deterministic)
}
func (m *TagField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagField.Merge(m, src)
}
func (m *TagField) XXX_Size() int {
	return xxx_messageInfo_TagField.Size(m)
}
func (m *TagField) XXX_DiscardUnknown() {
	xxx_messageInfo_TagField.DiscardUnknown(m)
}

var xxx_messageInfo_TagField proto.InternalMessageInfo

func (m *TagField) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type isTagField_Kind interface {
	isTagField_Kind()
}

type TagField_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,2,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type TagField_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type TagField_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type TagField_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,5,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type TagField_EnumValue_ struct {
	EnumValue *TagField_EnumValue `protobuf:"bytes,6,opt,name=enum_value,json=enumValue,proto3,oneof"`
}

func (*TagField_DoubleValue) isTagField_Kind() {}

func (*TagField_StringValue) isTagField_Kind() {}

func (*TagField_BoolValue) isTagField_Kind() {}

func (*TagField_TimestampValue) isTagField_Kind() {}

func (*TagField_EnumValue_) isTagField_Kind() {}

func (m *TagField) GetKind() isTagField_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *TagField) GetDoubleValue() float64 {
	if x, ok := m.GetKind().(*TagField_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *TagField) GetStringValue() string {
	if x, ok := m.GetKind().(*TagField_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *TagField) GetBoolValue() bool {
	if x, ok := m.GetKind().(*TagField_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *TagField) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetKind().(*TagField_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *TagField) GetEnumValue() *TagField_EnumValue {
	if x, ok := m.GetKind().(*TagField_EnumValue_); ok {
		return x.EnumValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagField) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagField_DoubleValue)(nil),
		(*TagField_StringValue)(nil),
		(*TagField_BoolValue)(nil),
		(*TagField_TimestampValue)(nil),
		(*TagField_EnumValue_)(nil),
	}
}

// Holds an enum value.
type TagField_EnumValue struct {
	// The display name of the enum value.
	DisplayName          string   `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagField_EnumValue) Reset()         { *m = TagField_EnumValue{} }
func (m *TagField_EnumValue) String() string { return proto.CompactTextString(m) }
func (*TagField_EnumValue) ProtoMessage()    {}
func (*TagField_EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{1, 0}
}

func (m *TagField_EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagField_EnumValue.Unmarshal(m, b)
}
func (m *TagField_EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagField_EnumValue.Marshal(b, m, deterministic)
}
func (m *TagField_EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagField_EnumValue.Merge(m, src)
}
func (m *TagField_EnumValue) XXX_Size() int {
	return xxx_messageInfo_TagField_EnumValue.Size(m)
}
func (m *TagField_EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TagField_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_TagField_EnumValue proto.InternalMessageInfo

func (m *TagField_EnumValue) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// A tag template defines the schema of the tags used to attach to Data Catalog
// resources. It defines the mapping of accepted field names and types that can
// be used within the tag. The tag template also controls the access to the tag.
type TagTemplate struct {
	// The resource name of the tag template in URL format. Example:
	//
	// * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id}
	//
	// Note that this TagTemplate and its child resources may not actually be
	// stored in the location in this name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The display name for this template. Defaults to an empty string.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. Map of tag template field IDs to the settings for the field.
	// This map is an exhaustive list of the allowed fields. This map must contain
	// at least one field and at most 500 fields.
	//
	// The keys to this map are tag template field IDs. Field IDs can contain
	// letters (both uppercase and lowercase), numbers (0-9) and underscores (_).
	// Field IDs must be at least 1 character long and at most
	// 64 characters long. Field IDs must start with a letter or underscore.
	Fields               map[string]*TagTemplateField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *TagTemplate) Reset()         { *m = TagTemplate{} }
func (m *TagTemplate) String() string { return proto.CompactTextString(m) }
func (*TagTemplate) ProtoMessage()    {}
func (*TagTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{2}
}

func (m *TagTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagTemplate.Unmarshal(m, b)
}
func (m *TagTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagTemplate.Marshal(b, m, deterministic)
}
func (m *TagTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagTemplate.Merge(m, src)
}
func (m *TagTemplate) XXX_Size() int {
	return xxx_messageInfo_TagTemplate.Size(m)
}
func (m *TagTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_TagTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_TagTemplate proto.InternalMessageInfo

func (m *TagTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TagTemplate) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TagTemplate) GetFields() map[string]*TagTemplateField {
	if m != nil {
		return m.Fields
	}
	return nil
}

// The template for an individual field within a tag template.
type TagTemplateField struct {
	// Output only. The resource name of the tag template field in URL format. Example:
	//
	// * projects/{project_id}/locations/{location}/tagTemplates/{tag_template}/fields/{field}
	//
	// Note that this TagTemplateField may not actually be stored in the location
	// in this name.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// The display name for this field. Defaults to an empty string.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The type of value this tag field can contain.
	Type *FieldType `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Whether this is a required field. Defaults to false.
	IsRequired           bool     `protobuf:"varint,3,opt,name=is_required,json=isRequired,proto3" json:"is_required,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagTemplateField) Reset()         { *m = TagTemplateField{} }
func (m *TagTemplateField) String() string { return proto.CompactTextString(m) }
func (*TagTemplateField) ProtoMessage()    {}
func (*TagTemplateField) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{3}
}

func (m *TagTemplateField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagTemplateField.Unmarshal(m, b)
}
func (m *TagTemplateField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagTemplateField.Marshal(b, m, deterministic)
}
func (m *TagTemplateField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagTemplateField.Merge(m, src)
}
func (m *TagTemplateField) XXX_Size() int {
	return xxx_messageInfo_TagTemplateField.Size(m)
}
func (m *TagTemplateField) XXX_DiscardUnknown() {
	xxx_messageInfo_TagTemplateField.DiscardUnknown(m)
}

var xxx_messageInfo_TagTemplateField proto.InternalMessageInfo

func (m *TagTemplateField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TagTemplateField) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TagTemplateField) GetType() *FieldType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *TagTemplateField) GetIsRequired() bool {
	if m != nil {
		return m.IsRequired
	}
	return false
}

type FieldType struct {
	// Required.
	//
	// Types that are valid to be assigned to TypeDecl:
	//	*FieldType_PrimitiveType_
	//	*FieldType_EnumType_
	TypeDecl             isFieldType_TypeDecl `protobuf_oneof:"type_decl"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FieldType) Reset()         { *m = FieldType{} }
func (m *FieldType) String() string { return proto.CompactTextString(m) }
func (*FieldType) ProtoMessage()    {}
func (*FieldType) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{4}
}

func (m *FieldType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldType.Unmarshal(m, b)
}
func (m *FieldType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldType.Marshal(b, m, deterministic)
}
func (m *FieldType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldType.Merge(m, src)
}
func (m *FieldType) XXX_Size() int {
	return xxx_messageInfo_FieldType.Size(m)
}
func (m *FieldType) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldType.DiscardUnknown(m)
}

var xxx_messageInfo_FieldType proto.InternalMessageInfo

type isFieldType_TypeDecl interface {
	isFieldType_TypeDecl()
}

type FieldType_PrimitiveType_ struct {
	PrimitiveType FieldType_PrimitiveType `protobuf:"varint,1,opt,name=primitive_type,json=primitiveType,proto3,enum=google.cloud.datacatalog.v1beta1.FieldType_PrimitiveType,oneof"`
}

type FieldType_EnumType_ struct {
	EnumType *FieldType_EnumType `protobuf:"bytes,2,opt,name=enum_type,json=enumType,proto3,oneof"`
}

func (*FieldType_PrimitiveType_) isFieldType_TypeDecl() {}

func (*FieldType_EnumType_) isFieldType_TypeDecl() {}

func (m *FieldType) GetTypeDecl() isFieldType_TypeDecl {
	if m != nil {
		return m.TypeDecl
	}
	return nil
}

func (m *FieldType) GetPrimitiveType() FieldType_PrimitiveType {
	if x, ok := m.GetTypeDecl().(*FieldType_PrimitiveType_); ok {
		return x.PrimitiveType
	}
	return FieldType_PRIMITIVE_TYPE_UNSPECIFIED
}

func (m *FieldType) GetEnumType() *FieldType_EnumType {
	if x, ok := m.GetTypeDecl().(*FieldType_EnumType_); ok {
		return x.EnumType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FieldType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FieldType_PrimitiveType_)(nil),
		(*FieldType_EnumType_)(nil),
	}
}

type FieldType_EnumType struct {
	// Required on create; optional on update. The set of allowed values for
	// this enum. This set must not be empty, the display names of the values in
	// this set must not be empty and the display names of the values must be
	// case-insensitively unique within this set. Currently, enum values can
	// only be added to the list of allowed values. Deletion and renaming of
	// enum values are not supported. Can have up to 500 allowed values.
	AllowedValues        []*FieldType_EnumType_EnumValue `protobuf:"bytes,1,rep,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FieldType_EnumType) Reset()         { *m = FieldType_EnumType{} }
func (m *FieldType_EnumType) String() string { return proto.CompactTextString(m) }
func (*FieldType_EnumType) ProtoMessage()    {}
func (*FieldType_EnumType) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{4, 0}
}

func (m *FieldType_EnumType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldType_EnumType.Unmarshal(m, b)
}
func (m *FieldType_EnumType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldType_EnumType.Marshal(b, m, deterministic)
}
func (m *FieldType_EnumType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldType_EnumType.Merge(m, src)
}
func (m *FieldType_EnumType) XXX_Size() int {
	return xxx_messageInfo_FieldType_EnumType.Size(m)
}
func (m *FieldType_EnumType) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldType_EnumType.DiscardUnknown(m)
}

var xxx_messageInfo_FieldType_EnumType proto.InternalMessageInfo

func (m *FieldType_EnumType) GetAllowedValues() []*FieldType_EnumType_EnumValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

type FieldType_EnumType_EnumValue struct {
	// Required. The display name of the enum value. Must not be an empty string.
	DisplayName          string   `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldType_EnumType_EnumValue) Reset()         { *m = FieldType_EnumType_EnumValue{} }
func (m *FieldType_EnumType_EnumValue) String() string { return proto.CompactTextString(m) }
func (*FieldType_EnumType_EnumValue) ProtoMessage()    {}
func (*FieldType_EnumType_EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{4, 0, 0}
}

func (m *FieldType_EnumType_EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldType_EnumType_EnumValue.Unmarshal(m, b)
}
func (m *FieldType_EnumType_EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldType_EnumType_EnumValue.Marshal(b, m, deterministic)
}
func (m *FieldType_EnumType_EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldType_EnumType_EnumValue.Merge(m, src)
}
func (m *FieldType_EnumType_EnumValue) XXX_Size() int {
	return xxx_messageInfo_FieldType_EnumType_EnumValue.Size(m)
}
func (m *FieldType_EnumType_EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldType_EnumType_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_FieldType_EnumType_EnumValue proto.InternalMessageInfo

func (m *FieldType_EnumType_EnumValue) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.datacatalog.v1beta1.FieldType_PrimitiveType", FieldType_PrimitiveType_name, FieldType_PrimitiveType_value)
	proto.RegisterType((*Tag)(nil), "google.cloud.datacatalog.v1beta1.Tag")
	proto.RegisterMapType((map[string]*TagField)(nil), "google.cloud.datacatalog.v1beta1.Tag.FieldsEntry")
	proto.RegisterType((*TagField)(nil), "google.cloud.datacatalog.v1beta1.TagField")
	proto.RegisterType((*TagField_EnumValue)(nil), "google.cloud.datacatalog.v1beta1.TagField.EnumValue")
	proto.RegisterType((*TagTemplate)(nil), "google.cloud.datacatalog.v1beta1.TagTemplate")
	proto.RegisterMapType((map[string]*TagTemplateField)(nil), "google.cloud.datacatalog.v1beta1.TagTemplate.FieldsEntry")
	proto.RegisterType((*TagTemplateField)(nil), "google.cloud.datacatalog.v1beta1.TagTemplateField")
	proto.RegisterType((*FieldType)(nil), "google.cloud.datacatalog.v1beta1.FieldType")
	proto.RegisterType((*FieldType_EnumType)(nil), "google.cloud.datacatalog.v1beta1.FieldType.EnumType")
	proto.RegisterType((*FieldType_EnumType_EnumValue)(nil), "google.cloud.datacatalog.v1beta1.FieldType.EnumType.EnumValue")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/tags.proto", fileDescriptor_6fd303c40f581309)
}

var fileDescriptor_6fd303c40f581309 = []byte{
	// 924 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x8f, 0xed, 0x34, 0x24, 0xcf, 0xdb, 0x12, 0x0d, 0x42, 0x84, 0x1c, 0xb6, 0x25, 0xa0, 0xd5,
	0x8a, 0x4a, 0xb6, 0xda, 0x5d, 0x09, 0x28, 0x12, 0x22, 0xd9, 0x7a, 0x37, 0x11, 0xdb, 0x36, 0x72,
	0xdd, 0x4a, 0x20, 0x24, 0x33, 0xb1, 0x67, 0x8d, 0x59, 0xdb, 0x63, 0xfc, 0xa7, 0x10, 0x45, 0x41,
	0xda, 0x03, 0x5f, 0x85, 0x6f, 0xc1, 0x87, 0xe0, 0xce, 0x81, 0x73, 0xaf, 0x5c, 0xb8, 0x20, 0x21,
	0xcf, 0x8c, 0x13, 0xa7, 0x7f, 0x68, 0x16, 0x71, 0xca, 0xbc, 0x37, 0xbf, 0xf7, 0x9b, 0x79, 0xbf,
	0xf7, 0xde, 0xc4, 0xb0, 0xeb, 0x51, 0xea, 0x05, 0x44, 0x77, 0x02, 0x9a, 0xbb, 0xba, 0x8b, 0x33,
	0xec, 0xe0, 0x0c, 0x07, 0xd4, 0xd3, 0x2f, 0xf6, 0x26, 0x24, 0xc3, 0x7b, 0x7a, 0x86, 0xbd, 0x54,
	0x8b, 0x13, 0x9a, 0x51, 0xb4, 0xc3, 0xc1, 0x1a, 0x03, 0x6b, 0x15, 0xb0, 0x26, 0xc0, 0xdd, 0x6d,
	0x41, 0x87, 0x63, 0x5f, 0x7f, 0xe1, 0x93, 0xc0, 0xb5, 0x27, 0xe4, 0x5b, 0x7c, 0xe1, 0xd3, 0x84,
	0x53, 0x74, 0xdf, 0xad, 0x00, 0x12, 0x92, 0xd2, 0x3c, 0x71, 0x88, 0xd8, 0x2a, 0x63, 0x99, 0x35,
	0xc9, 0x5f, 0xe8, 0x99, 0x1f, 0x92, 0x34, 0xc3, 0x61, 0xcc, 0x01, 0xbd, 0xdf, 0x14, 0x50, 0x2c,
	0xec, 0x21, 0x04, 0xf5, 0x08, 0x87, 0xa4, 0x23, 0xed, 0x48, 0x0f, 0x5b, 0x26, 0x5b, 0xa3, 0x6d,
	0x68, 0x66, 0x24, 0x8c, 0x03, 0x9c, 0x91, 0x8e, 0x5c, 0xf8, 0x07, 0xca, 0x1f, 0x7d, 0xd9, 0x5c,
	0x38, 0xd1, 0x47, 0xf0, 0x76, 0xb9, 0xb6, 0x5d, 0x3f, 0x8d, 0x03, 0x3c, 0xb5, 0x19, 0xcb, 0x46,
	0x89, 0x56, 0xcc, 0xb7, 0x4a, 0xc4, 0x21, 0x07, 0x1c, 0x17, 0xcc, 0x1d, 0x68, 0x38, 0x34, 0xc8,
	0xc3, 0xa8, 0x53, 0x2f, 0x90, 0xc3, 0x9a, 0x29, 0x6c, 0x74, 0x0c, 0x0d, 0x96, 0x63, 0xda, 0x51,
	0x76, 0x94, 0x87, 0xea, 0xfe, 0x9e, 0x76, 0x97, 0x3e, 0x9a, 0x85, 0x3d, 0xed, 0x29, 0x8b, 0x31,
	0xa2, 0x2c, 0x99, 0xf2, 0x4b, 0x0a, 0x96, 0x2e, 0x01, 0xb5, 0xb2, 0x87, 0xda, 0xa0, 0xbc, 0x24,
	0x53, 0x91, 0x65, 0xb1, 0x44, 0x9f, 0xc3, 0xc6, 0x05, 0x0e, 0x72, 0x9e, 0xa1, 0xba, 0xff, 0xe1,
	0x5a, 0xe7, 0x31, 0x4a, 0x93, 0x07, 0x1e, 0xc8, 0x1f, 0x4b, 0x07, 0xaf, 0xa4, 0xcb, 0xfe, 0x4f,
	0x70, 0xbf, 0x0a, 0xe7, 0x3c, 0x38, 0xf6, 0x53, 0xcd, 0xa1, 0xa1, 0x5e, 0x88, 0xfc, 0x75, 0x9c,
	0xd0, 0xef, 0x88, 0x93, 0xa5, 0xfa, 0x4c, 0xac, 0xe6, 0x7a, 0x40, 0x1d, 0x9c, 0xf9, 0x34, 0x4a,
	0xf5, 0x59, 0xb9, 0x9c, 0xeb, 0xa4, 0xb8, 0xec, 0xb3, 0x84, 0xe6, 0x71, 0xaa, 0xcf, 0x98, 0x61,
	0x7b, 0x85, 0xc5, 0x77, 0x7c, 0x52, 0x7a, 0xe7, 0xac, 0x93, 0xf4, 0x59, 0x86, 0xbd, 0xf9, 0xe0,
	0x0d, 0xd8, 0x48, 0x1d, 0x1a, 0x93, 0xde, 0x9f, 0x32, 0x34, 0xcb, 0x4b, 0xa2, 0x07, 0x70, 0x6f,
	0xa5, 0x34, 0xd2, 0xb2, 0x34, 0xaa, 0x5b, 0x29, 0xc9, 0xfb, 0x70, 0xcf, 0xa5, 0xf9, 0x24, 0x20,
	0xf6, 0x52, 0x0e, 0x69, 0x58, 0x33, 0x55, 0xee, 0x3d, 0x2f, 0x9c, 0x05, 0x28, 0xcd, 0x12, 0x3f,
	0xf2, 0x04, 0x48, 0x11, 0xd5, 0x53, 0xb9, 0x97, 0x83, 0xb6, 0x01, 0x26, 0x94, 0x06, 0x02, 0x52,
	0x14, 0xb8, 0x39, 0xac, 0x99, 0xad, 0xc2, 0xc7, 0x01, 0x06, 0xbc, 0xb9, 0x68, 0x43, 0x81, 0xda,
	0x60, 0xe2, 0x77, 0x4b, 0xf1, 0xcb, 0x76, 0xd5, 0xac, 0x12, 0x37, 0xac, 0x99, 0x5b, 0x8b, 0x20,
	0x4e, 0x73, 0x06, 0x40, 0xa2, 0x3c, 0x14, 0x0c, 0x0d, 0xc6, 0xf0, 0x78, 0xfd, 0xf2, 0x69, 0x46,
	0x94, 0x87, 0x8c, 0xa9, 0xb8, 0x1d, 0x29, 0x8d, 0xae, 0x06, 0xad, 0xc5, 0x0e, 0x7a, 0xef, 0x26,
	0xf5, 0x56, 0x84, 0x1b, 0x34, 0xa0, 0xfe, 0xd2, 0x8f, 0xdc, 0xde, 0xdf, 0x32, 0xa8, 0x16, 0xf6,
	0xac, 0x72, 0x38, 0x6e, 0x9a, 0xa8, 0xab, 0x74, 0xf2, 0x35, 0x3a, 0x74, 0x7e, 0x65, 0x00, 0x3e,
	0x59, 0x2b, 0xa3, 0xf2, 0xd4, 0xdb, 0x07, 0x21, 0xbc, 0x6b, 0x10, 0x86, 0xab, 0x83, 0xb0, 0xff,
	0x5a, 0xe7, 0x5e, 0x1b, 0x88, 0xf8, 0xb2, 0x1f, 0xc2, 0x83, 0x7f, 0x9f, 0x87, 0x85, 0x54, 0x4f,
	0xd6, 0x9d, 0x8b, 0x6c, 0x19, 0xc4, 0xbb, 0xde, 0x2e, 0x5f, 0x97, 0x79, 0xef, 0x57, 0x19, 0xda,
	0x57, 0x6f, 0x84, 0xde, 0x11, 0x45, 0x68, 0x2c, 0xbb, 0xfe, 0xe6, 0x4a, 0x5c, 0x2f, 0x2c, 0x3a,
	0x84, 0x7a, 0x36, 0x8d, 0x4b, 0x3d, 0x76, 0xef, 0xd6, 0x83, 0x1d, 0x69, 0x4d, 0x63, 0xc2, 0x95,
	0x67, 0xd1, 0x68, 0x1b, 0x54, 0x3f, 0xb5, 0x13, 0xf2, 0x7d, 0xee, 0x27, 0xc4, 0x65, 0x13, 0xd3,
	0x34, 0xc1, 0x4f, 0x4d, 0xe1, 0x39, 0xf8, 0x59, 0xba, 0xec, 0xbf, 0x92, 0x60, 0x77, 0x3d, 0xad,
	0x78, 0x5a, 0xe6, 0xff, 0x20, 0x18, 0xff, 0x2f, 0x49, 0xf5, 0x19, 0xfb, 0x9d, 0xf7, 0x7e, 0x57,
	0xa0, 0xb5, 0xc8, 0x00, 0x4d, 0x60, 0x2b, 0x4e, 0xfc, 0xd0, 0xcf, 0xfc, 0x0b, 0x62, 0x33, 0x19,
	0x0a, 0x85, 0xb6, 0xd6, 0x69, 0xc7, 0x05, 0x89, 0x36, 0x2e, 0x19, 0x0a, 0x6b, 0x58, 0x33, 0x37,
	0xe3, 0xaa, 0x03, 0x9d, 0x02, 0x1b, 0x3b, 0xbb, 0xa2, 0xf2, 0xe3, 0xd7, 0xa1, 0x2f, 0xc6, 0x54,
	0x30, 0x37, 0x89, 0x58, 0x77, 0x7f, 0x91, 0xa0, 0x59, 0x6e, 0x20, 0x02, 0x5b, 0x38, 0x08, 0xe8,
	0x0f, 0xc4, 0xe5, 0xaf, 0x44, 0xda, 0x91, 0xd8, 0x50, 0x7d, 0xf6, 0x5f, 0x8e, 0x59, 0x3e, 0x18,
	0xe6, 0xa6, 0x60, 0x65, 0x56, 0xda, 0x7d, 0x54, 0x7d, 0x32, 0x6e, 0x7d, 0x70, 0xe5, 0x95, 0xf6,
	0xea, 0x7d, 0x03, 0x9b, 0x2b, 0xfa, 0xa0, 0xfb, 0xd0, 0x1d, 0x9b, 0xa3, 0xa3, 0x91, 0x35, 0x3a,
	0x37, 0x6c, 0xeb, 0xcb, 0xb1, 0x61, 0x9f, 0x1d, 0x9f, 0x8e, 0x8d, 0x27, 0xa3, 0xa7, 0x23, 0xe3,
	0xb0, 0x5d, 0x43, 0x00, 0x8d, 0xc3, 0x93, 0xb3, 0xc1, 0x73, 0xa3, 0x2d, 0x15, 0xeb, 0x53, 0xcb,
	0x1c, 0x1d, 0x3f, 0x6b, 0xcb, 0xa8, 0x09, 0xf5, 0xc1, 0xc9, 0xc9, 0xf3, 0xb6, 0x82, 0x36, 0xa1,
	0x65, 0x8d, 0x8e, 0x8c, 0x53, 0xab, 0x7f, 0x34, 0x6e, 0xd7, 0x07, 0x2a, 0xb4, 0x0a, 0x69, 0x6d,
	0x97, 0x38, 0xc1, 0xe0, 0x47, 0xf8, 0xc0, 0xa1, 0xe1, 0x9d, 0x79, 0x8f, 0xa5, 0xaf, 0xbe, 0x10,
	0x18, 0x8f, 0x06, 0x38, 0xf2, 0x34, 0x9a, 0x78, 0xba, 0x47, 0x22, 0xf6, 0x24, 0xeb, 0xcb, 0xc6,
	0xbc, 0xfd, 0xeb, 0xe6, 0xd3, 0x8a, 0xef, 0x2f, 0x49, 0x9a, 0x34, 0x58, 0xe8, 0xa3, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x62, 0xf4, 0x79, 0x9d, 0x17, 0x09, 0x00, 0x00,
}
