// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/watchdog/v3alpha/abort_action.proto

package envoy_watchdog_v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A GuardDogAction that will terminate the process by killing the
// stuck thread. This would allow easier access to the call stack of the stuck
// thread since we would run signal handlers on that thread. By default
// this will be registered to run as the last watchdog action on KILL and
// MULTIKILL events if those are enabled.
type AbortActionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How long to wait for the thread to respond to the thread kill function
	// before killing the process from this action. This is a blocking action.
	// By default this is 5 seconds.
	WaitDuration *duration.Duration `protobuf:"bytes,1,opt,name=wait_duration,json=waitDuration,proto3" json:"wait_duration,omitempty"`
}

func (x *AbortActionConfig) Reset() {
	*x = AbortActionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_watchdog_v3alpha_abort_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbortActionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbortActionConfig) ProtoMessage() {}

func (x *AbortActionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_watchdog_v3alpha_abort_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbortActionConfig.ProtoReflect.Descriptor instead.
func (*AbortActionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_watchdog_v3alpha_abort_action_proto_rawDescGZIP(), []int{0}
}

func (x *AbortActionConfig) GetWaitDuration() *duration.Duration {
	if x != nil {
		return x.WaitDuration
	}
	return nil
}

var File_envoy_watchdog_v3alpha_abort_action_proto protoreflect.FileDescriptor

var file_envoy_watchdog_v3alpha_abort_action_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x53, 0x0a, 0x11, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x61, 0x69, 0x74, 0x5f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x61, 0x69, 0x74, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x4a, 0x0a, 0x24, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42,
	0x10, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x08, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_watchdog_v3alpha_abort_action_proto_rawDescOnce sync.Once
	file_envoy_watchdog_v3alpha_abort_action_proto_rawDescData = file_envoy_watchdog_v3alpha_abort_action_proto_rawDesc
)

func file_envoy_watchdog_v3alpha_abort_action_proto_rawDescGZIP() []byte {
	file_envoy_watchdog_v3alpha_abort_action_proto_rawDescOnce.Do(func() {
		file_envoy_watchdog_v3alpha_abort_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_watchdog_v3alpha_abort_action_proto_rawDescData)
	})
	return file_envoy_watchdog_v3alpha_abort_action_proto_rawDescData
}

var file_envoy_watchdog_v3alpha_abort_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_watchdog_v3alpha_abort_action_proto_goTypes = []interface{}{
	(*AbortActionConfig)(nil), // 0: envoy.watchdog.v3alpha.AbortActionConfig
	(*duration.Duration)(nil), // 1: google.protobuf.Duration
}
var file_envoy_watchdog_v3alpha_abort_action_proto_depIdxs = []int32{
	1, // 0: envoy.watchdog.v3alpha.AbortActionConfig.wait_duration:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_watchdog_v3alpha_abort_action_proto_init() }
func file_envoy_watchdog_v3alpha_abort_action_proto_init() {
	if File_envoy_watchdog_v3alpha_abort_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_watchdog_v3alpha_abort_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbortActionConfig); i {
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
			RawDescriptor: file_envoy_watchdog_v3alpha_abort_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_watchdog_v3alpha_abort_action_proto_goTypes,
		DependencyIndexes: file_envoy_watchdog_v3alpha_abort_action_proto_depIdxs,
		MessageInfos:      file_envoy_watchdog_v3alpha_abort_action_proto_msgTypes,
	}.Build()
	File_envoy_watchdog_v3alpha_abort_action_proto = out.File
	file_envoy_watchdog_v3alpha_abort_action_proto_rawDesc = nil
	file_envoy_watchdog_v3alpha_abort_action_proto_goTypes = nil
	file_envoy_watchdog_v3alpha_abort_action_proto_depIdxs = nil
}
