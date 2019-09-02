// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/cluster/outlier_detection.proto

package cluster

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// See the :ref:`architecture overview <arch_overview_outlier_detection>` for
// more information on outlier detection.
type OutlierDetection struct {
	// The number of consecutive 5xx responses or local origin errors that are mapped
	// to 5xx error codes before a consecutive 5xx ejection
	// occurs. Defaults to 5.
	Consecutive_5Xx *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx,proto3" json:"consecutive_5xx,omitempty"`
	// The time interval between ejection analysis sweeps. This can result in
	// both new ejections as well as hosts being returned to service. Defaults
	// to 10000ms or 10s.
	Interval *duration.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// The base time that a host is ejected for. The real time is equal to the
	// base time multiplied by the number of times the host has been ejected.
	// Defaults to 30000ms or 30s.
	BaseEjectionTime *duration.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime,proto3" json:"base_ejection_time,omitempty"`
	// The maximum % of an upstream cluster that can be ejected due to outlier
	// detection. Defaults to 10% but will eject at least one host regardless of the value.
	MaxEjectionPercent *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent,proto3" json:"max_ejection_percent,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive 5xx. This setting can be used to disable
	// ejection or to ramp it up slowly. Defaults to 100.
	EnforcingConsecutive_5Xx *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx,proto3" json:"enforcing_consecutive_5xx,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics. This setting can be used to
	// disable ejection or to ramp it up slowly. Defaults to 100.
	EnforcingSuccessRate *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate,proto3" json:"enforcing_success_rate,omitempty"`
	// The number of hosts in a cluster that must have enough request volume to
	// detect success rate outliers. If the number of hosts is less than this
	// setting, outlier detection via success rate statistics is not performed
	// for any host in the cluster. Defaults to 5.
	SuccessRateMinimumHosts *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts,proto3" json:"success_rate_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one
	// interval (as defined by the interval duration above) to include this host
	// in success rate based outlier detection. If the volume is lower than this
	// setting, outlier detection via success rate statistics is not performed
	// for that host. Defaults to 100.
	SuccessRateRequestVolume *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume,proto3" json:"success_rate_request_volume,omitempty"`
	// This factor is used to determine the ejection threshold for success rate
	// outlier ejection. The ejection threshold is the difference between the
	// mean success rate, and the product of this factor and the standard
	// deviation of the mean success rate: mean - (stdev *
	// success_rate_stdev_factor). This factor is divided by a thousand to get a
	// double. That is, if the desired factor is 1.9, the runtime value should
	// be 1900. Defaults to 1900.
	SuccessRateStdevFactor *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor,proto3" json:"success_rate_stdev_factor,omitempty"`
	// The number of consecutive gateway failures (502, 503, 504 status codes)
	// before a consecutive gateway failure ejection occurs. Defaults to 5.
	ConsecutiveGatewayFailure *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure,proto3" json:"consecutive_gateway_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive gateway failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 0.
	EnforcingConsecutiveGatewayFailure *wrappers.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure,proto3" json:"enforcing_consecutive_gateway_failure,omitempty"`
	// Determines whether to distinguish local origin failures from external errors. If set to true
	// the following configuration parameters are taken into account:
	// :ref:`consecutive_local_origin_failure<envoy_api_field_cluster.OutlierDetection.consecutive_local_origin_failure>`,
	// :ref:`enforcing_consecutive_local_origin_failure<envoy_api_field_cluster.OutlierDetection.enforcing_consecutive_local_origin_failure>`
	// and
	// :ref:`enforcing_local_origin_success_rate<envoy_api_field_cluster.OutlierDetection.enforcing_local_origin_success_rate>`.
	// Defaults to false.
	SplitExternalLocalOriginErrors bool `protobuf:"varint,12,opt,name=split_external_local_origin_errors,json=splitExternalLocalOriginErrors,proto3" json:"split_external_local_origin_errors,omitempty"`
	// The number of consecutive locally originated failures before ejection
	// occurs. Defaults to 5. Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	ConsecutiveLocalOriginFailure *wrappers.UInt32Value `protobuf:"bytes,13,opt,name=consecutive_local_origin_failure,json=consecutiveLocalOriginFailure,proto3" json:"consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive locally originated failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingConsecutiveLocalOriginFailure *wrappers.UInt32Value `protobuf:"bytes,14,opt,name=enforcing_consecutive_local_origin_failure,json=enforcingConsecutiveLocalOriginFailure,proto3" json:"enforcing_consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics for locally originated errors.
	// This setting can be used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingLocalOriginSuccessRate *wrappers.UInt32Value `protobuf:"bytes,15,opt,name=enforcing_local_origin_success_rate,json=enforcingLocalOriginSuccessRate,proto3" json:"enforcing_local_origin_success_rate,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}              `json:"-"`
	XXX_unrecognized                []byte                `json:"-"`
	XXX_sizecache                   int32                 `json:"-"`
}

