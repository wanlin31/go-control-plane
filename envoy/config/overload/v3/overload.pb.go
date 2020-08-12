// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/config/overload/v3/overload.proto

package envoy_config_overload_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type ResourceMonitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to ConfigType:
	//	*ResourceMonitor_TypedConfig
	//	*ResourceMonitor_HiddenEnvoyDeprecatedConfig
	ConfigType isResourceMonitor_ConfigType `protobuf_oneof:"config_type"`
}

func (x *ResourceMonitor) Reset() {
	*x = ResourceMonitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceMonitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceMonitor) ProtoMessage() {}

func (x *ResourceMonitor) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceMonitor.ProtoReflect.Descriptor instead.
func (*ResourceMonitor) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v3_overload_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceMonitor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ResourceMonitor) GetConfigType() isResourceMonitor_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *ResourceMonitor) GetTypedConfig() *any.Any {
	if x, ok := x.GetConfigType().(*ResourceMonitor_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// Deprecated: Do not use.
func (x *ResourceMonitor) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := x.GetConfigType().(*ResourceMonitor_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

type isResourceMonitor_ConfigType interface {
	isResourceMonitor_ConfigType()
}

type ResourceMonitor_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

type ResourceMonitor_HiddenEnvoyDeprecatedConfig struct {
	// Deprecated: Do not use.
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

func (*ResourceMonitor_TypedConfig) isResourceMonitor_ConfigType() {}

func (*ResourceMonitor_HiddenEnvoyDeprecatedConfig) isResourceMonitor_ConfigType() {}

type ThresholdTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ThresholdTrigger) Reset() {
	*x = ThresholdTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThresholdTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThresholdTrigger) ProtoMessage() {}

func (x *ThresholdTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThresholdTrigger.ProtoReflect.Descriptor instead.
func (*ThresholdTrigger) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v3_overload_proto_rawDescGZIP(), []int{1}
}

func (x *ThresholdTrigger) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ScaledTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScalingThreshold    float64 `protobuf:"fixed64,1,opt,name=scaling_threshold,json=scalingThreshold,proto3" json:"scaling_threshold,omitempty"`
	SaturationThreshold float64 `protobuf:"fixed64,2,opt,name=saturation_threshold,json=saturationThreshold,proto3" json:"saturation_threshold,omitempty"`
}

func (x *ScaledTrigger) Reset() {
	*x = ScaledTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScaledTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaledTrigger) ProtoMessage() {}

func (x *ScaledTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaledTrigger.ProtoReflect.Descriptor instead.
func (*ScaledTrigger) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v3_overload_proto_rawDescGZIP(), []int{2}
}

func (x *ScaledTrigger) GetScalingThreshold() float64 {
	if x != nil {
		return x.ScalingThreshold
	}
	return 0
}

func (x *ScaledTrigger) GetSaturationThreshold() float64 {
	if x != nil {
		return x.SaturationThreshold
	}
	return 0
}

type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to TriggerOneof:
	//	*Trigger_Threshold
	//	*Trigger_Scaled
	TriggerOneof isTrigger_TriggerOneof `protobuf_oneof:"trigger_oneof"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v3_overload_proto_rawDescGZIP(), []int{3}
}

func (x *Trigger) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *Trigger) GetTriggerOneof() isTrigger_TriggerOneof {
	if m != nil {
		return m.TriggerOneof
	}
	return nil
}

func (x *Trigger) GetThreshold() *ThresholdTrigger {
	if x, ok := x.GetTriggerOneof().(*Trigger_Threshold); ok {
		return x.Threshold
	}
	return nil
}

func (x *Trigger) GetScaled() *ScaledTrigger {
	if x, ok := x.GetTriggerOneof().(*Trigger_Scaled); ok {
		return x.Scaled
	}
	return nil
}

type isTrigger_TriggerOneof interface {
	isTrigger_TriggerOneof()
}

type Trigger_Threshold struct {
	Threshold *ThresholdTrigger `protobuf:"bytes,2,opt,name=threshold,proto3,oneof"`
}

type Trigger_Scaled struct {
	Scaled *ScaledTrigger `protobuf:"bytes,3,opt,name=scaled,proto3,oneof"`
}

func (*Trigger_Threshold) isTrigger_TriggerOneof() {}

func (*Trigger_Scaled) isTrigger_TriggerOneof() {}

type OverloadAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Triggers []*Trigger `protobuf:"bytes,2,rep,name=triggers,proto3" json:"triggers,omitempty"`
}

func (x *OverloadAction) Reset() {
	*x = OverloadAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverloadAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverloadAction) ProtoMessage() {}

func (x *OverloadAction) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverloadAction.ProtoReflect.Descriptor instead.
func (*OverloadAction) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v3_overload_proto_rawDescGZIP(), []int{4}
}

func (x *OverloadAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OverloadAction) GetTriggers() []*Trigger {
	if x != nil {
		return x.Triggers
	}
	return nil
}

type OverloadManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshInterval  *duration.Duration `protobuf:"bytes,1,opt,name=refresh_interval,json=refreshInterval,proto3" json:"refresh_interval,omitempty"`
	ResourceMonitors []*ResourceMonitor `protobuf:"bytes,2,rep,name=resource_monitors,json=resourceMonitors,proto3" json:"resource_monitors,omitempty"`
	Actions          []*OverloadAction  `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *OverloadManager) Reset() {
	*x = OverloadManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverloadManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverloadManager) ProtoMessage() {}

func (x *OverloadManager) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v3_overload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverloadManager.ProtoReflect.Descriptor instead.
func (*OverloadManager) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v3_overload_proto_rawDescGZIP(), []int{5}
}

