//go get -u google.golang.org/protobuf/cmd/protoc-gen-go && go install google.golang.org/protobuf/cmd/protoc-gen-go && go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

// export PATH="$PATH:$(go env GOPATH)/bin"
// protoc --go_out=./proto --go-grpc_out=./proto ./proto/hello.proto

// npm install grpc_tools_node_protoc_ts --save-dev
// npm install @grpc/grpc-js
// protoc --plugin=protoc-gen-ts=../node_modules/.bin/protoc-gen-ts --ts_out=grpc_js:. hello.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.3
// source: proto/hello.proto

package service

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

type ServerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerStatusRequest) Reset() {
	*x = ServerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusRequest) ProtoMessage() {}

func (x *ServerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatusRequest.ProtoReflect.Descriptor instead.
func (*ServerStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{0}
}

type ServerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uptime         uint64         `protobuf:"varint,1,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Memory         float32        `protobuf:"fixed32,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk           float32        `protobuf:"fixed32,3,opt,name=disk,proto3" json:"disk,omitempty"`
	Cpu            float32        `protobuf:"fixed32,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Load           []float32      `protobuf:"fixed32,5,rep,packed,name=load,proto3" json:"load,omitempty"`
	Network        *NetworkStatus `protobuf:"bytes,6,opt,name=network,proto3" json:"network,omitempty"`
	CurrentNetwork *NetworkStatus `protobuf:"bytes,7,opt,name=current_network,json=currentNetwork,proto3" json:"current_network,omitempty"`
	Ping           uint64         `protobuf:"varint,8,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *ServerStatusResponse) Reset() {
	*x = ServerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusResponse) ProtoMessage() {}

func (x *ServerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatusResponse.ProtoReflect.Descriptor instead.
func (*ServerStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{1}
}

func (x *ServerStatusResponse) GetUptime() uint64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *ServerStatusResponse) GetMemory() float32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *ServerStatusResponse) GetDisk() float32 {
	if x != nil {
		return x.Disk
	}
	return 0
}

func (x *ServerStatusResponse) GetCpu() float32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *ServerStatusResponse) GetLoad() []float32 {
	if x != nil {
		return x.Load
	}
	return nil
}

func (x *ServerStatusResponse) GetNetwork() *NetworkStatus {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *ServerStatusResponse) GetCurrentNetwork() *NetworkStatus {
	if x != nil {
		return x.CurrentNetwork
	}
	return nil
}

func (x *ServerStatusResponse) GetPing() uint64 {
	if x != nil {
		return x.Ping
	}
	return 0
}

type NetworkStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Upload   uint64 `protobuf:"varint,1,opt,name=upload,proto3" json:"upload,omitempty"`
	Download uint64 `protobuf:"varint,2,opt,name=download,proto3" json:"download,omitempty"`
}

func (x *NetworkStatus) Reset() {
	*x = NetworkStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkStatus) ProtoMessage() {}

func (x *NetworkStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkStatus.ProtoReflect.Descriptor instead.
func (*NetworkStatus) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkStatus) GetUpload() uint64 {
	if x != nil {
		return x.Upload
	}
	return 0
}

func (x *NetworkStatus) GetDownload() uint64 {
	if x != nil {
		return x.Download
	}
	return 0
}

var File_proto_hello_proto protoreflect.FileDescriptor

var file_proto_hello_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xf7, 0x01, 0x0a, 0x14, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x28, 0x0a,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x37, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x70, 0x69, 0x6e, 0x67, 0x22, 0x43, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x56, 0x0a, 0x13, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_proto_rawDescOnce sync.Once
	file_proto_hello_proto_rawDescData = file_proto_hello_proto_rawDesc
)

func file_proto_hello_proto_rawDescGZIP() []byte {
	file_proto_hello_proto_rawDescOnce.Do(func() {
		file_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_proto_rawDescData)
	})
	return file_proto_hello_proto_rawDescData
}

var file_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_hello_proto_goTypes = []interface{}{
	(*ServerStatusRequest)(nil),  // 0: ServerStatusRequest
	(*ServerStatusResponse)(nil), // 1: ServerStatusResponse
	(*NetworkStatus)(nil),        // 2: NetworkStatus
}
var file_proto_hello_proto_depIdxs = []int32{
	2, // 0: ServerStatusResponse.network:type_name -> NetworkStatus
	2, // 1: ServerStatusResponse.current_network:type_name -> NetworkStatus
	0, // 2: ServerStatusService.ServerStatus:input_type -> ServerStatusRequest
	1, // 3: ServerStatusService.ServerStatus:output_type -> ServerStatusResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_hello_proto_init() }
func file_proto_hello_proto_init() {
	if File_proto_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatusRequest); i {
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
		file_proto_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatusResponse); i {
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
		file_proto_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkStatus); i {
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
			RawDescriptor: file_proto_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_proto_goTypes,
		DependencyIndexes: file_proto_hello_proto_depIdxs,
		MessageInfos:      file_proto_hello_proto_msgTypes,
	}.Build()
	File_proto_hello_proto = out.File
	file_proto_hello_proto_rawDesc = nil
	file_proto_hello_proto_goTypes = nil
	file_proto_hello_proto_depIdxs = nil
}