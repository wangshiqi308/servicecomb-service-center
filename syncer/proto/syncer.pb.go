/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: syncer.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	syncer.proto

It has these top-level messages:
	PullRequest
	SyncData
	SyncService
	SyncInstance
	Expansion
	HealthCheck
	MappingEntry
*/
package proto

import (
	fmt "fmt"

	proto1 "github.com/golang/protobuf/proto"

	math "math"

	context "context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type SyncService_Status int32

const (
	SyncService_UNKNOWN SyncService_Status = 0
	SyncService_UP      SyncService_Status = 1
	SyncService_DOWN    SyncService_Status = 2
)

var SyncService_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "UP",
	2: "DOWN",
}
var SyncService_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"UP":      1,
	"DOWN":    2,
}

func (x SyncService_Status) String() string {
	return proto1.EnumName(SyncService_Status_name, int32(x))
}
func (SyncService_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type SyncInstance_Status int32

const (
	SyncInstance_UNKNOWN      SyncInstance_Status = 0
	SyncInstance_UP           SyncInstance_Status = 1
	SyncInstance_DOWN         SyncInstance_Status = 2
	SyncInstance_STARTING     SyncInstance_Status = 3
	SyncInstance_OUTOFSERVICE SyncInstance_Status = 4
)

var SyncInstance_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "UP",
	2: "DOWN",
	3: "STARTING",
	4: "OUTOFSERVICE",
}
var SyncInstance_Status_value = map[string]int32{
	"UNKNOWN":      0,
	"UP":           1,
	"DOWN":         2,
	"STARTING":     3,
	"OUTOFSERVICE": 4,
}

func (x SyncInstance_Status) String() string {
	return proto1.EnumName(SyncInstance_Status_name, int32(x))
}
func (SyncInstance_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type HealthCheck_Modes int32

const (
	HealthCheck_UNKNOWN HealthCheck_Modes = 0
	HealthCheck_PUSH    HealthCheck_Modes = 1
	HealthCheck_PULL    HealthCheck_Modes = 2
)

var HealthCheck_Modes_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUSH",
	2: "PULL",
}
var HealthCheck_Modes_value = map[string]int32{
	"UNKNOWN": 0,
	"PUSH":    1,
	"PULL":    2,
}

func (x HealthCheck_Modes) String() string {
	return proto1.EnumName(HealthCheck_Modes_name, int32(x))
}
func (HealthCheck_Modes) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type PullRequest struct {
	ServiceName string `protobuf:"bytes,1,opt,name=serviceName" json:"serviceName,omitempty"`
	Options     string `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	Time        string `protobuf:"bytes,3,opt,name=time" json:"time,omitempty"`
}

func (m *PullRequest) Reset()                    { *m = PullRequest{} }
func (m *PullRequest) String() string            { return proto1.CompactTextString(m) }
func (*PullRequest) ProtoMessage()               {}
func (*PullRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PullRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *PullRequest) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *PullRequest) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type SyncData struct {
	Services  []*SyncService  `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
	Instances []*SyncInstance `protobuf:"bytes,2,rep,name=instances" json:"instances,omitempty"`
}

func (m *SyncData) Reset()                    { *m = SyncData{} }
func (m *SyncData) String() string            { return proto1.CompactTextString(m) }
func (*SyncData) ProtoMessage()               {}
func (*SyncData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SyncData) GetServices() []*SyncService {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *SyncData) GetInstances() []*SyncInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type SyncService struct {
	ServiceId     string             `protobuf:"bytes,1,opt,name=serviceId" json:"serviceId,omitempty"`
	App           string             `protobuf:"bytes,2,opt,name=app" json:"app,omitempty"`
	Name          string             `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Version       string             `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	Status        SyncService_Status `protobuf:"varint,5,opt,name=status,enum=proto.SyncService_Status" json:"status,omitempty"`
	DomainProject string             `protobuf:"bytes,6,opt,name=domainProject" json:"domainProject,omitempty"`
	Environment   string             `protobuf:"bytes,7,opt,name=environment" json:"environment,omitempty"`
	PluginName    string             `protobuf:"bytes,8,opt,name=pluginName" json:"pluginName,omitempty"`
	Expansions    []*Expansion       `protobuf:"bytes,9,rep,name=expansions" json:"expansions,omitempty"`
}