func (x *OverloadManager) GetRefreshInterval() *duration.Duration {
	if x != nil {
		return x.RefreshInterval
	}
	return nil
}

func (x *OverloadManager) GetResourceMonitors() []*ResourceMonitor {
	if x != nil {
		return x.ResourceMonitors
	}
	return nil
}

func (x *OverloadManager) GetActions() []*OverloadAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_envoy_config_overload_v3_overload_proto protoreflect.FileDescriptor

var file_envoy_config_overload_v3_overload_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6f,
	0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x76, 0x33, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x62, 0x0a, 0x1e, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x5f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x42, 0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x1b, 0x68, 0x69, 0x64, 0x64,
	0x65, 0x6e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x34, 0x9a, 0xc5, 0x88, 0x1e, 0x2f, 0x0a, 0x2d,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65,
	0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x42, 0x0d, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x78, 0x0a, 0x10,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x42,
	0x17, 0xfa, 0x42, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x29,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x35, 0x9a, 0xc5, 0x88, 0x1e, 0x30, 0x0a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x64, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x11, 0x73, 0x63, 0x61, 0x6c,
	0x69, 0x6e, 0x67, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x42, 0x17, 0xfa, 0x42, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0xf0, 0x3f, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x10, 0x73, 0x63,
	0x61, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x4a,
	0x0a, 0x14, 0x73, 0x61, 0x74, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xfa, 0x42,
	0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x29, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x13, 0x73, 0x61, 0x74, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x22, 0xf9, 0x01, 0x0a, 0x07, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x33, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12,
	0x41, 0x0a, 0x06, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f,
	0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x64, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x64, 0x3a, 0x2c, 0x9a, 0xc5, 0x88, 0x1e, 0x27, 0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x42, 0x14, 0x0a, 0x0d, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xab, 0x01, 0x0a, 0x0e, 0x4f, 0x76, 0x65, 0x72, 0x6c,
	0x6f, 0x61, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x3a,
	0x33, 0x9a, 0xc5, 0x88, 0x1e, 0x2e, 0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb3, 0x02, 0x0a, 0x0f, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61,
	0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x10, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x60,
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x10,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x42, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x76, 0x65,
	0x72, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x34, 0x9a, 0xc5, 0x88, 0x1e, 0x2f, 0x0a, 0x2d, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c,
	0x6f, 0x61, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0x41, 0x0a, 0x26, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_overload_v3_overload_proto_rawDescOnce sync.Once
	file_envoy_config_overload_v3_overload_proto_rawDescData = file_envoy_config_overload_v3_overload_proto_rawDesc
)

