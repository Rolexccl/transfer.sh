// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A customer.
type Customer struct {
	// The resource name of the customer.
	// Customer resource names have the form:
	//
	// `customers/{customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the customer.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Optional, non-unique descriptive name of the customer.
	DescriptiveName *wrappers.StringValue `protobuf:"bytes,4,opt,name=descriptive_name,json=descriptiveName,proto3" json:"descriptive_name,omitempty"`
	// The currency in which the account operates.
	// A subset of the currency codes from the ISO 4217 standard is
	// supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The local timezone ID of the customer.
	TimeZone *wrappers.StringValue `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// The URL template for constructing a tracking URL out of parameters.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,7,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The URL template for appending params to the final URL
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,11,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// Whether auto-tagging is enabled for the customer.
	AutoTaggingEnabled *wrappers.BoolValue `protobuf:"bytes,8,opt,name=auto_tagging_enabled,json=autoTaggingEnabled,proto3" json:"auto_tagging_enabled,omitempty"`
	// Whether the Customer has a Partners program badge. If the Customer is not
	// associated with the Partners program, this will be false. For more
	// information, see https://support.google.com/partners/answer/3125774.
	HasPartnersBadge *wrappers.BoolValue `protobuf:"bytes,9,opt,name=has_partners_badge,json=hasPartnersBadge,proto3" json:"has_partners_badge,omitempty"`
	// Whether the customer is a manager.
	Manager *wrappers.BoolValue `protobuf:"bytes,12,opt,name=manager,proto3" json:"manager,omitempty"`
	// Whether the customer is a test account.
	TestAccount *wrappers.BoolValue `protobuf:"bytes,13,opt,name=test_account,json=testAccount,proto3" json:"test_account,omitempty"`
	// Call reporting setting for a customer.
	CallReportingSetting *CallReportingSetting `protobuf:"bytes,10,opt,name=call_reporting_setting,json=callReportingSetting,proto3" json:"call_reporting_setting,omitempty"`
	// Conversion tracking setting for a customer.
	ConversionTrackingSetting *ConversionTrackingSetting `protobuf:"bytes,14,opt,name=conversion_tracking_setting,json=conversionTrackingSetting,proto3" json:"conversion_tracking_setting,omitempty"`
	// Remarketing setting for a customer.
	RemarketingSetting   *RemarketingSetting `protobuf:"bytes,15,opt,name=remarketing_setting,json=remarketingSetting,proto3" json:"remarketing_setting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_67cce403a0f5fba7, []int{0}
}
func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (dst *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(dst, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Customer) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Customer) GetDescriptiveName() *wrappers.StringValue {
	if m != nil {
		return m.DescriptiveName
	}
	return nil
}

func (m *Customer) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *Customer) GetTimeZone() *wrappers.StringValue {
	if m != nil {
		return m.TimeZone
	}
	return nil
}

func (m *Customer) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Customer) GetFinalUrlSuffix() *wrappers.StringValue {
	if m != nil {
		return m.FinalUrlSuffix
	}
	return nil
}

func (m *Customer) GetAutoTaggingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.AutoTaggingEnabled
	}
	return nil
}

func (m *Customer) GetHasPartnersBadge() *wrappers.BoolValue {
	if m != nil {
		return m.HasPartnersBadge
	}
	return nil
}

func (m *Customer) GetManager() *wrappers.BoolValue {
	if m != nil {
		return m.Manager
	}
	return nil
}

func (m *Customer) GetTestAccount() *wrappers.BoolValue {
	if m != nil {
		return m.TestAccount
	}
	return nil
}

func (m *Customer) GetCallReportingSetting() *CallReportingSetting {
	if m != nil {
		return m.CallReportingSetting
	}
	return nil
}

func (m *Customer) GetConversionTrackingSetting() *ConversionTrackingSetting {
	if m != nil {
		return m.ConversionTrackingSetting
	}
	return nil
}

func (m *Customer) GetRemarketingSetting() *RemarketingSetting {
	if m != nil {
		return m.RemarketingSetting
	}
	return nil
}

