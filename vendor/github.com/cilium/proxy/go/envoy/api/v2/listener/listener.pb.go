// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/listener/listener.proto

package listener

import (
	fmt "fmt"
	auth "github.com/cilium/proxy/go/envoy/api/v2/auth"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type FilterChainMatch_ConnectionSourceType int32

const (
	// Any connection source matches.
	FilterChainMatch_ANY FilterChainMatch_ConnectionSourceType = 0
	// Match a connection originating from the same host.
	FilterChainMatch_LOCAL FilterChainMatch_ConnectionSourceType = 1
	// Match a connection originating from a different host.
	FilterChainMatch_EXTERNAL FilterChainMatch_ConnectionSourceType = 2
)

var FilterChainMatch_ConnectionSourceType_name = map[int32]string{
	0: "ANY",
	1: "LOCAL",
	2: "EXTERNAL",
}

var FilterChainMatch_ConnectionSourceType_value = map[string]int32{
	"ANY":      0,
	"LOCAL":    1,
	"EXTERNAL": 2,
}

func (x FilterChainMatch_ConnectionSourceType) String() string {
	return proto.EnumName(FilterChainMatch_ConnectionSourceType_name, int32(x))
}

func (FilterChainMatch_ConnectionSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{1, 0}
}

type Filter struct {
	// The name of the filter to instantiate. The name must match a
	// :ref:`supported filter <config_network_filters>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*Filter_Config
	//	*Filter_TypedConfig
	ConfigType           isFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isFilter_ConfigType interface {
	isFilter_ConfigType()
}

type Filter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type Filter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,4,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Filter_Config) isFilter_ConfigType() {}

func (*Filter_TypedConfig) isFilter_ConfigType() {}

func (m *Filter) GetConfigType() isFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *Filter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*Filter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *Filter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Filter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Filter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Filter_Config)(nil),
		(*Filter_TypedConfig)(nil),
	}
}

