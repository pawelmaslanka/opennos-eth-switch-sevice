// Copyright (c) 2020 Pawel Maslanka
// Contact: pawmas.pawelmaslanka@gmail.com
// This file is subject to the terms and conditions defined in
// file 'LICENSE.txt', which is part of this source code package.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.0-devel
// 	protoc        v3.11.0
// source: platform/port.proto

package platform

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	interfaces "opennos-eth-switch-service/mgmt/interfaces"
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

// Symbols defined in public import of interfaces/ethernet.proto.

type EthernetIntf = interfaces.EthernetIntf
type AddIpv4AddrToEthernetIntfRequest = interfaces.AddIpv4AddrToEthernetIntfRequest
type RemoveIpv4AddrFromEthernetIntfRequest = interfaces.RemoveIpv4AddrFromEthernetIntfRequest

type PortSpeed_Mode int32

const (
	PortSpeed_SPEED_10GB  PortSpeed_Mode = 0
	PortSpeed_SPEED_40GB  PortSpeed_Mode = 1
	PortSpeed_SPEED_100GB PortSpeed_Mode = 2
)

// Enum value maps for PortSpeed_Mode.
var (
	PortSpeed_Mode_name = map[int32]string{
		0: "SPEED_10GB",
		1: "SPEED_40GB",
		2: "SPEED_100GB",
	}
	PortSpeed_Mode_value = map[string]int32{
		"SPEED_10GB":  0,
		"SPEED_40GB":  1,
		"SPEED_100GB": 2,
	}
)

func (x PortSpeed_Mode) Enum() *PortSpeed_Mode {
	p := new(PortSpeed_Mode)
	*p = x
	return p
}

func (x PortSpeed_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PortSpeed_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_port_proto_enumTypes[0].Descriptor()
}

func (PortSpeed_Mode) Type() protoreflect.EnumType {
	return &file_platform_port_proto_enumTypes[0]
}

func (x PortSpeed_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PortSpeed_Mode.Descriptor instead.
func (PortSpeed_Mode) EnumDescriptor() ([]byte, []int) {
	return file_platform_port_proto_rawDescGZIP(), []int{0, 0}
}

type ChannelSpeed_Mode int32

const (
	ChannelSpeed_SPEED_10GB  ChannelSpeed_Mode = 0
	ChannelSpeed_SPEED_100GB ChannelSpeed_Mode = 1
)

// Enum value maps for ChannelSpeed_Mode.
var (
	ChannelSpeed_Mode_name = map[int32]string{
		0: "SPEED_10GB",
		1: "SPEED_100GB",
	}
	ChannelSpeed_Mode_value = map[string]int32{
		"SPEED_10GB":  0,
		"SPEED_100GB": 1,
	}
)

func (x ChannelSpeed_Mode) Enum() *ChannelSpeed_Mode {
	p := new(ChannelSpeed_Mode)
	*p = x
	return p
}

func (x ChannelSpeed_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelSpeed_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_port_proto_enumTypes[1].Descriptor()
}

func (ChannelSpeed_Mode) Type() protoreflect.EnumType {
	return &file_platform_port_proto_enumTypes[1]
}

func (x ChannelSpeed_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelSpeed_Mode.Descriptor instead.
func (ChannelSpeed_Mode) EnumDescriptor() ([]byte, []int) {
	return file_platform_port_proto_rawDescGZIP(), []int{1, 0}
}

type PortBreakoutRequest_NumChannels int32

const (
	PortBreakoutRequest_MODE_1x PortBreakoutRequest_NumChannels = 0
	PortBreakoutRequest_MODE_2x PortBreakoutRequest_NumChannels = 1
	PortBreakoutRequest_MODE_4x PortBreakoutRequest_NumChannels = 2
)

// Enum value maps for PortBreakoutRequest_NumChannels.
var (
	PortBreakoutRequest_NumChannels_name = map[int32]string{
		0: "MODE_1x",
		1: "MODE_2x",
		2: "MODE_4x",
	}
	PortBreakoutRequest_NumChannels_value = map[string]int32{
		"MODE_1x": 0,
		"MODE_2x": 1,
		"MODE_4x": 2,
	}
)

func (x PortBreakoutRequest_NumChannels) Enum() *PortBreakoutRequest_NumChannels {
	p := new(PortBreakoutRequest_NumChannels)
	*p = x
	return p
}

func (x PortBreakoutRequest_NumChannels) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PortBreakoutRequest_NumChannels) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_port_proto_enumTypes[2].Descriptor()
}

func (PortBreakoutRequest_NumChannels) Type() protoreflect.EnumType {
	return &file_platform_port_proto_enumTypes[2]
}