// Call reporting setting for a customer.
type CallReportingSetting struct {
	// Enable reporting of phone call events by redirecting them via Google
	// System.
	CallReportingEnabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=call_reporting_enabled,json=callReportingEnabled,proto3" json:"call_reporting_enabled,omitempty"`
	// Whether to enable call conversion reporting.
	CallConversionReportingEnabled *wrappers.BoolValue `protobuf:"bytes,2,opt,name=call_conversion_reporting_enabled,json=callConversionReportingEnabled,proto3" json:"call_conversion_reporting_enabled,omitempty"`
	// Customer-level call conversion action to attribute a call conversion to.
	// If not set a default conversion action is used. Only in effect when
	// call_conversion_reporting_enabled is set to true.
	CallConversionAction *wrappers.StringValue `protobuf:"bytes,9,opt,name=call_conversion_action,json=callConversionAction,proto3" json:"call_conversion_action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CallReportingSetting) Reset()         { *m = CallReportingSetting{} }
func (m *CallReportingSetting) String() string { return proto.CompactTextString(m) }
func (*CallReportingSetting) ProtoMessage()    {}
func (*CallReportingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_67cce403a0f5fba7, []int{1}
}
func (m *CallReportingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallReportingSetting.Unmarshal(m, b)
}
func (m *CallReportingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallReportingSetting.Marshal(b, m, deterministic)
}
func (dst *CallReportingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallReportingSetting.Merge(dst, src)
}
func (m *CallReportingSetting) XXX_Size() int {
	return xxx_messageInfo_CallReportingSetting.Size(m)
}
func (m *CallReportingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CallReportingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CallReportingSetting proto.InternalMessageInfo

func (m *CallReportingSetting) GetCallReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallConversionReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.CallConversionAction
	}
	return nil
}

// A collection of customer-wide settings related to Google Ads Conversion
// Tracking.
type ConversionTrackingSetting struct {
	// The conversion tracking id used for this account. This id is automatically
	// assigned after any conversion tracking feature is used. If the customer
	// doesn't use conversion tracking, this is 0. This field is read-only.
	ConversionTrackingId *wrappers.Int64Value `protobuf:"bytes,1,opt,name=conversion_tracking_id,json=conversionTrackingId,proto3" json:"conversion_tracking_id,omitempty"`
	// The conversion tracking id of the customer's manager. This is set when the
	// customer is opted into cross account conversion tracking, and it overrides
	// conversion_tracking_id. This field can only be managed through the Google
	// Ads UI. This field is read-only.
	CrossAccountConversionTrackingId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cross_account_conversion_tracking_id,json=crossAccountConversionTrackingId,proto3" json:"cross_account_conversion_tracking_id,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}             `json:"-"`
	XXX_unrecognized                 []byte               `json:"-"`
	XXX_sizecache                    int32                `json:"-"`
}

func (m *ConversionTrackingSetting) Reset()         { *m = ConversionTrackingSetting{} }
func (m *ConversionTrackingSetting) String() string { return proto.CompactTextString(m) }
func (*ConversionTrackingSetting) ProtoMessage()    {}
func (*ConversionTrackingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_67cce403a0f5fba7, []int{2}
}
func (m *ConversionTrackingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionTrackingSetting.Unmarshal(m, b)
}
func (m *ConversionTrackingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionTrackingSetting.Marshal(b, m, deterministic)
}
func (dst *ConversionTrackingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionTrackingSetting.Merge(dst, src)
}
func (m *ConversionTrackingSetting) XXX_Size() int {
	return xxx_messageInfo_ConversionTrackingSetting.Size(m)
}
func (m *ConversionTrackingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionTrackingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionTrackingSetting proto.InternalMessageInfo

func (m *ConversionTrackingSetting) GetConversionTrackingId() *wrappers.Int64Value {
	if m != nil {
		return m.ConversionTrackingId
	}
	return nil
}

func (m *ConversionTrackingSetting) GetCrossAccountConversionTrackingId() *wrappers.Int64Value {
	if m != nil {
		return m.CrossAccountConversionTrackingId
	}
	return nil
}

// Remarketing setting for a customer.
type RemarketingSetting struct {
	// The Google global site tag.
	GoogleGlobalSiteTag  *wrappers.StringValue `protobuf:"bytes,1,opt,name=google_global_site_tag,json=googleGlobalSiteTag,proto3" json:"google_global_site_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RemarketingSetting) Reset()         { *m = RemarketingSetting{} }
func (m *RemarketingSetting) String() string { return proto.CompactTextString(m) }
func (*RemarketingSetting) ProtoMessage()    {}
func (*RemarketingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_67cce403a0f5fba7, []int{3}
}
func (m *RemarketingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemarketingSetting.Unmarshal(m, b)
}
func (m *RemarketingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemarketingSetting.Marshal(b, m, deterministic)
}
func (dst *RemarketingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemarketingSetting.Merge(dst, src)
}
func (m *RemarketingSetting) XXX_Size() int {
	return xxx_messageInfo_RemarketingSetting.Size(m)
}
func (m *RemarketingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_RemarketingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_RemarketingSetting proto.InternalMessageInfo

func (m *RemarketingSetting) GetGoogleGlobalSiteTag() *wrappers.StringValue {
	if m != nil {
		return m.GoogleGlobalSiteTag
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "google.ads.googleads.v1.resources.Customer")
	proto.RegisterType((*CallReportingSetting)(nil), "google.ads.googleads.v1.resources.CallReportingSetting")
	proto.RegisterType((*ConversionTrackingSetting)(nil), "google.ads.googleads.v1.resources.ConversionTrackingSetting")
	proto.RegisterType((*RemarketingSetting)(nil), "google.ads.googleads.v1.resources.RemarketingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer.proto", fileDescriptor_customer_67cce403a0f5fba7)
}

var fileDescriptor_customer_67cce403a0f5fba7 = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x6e, 0xdb, 0x36,
	0x14, 0x87, 0x61, 0x75, 0x6b, 0x13, 0x26, 0x69, 0x03, 0xd6, 0x2b, 0xd4, 0xb4, 0x28, 0x92, 0x6c,
	0x05, 0x0a, 0x0c, 0x90, 0xe7, 0xad, 0xdb, 0x30, 0x6d, 0xbb, 0x50, 0x82, 0xcd, 0xeb, 0x30, 0x0c,
	0x9e, 0xed, 0xfa, 0xa2, 0x30, 0x20, 0xd0, 0xd4, 0xb1, 0x4a, 0x84, 0x22, 0x05, 0x92, 0xf2, 0xfe,
	0x60, 0xef, 0x32, 0x60, 0xbb, 0xdb, 0xa3, 0xec, 0x15, 0xf6, 0x06, 0x7b, 0x84, 0x5d, 0x15, 0xa2,
	0x44, 0x35, 0x81, 0x9c, 0xc8, 0x57, 0x26, 0xc4, 0xf3, 0x7d, 0x3f, 0x8a, 0x3e, 0xa4, 0xd0, 0x47,
	0xa9, 0x94, 0x29, 0x87, 0x01, 0x49, 0xf4, 0xa0, 0x1a, 0x96, 0xa3, 0xf5, 0x70, 0xa0, 0x40, 0xcb,
	0x42, 0x51, 0xd0, 0x03, 0x5a, 0x68, 0x23, 0x33, 0x50, 0x41, 0xae, 0xa4, 0x91, 0xf8, 0xa4, 0x2a,
	0x0b, 0x48, 0xa2, 0x83, 0x86, 0x08, 0xd6, 0xc3, 0xa0, 0x21, 0x8e, 0x9e, 0xd4, 0x52, 0x0b, 0x2c,
	0x8b, 0xd5, 0xe0, 0x67, 0x45, 0xf2, 0x1c, 0x94, 0xae, 0x14, 0x47, 0x8f, 0x5d, 0x68, 0xce, 0x06,
	0x44, 0x08, 0x69, 0x88, 0x61, 0x52, 0xd4, 0xb3, 0xa7, 0x7f, 0xec, 0xa0, 0x9d, 0xf3, 0x3a, 0x13,
	0xbf, 0x8f, 0x0e, 0x9c, 0x37, 0x16, 0x24, 0x03, 0xbf, 0x77, 0xdc, 0x7b, 0xb6, 0x3b, 0xd9, 0x77,
	0x0f, 0x7f, 0x24, 0x19, 0xe0, 0x0f, 0x91, 0xc7, 0x12, 0xff, 0xd6, 0x71, 0xef, 0xd9, 0xde, 0xc7,
	0x8f, 0xea, 0x45, 0x05, 0x2e, 0x3c, 0x78, 0x21, 0xcc, 0x67, 0xcf, 0xe7, 0x84, 0x17, 0x30, 0xf1,
	0x58, 0x82, 0x47, 0xe8, 0x30, 0x01, 0x4d, 0x15, 0xcb, 0x0d, 0x5b, 0xd7, 0xd2, 0x77, 0x2c, 0xfa,
	0xb8, 0x85, 0x4e, 0x8d, 0x62, 0x22, 0xad, 0xd8, 0x7b, 0x97, 0x28, 0x9b, 0x1a, 0xa1, 0x03, 0x5a,
	0x28, 0x05, 0x82, 0xfe, 0x1a, 0x53, 0x99, 0x80, 0xff, 0xee, 0x16, 0x96, 0x7d, 0x87, 0x9c, 0xcb,
	0x04, 0xf0, 0x17, 0x68, 0xd7, 0xb0, 0x0c, 0xe2, 0xdf, 0xa4, 0x00, 0xff, 0xf6, 0x16, 0xf8, 0x4e,
	0x59, 0xfe, 0x4a, 0x0a, 0xc0, 0x63, 0xf4, 0x9e, 0x51, 0x84, 0x5e, 0x30, 0x91, 0xc6, 0x85, 0xe2,
	0xb1, 0x81, 0x2c, 0xe7, 0xc4, 0x80, 0x7f, 0x67, 0x0b, 0xcd, 0x7d, 0x87, 0xbe, 0x54, 0x7c, 0x56,
	0x83, 0xf8, 0x5b, 0x74, 0xb8, 0x62, 0x82, 0x70, 0xab, 0xd3, 0xc5, 0x6a, 0xc5, 0x7e, 0xf1, 0xf7,
	0xb6, 0x90, 0xdd, 0xb5, 0xd4, 0x4b, 0xc5, 0xa7, 0x96, 0xc1, 0x3f, 0xa0, 0x3e, 0x29, 0x8c, 0x8c,
	0x0d, 0x49, 0xd3, 0x72, 0x75, 0x20, 0xc8, 0x92, 0x43, 0xe2, 0xef, 0x58, 0xd7, 0x51, 0xcb, 0x75,
	0x26, 0x25, 0xaf, 0x4c, 0xb8, 0xe4, 0x66, 0x15, 0xf6, 0x4d, 0x45, 0xe1, 0xef, 0x10, 0x7e, 0x4d,
	0x74, 0x9c, 0x13, 0x65, 0x04, 0x28, 0x1d, 0x2f, 0x49, 0x92, 0x82, 0xbf, 0xdb, 0xe9, 0x3a, 0x7c,
	0x4d, 0xf4, 0xb8, 0x86, 0xce, 0x4a, 0x06, 0x3f, 0x47, 0x77, 0x32, 0x22, 0x48, 0x0a, 0xca, 0xdf,
	0xef, 0xc4, 0x5d, 0x29, 0xfe, 0x1a, 0xed, 0x1b, 0xd0, 0x26, 0x26, 0x94, 0xca, 0x42, 0x18, 0xff,
	0xa0, 0x13, 0xdd, 0x2b, 0xeb, 0xa3, 0xaa, 0x1c, 0x67, 0xe8, 0x01, 0x25, 0x9c, 0xc7, 0x0a, 0x72,
	0xa9, 0x4c, 0xb9, 0x1d, 0x1a, 0x4c, 0xf9, 0xeb, 0x23, 0x2b, 0xfa, 0x3c, 0xe8, 0x3c, 0x4e, 0xc1,
	0x39, 0xe1, 0x7c, 0xe2, 0xf8, 0x69, 0x85, 0x4f, 0xfa, 0x74, 0xc3, 0x53, 0xfc, 0x3b, 0x7a, 0x44,
	0xa5, 0x58, 0x83, 0xd2, 0x4c, 0x8a, 0xb8, 0x69, 0x10, 0x97, 0x79, 0xd7, 0x66, 0x7e, 0xb5, 0x4d,
	0x66, 0x63, 0x99, 0xd5, 0x12, 0x17, 0xfc, 0x90, 0x5e, 0x37, 0x85, 0x57, 0xe8, 0xbe, 0x82, 0x8c,
	0xa8, 0x0b, 0xb8, 0xf2, 0xa6, 0xf7, 0x6c, 0xea, 0xa7, 0x5b, 0xa4, 0x4e, 0xde, 0xd2, 0x2e, 0x0e,
	0xab, 0xd6, 0xb3, 0xd3, 0xbf, 0x3c, 0xd4, 0xdf, 0xb4, 0x29, 0x78, 0xdc, 0xda, 0x6d, 0xd7, 0x7c,
	0xbd, 0xce, 0xbf, 0xed, 0xea, 0x86, 0xba, 0xf6, 0x03, 0x74, 0x62, 0x8d, 0x97, 0x76, 0xb5, 0x2d,
	0xf7, 0x3a, 0xe5, 0x4f, 0x4a, 0xc9, 0xdb, 0x3d, 0x6d, 0xc5, 0x4c, 0xea, 0x85, 0x5f, 0x8a, 0x21,
	0xb4, 0xbc, 0x14, 0xeb, 0x4e, 0xbf, 0xf9, 0x04, 0xf6, 0xaf, 0xda, 0x23, 0x4b, 0x9e, 0xfe, 0xdb,
	0x43, 0x0f, 0xaf, 0xfd, 0x1b, 0xf1, 0x4f, 0xe8, 0xc1, 0xa6, 0x4e, 0x61, 0x6e, 0xab, 0x6e, 0xbc,
	0x47, 0xfb, 0xed, 0x1e, 0x78, 0x91, 0xe0, 0x0b, 0xf4, 0x01, 0x55, 0x52, 0x6b, 0x77, 0x56, 0xe2,
	0x6b, 0x02, 0xbc, 0xee, 0x80, 0x63, 0x2b, 0xaa, 0x0f, 0xd1, 0xf9, 0x86, 0xb0, 0xd3, 0x14, 0xe1,
	0x76, 0xb7, 0x94, 0x6f, 0x55, 0x59, 0xe3, 0x94, 0xcb, 0x25, 0xe1, 0xb1, 0x66, 0x06, 0xca, 0x9b,
	0xa8, 0x7e, 0xab, 0x8e, 0x6b, 0xb1, 0x9a, 0x1c, 0x59, 0x74, 0xca, 0x0c, 0xcc, 0x48, 0x7a, 0xf6,
	0x7f, 0x0f, 0x3d, 0xa5, 0x32, 0xeb, 0xee, 0xde, 0xb3, 0x03, 0xf7, 0xd5, 0x1a, 0x97, 0xf2, 0x71,
	0xef, 0xd5, 0xf7, 0x35, 0x93, 0x4a, 0x4e, 0x44, 0x1a, 0x48, 0x95, 0x0e, 0x52, 0x10, 0x36, 0xda,
	0x7d, 0x6c, 0x73, 0xa6, 0x6f, 0xf8, 0xf6, 0x7e, 0xd9, 0x8c, 0xfe, 0xf4, 0x6e, 0x8d, 0xa2, 0xe8,
	0x6f, 0xef, 0x64, 0x54, 0x29, 0xa3, 0x44, 0x07, 0xd5, 0xb0, 0x1c, 0xcd, 0x87, 0xc1, 0xc4, 0x55,
	0xfe, 0xe3, 0x6a, 0x16, 0x51, 0xa2, 0x17, 0x4d, 0xcd, 0x62, 0x3e, 0x5c, 0x34, 0x35, 0xff, 0x79,
	0x4f, 0xab, 0x89, 0x30, 0x8c, 0x12, 0x1d, 0x86, 0x4d, 0x55, 0x18, 0xce, 0x87, 0x61, 0xd8, 0xd4,
	0x2d, 0x6f, 0xdb, 0xc5, 0x7e, 0xf2, 0x26, 0x00, 0x00, 0xff, 0xff, 0x13, 0x64, 0x5c, 0x10, 0x27,
	0x08, 0x00, 0x00,
}