// Specifies the match criteria for selecting a specific filter chain for a
// listener.
//
// In order for a filter chain to be selected, *ALL* of its criteria must be
// fulfilled by the incoming connection, properties of which are set by the
// networking stack and/or listener filters.
//
// The following order applies:
//
// 1. Destination port.
// 2. Destination IP address.
// 3. Server name (e.g. SNI for TLS protocol),
// 4. Transport protocol.
// 5. Application protocols (e.g. ALPN for TLS protocol).
// 6. Source type (e.g. any, local or external network).
// 7. Source IP address.
// 8. Source port.
//
// For criteria that allow ranges or wildcards, the most specific value in any
// of the configured filter chains that matches the incoming connection is going
// to be used (e.g. for SNI ``www.example.com`` the most specific match would be
// ``www.example.com``, then ``*.example.com``, then ``*.com``, then any filter
// chain without ``server_names`` requirements).
//
// [#comment: Implemented rules are kept in the preference order, with deprecated fields
// listed at the end, because that's how we want to list them in the docs.
//
// [#comment:TODO(PiotrSikora): Add support for configurable precedence of the rules]
type FilterChainMatch struct {
	// Optional destination port to consider when use_original_dst is set on the
	// listener in determining a filter chain match.
	DestinationPort *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	// If non-empty, an IP address and prefix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	PrefixRanges []*core.CidrRange `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges,proto3" json:"prefix_ranges,omitempty"`
	// If non-empty, an IP address and suffix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	// [#not-implemented-hide:]
	AddressSuffix string `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix,proto3" json:"address_suffix,omitempty"`
	// [#not-implemented-hide:]
	SuffixLen *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen,proto3" json:"suffix_len,omitempty"`
	// Specifies the connection source IP match type. Can be any, local or external network.
	SourceType FilterChainMatch_ConnectionSourceType `protobuf:"varint,12,opt,name=source_type,json=sourceType,proto3,enum=envoy.api.v2.listener.FilterChainMatch_ConnectionSourceType" json:"source_type,omitempty"`
	// The criteria is satisfied if the source IP address of the downstream
	// connection is contained in at least one of the specified subnets. If the
	// parameter is not specified or the list is empty, the source IP address is
	// ignored.
	SourcePrefixRanges []*core.CidrRange `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	// The criteria is satisfied if the source port of the downstream connection
	// is contained in at least one of the specified ports. If the parameter is
	// not specified, the source port is ignored.
	SourcePorts []uint32 `protobuf:"varint,7,rep,packed,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	// If non-empty, a list of server names (e.g. SNI for TLS protocol) to consider when determining
	// a filter chain match. Those values will be compared against the server names of a new
	// connection, when detected by one of the listener filters.
	//
	// The server name will be matched against all wildcard domains, i.e. ``www.example.com``
	// will be first matched against ``www.example.com``, then ``*.example.com``, then ``*.com``.
	//
	// Note that partial wildcards are not supported, and values like ``*w.example.com`` are invalid.
	//
	// .. attention::
	//
	//   See the :ref:`FAQ entry <faq_how_to_setup_sni>` on how to configure SNI for more
	//   information.
	ServerNames []string `protobuf:"bytes,11,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	// If non-empty, a transport protocol to consider when determining a filter chain match.
	// This value will be compared against the transport protocol of a new connection, when
	// it's detected by one of the listener filters.
	//
	// Suggested values include:
	//
	// * ``raw_buffer`` - default, used when no transport protocol is detected,
	// * ``tls`` - set by :ref:`envoy.listener.tls_inspector <config_listener_filters_tls_inspector>`
	//   when TLS protocol is detected.
	TransportProtocol string `protobuf:"bytes,9,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	// If non-empty, a list of application protocols (e.g. ALPN for TLS protocol) to consider when
	// determining a filter chain match. Those values will be compared against the application
	// protocols of a new connection, when detected by one of the listener filters.
	//
	// Suggested values include:
	//
	// * ``http/1.1`` - set by :ref:`envoy.listener.tls_inspector
	//   <config_listener_filters_tls_inspector>`,
	// * ``h2`` - set by :ref:`envoy.listener.tls_inspector <config_listener_filters_tls_inspector>`
	//
	// .. attention::
	//
	//   Currently, only :ref:`TLS Inspector <config_listener_filters_tls_inspector>` provides
	//   application protocol detection based on the requested
	//   `ALPN <https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation>`_ values.
	//
	//   However, the use of ALPN is pretty much limited to the HTTP/2 traffic on the Internet,
	//   and matching on values other than ``h2`` is going to lead to a lot of false negatives,
	//   unless all connecting clients are known to use ALPN.
	ApplicationProtocols []string `protobuf:"bytes,10,rep,name=application_protocols,json=applicationProtocols,proto3" json:"application_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterChainMatch) Reset()         { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()    {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{1}
}

func (m *FilterChainMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChainMatch.Unmarshal(m, b)
}
func (m *FilterChainMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChainMatch.Marshal(b, m, deterministic)
}
func (m *FilterChainMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChainMatch.Merge(m, src)
}
func (m *FilterChainMatch) XXX_Size() int {
	return xxx_messageInfo_FilterChainMatch.Size(m)
}
func (m *FilterChainMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChainMatch.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChainMatch proto.InternalMessageInfo

func (m *FilterChainMatch) GetDestinationPort() *wrappers.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourceType() FilterChainMatch_ConnectionSourceType {
	if m != nil {
		return m.SourceType
	}
	return FilterChainMatch_ANY
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []uint32 {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetServerNames() []string {
	if m != nil {
		return m.ServerNames
	}
	return nil
}

func (m *FilterChainMatch) GetTransportProtocol() string {
	if m != nil {
		return m.TransportProtocol
	}
	return ""
}

func (m *FilterChainMatch) GetApplicationProtocols() []string {
	if m != nil {
		return m.ApplicationProtocols
	}
	return nil
}

// A filter chain wraps a set of match criteria, an option TLS context, a set of filters, and
// various other parameters.
type FilterChain struct {
	// The criteria to use when matching a connection to this filter chain.
	FilterChainMatch *FilterChainMatch `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch,proto3" json:"filter_chain_match,omitempty"`
	// The TLS context for this filter chain.
	TlsContext *auth.DownstreamTlsContext `protobuf:"bytes,2,opt,name=tls_context,json=tlsContext,proto3" json:"tls_context,omitempty"`
	// A list of individual network filters that make up the filter chain for
	// connections established with the listener. Order matters as the filters are
	// processed sequentially as connection events happen. Note: If the filter
	// list is empty, the connection will close by default.
	Filters []*Filter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// Whether the listener should expect a PROXY protocol V1 header on new
	// connections. If this option is enabled, the listener will assume that that
	// remote address of the connection is the one specified in the header. Some
	// load balancers including the AWS ELB support this option. If the option is
	// absent or set to false, Envoy will use the physical peer address of the
	// connection as the remote address.
	UseProxyProto *wrappers.BoolValue `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	// [#not-implemented-hide:] filter chain metadata.
	Metadata *core.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// See :ref:`base.TransportSocket<envoy_api_msg_core.TransportSocket>` description.
	TransportSocket      *core.TransportSocket `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FilterChain) Reset()         { *m = FilterChain{} }