func (x PortBreakoutRequest_NumChannels) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PortBreakoutRequest_NumChannels.Descriptor instead.
func (PortBreakoutRequest_NumChannels) EnumDescriptor() ([]byte, []int) {
	return file_platform_port_proto_rawDescGZIP(), []int{2, 0}
}

type PortSpeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode PortSpeed_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=platform.PortSpeed_Mode" json:"mode,omitempty"`
}

func (x *PortSpeed) Reset() {
	*x = PortSpeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_port_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortSpeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortSpeed) ProtoMessage() {}

func (x *PortSpeed) ProtoReflect() protoreflect.Message {
	mi := &file_platform_port_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortSpeed.ProtoReflect.Descriptor instead.
func (*PortSpeed) Descriptor() ([]byte, []int) {
	return file_platform_port_proto_rawDescGZIP(), []int{0}
}

func (x *PortSpeed) GetMode() PortSpeed_Mode {
	if x != nil {
		return x.Mode
	}
	return PortSpeed_SPEED_10GB
}

type ChannelSpeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode ChannelSpeed_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=platform.ChannelSpeed_Mode" json:"mode,omitempty"`
}

func (x *ChannelSpeed) Reset() {
	*x = ChannelSpeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_port_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelSpeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelSpeed) ProtoMessage() {}

func (x *ChannelSpeed) ProtoReflect() protoreflect.Message {
	mi := &file_platform_port_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelSpeed.ProtoReflect.Descriptor instead.
func (*ChannelSpeed) Descriptor() ([]byte, []int) {
	return file_platform_port_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelSpeed) GetMode() ChannelSpeed_Mode {
	if x != nil {
		return x.Mode
	}
	return ChannelSpeed_SPEED_10GB
}

type PortBreakoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthIntf      *interfaces.EthernetIntf        `protobuf:"bytes,1,opt,name=ethIntf,proto3" json:"ethIntf,omitempty"`
	NumChannels  PortBreakoutRequest_NumChannels `protobuf:"varint,2,opt,name=numChannels,proto3,enum=platform.PortBreakoutRequest_NumChannels" json:"numChannels,omitempty"`
	ChannelSpeed *ChannelSpeed                   `protobuf:"bytes,3,opt,name=channelSpeed,proto3" json:"channelSpeed,omitempty"`
}

func (x *PortBreakoutRequest) Reset() {
	*x = PortBreakoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_port_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortBreakoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortBreakoutRequest) ProtoMessage() {}

func (x *PortBreakoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_port_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortBreakoutRequest.ProtoReflect.Descriptor instead.
func (*PortBreakoutRequest) Descriptor() ([]byte, []int) {
	return file_platform_port_proto_rawDescGZIP(), []int{2}
}

func (x *PortBreakoutRequest) GetEthIntf() *interfaces.EthernetIntf {
	if x != nil {
		return x.EthIntf
	}
	return nil
}

func (x *PortBreakoutRequest) GetNumChannels() PortBreakoutRequest_NumChannels {
	if x != nil {
		return x.NumChannels
	}
	return PortBreakoutRequest_MODE_1x
}

func (x *PortBreakoutRequest) GetChannelSpeed() *ChannelSpeed {
	if x != nil {
		return x.ChannelSpeed
	}
	return nil
}

type PortBreakoutChanSpeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthIntf      *interfaces.EthernetIntf `protobuf:"bytes,1,opt,name=ethIntf,proto3" json:"ethIntf,omitempty"`
	ChannelSpeed *ChannelSpeed            `protobuf:"bytes,2,opt,name=channelSpeed,proto3" json:"channelSpeed,omitempty"`
}

func (x *PortBreakoutChanSpeedRequest) Reset() {
	*x = PortBreakoutChanSpeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_port_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortBreakoutChanSpeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortBreakoutChanSpeedRequest) ProtoMessage() {}

