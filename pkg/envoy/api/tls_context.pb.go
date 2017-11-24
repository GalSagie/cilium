// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/tls_context.proto

package envoy_api_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TlsParameters_TlsProtocol int32

const (
	TlsParameters_TLS_AUTO TlsParameters_TlsProtocol = 0
	TlsParameters_TLSv1_0  TlsParameters_TlsProtocol = 1
	TlsParameters_TLSv1_1  TlsParameters_TlsProtocol = 2
	TlsParameters_TLSv1_2  TlsParameters_TlsProtocol = 3
	TlsParameters_TLSv1_3  TlsParameters_TlsProtocol = 4
)

var TlsParameters_TlsProtocol_name = map[int32]string{
	0: "TLS_AUTO",
	1: "TLSv1_0",
	2: "TLSv1_1",
	3: "TLSv1_2",
	4: "TLSv1_3",
}
var TlsParameters_TlsProtocol_value = map[string]int32{
	"TLS_AUTO": 0,
	"TLSv1_0":  1,
	"TLSv1_1":  2,
	"TLSv1_2":  3,
	"TLSv1_3":  4,
}

func (x TlsParameters_TlsProtocol) String() string {
	return proto.EnumName(TlsParameters_TlsProtocol_name, int32(x))
}
func (TlsParameters_TlsProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor13, []int{1, 0}
}

type DataSource struct {
	// Types that are valid to be assigned to Specifier:
	//	*DataSource_Filename
	//	*DataSource_Inline
	Specifier isDataSource_Specifier `protobuf_oneof:"specifier"`
}

func (m *DataSource) Reset()                    { *m = DataSource{} }
func (m *DataSource) String() string            { return proto.CompactTextString(m) }
func (*DataSource) ProtoMessage()               {}
func (*DataSource) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

type isDataSource_Specifier interface {
	isDataSource_Specifier()
}

type DataSource_Filename struct {
	Filename string `protobuf:"bytes,1,opt,name=filename,oneof"`
}
type DataSource_Inline struct {
	Inline []byte `protobuf:"bytes,2,opt,name=inline,proto3,oneof"`
}

func (*DataSource_Filename) isDataSource_Specifier() {}
func (*DataSource_Inline) isDataSource_Specifier()   {}

func (m *DataSource) GetSpecifier() isDataSource_Specifier {
	if m != nil {
		return m.Specifier
	}
	return nil
}

func (m *DataSource) GetFilename() string {
	if x, ok := m.GetSpecifier().(*DataSource_Filename); ok {
		return x.Filename
	}
	return ""
}

func (m *DataSource) GetInline() []byte {
	if x, ok := m.GetSpecifier().(*DataSource_Inline); ok {
		return x.Inline
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DataSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DataSource_OneofMarshaler, _DataSource_OneofUnmarshaler, _DataSource_OneofSizer, []interface{}{
		(*DataSource_Filename)(nil),
		(*DataSource_Inline)(nil),
	}
}

func _DataSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DataSource)
	// specifier
	switch x := m.Specifier.(type) {
	case *DataSource_Filename:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Filename)
	case *DataSource_Inline:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Inline)
	case nil:
	default:
		return fmt.Errorf("DataSource.Specifier has unexpected type %T", x)
	}
	return nil
}

func _DataSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DataSource)
	switch tag {
	case 1: // specifier.filename
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Specifier = &DataSource_Filename{x}
		return true, err
	case 2: // specifier.inline
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Specifier = &DataSource_Inline{x}
		return true, err
	default:
		return false, nil
	}
}