func (m *SyncService) Reset()                    { *m = SyncService{} }
func (m *SyncService) String() string            { return proto1.CompactTextString(m) }
func (*SyncService) ProtoMessage()               {}
func (*SyncService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SyncService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *SyncService) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *SyncService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SyncService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SyncService) GetStatus() SyncService_Status {
	if m != nil {
		return m.Status
	}
	return SyncService_UNKNOWN
}

func (m *SyncService) GetDomainProject() string {
	if m != nil {
		return m.DomainProject
	}
	return ""
}

func (m *SyncService) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *SyncService) GetPluginName() string {
	if m != nil {
		return m.PluginName
	}
	return ""
}

func (m *SyncService) GetExpansions() []*Expansion {
	if m != nil {
		return m.Expansions
	}
	return nil
}

type SyncInstance struct {
	InstanceId  string              `protobuf:"bytes,1,opt,name=instanceId" json:"instanceId,omitempty"`
	ServiceId   string              `protobuf:"bytes,2,opt,name=serviceId" json:"serviceId,omitempty"`
	Endpoints   []string            `protobuf:"bytes,3,rep,name=endpoints" json:"endpoints,omitempty"`
	HostName    string              `protobuf:"bytes,4,opt,name=hostName" json:"hostName,omitempty"`
	Status      SyncInstance_Status `protobuf:"varint,5,opt,name=status,enum=proto.SyncInstance_Status" json:"status,omitempty"`
	HealthCheck *HealthCheck        `protobuf:"bytes,6,opt,name=healthCheck" json:"healthCheck,omitempty"`
	Version     string              `protobuf:"bytes,7,opt,name=version" json:"version,omitempty"`
	PluginName  string              `protobuf:"bytes,8,opt,name=pluginName" json:"pluginName,omitempty"`
	Expansions  []*Expansion        `protobuf:"bytes,9,rep,name=expansions" json:"expansions,omitempty"`
}

func (m *SyncInstance) Reset()                    { *m = SyncInstance{} }
func (m *SyncInstance) String() string            { return proto1.CompactTextString(m) }
func (*SyncInstance) ProtoMessage()               {}
func (*SyncInstance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SyncInstance) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *SyncInstance) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *SyncInstance) GetEndpoints() []string {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *SyncInstance) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *SyncInstance) GetStatus() SyncInstance_Status {
	if m != nil {
		return m.Status
	}
	return SyncInstance_UNKNOWN
}

func (m *SyncInstance) GetHealthCheck() *HealthCheck {
	if m != nil {
		return m.HealthCheck
	}
	return nil
}

func (m *SyncInstance) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SyncInstance) GetPluginName() string {
	if m != nil {
		return m.PluginName
	}
	return ""
}

func (m *SyncInstance) GetExpansions() []*Expansion {
	if m != nil {
		return m.Expansions
	}
	return nil
}

type Expansion struct {
	Kind   string            `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Bytes  []byte            `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Expansion) Reset()                    { *m = Expansion{} }
func (m *Expansion) String() string            { return proto1.CompactTextString(m) }
func (*Expansion) ProtoMessage()               {}
func (*Expansion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Expansion) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Expansion) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *Expansion) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type HealthCheck struct {
	Mode     HealthCheck_Modes `protobuf:"varint,1,opt,name=mode,enum=proto.HealthCheck_Modes" json:"mode,omitempty"`
	Port     int32             `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Interval int32             `protobuf:"varint,3,opt,name=interval" json:"interval,omitempty"`
	Times    int32             `protobuf:"varint,4,opt,name=times" json:"times,omitempty"`
	Url      string            `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
}

func (m *HealthCheck) Reset()                    { *m = HealthCheck{} }
func (m *HealthCheck) String() string            { return proto1.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()               {}
func (*HealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HealthCheck) GetMode() HealthCheck_Modes {
	if m != nil {
		return m.Mode
	}
	return HealthCheck_UNKNOWN
}

func (m *HealthCheck) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *HealthCheck) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *HealthCheck) GetTimes() int32 {
	if m != nil {
		return m.Times
	}
	return 0
}

func (m *HealthCheck) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type MappingEntry struct {
	ClusterName string `protobuf:"bytes,1,opt,name=clusterName" json:"clusterName,omitempty"`
	//    Tenant tenant = 2;
	DomainProject string `protobuf:"bytes,2,opt,name=domainProject" json:"domainProject,omitempty"`
	OrgServiceID  string `protobuf:"bytes,3,opt,name=orgServiceID" json:"orgServiceID,omitempty"`
	OrgInstanceID string `protobuf:"bytes,4,opt,name=orgInstanceID" json:"orgInstanceID,omitempty"`
	CurServiceID  string `protobuf:"bytes,5,opt,name=curServiceID" json:"curServiceID,omitempty"`
	CurInstanceID string `protobuf:"bytes,6,opt,name=curInstanceID" json:"curInstanceID,omitempty"`
}