func (x *PortBreakoutChanSpeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_port_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortBreakoutChanSpeedRequest.ProtoReflect.Descriptor instead.
func (*PortBreakoutChanSpeedRequest) Descriptor() ([]byte, []int) {
	return file_platform_port_proto_rawDescGZIP(), []int{3}
}

func (x *PortBreakoutChanSpeedRequest) GetEthIntf() *interfaces.EthernetIntf {
	if x != nil {
		return x.EthIntf
	}
	return nil
}

func (x *PortBreakoutChanSpeedRequest) GetChannelSpeed() *ChannelSpeed {
	if x != nil {
		return x.ChannelSpeed
	}
	return nil
}

var File_platform_port_proto protoreflect.FileDescriptor

var file_platform_port_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x1a,
	0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x09, 0x50, 0x6f,
	0x72, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x37, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x50, 0x45, 0x45, 0x44, 0x5f, 0x31, 0x30, 0x47, 0x42, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x50, 0x45, 0x45, 0x44, 0x5f, 0x34, 0x30, 0x47, 0x42, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x50, 0x45, 0x45, 0x44, 0x5f, 0x31, 0x30, 0x30, 0x47, 0x42, 0x10, 0x02, 0x22, 0x68,
	0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x2f,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22,
	0x27, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x50, 0x45, 0x45, 0x44,
	0x5f, 0x31, 0x30, 0x47, 0x42, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x50, 0x45, 0x45, 0x44,
	0x5f, 0x31, 0x30, 0x30, 0x47, 0x42, 0x10, 0x01, 0x22, 0x88, 0x02, 0x0a, 0x13, 0x50, 0x6f, 0x72,
	0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x07, 0x65, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x66, 0x52, 0x07, 0x65, 0x74, 0x68,
	0x49, 0x6e, 0x74, 0x66, 0x12, 0x4b, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x75, 0x6d, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52,
	0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x65, 0x64, 0x22, 0x34, 0x0a,
	0x0b, 0x4e, 0x75, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x0b, 0x0a, 0x07,
	0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x31, 0x78, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x32, 0x78, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x34,
	0x78, 0x10, 0x02, 0x22, 0x8e, 0x01, 0x0a, 0x1c, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x6f, 0x75, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x65, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x66, 0x52,
	0x07, 0x65, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x66, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x42, 0x2a, 0x5a, 0x28, 0x6f, 0x70, 0x65, 0x6e, 0x6e, 0x6f, 0x73, 0x2d,
	0x65, 0x74, 0x68, 0x2d, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_port_proto_rawDescOnce sync.Once
	file_platform_port_proto_rawDescData = file_platform_port_proto_rawDesc
)

func file_platform_port_proto_rawDescGZIP() []byte {
	file_platform_port_proto_rawDescOnce.Do(func() {
		file_platform_port_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_port_proto_rawDescData)
	})
	return file_platform_port_proto_rawDescData
}

var file_platform_port_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_platform_port_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_platform_port_proto_goTypes = []interface{}{
	(PortSpeed_Mode)(0),                  // 0: platform.PortSpeed.Mode
	(ChannelSpeed_Mode)(0),               // 1: platform.ChannelSpeed.Mode
	(PortBreakoutRequest_NumChannels)(0), // 2: platform.PortBreakoutRequest.NumChannels
	(*PortSpeed)(nil),                    // 3: platform.PortSpeed
	(*ChannelSpeed)(nil),                 // 4: platform.ChannelSpeed
	(*PortBreakoutRequest)(nil),          // 5: platform.PortBreakoutRequest
	(*PortBreakoutChanSpeedRequest)(nil), // 6: platform.PortBreakoutChanSpeedRequest
	(*interfaces.EthernetIntf)(nil),      // 7: interfaces.EthernetIntf
}
var file_platform_port_proto_depIdxs = []int32{
	0, // 0: platform.PortSpeed.mode:type_name -> platform.PortSpeed.Mode
	1, // 1: platform.ChannelSpeed.mode:type_name -> platform.ChannelSpeed.Mode
	7, // 2: platform.PortBreakoutRequest.ethIntf:type_name -> interfaces.EthernetIntf
	2, // 3: platform.PortBreakoutRequest.numChannels:type_name -> platform.PortBreakoutRequest.NumChannels
	4, // 4: platform.PortBreakoutRequest.channelSpeed:type_name -> platform.ChannelSpeed
	7, // 5: platform.PortBreakoutChanSpeedRequest.ethIntf:type_name -> interfaces.EthernetIntf
	4, // 6: platform.PortBreakoutChanSpeedRequest.channelSpeed:type_name -> platform.ChannelSpeed
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_platform_port_proto_init() }
func file_platform_port_proto_init() {
	if File_platform_port_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_port_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortSpeed); i {
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
		file_platform_port_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelSpeed); i {
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
		file_platform_port_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortBreakoutRequest); i {
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
		file_platform_port_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortBreakoutChanSpeedRequest); i {
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
			RawDescriptor: file_platform_port_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_port_proto_goTypes,
		DependencyIndexes: file_platform_port_proto_depIdxs,
		EnumInfos:         file_platform_port_proto_enumTypes,
		MessageInfos:      file_platform_port_proto_msgTypes,
	}.Build()
	File_platform_port_proto = out.File
	file_platform_port_proto_rawDesc = nil
	file_platform_port_proto_goTypes = nil
	file_platform_port_proto_depIdxs = nil
}
