// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1beta/scan_config.proto

package websecurityscanner

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

// Type of user agents used for scanning.
type ScanConfig_UserAgent int32

const (
	// The user agent is unknown. Service will default to CHROME_LINUX.
	ScanConfig_USER_AGENT_UNSPECIFIED ScanConfig_UserAgent = 0
	// Chrome on Linux. This is the service default if unspecified.
	ScanConfig_CHROME_LINUX ScanConfig_UserAgent = 1
	// Chrome on Android.
	ScanConfig_CHROME_ANDROID ScanConfig_UserAgent = 2
	// Safari on IPhone.
	ScanConfig_SAFARI_IPHONE ScanConfig_UserAgent = 3
)

var ScanConfig_UserAgent_name = map[int32]string{
	0: "USER_AGENT_UNSPECIFIED",
	1: "CHROME_LINUX",
	2: "CHROME_ANDROID",
	3: "SAFARI_IPHONE",
}

var ScanConfig_UserAgent_value = map[string]int32{
	"USER_AGENT_UNSPECIFIED": 0,
	"CHROME_LINUX":           1,
	"CHROME_ANDROID":         2,
	"SAFARI_IPHONE":          3,
}

func (x ScanConfig_UserAgent) String() string {
	return proto.EnumName(ScanConfig_UserAgent_name, int32(x))
}

func (ScanConfig_UserAgent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b8e0c37c8759a6e, []int{0, 0}
}

// Cloud platforms supported by Cloud Web Security Scanner.
type ScanConfig_TargetPlatform int32

const (
	// The target platform is unknown. Requests with this enum value will be
	// rejected with INVALID_ARGUMENT error.
	ScanConfig_TARGET_PLATFORM_UNSPECIFIED ScanConfig_TargetPlatform = 0
	// Google App Engine service.
	ScanConfig_APP_ENGINE ScanConfig_TargetPlatform = 1
	// Google Compute Engine service.
	ScanConfig_COMPUTE ScanConfig_TargetPlatform = 2
)

var ScanConfig_TargetPlatform_name = map[int32]string{
	0: "TARGET_PLATFORM_UNSPECIFIED",
	1: "APP_ENGINE",
	2: "COMPUTE",
}

var ScanConfig_TargetPlatform_value = map[string]int32{
	"TARGET_PLATFORM_UNSPECIFIED": 0,
	"APP_ENGINE":                  1,
	"COMPUTE":                     2,
}

func (x ScanConfig_TargetPlatform) String() string {
	return proto.EnumName(ScanConfig_TargetPlatform_name, int32(x))
}

func (ScanConfig_TargetPlatform) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b8e0c37c8759a6e, []int{0, 1}
}

// Scan risk levels supported by Cloud Web Security Scanner. LOW impact
// scanning will minimize requests with the potential to modify data. To
// achieve the maximum scan coverage, NORMAL risk level is recommended.
type ScanConfig_RiskLevel int32

const (
	// Use default, which is NORMAL.
	ScanConfig_RISK_LEVEL_UNSPECIFIED ScanConfig_RiskLevel = 0
	// Normal scanning (Recommended)
	ScanConfig_NORMAL ScanConfig_RiskLevel = 1
	// Lower impact scanning
	ScanConfig_LOW ScanConfig_RiskLevel = 2
)

var ScanConfig_RiskLevel_name = map[int32]string{
	0: "RISK_LEVEL_UNSPECIFIED",
	1: "NORMAL",
	2: "LOW",
}

var ScanConfig_RiskLevel_value = map[string]int32{
	"RISK_LEVEL_UNSPECIFIED": 0,
	"NORMAL":                 1,
	"LOW":                    2,
}

func (x ScanConfig_RiskLevel) String() string {
	return proto.EnumName(ScanConfig_RiskLevel_name, int32(x))
}

func (ScanConfig_RiskLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b8e0c37c8759a6e, []int{0, 2}
}

// Controls export of scan configurations and results to Cloud Security
// Command Center.
type ScanConfig_ExportToSecurityCommandCenter int32

const (
	// Use default, which is ENABLED.
	ScanConfig_EXPORT_TO_SECURITY_COMMAND_CENTER_UNSPECIFIED ScanConfig_ExportToSecurityCommandCenter = 0
	// Export results of this scan to Cloud Security Command Center.
	ScanConfig_ENABLED ScanConfig_ExportToSecurityCommandCenter = 1
	// Do not export results of this scan to Cloud Security Command Center.
	ScanConfig_DISABLED ScanConfig_ExportToSecurityCommandCenter = 2
)