func (m *FilterChain) String() string { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()    {}
func (*FilterChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{2}
}

func (m *FilterChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChain.Unmarshal(m, b)
}
func (m *FilterChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChain.Marshal(b, m, deterministic)
}
func (m *FilterChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChain.Merge(m, src)
}
func (m *FilterChain) XXX_Size() int {
	return xxx_messageInfo_FilterChain.Size(m)
}
func (m *FilterChain) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChain.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChain proto.InternalMessageInfo

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

func (m *FilterChain) GetTlsContext() *auth.DownstreamTlsContext {
	if m != nil {
		return m.TlsContext
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *wrappers.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

type ListenerFilter struct {
	// The name of the filter to instantiate. The name must match a
	// :ref:`supported filter <config_listener_filters>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being instantiated.
	// See the supported filters for further documentation.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*ListenerFilter_Config
	//	*ListenerFilter_TypedConfig
	ConfigType           isListenerFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListenerFilter) Reset()         { *m = ListenerFilter{} }
func (m *ListenerFilter) String() string { return proto.CompactTextString(m) }
func (*ListenerFilter) ProtoMessage()    {}
func (*ListenerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{3}
}

func (m *ListenerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilter.Unmarshal(m, b)
}
func (m *ListenerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilter.Marshal(b, m, deterministic)
}
func (m *ListenerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilter.Merge(m, src)
}
func (m *ListenerFilter) XXX_Size() int {
	return xxx_messageInfo_ListenerFilter.Size(m)
}
func (m *ListenerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilter proto.InternalMessageInfo

func (m *ListenerFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListenerFilter_ConfigType interface {
	isListenerFilter_ConfigType()
}

type ListenerFilter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type ListenerFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ListenerFilter_Config) isListenerFilter_ConfigType() {}

func (*ListenerFilter_TypedConfig) isListenerFilter_ConfigType() {}

func (m *ListenerFilter) GetConfigType() isListenerFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ListenerFilter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ListenerFilter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ListenerFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ListenerFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerFilter_Config)(nil),
		(*ListenerFilter_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.api.v2.listener.FilterChainMatch_ConnectionSourceType", FilterChainMatch_ConnectionSourceType_name, FilterChainMatch_ConnectionSourceType_value)
	proto.RegisterType((*Filter)(nil), "envoy.api.v2.listener.Filter")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.api.v2.listener.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.api.v2.listener.FilterChain")
	proto.RegisterType((*ListenerFilter)(nil), "envoy.api.v2.listener.ListenerFilter")
}

func init() {
	proto.RegisterFile("envoy/api/v2/listener/listener.proto", fileDescriptor_0ced541f18749edd)
}

var fileDescriptor_0ced541f18749edd = []byte{
	// 901 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0x2b, 0x35,
	0x14, 0xee, 0x64, 0xd2, 0x34, 0xf1, 0x24, 0xed, 0x60, 0xe5, 0xaa, 0x43, 0xb9, 0x3f, 0x21, 0x80,
	0x88, 0x90, 0x98, 0x11, 0xa9, 0x10, 0x10, 0x60, 0x91, 0x09, 0x45, 0xe5, 0x2a, 0x0d, 0xd1, 0xa4,
	0xb7, 0x02, 0x36, 0x23, 0x77, 0xe2, 0xa4, 0x16, 0x13, 0x7b, 0x64, 0x3b, 0xb9, 0x8d, 0xd8, 0xf1,
	0x08, 0x3c, 0x05, 0xba, 0xe2, 0x05, 0x60, 0xc5, 0x92, 0x37, 0x60, 0xc7, 0x82, 0x1d, 0xbc, 0xc4,
	0x45, 0xf6, 0xcc, 0xa4, 0x4d, 0x1a, 0xc1, 0x15, 0x8b, 0xbb, 0xb3, 0xcf, 0xf7, 0x7d, 0xc7, 0xc7,
	0xdf, 0x39, 0x36, 0x78, 0x13, 0xd3, 0x05, 0x5b, 0x7a, 0x28, 0x21, 0xde, 0xa2, 0xed, 0xc5, 0x44,
	0x48, 0x4c, 0x31, 0x5f, 0x2d, 0xdc, 0x84, 0x33, 0xc9, 0xe0, 0x3d, 0xcd, 0x72, 0x51, 0x42, 0xdc,
	0x45, 0xdb, 0xcd, 0xc1, 0xa3, 0x47, 0x6b, 0xe2, 0x88, 0x71, 0xec, 0xa1, 0xf1, 0x98, 0x63, 0x21,
	0x52, 0xdd, 0xd1, 0xfd, 0x35, 0x02, 0x9a, 0xcb, 0x2b, 0x2f, 0xc2, 0x5c, 0x6e, 0x45, 0xb5, 0xfc,
	0x12, 0x09, 0x9c, 0xa1, 0xaf, 0x4e, 0x19, 0x9b, 0xc6, 0xd8, 0xd3, 0xbb, 0xcb, 0xf9, 0xc4, 0x43,
	0x74, 0x99, 0x0b, 0x37, 0x21, 0x21, 0xf9, 0x3c, 0xca, 0xd3, 0x3e, 0xdc, 0x44, 0x9f, 0x72, 0x94,
	0x24, 0x98, 0xe7, 0x45, 0x1d, 0x2e, 0x50, 0x4c, 0xc6, 0x48, 0x62, 0x2f, 0x5f, 0x64, 0x40, 0x7d,
	0xca, 0xa6, 0x4c, 0x2f, 0x3d, 0xb5, 0x4a, 0xa3, 0xcd, 0x67, 0x06, 0x28, 0x7d, 0x4e, 0x62, 0x89,
	0x39, 0x7c, 0x00, 0x8a, 0x14, 0xcd, 0xb0, 0x63, 0x34, 0x8c, 0x56, 0xc5, 0xaf, 0xfc, 0xf2, 0xd7,
	0xaf, 0x66, 0x91, 0x17, 0x1a, 0x46, 0xa0, 0xc3, 0xf0, 0x3d, 0x50, 0x8a, 0x18, 0x9d, 0x90, 0xa9,
	0x53, 0x68, 0x18, 0x2d, 0xab, 0x7d, 0xe8, 0xa6, 0x95, 0xb8, 0x79, 0x25, 0xee, 0x48, 0xd7, 0x79,
	0xba, 0x13, 0x64, 0x44, 0xf8, 0x11, 0xa8, 0xca, 0x65, 0x82, 0xc7, 0x61, 0x26, 0x2c, 0x6a, 0x61,
	0xfd, 0x8e, 0xb0, 0x4b, 0x97, 0xa7, 0x3b, 0x81, 0xa5, 0xb9, 0x3d, 0x4d, 0xf5, 0x6b, 0xc0, 0x4a,
	0x45, 0xa1, 0x8a, 0x3e, 0x2e, 0x96, 0x4d, 0xbb, 0xd8, 0xfc, 0x7d, 0x17, 0xd8, 0x69, 0xb1, 0xbd,
	0x2b, 0x44, 0xe8, 0x19, 0x92, 0xd1, 0x15, 0xbc, 0x00, 0xf6, 0x18, 0x0b, 0x49, 0x28, 0x92, 0x84,
	0xd1, 0x30, 0x61, 0x5c, 0x3a, 0x65, 0x7d, 0xd0, 0xfd, 0x3b, 0x07, 0x3d, 0xf9, 0x82, 0xca, 0xe3,
	0xf6, 0x05, 0x8a, 0xe7, 0xd8, 0xaf, 0xa9, 0x0b, 0x96, 0xdf, 0x29, 0x39, 0xcf, 0x9f, 0x9b, 0x2d,
	0x23, 0x38, 0xb8, 0x95, 0x64, 0xc8, 0xb8, 0x84, 0x5d, 0x50, 0x4b, 0x38, 0x9e, 0x90, 0xeb, 0x90,
	0x23, 0x3a, 0xc5, 0xc2, 0x31, 0x1b, 0xa6, 0x4e, 0xba, 0x36, 0x2d, 0xaa, 0xaf, 0x6e, 0x8f, 0x8c,
	0x79, 0xa0, 0x48, 0x41, 0x35, 0x95, 0xe8, 0x8d, 0x80, 0x6f, 0x81, 0xfd, 0x6c, 0x62, 0x42, 0x31,
	0x9f, 0x4c, 0xc8, 0xb5, 0x76, 0xa0, 0x12, 0xd4, 0xb2, 0xe8, 0x48, 0x07, 0xe1, 0xc7, 0x00, 0xa4,
	0x70, 0x18, 0x63, 0xea, 0xec, 0xfe, 0x77, 0xed, 0x41, 0x25, 0xe5, 0xf7, 0x31, 0x85, 0x04, 0x58,
	0x82, 0xcd, 0x79, 0x84, 0xb5, 0x51, 0x4e, 0xb5, 0x61, 0xb4, 0xf6, 0xdb, 0x9f, 0xb8, 0x5b, 0x47,
	0xda, 0xdd, 0x34, 0xcf, 0xed, 0x31, 0x4a, 0x71, 0xa4, 0xee, 0x3c, 0xd2, 0x49, 0xce, 0x97, 0x09,
	0xf6, 0x81, 0x72, 0x66, 0xf7, 0x7b, 0xa3, 0x60, 0x1b, 0x01, 0x10, 0xab, 0x38, 0x1c, 0x80, 0x7a,
	0x76, 0xd4, 0xba, 0x31, 0xa5, 0x17, 0x30, 0x06, 0xa6, 0xca, 0xe1, 0x6d, 0x7b, 0xde, 0x07, 0xd5,
	0x3c, 0x1f, 0xe3, 0x52, 0x38, 0x7b, 0x0d, 0xb3, 0x55, 0xf3, 0xa1, 0x3a, 0xbd, 0xf6, 0x83, 0x01,
	0x9a, 0x37, 0xcd, 0xc9, 0xae, 0xa8, 0xfa, 0x22, 0xe0, 0xeb, 0xa0, 0x2a, 0x30, 0x5f, 0x60, 0x1e,
	0xaa, 0xb9, 0x14, 0x8e, 0xd5, 0x30, 0x5b, 0x95, 0xc0, 0x4a, 0x63, 0x03, 0x15, 0x82, 0xef, 0x02,
	0x28, 0x39, 0xa2, 0x42, 0xe5, 0x0d, 0xb5, 0x83, 0x11, 0x8b, 0x9d, 0x8a, 0x36, 0xff, 0x95, 0x15,
	0x32, 0xcc, 0x00, 0x78, 0x0c, 0xee, 0xa1, 0x24, 0x89, 0x49, 0x94, 0x8d, 0x50, 0x16, 0x17, 0x0e,
	0xd0, 0xa9, 0xeb, 0xb7, 0xc0, 0x5c, 0x23, 0x9a, 0x1f, 0x82, 0xfa, 0x36, 0xf7, 0xe0, 0x1e, 0x30,
	0xbb, 0x83, 0xaf, 0xed, 0x1d, 0x58, 0x01, 0xbb, 0xfd, 0x2f, 0x7b, 0xdd, 0xbe, 0x6d, 0xc0, 0x2a,
	0x28, 0x9f, 0x7c, 0x75, 0x7e, 0x12, 0x0c, 0xba, 0x7d, 0xbb, 0xf0, 0xb8, 0x58, 0x36, 0xec, 0x42,
	0x60, 0x09, 0x4a, 0xc2, 0x31, 0x9b, 0x21, 0x42, 0x45, 0xf3, 0x67, 0x13, 0x58, 0xb7, 0x9a, 0x03,
	0x9f, 0x00, 0x38, 0xd1, 0xdb, 0x30, 0x52, 0xfb, 0x70, 0xa6, 0xba, 0xa5, 0x5f, 0xa6, 0xd5, 0x7e,
	0xfb, 0x05, 0x9b, 0x1b, 0xd8, 0x93, 0xcd, 0xb7, 0x72, 0x0a, 0x2c, 0x19, 0x0b, 0xf5, 0x1c, 0x25,
	0xbe, 0x96, 0xd9, 0x43, 0xde, 0xc8, 0xa7, 0xfe, 0x31, 0xf7, 0x33, 0xf6, 0x94, 0x0a, 0xc9, 0x31,
	0x9a, 0x9d, 0xc7, 0xa2, 0x97, 0xd2, 0x03, 0x20, 0x57, 0x6b, 0xf8, 0x29, 0xd8, 0x4b, 0xb3, 0xe7,
	0xef, 0xe2, 0xc1, 0xbf, 0x56, 0xe5, 0x17, 0x7f, 0xfb, 0xe3, 0xd1, 0x4e, 0x90, 0x6b, 0xa0, 0x0f,
	0x0e, 0xe6, 0x42, 0xcd, 0x11, 0xbb, 0x5e, 0xa6, 0x7e, 0x67, 0x9f, 0xc3, 0xd1, 0x9d, 0xb9, 0xf7,
	0x19, 0x8b, 0xd3, 0xa9, 0xaf, 0xcd, 0x05, 0x1e, 0x2a, 0x85, 0x6e, 0x02, 0xfc, 0x00, 0x94, 0x67,
	0x58, 0xa2, 0x31, 0x92, 0x28, 0x7b, 0x34, 0xaf, 0x6d, 0x19, 0xc1, 0xb3, 0x8c, 0x12, 0xac, 0xc8,
	0xf0, 0x0c, 0xd8, 0x37, 0xd3, 0x21, 0x58, 0xf4, 0x2d, 0x96, 0x4e, 0x49, 0x27, 0x68, 0x6e, 0x49,
	0x70, 0x9e, 0x53, 0x47, 0x9a, 0x19, 0x1c, 0xc8, 0xf5, 0x40, 0xf3, 0x27, 0x03, 0xec, 0xf7, 0xb3,
	0xeb, 0xbe, 0xb4, 0xaf, 0xd4, 0xfc, 0xbf, 0x5f, 0xa9, 0xff, 0xdd, 0x8f, 0x7f, 0x3e, 0x34, 0xc0,
	0x1b, 0x84, 0xa5, 0x77, 0xd5, 0x2d, 0xd8, 0xde, 0x3b, 0xbf, 0x96, 0x5f, 0x4b, 0x1b, 0x3e, 0x34,
	0xbe, 0x29, 0xe7, 0xd0, 0xb3, 0xc2, 0xe1, 0x89, 0x96, 0x74, 0x13, 0xe2, 0x5e, 0xb4, 0xdd, 0x9c,
	0x38, 0x18, 0xfd, 0x5d, 0x38, 0xd2, 0x48, 0xa7, 0xd3, 0x4d, 0x48, 0xa7, 0x73, 0xd1, 0xee, 0x74,
	0x6e, 0xc0, 0xcb, 0x92, 0x2e, 0xf4, 0xf8, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x63, 0x9a,
	0x3c, 0x99, 0x07, 0x00, 0x00,
}