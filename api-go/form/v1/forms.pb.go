// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: form/v1/forms.proto

package form

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Question_Type int32

const (
	Question_TYPE_UNSPECIFIED Question_Type = 0
	Question_TYPE_TEXT        Question_Type = 1
	Question_TYPE_RADIO       Question_Type = 2
	Question_TYPE_CHECKBOX    Question_Type = 3
)

// Enum value maps for Question_Type.
var (
	Question_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_TEXT",
		2: "TYPE_RADIO",
		3: "TYPE_CHECKBOX",
	}
	Question_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_TEXT":        1,
		"TYPE_RADIO":       2,
		"TYPE_CHECKBOX":    3,
	}
)

func (x Question_Type) Enum() *Question_Type {
	p := new(Question_Type)
	*p = x
	return p
}

func (x Question_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Question_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_form_v1_forms_proto_enumTypes[0].Descriptor()
}

func (Question_Type) Type() protoreflect.EnumType {
	return &file_form_v1_forms_proto_enumTypes[0]
}

func (x Question_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Question_Type.Descriptor instead.
func (Question_Type) EnumDescriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{1, 0}
}

type Form struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version   uint32                 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Title     string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Questions []*Question            `protobuf:"bytes,4,rep,name=questions,proto3" json:"questions,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Form) Reset() {
	*x = Form{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Form) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Form) ProtoMessage() {}

func (x *Form) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Form.ProtoReflect.Descriptor instead.
func (*Form) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{0}
}

func (x *Form) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Form) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Form) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Form) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

func (x *Form) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type    Question_Type `protobuf:"varint,2,opt,name=type,proto3,enum=form.v1.Question_Type" json:"type,omitempty"`
	Options []string      `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{1}
}

func (x *Question) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Question) GetType() Question_Type {
	if x != nil {
		return x.Type
	}
	return Question_TYPE_UNSPECIFIED
}

func (x *Question) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

type ResponsePagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ResponsePagination) Reset() {
	*x = ResponsePagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePagination) ProtoMessage() {}

func (x *ResponsePagination) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePagination.ProtoReflect.Descriptor instead.
func (*ResponsePagination) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{2}
}

func (x *ResponsePagination) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIdRequest) Reset() {
	*x = GetByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdRequest) ProtoMessage() {}

func (x *GetByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdRequest.ProtoReflect.Descriptor instead.
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{3}
}

func (x *GetByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *Form `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *GetByIdResponse) Reset() {
	*x = GetByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdResponse) ProtoMessage() {}

func (x *GetByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdResponse.ProtoReflect.Descriptor instead.
func (*GetByIdResponse) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{4}
}

func (x *GetByIdResponse) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string                      `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Questions []*CreateQuestionParameters `protobuf:"bytes,2,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRequest) GetQuestions() []*CreateQuestionParameters {
	if x != nil {
		return x.Questions
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *Form `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{6}
}