var ScanConfig_ExportToSecurityCommandCenter_name = map[int32]string{
	0: "EXPORT_TO_SECURITY_COMMAND_CENTER_UNSPECIFIED",
	1: "ENABLED",
	2: "DISABLED",
}

var ScanConfig_ExportToSecurityCommandCenter_value = map[string]int32{
	"EXPORT_TO_SECURITY_COMMAND_CENTER_UNSPECIFIED": 0,
	"ENABLED":  1,
	"DISABLED": 2,
}

func (x ScanConfig_ExportToSecurityCommandCenter) String() string {
	return proto.EnumName(ScanConfig_ExportToSecurityCommandCenter_name, int32(x))
}

func (ScanConfig_ExportToSecurityCommandCenter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b8e0c37c8759a6e, []int{0, 3}
}

// A ScanConfig resource contains the configurations to launch a scan.
type ScanConfig struct {
	// The resource name of the ScanConfig. The name follows the format of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are
	// generated by the system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The user provided display name of the ScanConfig.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The maximum QPS during scanning. A valid value ranges from 5 to 20
	// inclusively. If the field is unspecified or its value is set 0, server will
	// default to 15. Other values outside of [5, 20] range will be rejected with
	// INVALID_ARGUMENT error.
	MaxQps int32 `protobuf:"varint,3,opt,name=max_qps,json=maxQps,proto3" json:"max_qps,omitempty"`
	// Required. The starting URLs from which the scanner finds site pages.
	StartingUrls []string `protobuf:"bytes,4,rep,name=starting_urls,json=startingUrls,proto3" json:"starting_urls,omitempty"`
	// The authentication configuration. If specified, service will use the
	// authentication configuration during scanning.
	Authentication *ScanConfig_Authentication `protobuf:"bytes,5,opt,name=authentication,proto3" json:"authentication,omitempty"`
	// The user agent used during scanning.
	UserAgent ScanConfig_UserAgent `protobuf:"varint,6,opt,name=user_agent,json=userAgent,proto3,enum=google.cloud.websecurityscanner.v1beta.ScanConfig_UserAgent" json:"user_agent,omitempty"`
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns []string `protobuf:"bytes,7,rep,name=blacklist_patterns,json=blacklistPatterns,proto3" json:"blacklist_patterns,omitempty"`
	// The schedule of the ScanConfig.
	Schedule *ScanConfig_Schedule `protobuf:"bytes,8,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be
	// used as a default.
	TargetPlatforms []ScanConfig_TargetPlatform `protobuf:"varint,9,rep,packed,name=target_platforms,json=targetPlatforms,proto3,enum=google.cloud.websecurityscanner.v1beta.ScanConfig_TargetPlatform" json:"target_platforms,omitempty"`
	// Controls export of scan configurations and results to Cloud Security
	// Command Center.
	ExportToSecurityCommandCenter ScanConfig_ExportToSecurityCommandCenter `protobuf:"varint,10,opt,name=export_to_security_command_center,json=exportToSecurityCommandCenter,proto3,enum=google.cloud.websecurityscanner.v1beta.ScanConfig_ExportToSecurityCommandCenter" json:"export_to_security_command_center,omitempty"`
	// Latest ScanRun if available.
	LatestRun *ScanRun `protobuf:"bytes,11,opt,name=latest_run,json=latestRun,proto3" json:"latest_run,omitempty"`
	// The risk level selected for the scan
	RiskLevel            ScanConfig_RiskLevel `protobuf:"varint,12,opt,name=risk_level,json=riskLevel,proto3,enum=google.cloud.websecurityscanner.v1beta.ScanConfig_RiskLevel" json:"risk_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ScanConfig) Reset()         { *m = ScanConfig{} }
func (m *ScanConfig) String() string { return proto.CompactTextString(m) }
func (*ScanConfig) ProtoMessage()    {}
func (*ScanConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b8e0c37c8759a6e, []int{0}
}

func (m *ScanConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanConfig.Unmarshal(m, b)
}
func (m *ScanConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanConfig.Marshal(b, m, deterministic)
}
func (m *ScanConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanConfig.Merge(m, src)
}
func (m *ScanConfig) XXX_Size() int {
	return xxx_messageInfo_ScanConfig.Size(m)
}
func (m *ScanConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ScanConfig proto.InternalMessageInfo

func (m *ScanConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScanConfig) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ScanConfig) GetMaxQps() int32 {
	if m != nil {
		return m.MaxQps
	}
	return 0
}

func (m *ScanConfig) GetStartingUrls() []string {
	if m != nil {
		return m.StartingUrls
	}
	return nil
}

func (m *ScanConfig) GetAuthentication() *ScanConfig_Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *ScanConfig) GetUserAgent() ScanConfig_UserAgent {
	if m != nil {
		return m.UserAgent
	}
	return ScanConfig_USER_AGENT_UNSPECIFIED
}