func (m *MappingEntry) Reset()                    { *m = MappingEntry{} }
func (m *MappingEntry) String() string            { return proto1.CompactTextString(m) }
func (*MappingEntry) ProtoMessage()               {}
func (*MappingEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MappingEntry) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *MappingEntry) GetDomainProject() string {
	if m != nil {
		return m.DomainProject
	}
	return ""
}

func (m *MappingEntry) GetOrgServiceID() string {
	if m != nil {
		return m.OrgServiceID
	}
	return ""
}

func (m *MappingEntry) GetOrgInstanceID() string {
	if m != nil {
		return m.OrgInstanceID
	}
	return ""
}

func (m *MappingEntry) GetCurServiceID() string {
	if m != nil {
		return m.CurServiceID
	}
	return ""
}

func (m *MappingEntry) GetCurInstanceID() string {
	if m != nil {
		return m.CurInstanceID
	}
	return ""
}

func init() {
	proto1.RegisterType((*PullRequest)(nil), "proto.PullRequest")
	proto1.RegisterType((*SyncData)(nil), "proto.SyncData")
	proto1.RegisterType((*SyncService)(nil), "proto.SyncService")
	proto1.RegisterType((*SyncInstance)(nil), "proto.SyncInstance")
	proto1.RegisterType((*Expansion)(nil), "proto.Expansion")
	proto1.RegisterType((*HealthCheck)(nil), "proto.HealthCheck")
	proto1.RegisterType((*MappingEntry)(nil), "proto.MappingEntry")
	proto1.RegisterEnum("proto.SyncService_Status", SyncService_Status_name, SyncService_Status_value)
	proto1.RegisterEnum("proto.SyncInstance_Status", SyncInstance_Status_name, SyncInstance_Status_value)
	proto1.RegisterEnum("proto.HealthCheck_Modes", HealthCheck_Modes_name, HealthCheck_Modes_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Sync service

type SyncClient interface {
	Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*SyncData, error)
}

type syncClient struct {
	cc *grpc.ClientConn
}

func NewSyncClient(cc *grpc.ClientConn) SyncClient {
	return &syncClient{cc}
}

func (c *syncClient) Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*SyncData, error) {
	out := new(SyncData)
	err := grpc.Invoke(ctx, "/proto.Sync/Pull", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sync service

type SyncServer interface {
	Pull(context.Context, *PullRequest) (*SyncData, error)
}

func RegisterSyncServer(s *grpc.Server, srv SyncServer) {
	s.RegisterService(&_Sync_serviceDesc, srv)
}

func _Sync_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sync/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServer).Pull(ctx, req.(*PullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Sync",
	HandlerType: (*SyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pull",
			Handler:    _Sync_Pull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syncer.proto",
}

func init() { proto1.RegisterFile("syncer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xdb, 0x48,
	0x0c, 0x8e, 0x64, 0xf9, 0x8f, 0xf2, 0x66, 0x85, 0xd9, 0x3d, 0x68, 0x8d, 0x60, 0x61, 0x08, 0x0b,
	0xac, 0x0f, 0xbb, 0x46, 0xe3, 0xa6, 0x40, 0xdb, 0x5b, 0x11, 0xbb, 0x89, 0xd1, 0xc4, 0x31, 0xc6,
	0x71, 0x7b, 0xea, 0x41, 0x91, 0x07, 0xb6, 0x1a, 0x79, 0xa4, 0x6a, 0x46, 0x46, 0xfd, 0x40, 0xed,
	0x5b, 0xf4, 0x61, 0x7a, 0xed, 0x53, 0x14, 0xf3, 0x63, 0x7b, 0xe4, 0xe4, 0xd0, 0x43, 0x4f, 0x26,
	0x3f, 0x52, 0x1c, 0x92, 0x1f, 0x49, 0x43, 0x8b, 0x6d, 0x68, 0x44, 0xf2, 0x5e, 0x96, 0xa7, 0x3c,
	0x45, 0x55, 0xf9, 0x13, 0xbc, 0x07, 0x77, 0x52, 0x24, 0x09, 0x26, 0x1f, 0x0b, 0xc2, 0x38, 0xea,
	0x80, 0xcb, 0x48, 0xbe, 0x8e, 0x23, 0x32, 0x0e, 0x57, 0xc4, 0xb7, 0x3a, 0x56, 0xb7, 0x89, 0x4d,
	0x08, 0xf9, 0x50, 0x4f, 0x33, 0x1e, 0xa7, 0x94, 0xf9, 0xb6, 0xb4, 0x6e, 0x55, 0x84, 0xc0, 0xe1,
	0xf1, 0x8a, 0xf8, 0x15, 0x09, 0x4b, 0x39, 0x58, 0x41, 0x63, 0xba, 0xa1, 0xd1, 0x20, 0xe4, 0x21,
	0xea, 0x41, 0x43, 0x07, 0x62, 0xbe, 0xd5, 0xa9, 0x74, 0xdd, 0x3e, 0x52, 0xb9, 0xf4, 0x84, 0xcb,
	0x54, 0x99, 0xf0, 0xce, 0x07, 0x9d, 0x42, 0x33, 0xa6, 0x8c, 0x87, 0x54, 0x7c, 0x60, 0xcb, 0x0f,
	0xfe, 0x30, 0x3e, 0x18, 0x69, 0x1b, 0xde, 0x7b, 0x05, 0xdf, 0x6c, 0x70, 0x8d, 0x60, 0xe8, 0x04,
	0x9a, 0x3a, 0xdc, 0x68, 0xae, 0x8b, 0xd9, 0x03, 0xc8, 0x83, 0x4a, 0x98, 0x65, 0xba, 0x0c, 0x21,
	0x8a, 0x12, 0x68, 0xb8, 0x2f, 0x81, 0xea, 0x82, 0xd7, 0x24, 0x67, 0x71, 0x4a, 0x7d, 0x47, 0x15,
	0xac, 0x55, 0x74, 0x0a, 0x35, 0xc6, 0x43, 0x5e, 0x30, 0xbf, 0xda, 0xb1, 0xba, 0xc7, 0xfd, 0xbf,
	0x1e, 0x96, 0xd3, 0x9b, 0x4a, 0x07, 0xac, 0x1d, 0xd1, 0x3f, 0xf0, 0xdb, 0x3c, 0x5d, 0x85, 0x31,
	0x9d, 0xe4, 0xe9, 0x07, 0x12, 0x71, 0xbf, 0x26, 0x43, 0x96, 0x41, 0xc1, 0x02, 0xa1, 0xeb, 0x38,
	0x4f, 0xe9, 0x8a, 0x50, 0xee, 0xd7, 0x15, 0x0b, 0x06, 0x84, 0xfe, 0x06, 0xc8, 0x92, 0x62, 0x11,
	0x53, 0x49, 0x53, 0x43, 0x3a, 0x18, 0x08, 0x7a, 0x02, 0x40, 0x3e, 0x65, 0x21, 0x65, 0x92, 0xa8,
	0xa6, 0x6c, 0x9e, 0xa7, 0xd3, 0x1b, 0x6e, 0x0d, 0xd8, 0xf0, 0x09, 0xfe, 0x85, 0x9a, 0xca, 0x15,
	0xb9, 0x50, 0x9f, 0x8d, 0xdf, 0x8c, 0x6f, 0xde, 0x8d, 0xbd, 0x23, 0x54, 0x03, 0x7b, 0x36, 0xf1,
	0x2c, 0xd4, 0x00, 0x67, 0x20, 0x10, 0x3b, 0xf8, 0x5c, 0x81, 0x96, 0xd9, 0x7f, 0x91, 0xcb, 0x96,
	0x81, 0x5d, 0x97, 0x0d, 0xa4, 0x4c, 0x82, 0x7d, 0x48, 0xc2, 0x09, 0x34, 0x09, 0x9d, 0x67, 0x69,
	0x4c, 0x39, 0xf3, 0x2b, 0x9d, 0x8a, 0xb0, 0xee, 0x00, 0xd4, 0x86, 0xc6, 0x32, 0x65, 0x5c, 0x56,
	0xa9, 0xba, 0xbf, 0xd3, 0x51, 0xff, 0xa0, 0xfd, 0xed, 0x47, 0x86, 0xe3, 0xb0, 0xff, 0x67, 0xe0,
	0x2e, 0x49, 0x98, 0xf0, 0xe5, 0xf9, 0x92, 0x44, 0xf7, 0xb2, 0xfb, 0xfb, 0x31, 0xbc, 0xdc, 0x5b,
	0xb0, 0xe9, 0x66, 0x8e, 0x40, 0xbd, 0x3c, 0x02, 0xbf, 0x9e, 0x87, 0x8b, 0x9f, 0xe4, 0x01, 0xb5,
	0xa0, 0x31, 0xbd, 0x7d, 0x85, 0x6f, 0x47, 0xe3, 0x0b, 0xaf, 0x82, 0x3c, 0x68, 0xdd, 0xcc, 0x6e,
	0x6f, 0x5e, 0x4f, 0x87, 0xf8, 0xed, 0xe8, 0x7c, 0xe8, 0x39, 0xc1, 0x17, 0x0b, 0x9a, 0xbb, 0x27,
	0xc4, 0x64, 0xdf, 0xc7, 0x74, 0x4b, 0x8f, 0x94, 0xd1, 0x9f, 0x50, 0xbd, 0xdb, 0x70, 0xa2, 0x16,
	0xb9, 0x85, 0x95, 0x82, 0xce, 0xa0, 0x96, 0x84, 0x77, 0x24, 0x51, 0x6c, 0xb8, 0xfd, 0x93, 0xc3,
	0x74, 0x7b, 0x57, 0xd2, 0x3c, 0xa4, 0x3c, 0xdf, 0x60, 0xed, 0xdb, 0x7e, 0x01, 0xae, 0x01, 0x8b,
	0xd5, 0xba, 0x27, 0x1b, 0xfd, 0x9a, 0x10, 0xc5, 0x63, 0xeb, 0x30, 0x29, 0x88, 0x9e, 0x00, 0xa5,
	0xbc, 0xb4, 0x9f, 0x5b, 0xc1, 0x57, 0x0b, 0x5c, 0xa3, 0xf5, 0xe8, 0x3f, 0x70, 0x56, 0xe9, 0x5c,
	0x1d, 0x9f, 0xe3, 0xbe, 0xff, 0x90, 0x9c, 0xde, 0x75, 0x3a, 0x27, 0x0c, 0x4b, 0x2f, 0x51, 0x58,
	0x96, 0xe6, 0x5c, 0x86, 0xad, 0x62, 0x29, 0x8b, 0xa9, 0x89, 0x29, 0x27, 0xf9, 0x3a, 0x4c, 0xe4,
	0x2a, 0x57, 0xf1, 0x4e, 0x17, 0x79, 0x88, 0xcb, 0xc4, 0xe4, 0x38, 0x55, 0xb1, 0x52, 0x44, 0xbe,
	0x45, 0x9e, 0xc8, 0x41, 0x6a, 0x62, 0x21, 0x06, 0x5d, 0xa8, 0xca, 0x67, 0xca, 0x34, 0x34, 0xc0,
	0x99, 0xcc, 0xa6, 0x97, 0x8a, 0x88, 0xc9, 0xec, 0xea, 0xca, 0xb3, 0x83, 0xef, 0x16, 0xb4, 0xae,
	0xc3, 0x2c, 0x8b, 0xe9, 0x42, 0x15, 0xdf, 0x01, 0x37, 0x4a, 0x0a, 0xc6, 0x49, 0x6e, 0x1e, 0x51,
	0x03, 0x7a, 0x78, 0x06, 0xec, 0xc7, 0xce, 0x40, 0x00, 0xad, 0x34, 0x5f, 0xe8, 0x4b, 0x32, 0x1a,
	0xe8, 0xab, 0x54, 0xc2, 0x44, 0xa4, 0x34, 0x5f, 0x6c, 0xc7, 0x7d, 0x34, 0xd0, 0x5b, 0x52, 0x06,
	0x45, 0xa4, 0xa8, 0xc8, 0xf7, 0x91, 0x54, 0x9d, 0x25, 0x4c, 0x44, 0x8a, 0x8a, 0xdc, 0x88, 0xa4,
	0x4f, 0x53, 0x09, 0xec, 0x3f, 0x03, 0x47, 0xec, 0x17, 0xfa, 0x1f, 0x1c, 0xf1, 0xbf, 0x81, 0xb6,
	0xbb, 0x63, 0xfc, 0x89, 0xb4, 0x7f, 0x37, 0x16, 0x51, 0x5c, 0xfe, 0xe0, 0xe8, 0xae, 0x26, 0x91,
	0xa7, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0xe0, 0xda, 0xd0, 0x84, 0x06, 0x00, 0x00,
}