func file_envoy_config_overload_v3_overload_proto_rawDescGZIP() []byte {
	file_envoy_config_overload_v3_overload_proto_rawDescOnce.Do(func() {
		file_envoy_config_overload_v3_overload_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_overload_v3_overload_proto_rawDescData)
	})
	return file_envoy_config_overload_v3_overload_proto_rawDescData
}

var file_envoy_config_overload_v3_overload_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_envoy_config_overload_v3_overload_proto_goTypes = []interface{}{
	(*ResourceMonitor)(nil),   // 0: envoy.config.overload.v3.ResourceMonitor
	(*ThresholdTrigger)(nil),  // 1: envoy.config.overload.v3.ThresholdTrigger
	(*ScaledTrigger)(nil),     // 2: envoy.config.overload.v3.ScaledTrigger
	(*Trigger)(nil),           // 3: envoy.config.overload.v3.Trigger
	(*OverloadAction)(nil),    // 4: envoy.config.overload.v3.OverloadAction
	(*OverloadManager)(nil),   // 5: envoy.config.overload.v3.OverloadManager
	(*any.Any)(nil),           // 6: google.protobuf.Any
	(*_struct.Struct)(nil),    // 7: google.protobuf.Struct
	(*duration.Duration)(nil), // 8: google.protobuf.Duration
}
var file_envoy_config_overload_v3_overload_proto_depIdxs = []int32{
	6, // 0: envoy.config.overload.v3.ResourceMonitor.typed_config:type_name -> google.protobuf.Any
	7, // 1: envoy.config.overload.v3.ResourceMonitor.hidden_envoy_deprecated_config:type_name -> google.protobuf.Struct
	1, // 2: envoy.config.overload.v3.Trigger.threshold:type_name -> envoy.config.overload.v3.ThresholdTrigger
	2, // 3: envoy.config.overload.v3.Trigger.scaled:type_name -> envoy.config.overload.v3.ScaledTrigger
	3, // 4: envoy.config.overload.v3.OverloadAction.triggers:type_name -> envoy.config.overload.v3.Trigger
	8, // 5: envoy.config.overload.v3.OverloadManager.refresh_interval:type_name -> google.protobuf.Duration
	0, // 6: envoy.config.overload.v3.OverloadManager.resource_monitors:type_name -> envoy.config.overload.v3.ResourceMonitor
	4, // 7: envoy.config.overload.v3.OverloadManager.actions:type_name -> envoy.config.overload.v3.OverloadAction
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_envoy_config_overload_v3_overload_proto_init() }
func file_envoy_config_overload_v3_overload_proto_init() {
	if File_envoy_config_overload_v3_overload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_overload_v3_overload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceMonitor); i {
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
		file_envoy_config_overload_v3_overload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThresholdTrigger); i {
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
		file_envoy_config_overload_v3_overload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScaledTrigger); i {
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
		file_envoy_config_overload_v3_overload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger); i {
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
		file_envoy_config_overload_v3_overload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverloadAction); i {
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
		file_envoy_config_overload_v3_overload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverloadManager); i {
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
	file_envoy_config_overload_v3_overload_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ResourceMonitor_TypedConfig)(nil),
		(*ResourceMonitor_HiddenEnvoyDeprecatedConfig)(nil),
	}
	file_envoy_config_overload_v3_overload_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Trigger_Threshold)(nil),
		(*Trigger_Scaled)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_overload_v3_overload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_overload_v3_overload_proto_goTypes,
		DependencyIndexes: file_envoy_config_overload_v3_overload_proto_depIdxs,
		MessageInfos:      file_envoy_config_overload_v3_overload_proto_msgTypes,
	}.Build()
	File_envoy_config_overload_v3_overload_proto = out.File
	file_envoy_config_overload_v3_overload_proto_rawDesc = nil
	file_envoy_config_overload_v3_overload_proto_goTypes = nil
	file_envoy_config_overload_v3_overload_proto_depIdxs = nil
}