func (m *ScanConfig) GetBlacklistPatterns() []string {
	if m != nil {
		return m.BlacklistPatterns
	}
	return nil
}

func (m *ScanConfig) GetSchedule() *ScanConfig_Schedule {
	if m != nil {
		return m.Schedule
	}
	return nil
}

func (m *ScanConfig) GetTargetPlatforms() []ScanConfig_TargetPlatform {
	if m != nil {
		return m.TargetPlatforms
	}
	return nil
}

func (m *ScanConfig) GetExportToSecurityCommandCenter() ScanConfig_ExportToSecurityCommandCenter {
	if m != nil {
		return m.ExportToSecurityCommandCenter
	}
	return ScanConfig_EXPORT_TO_SECURITY_COMMAND_CENTER_UNSPECIFIED
}

func (m *ScanConfig) GetLatestRun() *ScanRun {
	if m != nil {
		return m.LatestRun
	}
	return nil
}

func (m *ScanConfig) GetRiskLevel() ScanConfig_RiskLevel {
	if m != nil {
		return m.RiskLevel
	}
	return ScanConfig_RISK_LEVEL_UNSPECIFIED
}

// Scan authentication configuration.
type ScanConfig_Authentication struct {
	// Required.
	// Authentication configuration
	//
	// Types that are valid to be assigned to Authentication:
	//	*ScanConfig_Authentication_GoogleAccount_
	//	*ScanConfig_Authentication_CustomAccount_
	Authentication       isScanConfig_Authentication_Authentication `protobuf_oneof:"authentication"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *ScanConfig_Authentication) Reset()         { *m = ScanConfig_Authentication{} }
func (m *ScanConfig_Authentication) String() string { return proto.CompactTextString(m) }
func (*ScanConfig_Authentication) ProtoMessage()    {}
func (*ScanConfig_Authentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b8e0c37c8759a6e, []int{0, 0}
}

func (m *ScanConfig_Authentication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanConfig_Authentication.Unmarshal(m, b)
}
func (m *ScanConfig_Authentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanConfig_Authentication.Marshal(b, m, deterministic)
}
func (m *ScanConfig_Authentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanConfig_Authentication.Merge(m, src)
}
func (m *ScanConfig_Authentication) XXX_Size() int {
	return xxx_messageInfo_ScanConfig_Authentication.Size(m)
}
func (m *ScanConfig_Authentication) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanConfig_Authentication.DiscardUnknown(m)
}

var xxx_messageInfo_ScanConfig_Authentication proto.InternalMessageInfo

type isScanConfig_Authentication_Authentication interface {
	isScanConfig_Authentication_Authentication()
}

type ScanConfig_Authentication_GoogleAccount_ struct {
	GoogleAccount *ScanConfig_Authentication_GoogleAccount `protobuf:"bytes,1,opt,name=google_account,json=googleAccount,proto3,oneof"`
}

type ScanConfig_Authentication_CustomAccount_ struct {
	CustomAccount *ScanConfig_Authentication_CustomAccount `protobuf:"bytes,2,opt,name=custom_account,json=customAccount,proto3,oneof"`
}

func (*ScanConfig_Authentication_GoogleAccount_) isScanConfig_Authentication_Authentication() {}

func (*ScanConfig_Authentication_CustomAccount_) isScanConfig_Authentication_Authentication() {}

func (m *ScanConfig_Authentication) GetAuthentication() isScanConfig_Authentication_Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *ScanConfig_Authentication) GetGoogleAccount() *ScanConfig_Authentication_GoogleAccount {
	if x, ok := m.GetAuthentication().(*ScanConfig_Authentication_GoogleAccount_); ok {
		return x.GoogleAccount
	}
	return nil
}

func (m *ScanConfig_Authentication) GetCustomAccount() *ScanConfig_Authentication_CustomAccount {
	if x, ok := m.GetAuthentication().(*ScanConfig_Authentication_CustomAccount_); ok {
		return x.CustomAccount
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ScanConfig_Authentication) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ScanConfig_Authentication_GoogleAccount_)(nil),
		(*ScanConfig_Authentication_CustomAccount_)(nil),
	}
}

// Describes authentication configuration that uses a Google account.
type ScanConfig_Authentication_GoogleAccount struct {
	// Required. The user name of the Google account.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Required. Input only. The password of the Google account. The credential is stored encrypted
	// and not returned in any response nor included in audit logs.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanConfig_Authentication_GoogleAccount) Reset() {
	*m = ScanConfig_Authentication_GoogleAccount{}
}
func (m *ScanConfig_Authentication_GoogleAccount) String() string { return proto.CompactTextString(m) }
func (*ScanConfig_Authentication_GoogleAccount) ProtoMessage()    {}
func (*ScanConfig_Authentication_GoogleAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b8e0c37c8759a6e, []int{0, 0, 0}
}

func (m *ScanConfig_Authentication_GoogleAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanConfig_Authentication_GoogleAccount.Unmarshal(m, b)
}
func (m *ScanConfig_Authentication_GoogleAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanConfig_Authentication_GoogleAccount.Marshal(b, m, deterministic)
}
func (m *ScanConfig_Authentication_GoogleAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanConfig_Authentication_GoogleAccount.Merge(m, src)
}
func (m *ScanConfig_Authentication_GoogleAccount) XXX_Size() int {
	return xxx_messageInfo_ScanConfig_Authentication_GoogleAccount.Size(m)
}
func (m *ScanConfig_Authentication_GoogleAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanConfig_Authentication_GoogleAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ScanConfig_Authentication_GoogleAccount proto.InternalMessageInfo

func (m *ScanConfig_Authentication_GoogleAccount) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ScanConfig_Authentication_GoogleAccount) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Describes authentication configuration that uses a custom account.
type ScanConfig_Authentication_CustomAccount struct {
	// Required. The user name of the custom account.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Required. Input only. The password of the custom account. The credential is stored encrypted
	// and not returned in any response nor included in audit logs.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Required. The login form URL of the website.
	LoginUrl             string   `protobuf:"bytes,3,opt,name=login_url,json=loginUrl,proto3" json:"login_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanConfig_Authentication_CustomAccount) Reset() {
	*m = ScanConfig_Authentication_CustomAccount{}
}
func (m *ScanConfig_Authentication_CustomAccount) String() string { return proto.CompactTextString(m) }
func (*ScanConfig_Authentication_CustomAccount) ProtoMessage()    {}
func (*ScanConfig_Authentication_CustomAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b8e0c37c8759a6e, []int{0, 0, 1}
}