func _DataSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DataSource)
	// specifier
	switch x := m.Specifier.(type) {
	case *DataSource_Filename:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Filename)))
		n += len(x.Filename)
	case *DataSource_Inline:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Inline)))
		n += len(x.Inline)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TlsParameters struct {
	// Allowed TLS protocols.
	TlsMinimumProtocolVersion TlsParameters_TlsProtocol `protobuf:"varint,1,opt,name=tls_minimum_protocol_version,json=tlsMinimumProtocolVersion,enum=envoy.api.v2.TlsParameters_TlsProtocol" json:"tls_minimum_protocol_version,omitempty"`
	TlsMaximumProtocolVersion TlsParameters_TlsProtocol `protobuf:"varint,2,opt,name=tls_maximum_protocol_version,json=tlsMaximumProtocolVersion,enum=envoy.api.v2.TlsParameters_TlsProtocol" json:"tls_maximum_protocol_version,omitempty"`
	// If specified, the TLS listener will only support the specified cipher list.
	CipherSuites []string `protobuf:"bytes,3,rep,name=cipher_suites,json=cipherSuites" json:"cipher_suites,omitempty"`
	// If specified, the TLS connection will only support the specified ECDH
	// curves. If not specified, the default curves (X25519, P-256) will be used.
	EcdhCurves []string `protobuf:"bytes,4,rep,name=ecdh_curves,json=ecdhCurves" json:"ecdh_curves,omitempty"`
}