func (x *CreateResponse) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type CreateQuestionParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type    Question_Type `protobuf:"varint,2,opt,name=type,proto3,enum=form.v1.Question_Type" json:"type,omitempty"`
	Options []string      `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *CreateQuestionParameters) Reset() {
	*x = CreateQuestionParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestionParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestionParameters) ProtoMessage() {}

func (x *CreateQuestionParameters) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestionParameters.ProtoReflect.Descriptor instead.
func (*CreateQuestionParameters) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{7}
}

func (x *CreateQuestionParameters) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateQuestionParameters) GetType() Question_Type {
	if x != nil {
		return x.Type
	}
	return Question_TYPE_UNSPECIFIED
}

func (x *CreateQuestionParameters) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{8}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Forms      []*Form             `protobuf:"bytes,1,rep,name=forms,proto3" json:"forms,omitempty"`
	Pagination *ResponsePagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{9}
}

func (x *ListResponse) GetForms() []*Form {
	if x != nil {
		return x.Forms
	}
	return nil
}

func (x *ListResponse) GetPagination() *ResponsePagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the form to update
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The new form to replace the old form with
	NewForm *CreateRequest `protobuf:"bytes,2,opt,name=new_form,json=newForm,proto3" json:"new_form,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetNewForm() *CreateRequest {
	if x != nil {
		return x.NewForm
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *Form `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_v1_forms_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_form_v1_forms_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_form_v1_forms_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateResponse) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

var File_form_v1_forms_proto protoreflect.FileDescriptor

var file_form_v1_forms_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb2, 0x01, 0x0a, 0x04, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4e, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x41, 0x44, 0x49, 0x4f, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x42, 0x4f, 0x58, 0x10, 0x03, 0x22, 0x2a, 0x0a,
	0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72,
	0x6d, 0x22, 0x66, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x09,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x33, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x76,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x70, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6f, 0x72, 0x6d, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x46, 0x6f, 0x72, 0x6d, 0x22, 0x33, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d,
	0x32, 0xf6, 0x01, 0x0a, 0x0b, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x6c, 0x65, 0x65, 0x65, 0x6f,
	0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x6f, 0x72, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_form_v1_forms_proto_rawDescOnce sync.Once
	file_form_v1_forms_proto_rawDescData = file_form_v1_forms_proto_rawDesc
)

func file_form_v1_forms_proto_rawDescGZIP() []byte {
	file_form_v1_forms_proto_rawDescOnce.Do(func() {
		file_form_v1_forms_proto_rawDescData = protoimpl.X.CompressGZIP(file_form_v1_forms_proto_rawDescData)
	})
	return file_form_v1_forms_proto_rawDescData
}

var file_form_v1_forms_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_form_v1_forms_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_form_v1_forms_proto_goTypes = []interface{}{
	(Question_Type)(0),               // 0: form.v1.Question.Type
	(*Form)(nil),                     // 1: form.v1.Form
	(*Question)(nil),                 // 2: form.v1.Question
	(*ResponsePagination)(nil),       // 3: form.v1.ResponsePagination
	(*GetByIdRequest)(nil),           // 4: form.v1.GetByIdRequest
	(*GetByIdResponse)(nil),          // 5: form.v1.GetByIdResponse
	(*CreateRequest)(nil),            // 6: form.v1.CreateRequest
	(*CreateResponse)(nil),           // 7: form.v1.CreateResponse
	(*CreateQuestionParameters)(nil), // 8: form.v1.CreateQuestionParameters
	(*ListRequest)(nil),              // 9: form.v1.ListRequest
	(*ListResponse)(nil),             // 10: form.v1.ListResponse
	(*UpdateRequest)(nil),            // 11: form.v1.UpdateRequest
	(*UpdateResponse)(nil),           // 12: form.v1.UpdateResponse
	(*timestamppb.Timestamp)(nil),    // 13: google.protobuf.Timestamp
}
var file_form_v1_forms_proto_depIdxs = []int32{
	2,  // 0: form.v1.Form.questions:type_name -> form.v1.Question
	13, // 1: form.v1.Form.created_at:type_name -> google.protobuf.Timestamp
	0,  // 2: form.v1.Question.type:type_name -> form.v1.Question.Type
	1,  // 3: form.v1.GetByIdResponse.form:type_name -> form.v1.Form
	8,  // 4: form.v1.CreateRequest.questions:type_name -> form.v1.CreateQuestionParameters
	1,  // 5: form.v1.CreateResponse.form:type_name -> form.v1.Form
	0,  // 6: form.v1.CreateQuestionParameters.type:type_name -> form.v1.Question.Type
	1,  // 7: form.v1.ListResponse.forms:type_name -> form.v1.Form
	3,  // 8: form.v1.ListResponse.pagination:type_name -> form.v1.ResponsePagination
	6,  // 9: form.v1.UpdateRequest.new_form:type_name -> form.v1.CreateRequest
	1,  // 10: form.v1.UpdateResponse.form:type_name -> form.v1.Form
	4,  // 11: form.v1.FormService.GetById:input_type -> form.v1.GetByIdRequest
	6,  // 12: form.v1.FormService.Create:input_type -> form.v1.CreateRequest
	9,  // 13: form.v1.FormService.List:input_type -> form.v1.ListRequest
	11, // 14: form.v1.FormService.Update:input_type -> form.v1.UpdateRequest
	5,  // 15: form.v1.FormService.GetById:output_type -> form.v1.GetByIdResponse
	7,  // 16: form.v1.FormService.Create:output_type -> form.v1.CreateResponse
	10, // 17: form.v1.FormService.List:output_type -> form.v1.ListResponse
	12, // 18: form.v1.FormService.Update:output_type -> form.v1.UpdateResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_form_v1_forms_proto_init() }
func file_form_v1_forms_proto_init() {
	if File_form_v1_forms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_form_v1_forms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Form); i {
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
		file_form_v1_forms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_form_v1_forms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePagination); i {
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
		file_form_v1_forms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdRequest); i {
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
		file_form_v1_forms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdResponse); i {
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
		file_form_v1_forms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_form_v1_forms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_form_v1_forms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuestionParameters); i {
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
		file_form_v1_forms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_form_v1_forms_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_form_v1_forms_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_form_v1_forms_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_form_v1_forms_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_form_v1_forms_proto_goTypes,
		DependencyIndexes: file_form_v1_forms_proto_depIdxs,
		EnumInfos:         file_form_v1_forms_proto_enumTypes,
		MessageInfos:      file_form_v1_forms_proto_msgTypes,
	}.Build()
	File_form_v1_forms_proto = out.File
	file_form_v1_forms_proto_rawDesc = nil
	file_form_v1_forms_proto_goTypes = nil
	file_form_v1_forms_proto_depIdxs = nil
}