func (m *OutlierDetection) Reset()         { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()    {}
func (*OutlierDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_56cd87362a3f00c9, []int{0}
}

func (m *OutlierDetection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierDetection.Unmarshal(m, b)
}
func (m *OutlierDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierDetection.Marshal(b, m, deterministic)
}
func (m *OutlierDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierDetection.Merge(m, src)
}
func (m *OutlierDetection) XXX_Size() int {
	return xxx_messageInfo_OutlierDetection.Size(m)
}
func (m *OutlierDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierDetection.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierDetection proto.InternalMessageInfo

func (m *OutlierDetection) GetConsecutive_5Xx() *wrappers.UInt32Value {
	if m != nil {
		return m.Consecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection) GetBaseEjectionTime() *duration.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection) GetMaxEjectionPercent() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutive_5Xx() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingSuccessRate() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateMinimumHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateRequestVolume() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateRequestVolume
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateStdevFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateStdevFactor
	}
	return nil
}

func (m *OutlierDetection) GetConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.ConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetSplitExternalLocalOriginErrors() bool {
	if m != nil {
		return m.SplitExternalLocalOriginErrors
	}
	return false
}

func (m *OutlierDetection) GetConsecutiveLocalOriginFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.ConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveLocalOriginFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingLocalOriginSuccessRate() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingLocalOriginSuccessRate
	}
	return nil
}

func init() {
	proto.RegisterType((*OutlierDetection)(nil), "envoy.api.v2.cluster.OutlierDetection")
}

func init() {
	proto.RegisterFile("envoy/api/v2/cluster/outlier_detection.proto", fileDescriptor_56cd87362a3f00c9)
}