func (m *TlsParameters) Reset()                    { *m = TlsParameters{} }
func (m *TlsParameters) String() string            { return proto.CompactTextString(m) }
func (*TlsParameters) ProtoMessage()               {}
func (*TlsParameters) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *TlsParameters) GetTlsMinimumProtocolVersion() TlsParameters_TlsProtocol {
	if m != nil {
		return m.TlsMinimumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (m *TlsParameters) GetTlsMaximumProtocolVersion() TlsParameters_TlsProtocol {
	if m != nil {
		return m.TlsMaximumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (m *TlsParameters) GetCipherSuites() []string {
	if m != nil {
		return m.CipherSuites
	}
	return nil
}

func (m *TlsParameters) GetEcdhCurves() []string {
	if m != nil {
		return m.EcdhCurves
	}
	return nil
}

// TLS certs can be loaded from file or delivered inline [V2-API-DIFF]. Individual fields may
// be loaded from either.
type TlsCertificate struct {
	CertificateChain           *DataSource   `protobuf:"bytes,1,opt,name=certificate_chain,json=certificateChain" json:"certificate_chain,omitempty"`
	PrivateKey                 *DataSource   `protobuf:"bytes,2,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	Password                   *DataSource   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	OcspStaple                 *DataSource   `protobuf:"bytes,4,opt,name=ocsp_staple,json=ocspStaple" json:"ocsp_staple,omitempty"`
	SignedCertificateTimestamp []*DataSource `protobuf:"bytes,5,rep,name=signed_certificate_timestamp,json=signedCertificateTimestamp" json:"signed_certificate_timestamp,omitempty"`
}

func (m *TlsCertificate) Reset()                    { *m = TlsCertificate{} }
func (m *TlsCertificate) String() string            { return proto.CompactTextString(m) }
func (*TlsCertificate) ProtoMessage()               {}
func (*TlsCertificate) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *TlsCertificate) GetCertificateChain() *DataSource {
	if m != nil {
		return m.CertificateChain
	}
	return nil
}

func (m *TlsCertificate) GetPrivateKey() *DataSource {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *TlsCertificate) GetPassword() *DataSource {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *TlsCertificate) GetOcspStaple() *DataSource {
	if m != nil {
		return m.OcspStaple
	}
	return nil
}

func (m *TlsCertificate) GetSignedCertificateTimestamp() []*DataSource {
	if m != nil {
		return m.SignedCertificateTimestamp
	}
	return nil
}

type CertificateValidationContext struct {
	// TLS certificate data containing certificate authority certificates to use
	// in verifying a presented certificate. If not specified and a certificate is
	// presented it will not be verified.
	TrustedCa *DataSource `protobuf:"bytes,1,opt,name=trusted_ca,json=trustedCa" json:"trusted_ca,omitempty"`
	// If specified, Envoy will verify (pin) hex-encoded SHA-256 hash of
	// the presented certificate.
	VerifyCertificateHash []string `protobuf:"bytes,2,rep,name=verify_certificate_hash,json=verifyCertificateHash" json:"verify_certificate_hash,omitempty"`
	// If specified, Envoy will verify (pin) base64-encoded SHA-256 hash of
	// the Subject Public Key Information (SPKI) of the presented certificate.
	// This is the same format as used in HTTP Public Key Pinning.
	VerifySpkiSha256 []string `protobuf:"bytes,3,rep,name=verify_spki_sha256,json=verifySpkiSha256" json:"verify_spki_sha256,omitempty"`
	// An optional list of subject alt names. If specified, Envoy will verify that
	// the certificate’s subject alt name matches one of the specified values.
	VerifySubjectAltName []string `protobuf:"bytes,4,rep,name=verify_subject_alt_name,json=verifySubjectAltName" json:"verify_subject_alt_name,omitempty"`
	// Must present a signed time-stamped OCSP response.
	RequireOcspStaple *google_protobuf.BoolValue `protobuf:"bytes,5,opt,name=require_ocsp_staple,json=requireOcspStaple" json:"require_ocsp_staple,omitempty"`
	// Must present signed certificate time-stamp.
	RequireSignedCertificateTimestamp *google_protobuf.BoolValue `protobuf:"bytes,6,opt,name=require_signed_certificate_timestamp,json=requireSignedCertificateTimestamp" json:"require_signed_certificate_timestamp,omitempty"`
}

func (m *CertificateValidationContext) Reset()                    { *m = CertificateValidationContext{} }
func (m *CertificateValidationContext) String() string            { return proto.CompactTextString(m) }
func (*CertificateValidationContext) ProtoMessage()               {}
func (*CertificateValidationContext) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *CertificateValidationContext) GetTrustedCa() *DataSource {
	if m != nil {
		return m.TrustedCa
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifyCertificateHash() []string {
	if m != nil {
		return m.VerifyCertificateHash
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifySpkiSha256() []string {
	if m != nil {
		return m.VerifySpkiSha256
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifySubjectAltName() []string {
	if m != nil {
		return m.VerifySubjectAltName
	}
	return nil
}

func (m *CertificateValidationContext) GetRequireOcspStaple() *google_protobuf.BoolValue {
	if m != nil {
		return m.RequireOcspStaple
	}
	return nil
}

func (m *CertificateValidationContext) GetRequireSignedCertificateTimestamp() *google_protobuf.BoolValue {
	if m != nil {
		return m.RequireSignedCertificateTimestamp
	}
	return nil
}

// TLS context shared by both client and server TLS contexts.
type CommonTlsContext struct {
	// TLS protocol versions, cipher suites etc.
	TlsParams *TlsParameters `protobuf:"bytes,1,opt,name=tls_params,json=tlsParams" json:"tls_params,omitempty"`
	// Multiple TLS certificates can be associated with the same context,
	// e.g. to allow both RSA and ECDSA certificates [V2-API-DIFF].
	TlsCertificates []*TlsCertificate `protobuf:"bytes,2,rep,name=tls_certificates,json=tlsCertificates" json:"tls_certificates,omitempty"`
	// How to validate peer certificates.
	ValidationContext *CertificateValidationContext `protobuf:"bytes,3,opt,name=validation_context,json=validationContext" json:"validation_context,omitempty"`
	// Protocols to negotiate over ALPN
	AlpnProtocols []string                       `protobuf:"bytes,4,rep,name=alpn_protocols,json=alpnProtocols" json:"alpn_protocols,omitempty"`
	DeprecatedV1  *CommonTlsContext_DeprecatedV1 `protobuf:"bytes,5,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
}

func (m *CommonTlsContext) Reset()                    { *m = CommonTlsContext{} }
func (m *CommonTlsContext) String() string            { return proto.CompactTextString(m) }
func (*CommonTlsContext) ProtoMessage()               {}
func (*CommonTlsContext) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *CommonTlsContext) GetTlsParams() *TlsParameters {
	if m != nil {
		return m.TlsParams
	}
	return nil
}

func (m *CommonTlsContext) GetTlsCertificates() []*TlsCertificate {
	if m != nil {
		return m.TlsCertificates
	}
	return nil
}

func (m *CommonTlsContext) GetValidationContext() *CertificateValidationContext {
	if m != nil {
		return m.ValidationContext
	}
	return nil
}

func (m *CommonTlsContext) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

func (m *CommonTlsContext) GetDeprecatedV1() *CommonTlsContext_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

// These fields are deprecated and only are used during the interim v1 -> v2
// transition period for internal purposes. They should not be used outside of
// the Envoy binary.
type CommonTlsContext_DeprecatedV1 struct {
	AltAlpnProtocols string `protobuf:"bytes,1,opt,name=alt_alpn_protocols,json=altAlpnProtocols" json:"alt_alpn_protocols,omitempty"`
}

func (m *CommonTlsContext_DeprecatedV1) Reset()         { *m = CommonTlsContext_DeprecatedV1{} }
func (m *CommonTlsContext_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*CommonTlsContext_DeprecatedV1) ProtoMessage()    {}
func (*CommonTlsContext_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{4, 0}
}

func (m *CommonTlsContext_DeprecatedV1) GetAltAlpnProtocols() string {
	if m != nil {
		return m.AltAlpnProtocols
	}
	return ""
}

type UpstreamTlsContext struct {
	CommonTlsContext *CommonTlsContext `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext" json:"common_tls_context,omitempty"`
	// SNI string to use when creating TLS backend connections.
	Sni string `protobuf:"bytes,2,opt,name=sni" json:"sni,omitempty"`
}

func (m *UpstreamTlsContext) Reset()                    { *m = UpstreamTlsContext{} }
func (m *UpstreamTlsContext) String() string            { return proto.CompactTextString(m) }
func (*UpstreamTlsContext) ProtoMessage()               {}
func (*UpstreamTlsContext) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *UpstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if m != nil {
		return m.CommonTlsContext
	}
	return nil
}

func (m *UpstreamTlsContext) GetSni() string {
	if m != nil {
		return m.Sni
	}
	return ""
}

// [V2-API-DIFF] This has been reworked to support alternative modes of
// certificate/key delivery, for consistency with the upstream TLS context and
// to segregate the client/server aspects of the TLS context.
type DownstreamTlsContext struct {
	CommonTlsContext *CommonTlsContext `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext" json:"common_tls_context,omitempty"`
	// If specified, Envoy will reject connections without a valid client
	// certificate.
	RequireClientCertificate *google_protobuf.BoolValue `protobuf:"bytes,2,opt,name=require_client_certificate,json=requireClientCertificate" json:"require_client_certificate,omitempty"`
	// If specified, Envoy will reject connections without a valid and matching SNI.
	RequireSni *google_protobuf.BoolValue `protobuf:"bytes,3,opt,name=require_sni,json=requireSni" json:"require_sni,omitempty"`
}

func (m *DownstreamTlsContext) Reset()                    { *m = DownstreamTlsContext{} }
func (m *DownstreamTlsContext) String() string            { return proto.CompactTextString(m) }
func (*DownstreamTlsContext) ProtoMessage()               {}
func (*DownstreamTlsContext) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *DownstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if m != nil {
		return m.CommonTlsContext
	}
	return nil
}

func (m *DownstreamTlsContext) GetRequireClientCertificate() *google_protobuf.BoolValue {
	if m != nil {
		return m.RequireClientCertificate
	}
	return nil
}

func (m *DownstreamTlsContext) GetRequireSni() *google_protobuf.BoolValue {
	if m != nil {
		return m.RequireSni
	}
	return nil
}

func init() {
	proto.RegisterType((*DataSource)(nil), "envoy.api.v2.DataSource")
	proto.RegisterType((*TlsParameters)(nil), "envoy.api.v2.TlsParameters")
	proto.RegisterType((*TlsCertificate)(nil), "envoy.api.v2.TlsCertificate")
	proto.RegisterType((*CertificateValidationContext)(nil), "envoy.api.v2.CertificateValidationContext")
	proto.RegisterType((*CommonTlsContext)(nil), "envoy.api.v2.CommonTlsContext")
	proto.RegisterType((*CommonTlsContext_DeprecatedV1)(nil), "envoy.api.v2.CommonTlsContext.DeprecatedV1")
	proto.RegisterType((*UpstreamTlsContext)(nil), "envoy.api.v2.UpstreamTlsContext")
	proto.RegisterType((*DownstreamTlsContext)(nil), "envoy.api.v2.DownstreamTlsContext")
	proto.RegisterEnum("envoy.api.v2.TlsParameters_TlsProtocol", TlsParameters_TlsProtocol_name, TlsParameters_TlsProtocol_value)
}

func init() { proto.RegisterFile("api/tls_context.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x6e, 0x92, 0x6e, 0x69, 0x4e, 0xd2, 0xe2, 0x0e, 0xbb, 0xc2, 0x84, 0x68, 0x29, 0x01, 0x44,
	0x05, 0x28, 0xa5, 0x59, 0x76, 0x11, 0x3f, 0x37, 0xdd, 0x14, 0x51, 0x41, 0xa1, 0x8b, 0x93, 0xad,
	0x80, 0x1b, 0x6b, 0xea, 0x9c, 0x34, 0x43, 0xc7, 0x9e, 0x61, 0x66, 0xec, 0x6e, 0x1f, 0x89, 0x0b,
	0x9e, 0x89, 0xf7, 0x80, 0x1b, 0xe4, 0xf1, 0x38, 0x71, 0xb2, 0x90, 0x4a, 0x48, 0xdc, 0xe5, 0xfc,
	0x7d, 0xdf, 0xc9, 0x9c, 0xef, 0x1c, 0xc3, 0x03, 0x2a, 0xd9, 0xa1, 0xe1, 0x3a, 0x8c, 0x44, 0x62,
	0xf0, 0x85, 0xe9, 0x4b, 0x25, 0x8c, 0x20, 0x6d, 0x4c, 0x32, 0x71, 0xdb, 0xa7, 0x92, 0xf5, 0xb3,
	0x41, 0xe7, 0xe1, 0x95, 0x10, 0x57, 0x1c, 0x0f, 0x6d, 0xec, 0x32, 0x9d, 0x1e, 0xde, 0x28, 0x2a,
	0x25, 0x2a, 0x5d, 0x64, 0xf7, 0x7e, 0x00, 0x38, 0xa1, 0x86, 0x8e, 0x44, 0xaa, 0x22, 0x24, 0x5d,
	0xd8, 0x9e, 0x32, 0x8e, 0x09, 0x8d, 0xd1, 0xaf, 0xed, 0xd7, 0x0e, 0x9a, 0xa7, 0x1b, 0xc1, 0xdc,
	0x43, 0x7c, 0xd8, 0x62, 0x09, 0x67, 0x09, 0xfa, 0xf5, 0xfd, 0xda, 0x41, 0xfb, 0x74, 0x23, 0x70,
	0xf6, 0xd3, 0x16, 0x34, 0xb5, 0xc4, 0x88, 0x4d, 0x19, 0xaa, 0xde, 0x9f, 0x75, 0xd8, 0x19, 0x73,
	0xfd, 0x8c, 0x2a, 0x1a, 0xa3, 0x41, 0xa5, 0xc9, 0x0c, 0xba, 0x79, 0x9f, 0x31, 0x4b, 0x58, 0x9c,
	0xc6, 0xa1, 0x65, 0x8e, 0x04, 0x0f, 0x33, 0x54, 0x9a, 0x89, 0xc4, 0x52, 0xed, 0x0e, 0xde, 0xef,
	0x57, 0x3b, 0xef, 0x2f, 0x41, 0x58, 0xcb, 0x95, 0x05, 0x6f, 0x18, 0xae, 0xbf, 0x2b, 0xb0, 0x4a,
	0xdf, 0x45, 0x81, 0x34, 0x67, 0xa2, 0x2f, 0xfe, 0x99, 0xa9, 0xfe, 0x1f, 0x98, 0x0a, 0xac, 0x55,
	0xa6, 0x77, 0x60, 0x27, 0x62, 0x72, 0x86, 0x2a, 0xd4, 0x29, 0x33, 0xa8, 0xfd, 0xc6, 0x7e, 0xe3,
	0xa0, 0x19, 0xb4, 0x0b, 0xe7, 0xc8, 0xfa, 0xc8, 0x5b, 0xd0, 0xc2, 0x68, 0x32, 0x0b, 0xa3, 0x54,
	0x65, 0xa8, 0xfd, 0x4d, 0x9b, 0x02, 0xb9, 0x6b, 0x68, 0x3d, 0xbd, 0x73, 0x68, 0x55, 0xf8, 0x48,
	0x1b, 0xb6, 0xc7, 0x67, 0xa3, 0xf0, 0xf8, 0xf9, 0xf8, 0xdc, 0xdb, 0x20, 0x2d, 0x78, 0x65, 0x7c,
	0x36, 0xca, 0x8e, 0xc2, 0x8f, 0xbd, 0xda, 0xc2, 0x38, 0xf2, 0xea, 0x0b, 0x63, 0xe0, 0x35, 0x16,
	0xc6, 0x23, 0x6f, 0xb3, 0xf7, 0x47, 0x1d, 0x76, 0xc7, 0x5c, 0x0f, 0x51, 0x19, 0x36, 0x65, 0x11,
	0x35, 0x48, 0xbe, 0x82, 0xbd, 0x68, 0x61, 0x86, 0xd1, 0x8c, 0xb2, 0xe2, 0xc9, 0x5b, 0x03, 0x7f,
	0xf9, 0x21, 0x16, 0x4a, 0x08, 0xbc, 0x4a, 0xc9, 0x30, 0xaf, 0x20, 0x9f, 0x41, 0x4b, 0x2a, 0x96,
	0xe5, 0x10, 0xd7, 0x78, 0x6b, 0x5f, 0x72, 0x1d, 0x00, 0xb8, 0xe4, 0x6f, 0xf1, 0x96, 0x7c, 0x02,
	0xdb, 0x92, 0x6a, 0x7d, 0x23, 0xd4, 0xc4, 0x6f, 0xdc, 0x51, 0x37, 0xcf, 0xcc, 0x09, 0x45, 0xa4,
	0x65, 0xa8, 0x0d, 0x95, 0x1c, 0xfd, 0xcd, 0xbb, 0x08, 0xf3, 0xe4, 0x91, 0xcd, 0x25, 0x3f, 0x43,
	0x57, 0xb3, 0xab, 0x04, 0x27, 0x61, 0xf5, 0x9f, 0x1b, 0x16, 0xa3, 0x36, 0x34, 0x96, 0xfe, 0xbd,
	0xfd, 0xc6, 0x5a, 0xac, 0x4e, 0x51, 0x5d, 0x79, 0xc5, 0x71, 0x59, 0xdb, 0xfb, 0xbd, 0x01, 0xdd,
	0x4a, 0xe0, 0x82, 0x72, 0x36, 0xa1, 0x86, 0x89, 0x64, 0x58, 0xac, 0x21, 0xf9, 0x14, 0xc0, 0xa8,
	0x54, 0x9b, 0x9c, 0x9d, 0xde, 0xf9, 0xd0, 0x4d, 0x97, 0x3b, 0xa4, 0xe4, 0x09, 0xbc, 0x9e, 0xa1,
	0x62, 0xd3, 0xdb, 0xa5, 0xae, 0x67, 0x54, 0xcf, 0xfc, 0xba, 0x55, 0xce, 0x83, 0x22, 0x5c, 0x61,
	0x3f, 0xa5, 0x7a, 0x46, 0x3e, 0x02, 0xe2, 0xea, 0xb4, 0xbc, 0x66, 0xa1, 0x9e, 0xd1, 0xc1, 0xe3,
	0x27, 0x4e, 0x8f, 0x5e, 0x11, 0x19, 0xc9, 0x6b, 0x36, 0xb2, 0x7e, 0xf2, 0x78, 0xce, 0xa2, 0xd3,
	0xcb, 0x5f, 0x30, 0x32, 0x21, 0xe5, 0x26, 0xb4, 0x2b, 0x5f, 0xe8, 0xf3, 0xbe, 0x2b, 0x29, 0xa2,
	0xc7, 0xdc, 0x7c, 0x9f, 0x2f, 0xff, 0x37, 0xf0, 0x9a, 0xc2, 0x5f, 0x53, 0xa6, 0x30, 0xac, 0x4e,
	0xe5, 0x9e, 0xfd, 0x7b, 0x9d, 0x7e, 0x71, 0x66, 0xfa, 0xe5, 0x99, 0xe9, 0x3f, 0x15, 0x82, 0x5f,
	0x50, 0x9e, 0x62, 0xb0, 0xe7, 0xca, 0xce, 0x17, 0xe3, 0xb9, 0x86, 0x77, 0x4b, 0xac, 0xb5, 0x63,
	0xda, 0xba, 0x13, 0xfc, 0x6d, 0x87, 0x33, 0xfa, 0xf7, 0x79, 0xfd, 0xd6, 0x00, 0x6f, 0x28, 0xe2,
	0x58, 0x24, 0xf9, 0x5e, 0xb8, 0x19, 0x7d, 0x0e, 0x90, 0xdf, 0x09, 0x99, 0xaf, 0xbd, 0x76, 0x33,
	0x7a, 0x73, 0xcd, 0x55, 0x08, 0x9a, 0xc6, 0x99, 0x9a, 0x7c, 0x0d, 0x9e, 0xbd, 0xba, 0x0b, 0x32,
	0x6d, 0xe7, 0xd3, 0x1a, 0x74, 0x5f, 0x42, 0xa8, 0x74, 0x14, 0xbc, 0x6a, 0x96, 0x6c, 0x4d, 0x7e,
	0x02, 0x92, 0xcd, 0xd5, 0x53, 0x5e, 0x71, 0xb7, 0x20, 0x1f, 0x2c, 0x43, 0xad, 0x13, 0x5c, 0xb0,
	0x97, 0xbd, 0xa4, 0xc1, 0xf7, 0x60, 0x97, 0x72, 0x99, 0xcc, 0x0f, 0x60, 0x79, 0x7b, 0x76, 0x72,
	0x6f, 0x79, 0x6e, 0x34, 0x79, 0x06, 0x3b, 0x13, 0x94, 0x0a, 0x73, 0xdc, 0x49, 0x98, 0x1d, 0xb9,
	0x71, 0x7e, 0xb8, 0x42, 0xbe, 0xf2, 0x7a, 0xfd, 0x93, 0x79, 0xcd, 0xc5, 0x51, 0xd0, 0x9e, 0x54,
	0xac, 0xce, 0x97, 0xd0, 0xae, 0x46, 0x73, 0x6d, 0xe6, 0xf2, 0x5a, 0x69, 0xc6, 0x7e, 0x5b, 0x02,
	0x8f, 0x72, 0x73, 0x5c, 0xed, 0xa7, 0x67, 0x80, 0x3c, 0x97, 0xda, 0x28, 0xa4, 0x71, 0x65, 0x58,
	0x67, 0x40, 0x22, 0xdb, 0x42, 0x58, 0xf9, 0xda, 0xb9, 0xa1, 0x3d, 0x5c, 0xdf, 0x6a, 0xe0, 0x45,
	0xab, 0xa3, 0xf7, 0xa0, 0xa1, 0x13, 0x66, 0xef, 0x57, 0x33, 0xc8, 0x7f, 0xf6, 0xfe, 0xaa, 0xc1,
	0xfd, 0x13, 0x71, 0x93, 0xfc, 0xcf, 0xc4, 0x3f, 0x42, 0xa7, 0x54, 0x7d, 0xc4, 0x19, 0x26, 0xa6,
	0x2a, 0x21, 0x77, 0x4f, 0xd7, 0x69, 0xdd, 0x77, 0xd5, 0x43, 0x5b, 0x5c, 0xbd, 0xf0, 0x5f, 0x40,
	0x6b, 0xbe, 0x4f, 0x09, 0x73, 0x0a, 0x5a, 0x07, 0x05, 0xe5, 0xda, 0x24, 0xec, 0x72, 0xcb, 0xc6,
	0x1f, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xac, 0x2f, 0x9d, 0xa0, 0x4f, 0x08, 0x00, 0x00,
}