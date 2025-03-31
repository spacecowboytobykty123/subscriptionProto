// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: subscription/subscription.proto

package subs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PlanId        int32                  `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubsRequest) Reset() {
	*x = SubsRequest{}
	mi := &file_subscription_subscription_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubsRequest) ProtoMessage() {}

func (x *SubsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubsRequest.ProtoReflect.Descriptor instead.
func (*SubsRequest) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *SubsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SubsRequest) GetPlanId() int32 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

type SubsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SubId         int64                  `protobuf:"varint,1,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubsResponse) Reset() {
	*x = SubsResponse{}
	mi := &file_subscription_subscription_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubsResponse) ProtoMessage() {}

func (x *SubsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubsResponse.ProtoReflect.Descriptor instead.
func (*SubsResponse) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *SubsResponse) GetSubId() int64 {
	if x != nil {
		return x.SubId
	}
	return 0
}

func (x *SubsResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ChangePlanRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NewPlanId     int32                  `protobuf:"varint,2,opt,name=new_plan_id,json=newPlanId,proto3" json:"new_plan_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePlanRequest) Reset() {
	*x = ChangePlanRequest{}
	mi := &file_subscription_subscription_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePlanRequest) ProtoMessage() {}

func (x *ChangePlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePlanRequest.ProtoReflect.Descriptor instead.
func (*ChangePlanRequest) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *ChangePlanRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChangePlanRequest) GetNewPlanId() int32 {
	if x != nil {
		return x.NewPlanId
	}
	return 0
}

type ChangePlanResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int32                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePlanResponse) Reset() {
	*x = ChangePlanResponse{}
	mi := &file_subscription_subscription_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePlanResponse) ProtoMessage() {}

func (x *ChangePlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePlanResponse.ProtoReflect.Descriptor instead.
func (*ChangePlanResponse) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{3}
}

func (x *ChangePlanResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UnSubsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnSubsRequest) Reset() {
	*x = UnSubsRequest{}
	mi := &file_subscription_subscription_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnSubsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnSubsRequest) ProtoMessage() {}

func (x *UnSubsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnSubsRequest.ProtoReflect.Descriptor instead.
func (*UnSubsRequest) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *UnSubsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UnSubsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnSubsResponse) Reset() {
	*x = UnSubsResponse{}
	mi := &file_subscription_subscription_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnSubsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnSubsResponse) ProtoMessage() {}

func (x *UnSubsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnSubsResponse.ProtoReflect.Descriptor instead.
func (*UnSubsResponse) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *UnSubsResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetSubRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSubRequest) Reset() {
	*x = GetSubRequest{}
	mi := &file_subscription_subscription_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubRequest) ProtoMessage() {}

func (x *GetSubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubRequest.ProtoReflect.Descriptor instead.
func (*GetSubRequest) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{6}
}

func (x *GetSubRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetSubResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	UserId         int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PlanId         int32                  `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	PlanName       string                 `protobuf:"bytes,3,opt,name=plan_name,json=planName,proto3" json:"plan_name,omitempty"`
	RemainingLimit int32                  `protobuf:"varint,4,opt,name=remaining_limit,json=remainingLimit,proto3" json:"remaining_limit,omitempty"`
	ExpiresAt      string                 `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetSubResponse) Reset() {
	*x = GetSubResponse{}
	mi := &file_subscription_subscription_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubResponse) ProtoMessage() {}

func (x *GetSubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubResponse.ProtoReflect.Descriptor instead.
func (*GetSubResponse) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{7}
}

func (x *GetSubResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetSubResponse) GetPlanId() int32 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

func (x *GetSubResponse) GetPlanName() string {
	if x != nil {
		return x.PlanName
	}
	return ""
}

func (x *GetSubResponse) GetRemainingLimit() int32 {
	if x != nil {
		return x.RemainingLimit
	}
	return 0
}

func (x *GetSubResponse) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

type CheckSubsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckSubsRequest) Reset() {
	*x = CheckSubsRequest{}
	mi := &file_subscription_subscription_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckSubsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckSubsRequest) ProtoMessage() {}

func (x *CheckSubsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckSubsRequest.ProtoReflect.Descriptor instead.
func (*CheckSubsRequest) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{8}
}

func (x *CheckSubsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckSubsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SubStatus     bool                   `protobuf:"varint,1,opt,name=sub_status,json=subStatus,proto3" json:"sub_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckSubsResponse) Reset() {
	*x = CheckSubsResponse{}
	mi := &file_subscription_subscription_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckSubsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckSubsResponse) ProtoMessage() {}

func (x *CheckSubsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckSubsResponse.ProtoReflect.Descriptor instead.
func (*CheckSubsResponse) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{9}
}

func (x *CheckSubsResponse) GetSubStatus() bool {
	if x != nil {
		return x.SubStatus
	}
	return false
}

type PlansRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlansRequest) Reset() {
	*x = PlansRequest{}
	mi := &file_subscription_subscription_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlansRequest) ProtoMessage() {}

func (x *PlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlansRequest.ProtoReflect.Descriptor instead.
func (*PlansRequest) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{10}
}

type PlansResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Plans         []*Plan                `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlansResponse) Reset() {
	*x = PlansResponse{}
	mi := &file_subscription_subscription_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlansResponse) ProtoMessage() {}

func (x *PlansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlansResponse.ProtoReflect.Descriptor instead.
func (*PlansResponse) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{11}
}

func (x *PlansResponse) GetPlans() []*Plan {
	if x != nil {
		return x.Plans
	}
	return nil
}

type Plan struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlanId        int32                  `protobuf:"varint,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RentalLimit   int32                  `protobuf:"varint,3,opt,name=rental_limit,json=rentalLimit,proto3" json:"rental_limit,omitempty"`
	Price         int32                  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Duration      string                 `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"` // "1/3/6/12 месяцев"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Plan) Reset() {
	*x = Plan{}
	mi := &file_subscription_subscription_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_subscription_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_subscription_subscription_proto_rawDescGZIP(), []int{12}
}

func (x *Plan) GetPlanId() int32 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

func (x *Plan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plan) GetRentalLimit() int32 {
	if x != nil {
		return x.RentalLimit
	}
	return 0
}

func (x *Plan) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Plan) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

var File_subscription_subscription_proto protoreflect.FileDescriptor

var file_subscription_subscription_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3f, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64,
	0x22, 0x3d, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x75, 0x62, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x4c, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x28, 0x0a, 0x0d, 0x55,
	0x6e, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x28, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x41, 0x74, 0x22, 0x2b, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x75, 0x62, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x32, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x75, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x22,
	0x88, 0x01, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xd9, 0x03, 0x0a, 0x0c, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x75, 0x62, 0x73, 0x50, 0x6c, 0x61,
	0x6e, 0x12, 0x1f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x1b, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x1a, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x74, 0x73, 0x75, 0x62, 0x73, 0x2e,
	0x76, 0x31, 0x3b, 0x73, 0x75, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_subscription_subscription_proto_rawDescOnce sync.Once
	file_subscription_subscription_proto_rawDescData []byte
)

func file_subscription_subscription_proto_rawDescGZIP() []byte {
	file_subscription_subscription_proto_rawDescOnce.Do(func() {
		file_subscription_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_subscription_subscription_proto_rawDesc), len(file_subscription_subscription_proto_rawDesc)))
	})
	return file_subscription_subscription_proto_rawDescData
}

var file_subscription_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_subscription_subscription_proto_goTypes = []any{
	(*SubsRequest)(nil),        // 0: subscription.SubsRequest
	(*SubsResponse)(nil),       // 1: subscription.SubsResponse
	(*ChangePlanRequest)(nil),  // 2: subscription.ChangePlanRequest
	(*ChangePlanResponse)(nil), // 3: subscription.ChangePlanResponse
	(*UnSubsRequest)(nil),      // 4: subscription.UnSubsRequest
	(*UnSubsResponse)(nil),     // 5: subscription.UnSubsResponse
	(*GetSubRequest)(nil),      // 6: subscription.GetSubRequest
	(*GetSubResponse)(nil),     // 7: subscription.GetSubResponse
	(*CheckSubsRequest)(nil),   // 8: subscription.CheckSubsRequest
	(*CheckSubsResponse)(nil),  // 9: subscription.CheckSubsResponse
	(*PlansRequest)(nil),       // 10: subscription.PlansRequest
	(*PlansResponse)(nil),      // 11: subscription.PlansResponse
	(*Plan)(nil),               // 12: subscription.Plan
}
var file_subscription_subscription_proto_depIdxs = []int32{
	12, // 0: subscription.PlansResponse.plans:type_name -> subscription.Plan
	0,  // 1: subscription.Subscription.Subscribe:input_type -> subscription.SubsRequest
	2,  // 2: subscription.Subscription.ChangeSubsPlan:input_type -> subscription.ChangePlanRequest
	4,  // 3: subscription.Subscription.Unsubscribe:input_type -> subscription.UnSubsRequest
	6,  // 4: subscription.Subscription.GetSubDetails:input_type -> subscription.GetSubRequest
	8,  // 5: subscription.Subscription.CheckSubscription:input_type -> subscription.CheckSubsRequest
	10, // 6: subscription.Subscription.ListPlans:input_type -> subscription.PlansRequest
	1,  // 7: subscription.Subscription.Subscribe:output_type -> subscription.SubsResponse
	3,  // 8: subscription.Subscription.ChangeSubsPlan:output_type -> subscription.ChangePlanResponse
	5,  // 9: subscription.Subscription.Unsubscribe:output_type -> subscription.UnSubsResponse
	7,  // 10: subscription.Subscription.GetSubDetails:output_type -> subscription.GetSubResponse
	9,  // 11: subscription.Subscription.CheckSubscription:output_type -> subscription.CheckSubsResponse
	11, // 12: subscription.Subscription.ListPlans:output_type -> subscription.PlansResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_subscription_subscription_proto_init() }
func file_subscription_subscription_proto_init() {
	if File_subscription_subscription_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_subscription_subscription_proto_rawDesc), len(file_subscription_subscription_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscription_subscription_proto_goTypes,
		DependencyIndexes: file_subscription_subscription_proto_depIdxs,
		MessageInfos:      file_subscription_subscription_proto_msgTypes,
	}.Build()
	File_subscription_subscription_proto = out.File
	file_subscription_subscription_proto_goTypes = nil
	file_subscription_subscription_proto_depIdxs = nil
}
