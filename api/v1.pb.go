// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: api/v1.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckVulnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets  []string `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`                           // IP address
	TcpPorts []int32  `protobuf:"varint,2,rep,packed,name=tcp_ports,json=tcpPorts,proto3" json:"tcp_ports,omitempty"` // only TCP ports
}

func (x *CheckVulnRequest) Reset() {
	*x = CheckVulnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckVulnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckVulnRequest) ProtoMessage() {}

func (x *CheckVulnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckVulnRequest.ProtoReflect.Descriptor instead.
func (*CheckVulnRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_rawDescGZIP(), []int{0}
}

func (x *CheckVulnRequest) GetTargets() []string {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *CheckVulnRequest) GetTcpPorts() []int32 {
	if x != nil {
		return x.TcpPorts
	}
	return nil
}

type CheckVulnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*TargetResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *CheckVulnResponse) Reset() {
	*x = CheckVulnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckVulnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckVulnResponse) ProtoMessage() {}

func (x *CheckVulnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckVulnResponse.ProtoReflect.Descriptor instead.
func (*CheckVulnResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_rawDescGZIP(), []int{1}
}

func (x *CheckVulnResponse) GetResults() []*TargetResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type TargetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   string     `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"` // target IP
	Services []*Service `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *TargetResult) Reset() {
	*x = TargetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetResult) ProtoMessage() {}

func (x *TargetResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetResult.ProtoReflect.Descriptor instead.
func (*TargetResult) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_rawDescGZIP(), []int{2}
}

func (x *TargetResult) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *TargetResult) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string           `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	TcpPort int32            `protobuf:"varint,3,opt,name=tcp_port,json=tcpPort,proto3" json:"tcp_port,omitempty"`
	Vulns   []*Vulnerability `protobuf:"bytes,4,rep,name=vulns,proto3" json:"vulns,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_rawDescGZIP(), []int{3}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Service) GetTcpPort() int32 {
	if x != nil {
		return x.TcpPort
	}
	return 0
}

func (x *Service) GetVulns() []*Vulnerability {
	if x != nil {
		return x.Vulns
	}
	return nil
}

type Vulnerability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string  `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	CvssScore  float32 `protobuf:"fixed32,2,opt,name=cvss_score,json=cvssScore,proto3" json:"cvss_score,omitempty"`
}

func (x *Vulnerability) Reset() {
	*x = Vulnerability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vulnerability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vulnerability) ProtoMessage() {}

func (x *Vulnerability) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vulnerability.ProtoReflect.Descriptor instead.
func (*Vulnerability) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_rawDescGZIP(), []int{4}
}

func (x *Vulnerability) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Vulnerability) GetCvssScore() float32 {
	if x != nil {
		return x.CvssScore
	}
	return 0
}

var File_api_v1_proto protoreflect.FileDescriptor

var file_api_v1_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x6e, 0x65, 0x74, 0x76, 0x75, 0x6c, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x49, 0x0a, 0x10, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x56, 0x75, 0x6c, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x63, 0x70, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x74, 0x63, 0x70,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x47, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x75,
	0x6c, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65,
	0x74, 0x76, 0x75, 0x6c, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x57,
	0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x65, 0x74, 0x76, 0x75,
	0x6c, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x63, 0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x63, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x05,
	0x76, 0x75, 0x6c, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65,
	0x74, 0x76, 0x75, 0x6c, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x75, 0x6c, 0x6e, 0x73, 0x22, 0x4e, 0x0a,
	0x0d, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x76, 0x73, 0x73, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x63, 0x76, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x32, 0x5c, 0x0a,
	0x0e, 0x4e, 0x65, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x75, 0x6c, 0x6e, 0x12, 0x1c, 0x2e, 0x6e,
	0x65, 0x74, 0x76, 0x75, 0x6c, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56,
	0x75, 0x6c, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x65, 0x74,
	0x76, 0x75, 0x6c, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x75, 0x6c,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_proto_rawDescOnce sync.Once
	file_api_v1_proto_rawDescData = file_api_v1_proto_rawDesc
)

func file_api_v1_proto_rawDescGZIP() []byte {
	file_api_v1_proto_rawDescOnce.Do(func() {
		file_api_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_proto_rawDescData)
	})
	return file_api_v1_proto_rawDescData
}

var file_api_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_proto_goTypes = []interface{}{
	(*CheckVulnRequest)(nil),  // 0: netvuln.v1.CheckVulnRequest
	(*CheckVulnResponse)(nil), // 1: netvuln.v1.CheckVulnResponse
	(*TargetResult)(nil),      // 2: netvuln.v1.TargetResult
	(*Service)(nil),           // 3: netvuln.v1.Service
	(*Vulnerability)(nil),     // 4: netvuln.v1.Vulnerability
}
var file_api_v1_proto_depIdxs = []int32{
	2, // 0: netvuln.v1.CheckVulnResponse.results:type_name -> netvuln.v1.TargetResult
	3, // 1: netvuln.v1.TargetResult.services:type_name -> netvuln.v1.Service
	4, // 2: netvuln.v1.Service.vulns:type_name -> netvuln.v1.Vulnerability
	0, // 3: netvuln.v1.NetVulnService.CheckVuln:input_type -> netvuln.v1.CheckVulnRequest
	1, // 4: netvuln.v1.NetVulnService.CheckVuln:output_type -> netvuln.v1.CheckVulnResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_proto_init() }
func file_api_v1_proto_init() {
	if File_api_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckVulnRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckVulnResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vulnerability); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_proto_goTypes,
		DependencyIndexes: file_api_v1_proto_depIdxs,
		MessageInfos:      file_api_v1_proto_msgTypes,
	}.Build()
	File_api_v1_proto = out.File
	file_api_v1_proto_rawDesc = nil
	file_api_v1_proto_goTypes = nil
	file_api_v1_proto_depIdxs = nil
}