var fileDescriptor_56cd87362a3f00c9 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdd, 0x6e, 0xd3, 0x3c,
	0x1c, 0xc6, 0xdf, 0x74, 0x1f, 0xef, 0xe6, 0xc1, 0x36, 0x59, 0x65, 0x4b, 0x37, 0x18, 0x53, 0x11,
	0x68, 0x9a, 0x50, 0x22, 0x75, 0xda, 0x05, 0xac, 0x5b, 0xc7, 0x87, 0x80, 0x4d, 0x2d, 0x0c, 0x21,
	0x40, 0x96, 0x97, 0xfe, 0x1b, 0x8c, 0x92, 0x38, 0xd8, 0x4e, 0x96, 0x71, 0xca, 0xcd, 0xa0, 0x71,
	0x07, 0x1c, 0x71, 0x2d, 0x9c, 0xc1, 0x55, 0xa0, 0x38, 0xfd, 0x48, 0xbb, 0x20, 0x9a, 0x33, 0xab,
	0x7e, 0x9e, 0xdf, 0xf3, 0xf8, 0x1f, 0xa7, 0x41, 0x0f, 0x21, 0x88, 0xf9, 0xa5, 0x4d, 0x43, 0x66,
	0xc7, 0x0d, 0xdb, 0xf1, 0x22, 0xa9, 0x40, 0xd8, 0x3c, 0x52, 0x1e, 0x03, 0x41, 0xba, 0xa0, 0xc0,
	0x51, 0x8c, 0x07, 0x56, 0x28, 0xb8, 0xe2, 0xb8, 0xaa, 0xd5, 0x16, 0x0d, 0x99, 0x15, 0x37, 0xac,
	0xbe, 0x7a, 0x63, 0xcb, 0xe5, 0xdc, 0xf5, 0xc0, 0xd6, 0x9a, 0xf3, 0xa8, 0x67, 0x77, 0x23, 0x41,
	0x47, 0xae, 0xeb, 0xfb, 0x17, 0x82, 0x86, 0x21, 0x08, 0xd9, 0xdf, 0x5f, 0x8f, 0xa9, 0xc7, 0xba,
	0x54, 0x81, 0x3d, 0x58, 0xf4, 0x37, 0xaa, 0x2e, 0x77, 0xb9, 0x5e, 0xda, 0xe9, 0x2a, 0xfb, 0xb5,
	0xfe, 0x6d, 0x09, 0xad, 0x9e, 0x64, 0x05, 0x8f, 0x06, 0xfd, 0x70, 0x0b, 0xad, 0x38, 0x3c, 0x90,
	0xe0, 0x44, 0x8a, 0xc5, 0x40, 0xf6, 0x93, 0xc4, 0x34, 0xb6, 0x8d, 0x9d, 0xa5, 0xc6, 0x6d, 0x2b,
	0x4b, 0xb7, 0x06, 0xe9, 0xd6, 0xab, 0x27, 0x81, 0xda, 0x6b, 0x9c, 0x51, 0x2f, 0x82, 0xf6, 0x72,
	0xce, 0xb4, 0x9f, 0x24, 0xf8, 0x00, 0x2d, 0xb0, 0x40, 0x81, 0x88, 0xa9, 0x67, 0x56, 0xb4, 0xbf,
	0x76, 0xcd, 0x7f, 0xd4, 0x3f, 0x5d, 0x13, 0x7d, 0xff, 0xf5, 0x63, 0x66, 0xee, 0xca, 0xa8, 0xec,
	0xfe, 0xd7, 0x1e, 0xda, 0x70, 0x07, 0xe1, 0x73, 0x2a, 0x81, 0xc0, 0xc7, 0xac, 0x1a, 0x51, 0xcc,
	0x07, 0x73, 0xa6, 0x0c, 0x6c, 0x35, 0x05, 0xb4, 0xfa, 0xfe, 0x97, 0xcc, 0x07, 0xfc, 0x06, 0x55,
	0x7d, 0x9a, 0x8c, 0x98, 0x21, 0x08, 0x07, 0x02, 0x65, 0xce, 0xfe, 0xfb, 0x8c, 0xcd, 0xc5, 0x94,
	0x3c, 0xbb, 0x5b, 0x31, 0xbb, 0x6d, 0xec, 0xd3, 0x64, 0xc0, 0x3d, 0xcd, 0x10, 0xd8, 0x41, 0x35,
	0x08, 0x7a, 0x5c, 0x38, 0x2c, 0x70, 0xc9, 0xe4, 0x0c, 0xe7, 0xca, 0xf1, 0xd7, 0x87, 0xa4, 0xc3,
	0xf1, 0xb9, 0xbe, 0x47, 0x6b, 0xa3, 0x10, 0x19, 0x39, 0x0e, 0x48, 0x49, 0x04, 0x55, 0x60, 0xce,
	0x97, 0x4b, 0xa8, 0x0e, 0x31, 0x9d, 0x8c, 0xd2, 0xa6, 0x2a, 0x1d, 0xcf, 0x46, 0x1e, 0x4a, 0x7c,
	0x16, 0x30, 0x3f, 0xf2, 0xc9, 0x07, 0x2e, 0x95, 0x34, 0xff, 0x9f, 0xe2, 0x22, 0xac, 0xcb, 0x11,
	0xee, 0x79, 0xe6, 0x7e, 0x9c, 0x9a, 0xf1, 0x5b, 0xb4, 0x39, 0x86, 0x16, 0xf0, 0x29, 0x02, 0xa9,
	0x48, 0xcc, 0xbd, 0xc8, 0x07, 0x73, 0x61, 0x0a, 0xb6, 0x99, 0x63, 0xb7, 0x33, 0xfb, 0x99, 0x76,
	0xe3, 0xd7, 0xa8, 0x36, 0x06, 0x97, 0xaa, 0x0b, 0x31, 0xe9, 0x51, 0x47, 0x71, 0x61, 0x2e, 0x4e,
	0x81, 0x5e, 0xcb, 0xa1, 0x3b, 0xa9, 0xf9, 0x58, 0x7b, 0xf1, 0x3b, 0xb4, 0x99, 0x7f, 0x94, 0x2e,
	0x55, 0x70, 0x41, 0x2f, 0x49, 0x8f, 0x32, 0x2f, 0x12, 0x60, 0xa2, 0x29, 0xd0, 0xb5, 0x1c, 0xe0,
	0x51, 0xe6, 0x3f, 0xce, 0xec, 0xf8, 0x33, 0xba, 0x5f, 0x7c, 0x65, 0x26, 0x73, 0x96, 0xca, 0x3d,
	0xdc, 0x7a, 0xd1, 0xf5, 0x99, 0xc8, 0x7e, 0x8a, 0xea, 0x32, 0xf4, 0x98, 0x22, 0x90, 0x28, 0x10,
	0x01, 0xf5, 0x88, 0xc7, 0x1d, 0xea, 0x11, 0x2e, 0x98, 0xcb, 0x02, 0x02, 0x42, 0x70, 0x21, 0xcd,
	0x1b, 0xdb, 0xc6, 0xce, 0x42, 0x7b, 0x4b, 0x2b, 0x5b, 0x7d, 0xe1, 0xb3, 0x54, 0x77, 0xa2, 0x65,
	0x2d, 0xad, 0xc2, 0x80, 0xb6, 0xf3, 0xed, 0xc7, 0x40, 0x83, 0x23, 0xdc, 0x9c, 0x62, 0x54, 0x77,
	0x72, 0x94, 0x5c, 0xca, 0xa0, 0xf2, 0x17, 0x03, 0xed, 0x16, 0xcf, 0xab, 0x30, 0x71, 0xb9, 0xdc,
	0xd0, 0x1e, 0x14, 0x0d, 0xad, 0xa0, 0x45, 0x84, 0xee, 0x8d, 0x4a, 0x8c, 0x05, 0x8f, 0xbd, 0x8f,
	0x2b, 0xe5, 0xd2, 0xef, 0x0e, 0x99, 0xb9, 0xc8, 0xdc, 0xab, 0xd9, 0x94, 0x5f, 0x7f, 0x6e, 0x19,
	0xa8, 0xce, 0xb8, 0xa5, 0xbf, 0x1d, 0xa1, 0xe0, 0xc9, 0xa5, 0x55, 0xf4, 0x19, 0x69, 0xde, 0x9a,
	0xfc, 0x53, 0x3f, 0x4d, 0xb3, 0x4f, 0x8d, 0xab, 0xca, 0x5a, 0x4b, 0xeb, 0x0f, 0x42, 0x66, 0x9d,
	0x35, 0xac, 0xc3, 0x4c, 0xff, 0xa2, 0xf3, 0xfb, 0x6f, 0x1b, 0xe7, 0xf3, 0xba, 0xf6, 0xde, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x97, 0xe4, 0x41, 0xde, 0x06, 0x00, 0x00,
}