func (m *ScanConfig_Authentication_CustomAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanConfig_Authentication_CustomAccount.Unmarshal(m, b)
}
func (m *ScanConfig_Authentication_CustomAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanConfig_Authentication_CustomAccount.Marshal(b, m, deterministic)
}
func (m *ScanConfig_Authentication_CustomAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanConfig_Authentication_CustomAccount.Merge(m, src)
}
func (m *ScanConfig_Authentication_CustomAccount) XXX_Size() int {
	return xxx_messageInfo_ScanConfig_Authentication_CustomAccount.Size(m)
}
func (m *ScanConfig_Authentication_CustomAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanConfig_Authentication_CustomAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ScanConfig_Authentication_CustomAccount proto.InternalMessageInfo

func (m *ScanConfig_Authentication_CustomAccount) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ScanConfig_Authentication_CustomAccount) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ScanConfig_Authentication_CustomAccount) GetLoginUrl() string {
	if m != nil {
		return m.LoginUrl
	}
	return ""
}

// Scan schedule configuration.
type ScanConfig_Schedule struct {
	// A timestamp indicates when the next run will be scheduled. The value is
	// refreshed by the server after each run. If unspecified, it will default
	// to current server time, which means the scan will be scheduled to start
	// immediately.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Required. The duration of time between executions in days.
	IntervalDurationDays int32    `protobuf:"varint,2,opt,name=interval_duration_days,json=intervalDurationDays,proto3" json:"interval_duration_days,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanConfig_Schedule) Reset()         { *m = ScanConfig_Schedule{} }
func (m *ScanConfig_Schedule) String() string { return proto.CompactTextString(m) }
func (*ScanConfig_Schedule) ProtoMessage()    {}
func (*ScanConfig_Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b8e0c37c8759a6e, []int{0, 1}
}

func (m *ScanConfig_Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanConfig_Schedule.Unmarshal(m, b)
}
func (m *ScanConfig_Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanConfig_Schedule.Marshal(b, m, deterministic)
}
func (m *ScanConfig_Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanConfig_Schedule.Merge(m, src)
}
func (m *ScanConfig_Schedule) XXX_Size() int {
	return xxx_messageInfo_ScanConfig_Schedule.Size(m)
}
func (m *ScanConfig_Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanConfig_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_ScanConfig_Schedule proto.InternalMessageInfo

func (m *ScanConfig_Schedule) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *ScanConfig_Schedule) GetIntervalDurationDays() int32 {
	if m != nil {
		return m.IntervalDurationDays
	}
	return 0
}

func init() {
	proto.RegisterEnum("google.cloud.websecurityscanner.v1beta.ScanConfig_UserAgent", ScanConfig_UserAgent_name, ScanConfig_UserAgent_value)
	proto.RegisterEnum("google.cloud.websecurityscanner.v1beta.ScanConfig_TargetPlatform", ScanConfig_TargetPlatform_name, ScanConfig_TargetPlatform_value)
	proto.RegisterEnum("google.cloud.websecurityscanner.v1beta.ScanConfig_RiskLevel", ScanConfig_RiskLevel_name, ScanConfig_RiskLevel_value)
	proto.RegisterEnum("google.cloud.websecurityscanner.v1beta.ScanConfig_ExportToSecurityCommandCenter", ScanConfig_ExportToSecurityCommandCenter_name, ScanConfig_ExportToSecurityCommandCenter_value)
	proto.RegisterType((*ScanConfig)(nil), "google.cloud.websecurityscanner.v1beta.ScanConfig")
	proto.RegisterType((*ScanConfig_Authentication)(nil), "google.cloud.websecurityscanner.v1beta.ScanConfig.Authentication")
	proto.RegisterType((*ScanConfig_Authentication_GoogleAccount)(nil), "google.cloud.websecurityscanner.v1beta.ScanConfig.Authentication.GoogleAccount")
	proto.RegisterType((*ScanConfig_Authentication_CustomAccount)(nil), "google.cloud.websecurityscanner.v1beta.ScanConfig.Authentication.CustomAccount")
	proto.RegisterType((*ScanConfig_Schedule)(nil), "google.cloud.websecurityscanner.v1beta.ScanConfig.Schedule")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1beta/scan_config.proto", fileDescriptor_0b8e0c37c8759a6e)
}

var fileDescriptor_0b8e0c37c8759a6e = []byte{
	// 1045 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5b, 0x6f, 0x1b, 0x45,
	0x14, 0xee, 0xda, 0xb9, 0xd8, 0xe3, 0x4b, 0xb6, 0x23, 0x54, 0x8c, 0x51, 0x15, 0xe3, 0x87, 0xc8,
	0x42, 0xad, 0x57, 0x0d, 0x42, 0x02, 0x5a, 0x09, 0x6d, 0xec, 0x4d, 0x62, 0x61, 0xaf, 0x97, 0xf1,
	0xba, 0x09, 0x10, 0x69, 0x18, 0xaf, 0x27, 0x9b, 0x25, 0x7b, 0x63, 0x66, 0x36, 0x4d, 0x54, 0xf5,
	0x95, 0x1f, 0xc1, 0xcf, 0x81, 0x37, 0x7e, 0x45, 0x1e, 0x78, 0xe2, 0x57, 0xa0, 0xbd, 0x39, 0x76,
	0x8a, 0x4a, 0x9b, 0xf6, 0x6d, 0xe6, 0x5c, 0xbe, 0xef, 0xcc, 0xf1, 0x39, 0x9f, 0x17, 0x7c, 0x65,
	0x07, 0x81, 0xed, 0x52, 0xc5, 0x72, 0x83, 0x68, 0xae, 0xbc, 0xa0, 0x33, 0x4e, 0xad, 0x88, 0x39,
	0xe2, 0x8a, 0x5b, 0xc4, 0xf7, 0x29, 0x53, 0x2e, 0x9e, 0xcc, 0xa8, 0x20, 0x4a, 0x7c, 0xc5, 0x56,
	0xe0, 0x9f, 0x3a, 0x76, 0x37, 0x64, 0x81, 0x08, 0xe0, 0x4e, 0x9a, 0xd9, 0x4d, 0x32, 0xbb, 0xaf,
	0x67, 0x76, 0xd3, 0xcc, 0xe6, 0x76, 0xc6, 0x40, 0x42, 0x47, 0x39, 0x75, 0xa8, 0x3b, 0xc7, 0x33,
	0x7a, 0x46, 0x2e, 0x9c, 0x80, 0xa5, 0x40, 0xcd, 0x4f, 0x96, 0x02, 0x18, 0xe5, 0x41, 0xc4, 0x2c,
	0x9a, 0xb9, 0xbe, 0x7c, 0x97, 0xea, 0x58, 0xe4, 0x67, 0x69, 0x39, 0x65, 0x72, 0x9b, 0x45, 0xa7,
	0x8a, 0x70, 0x3c, 0xca, 0x05, 0xf1, 0xc2, 0x34, 0xa0, 0xfd, 0xe7, 0x16, 0x00, 0x13, 0x8b, 0xf8,
	0xbd, 0xe4, 0x41, 0x10, 0x82, 0x35, 0x9f, 0x78, 0xb4, 0x21, 0xb5, 0xa4, 0x4e, 0x19, 0x25, 0x67,
	0xb8, 0x03, 0xaa, 0x73, 0x87, 0x87, 0x2e, 0xb9, 0xc2, 0x89, 0xaf, 0x10, 0xfb, 0xf6, 0x8a, 0xd7,
	0x6a, 0x01, 0x55, 0x32, 0x87, 0x1e, 0xc7, 0x7d, 0x0c, 0x36, 0x3d, 0x72, 0x89, 0x7f, 0x0d, 0x79,
	0xa3, 0xd8, 0x92, 0x3a, 0xeb, 0x68, 0xc3, 0x23, 0x97, 0xdf, 0x87, 0x1c, 0x76, 0x40, 0x8d, 0x0b,
	0xc2, 0x84, 0xe3, 0xdb, 0x38, 0x62, 0x2e, 0x6f, 0xac, 0xb5, 0x8a, 0x39, 0x42, 0x35, 0xf7, 0x4c,
	0x99, 0xcb, 0xa1, 0x03, 0xea, 0x24, 0x12, 0x67, 0xd4, 0x17, 0x8e, 0x45, 0x84, 0x13, 0xf8, 0x8d,
	0xf5, 0x96, 0xd4, 0xa9, 0xec, 0xaa, 0xdd, 0xb7, 0x6b, 0x71, 0xf7, 0xe6, 0x29, 0x5d, 0x75, 0x05,
	0x08, 0xdd, 0x02, 0x86, 0x3f, 0x01, 0x10, 0x71, 0xca, 0x30, 0xb1, 0xa9, 0x2f, 0x1a, 0x1b, 0x2d,
	0xa9, 0x53, 0xdf, 0x7d, 0x76, 0x07, 0x9a, 0x29, 0xa7, 0x4c, 0x8d, 0x31, 0x50, 0x39, 0xca, 0x8f,
	0xf0, 0x31, 0x80, 0x33, 0x97, 0x58, 0xe7, 0xae, 0xc3, 0x05, 0x0e, 0x89, 0x10, 0x94, 0xf9, 0xbc,
	0xb1, 0x19, 0x3f, 0x1b, 0xdd, 0x5f, 0x78, 0x8c, 0xcc, 0x01, 0x8f, 0x40, 0x89, 0x5b, 0x67, 0x74,
	0x1e, 0xb9, 0xb4, 0x51, 0x4a, 0x1e, 0xfc, 0xf4, 0x0e, 0x95, 0x4c, 0x32, 0x08, 0xb4, 0x00, 0x83,
	0x2e, 0x90, 0x05, 0x61, 0x36, 0x15, 0x38, 0x74, 0x89, 0x38, 0x0d, 0x98, 0xc7, 0x1b, 0xe5, 0x56,
	0xb1, 0x53, 0xbf, 0x53, 0x47, 0xcd, 0x04, 0xca, 0xc8, 0x90, 0xd0, 0x96, 0x58, 0xb9, 0x73, 0xf8,
	0xbb, 0x04, 0x3e, 0xa3, 0x97, 0x61, 0xc0, 0x04, 0x16, 0x01, 0xce, 0xf1, 0xb0, 0x15, 0x78, 0x1e,
	0xf1, 0xe7, 0xd8, 0xa2, 0xbe, 0xa0, 0xac, 0x01, 0x92, 0x56, 0x1b, 0x77, 0xe0, 0xd7, 0x12, 0x6c,
	0x33, 0x98, 0x64, 0x91, 0xbd, 0x14, 0xb8, 0x97, 0xe0, 0xa2, 0x87, 0xf4, 0x4d, 0x6e, 0xa8, 0x03,
	0xe0, 0x12, 0x41, 0xb9, 0x88, 0xb7, 0xa3, 0x51, 0x49, 0xba, 0xac, 0xbc, 0x4b, 0x11, 0x28, 0xf2,
	0x51, 0x39, 0x85, 0x40, 0x51, 0x32, 0x3f, 0xcc, 0xe1, 0xe7, 0xd8, 0xa5, 0x17, 0xd4, 0x6d, 0x54,
	0xef, 0x3c, 0x3f, 0xc8, 0xe1, 0xe7, 0xc3, 0x18, 0x03, 0x95, 0x59, 0x7e, 0x6c, 0xfe, 0x5d, 0x04,
	0xf5, 0xd5, 0xf9, 0x85, 0x97, 0xa0, 0x9e, 0x82, 0x63, 0x62, 0x59, 0x41, 0xe4, 0x8b, 0x64, 0x47,
	0x2b, 0xbb, 0xe3, 0xf7, 0x5e, 0x8d, 0xee, 0x41, 0x02, 0xa0, 0xa6, 0xb0, 0x87, 0xf7, 0x50, 0xcd,
	0x5e, 0x36, 0xc4, 0xcc, 0x56, 0xc4, 0x45, 0xe0, 0x2d, 0x98, 0x0b, 0x1f, 0x8a, 0xb9, 0x97, 0xe0,
	0x2e, 0x31, 0x5b, 0xcb, 0x86, 0xa6, 0x09, 0x6a, 0x2b, 0xb5, 0xc1, 0x6d, 0x50, 0x8a, 0x97, 0xec,
	0x46, 0xa2, 0x52, 0x11, 0x59, 0x18, 0x61, 0x1b, 0x94, 0x42, 0xc2, 0xf9, 0x8b, 0x80, 0xcd, 0x33,
	0x9d, 0xda, 0xb8, 0x56, 0x0b, 0xd7, 0xea, 0x1a, 0x5a, 0xd8, 0x9b, 0x17, 0xa0, 0xb6, 0xc2, 0xfb,
	0x41, 0x50, 0x61, 0x0b, 0x94, 0xdd, 0xc0, 0x76, 0xfc, 0x58, 0xe1, 0x12, 0xfd, 0xcb, 0x51, 0x12,
	0xeb, 0x94, 0xb9, 0x7b, 0xf2, 0x6d, 0x71, 0x6b, 0xfe, 0x26, 0x81, 0x52, 0xbe, 0xb5, 0xf0, 0x5b,
	0x50, 0xcb, 0xf7, 0x16, 0xc7, 0x2a, 0x9d, 0xfd, 0xbe, 0xcd, 0xbc, 0xcb, 0xb9, 0x84, 0x77, 0xcd,
	0x5c, 0xc2, 0x51, 0x35, 0x4f, 0x88, 0x4d, 0xf0, 0x6b, 0xf0, 0xc0, 0x89, 0x47, 0xfd, 0x82, 0xb8,
	0x78, 0x1e, 0xb1, 0x84, 0x02, 0xcf, 0xc9, 0x15, 0x4f, 0x6a, 0x5e, 0x4f, 0xcb, 0xf9, 0x28, 0x0f,
	0xe9, 0x67, 0x11, 0x7d, 0x72, 0xc5, 0xdb, 0x3f, 0x83, 0xf2, 0x42, 0xc7, 0x60, 0x13, 0x3c, 0x98,
	0x4e, 0x34, 0x84, 0xd5, 0x03, 0x4d, 0x37, 0xf1, 0x54, 0x9f, 0x18, 0x5a, 0x6f, 0xb0, 0x3f, 0xd0,
	0xfa, 0xf2, 0x3d, 0x28, 0x83, 0x6a, 0xef, 0x10, 0x8d, 0x47, 0x1a, 0x1e, 0x0e, 0xf4, 0xe9, 0xb1,
	0x2c, 0x41, 0x08, 0xea, 0x99, 0x45, 0xd5, 0xfb, 0x68, 0x3c, 0xe8, 0xcb, 0x05, 0x78, 0x1f, 0xd4,
	0x26, 0xea, 0xbe, 0x8a, 0x06, 0x78, 0x60, 0x1c, 0x8e, 0x75, 0x4d, 0x2e, 0xb6, 0x75, 0x50, 0x5f,
	0x95, 0x0f, 0xb8, 0x0d, 0x3e, 0x35, 0x55, 0x74, 0xa0, 0x99, 0xd8, 0x18, 0xaa, 0xe6, 0xfe, 0x18,
	0x8d, 0x6e, 0x71, 0xd5, 0x01, 0x50, 0x0d, 0x03, 0x6b, 0xfa, 0xc1, 0x40, 0xd7, 0x64, 0x09, 0x56,
	0xc0, 0x66, 0x6f, 0x3c, 0x32, 0xa6, 0xa6, 0x26, 0x17, 0xda, 0xcf, 0x40, 0x79, 0xb1, 0x39, 0x71,
	0xc5, 0x68, 0x30, 0xf9, 0x0e, 0x0f, 0xb5, 0xe7, 0xda, 0xf0, 0x16, 0x0a, 0x00, 0x1b, 0xfa, 0x18,
	0x8d, 0xd4, 0xa1, 0x2c, 0xc1, 0x4d, 0x50, 0x1c, 0x8e, 0x8f, 0xe4, 0x42, 0xdb, 0x03, 0x0f, 0xdf,
	0x28, 0x26, 0xf0, 0x09, 0x78, 0xac, 0x1d, 0x1b, 0x63, 0x64, 0x62, 0x73, 0x8c, 0x27, 0x5a, 0x6f,
	0x8a, 0x06, 0xe6, 0x0f, 0xb8, 0x37, 0x1e, 0x8d, 0x54, 0xbd, 0x8f, 0x7b, 0x9a, 0x6e, 0x6a, 0xe8,
	0x16, 0x51, 0x05, 0x6c, 0x6a, 0xba, 0xba, 0x37, 0xd4, 0xfa, 0xb2, 0x04, 0xab, 0xa0, 0xd4, 0x1f,
	0x4c, 0xd2, 0x5b, 0xe1, 0x1b, 0xfc, 0x8f, 0x7a, 0x02, 0x1e, 0xfd, 0xc7, 0x82, 0xa4, 0x3f, 0x2d,
	0x09, 0x1d, 0xde, 0xb5, 0x02, 0x4f, 0x59, 0xfa, 0x23, 0x7e, 0x14, 0xb2, 0xe0, 0x17, 0x6a, 0x09,
	0xae, 0xbc, 0xcc, 0x4e, 0xaf, 0x92, 0xff, 0xf6, 0xd4, 0xcd, 0x95, 0x97, 0x4b, 0x9f, 0x21, 0xaf,
	0xf6, 0xfe, 0x90, 0xc0, 0xe7, 0x56, 0xe0, 0xbd, 0xe5, 0x42, 0xee, 0x6d, 0xdd, 0x10, 0x19, 0xf1,
	0x54, 0x19, 0xd2, 0x8f, 0xc7, 0x59, 0xaa, 0x1d, 0xb8, 0xc4, 0xb7, 0xbb, 0x01, 0xb3, 0x15, 0x9b,
	0xfa, 0xc9, 0xcc, 0x29, 0x37, 0x55, 0xfe, 0xdf, 0xe7, 0xc7, 0xd3, 0xd7, 0x3d, 0x7f, 0x15, 0x76,
	0xd2, 0x1d, 0x3e, 0xe9, 0xc5, 0xb9, 0x27, 0x47, 0x74, 0x96, 0xb7, 0x7c, 0x92, 0x46, 0x9c, 0x3c,
	0x4f, 0x72, 0x67, 0x1b, 0x09, 0xdb, 0x17, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xb3, 0x1f,
	0xde, 0x89, 0x09, 0x00, 0x00,
}
