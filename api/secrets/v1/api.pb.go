// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: secrets/v1/api.proto

package secretsv1

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

// Secret represents a single secret version.
type Secret struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`          // The plain text key of the secret.
	Value         []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`      // The encrypted value of the secret.
	Version       uint32                 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"` // The version of the represented secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Secret) Reset() {
	*x = Secret{}
	mi := &file_secrets_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *Secret) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Secret) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Secret) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// SecretMetadata represents a secret and all its versions.
type SecretMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`                   // The plain text key of the secret.
	Versions      []uint32               `protobuf:"varint,2,rep,packed,name=versions,proto3" json:"versions,omitempty"` // All versions of the secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecretMetadata) Reset() {
	*x = SecretMetadata{}
	mi := &file_secrets_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecretMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretMetadata) ProtoMessage() {}

func (x *SecretMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretMetadata.ProtoReflect.Descriptor instead.
func (*SecretMetadata) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *SecretMetadata) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SecretMetadata) GetVersions() []uint32 {
	if x != nil {
		return x.Versions
	}
	return nil
}

// Request message for `secrets.v1.Kv2Service/CreateSecret`.
type CreateSecretRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`     // The plain text key of the secret.
	Value         []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // The encoded value of the secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSecretRequest) Reset() {
	*x = CreateSecretRequest{}
	mi := &file_secrets_v1_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSecretRequest) ProtoMessage() {}

func (x *CreateSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSecretRequest.ProtoReflect.Descriptor instead.
func (*CreateSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CreateSecretRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Response message for `secrets.v1.Kv2Service/CreateSecret`.
type CreateSecretResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Secret        *SecretMetadata        `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"` // The metadata of the created secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSecretResponse) Reset() {
	*x = CreateSecretResponse{}
	mi := &file_secrets_v1_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSecretResponse) ProtoMessage() {}

func (x *CreateSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSecretResponse.ProtoReflect.Descriptor instead.
func (*CreateSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSecretResponse) GetSecret() *SecretMetadata {
	if x != nil {
		return x.Secret
	}
	return nil
}

// Request message for `secrets.v1.Kv2Service/GetSecret`.
type GetSecretRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // The plain text key of the secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSecretRequest) Reset() {
	*x = GetSecretRequest{}
	mi := &file_secrets_v1_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretRequest) ProtoMessage() {}

func (x *GetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretRequest.ProtoReflect.Descriptor instead.
func (*GetSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Response message for `secrets.v1.Kv2Service/GetSecret`.
type GetSecretResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Secret        *Secret                `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"` // The requested secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSecretResponse) Reset() {
	*x = GetSecretResponse{}
	mi := &file_secrets_v1_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretResponse) ProtoMessage() {}

func (x *GetSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretResponse.ProtoReflect.Descriptor instead.
func (*GetSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetSecretResponse) GetSecret() *Secret {
	if x != nil {
		return x.Secret
	}
	return nil
}

// Request message for `secrets.v1.Kv2Service/UpdateSecret`.
type UpdateSecretRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`     // The plain text key of the secret.
	Value         []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // The encoded value of the secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSecretRequest) Reset() {
	*x = UpdateSecretRequest{}
	mi := &file_secrets_v1_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSecretRequest) ProtoMessage() {}

func (x *UpdateSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSecretRequest.ProtoReflect.Descriptor instead.
func (*UpdateSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateSecretRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Response message for `secrets.v1.Kv2Service/UpdateSecret`.
type UpdateSecretResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       uint32                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"` // The version of the updated secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSecretResponse) Reset() {
	*x = UpdateSecretResponse{}
	mi := &file_secrets_v1_api_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSecretResponse) ProtoMessage() {}

func (x *UpdateSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSecretResponse.ProtoReflect.Descriptor instead.
func (*UpdateSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateSecretResponse) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// Request message for `secrets.v1.Kv2Service/DeleteSecret`.
type DeleteSecretRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // The plain text key of the secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSecretRequest) Reset() {
	*x = DeleteSecretRequest{}
	mi := &file_secrets_v1_api_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSecretRequest) ProtoMessage() {}

func (x *DeleteSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSecretRequest.ProtoReflect.Descriptor instead.
func (*DeleteSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Empty message. Check for an error code to determine success.
type DeleteSecretResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSecretResponse) Reset() {
	*x = DeleteSecretResponse{}
	mi := &file_secrets_v1_api_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSecretResponse) ProtoMessage() {}

func (x *DeleteSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSecretResponse.ProtoReflect.Descriptor instead.
func (*DeleteSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{9}
}

// Request message for `secrets.v1.Kv2Service/RevertSecret`.
type RevertSecretRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // The plain text key of the secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevertSecretRequest) Reset() {
	*x = RevertSecretRequest{}
	mi := &file_secrets_v1_api_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevertSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevertSecretRequest) ProtoMessage() {}

func (x *RevertSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevertSecretRequest.ProtoReflect.Descriptor instead.
func (*RevertSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{10}
}

func (x *RevertSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Request message for `secrets.v1.Kv2Service/RevertSecret`.
type RevertSecretResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       uint32                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"` // The current version of the reverted secret.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevertSecretResponse) Reset() {
	*x = RevertSecretResponse{}
	mi := &file_secrets_v1_api_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevertSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevertSecretResponse) ProtoMessage() {}

func (x *RevertSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevertSecretResponse.ProtoReflect.Descriptor instead.
func (*RevertSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{11}
}

func (x *RevertSecretResponse) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// Empty message. No request message needed to list secrets.
type ListSecretsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSecretsRequest) Reset() {
	*x = ListSecretsRequest{}
	mi := &file_secrets_v1_api_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSecretsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSecretsRequest) ProtoMessage() {}

func (x *ListSecretsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSecretsRequest.ProtoReflect.Descriptor instead.
func (*ListSecretsRequest) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{12}
}

// Response message for `secrets.v1.Kv2Service/ListSecrets`.
type ListSecretsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Secrets       []*SecretMetadata      `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"` // The secret metadata for all available secrets.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSecretsResponse) Reset() {
	*x = ListSecretsResponse{}
	mi := &file_secrets_v1_api_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSecretsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSecretsResponse) ProtoMessage() {}

func (x *ListSecretsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSecretsResponse.ProtoReflect.Descriptor instead.
func (*ListSecretsResponse) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{13}
}

func (x *ListSecretsResponse) GetSecrets() []*SecretMetadata {
	if x != nil {
		return x.Secrets
	}
	return nil
}

// Request message for `secrets.v1.Kv2Service/Backup`.
type BackupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          *string                `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"` // The name of the backup. Defaults to `kv2.db`.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackupRequest) Reset() {
	*x = BackupRequest{}
	mi := &file_secrets_v1_api_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupRequest) ProtoMessage() {}

func (x *BackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupRequest.ProtoReflect.Descriptor instead.
func (*BackupRequest) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{14}
}

func (x *BackupRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

// Empty message. Check for an error code to determine success.
type BackupResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackupResponse) Reset() {
	*x = BackupResponse{}
	mi := &file_secrets_v1_api_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupResponse) ProtoMessage() {}

func (x *BackupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_v1_api_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupResponse.ProtoReflect.Descriptor instead.
func (*BackupResponse) Descriptor() ([]byte, []int) {
	return file_secrets_v1_api_proto_rawDescGZIP(), []int{15}
}

var File_secrets_v1_api_proto protoreflect.FileDescriptor

const file_secrets_v1_api_proto_rawDesc = "" +
	"\n" +
	"\x14secrets/v1/api.proto\x12\n" +
	"secrets.v1\"J\n" +
	"\x06Secret\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\fR\x05value\x12\x18\n" +
	"\aversion\x18\x03 \x01(\rR\aversion\">\n" +
	"\x0eSecretMetadata\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x1a\n" +
	"\bversions\x18\x02 \x03(\rR\bversions\"=\n" +
	"\x13CreateSecretRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\fR\x05value\"J\n" +
	"\x14CreateSecretResponse\x122\n" +
	"\x06secret\x18\x01 \x01(\v2\x1a.secrets.v1.SecretMetadataR\x06secret\"$\n" +
	"\x10GetSecretRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\"?\n" +
	"\x11GetSecretResponse\x12*\n" +
	"\x06secret\x18\x01 \x01(\v2\x12.secrets.v1.SecretR\x06secret\"=\n" +
	"\x13UpdateSecretRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\fR\x05value\"0\n" +
	"\x14UpdateSecretResponse\x12\x18\n" +
	"\aversion\x18\x01 \x01(\rR\aversion\"'\n" +
	"\x13DeleteSecretRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\"\x16\n" +
	"\x14DeleteSecretResponse\"'\n" +
	"\x13RevertSecretRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\"0\n" +
	"\x14RevertSecretResponse\x12\x18\n" +
	"\aversion\x18\x01 \x01(\rR\aversion\"\x14\n" +
	"\x12ListSecretsRequest\"K\n" +
	"\x13ListSecretsResponse\x124\n" +
	"\asecrets\x18\x01 \x03(\v2\x1a.secrets.v1.SecretMetadataR\asecrets\"1\n" +
	"\rBackupRequest\x12\x17\n" +
	"\x04name\x18\x01 \x01(\tH\x00R\x04name\x88\x01\x01B\a\n" +
	"\x05_name\"\x10\n" +
	"\x0eBackupResponse2\xb3\x04\n" +
	"\n" +
	"Kv2Service\x12Q\n" +
	"\fCreateSecret\x12\x1f.secrets.v1.CreateSecretRequest\x1a .secrets.v1.CreateSecretResponse\x12H\n" +
	"\tGetSecret\x12\x1c.secrets.v1.GetSecretRequest\x1a\x1d.secrets.v1.GetSecretResponse\x12Q\n" +
	"\fUpdateSecret\x12\x1f.secrets.v1.UpdateSecretRequest\x1a .secrets.v1.UpdateSecretResponse\x12Q\n" +
	"\fDeleteSecret\x12\x1f.secrets.v1.DeleteSecretRequest\x1a .secrets.v1.DeleteSecretResponse\x12Q\n" +
	"\fRevertSecret\x12\x1f.secrets.v1.RevertSecretRequest\x1a .secrets.v1.RevertSecretResponse\x12N\n" +
	"\vListSecrets\x12\x1e.secrets.v1.ListSecretsRequest\x1a\x1f.secrets.v1.ListSecretsResponse\x12?\n" +
	"\x06Backup\x12\x19.secrets.v1.BackupRequest\x1a\x1a.secrets.v1.BackupResponseB\x90\x01\n" +
	"\x0ecom.secrets.v1B\bApiProtoP\x01Z+git.huggins.io/kv2/api/secrets/v1;secretsv1\xa2\x02\x03SXX\xaa\x02\n" +
	"Secrets.V1\xca\x02\n" +
	"Secrets\\V1\xe2\x02\x16Secrets\\V1\\GPBMetadata\xea\x02\vSecrets::V1b\x06proto3"

var (
	file_secrets_v1_api_proto_rawDescOnce sync.Once
	file_secrets_v1_api_proto_rawDescData []byte
)

func file_secrets_v1_api_proto_rawDescGZIP() []byte {
	file_secrets_v1_api_proto_rawDescOnce.Do(func() {
		file_secrets_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_secrets_v1_api_proto_rawDesc), len(file_secrets_v1_api_proto_rawDesc)))
	})
	return file_secrets_v1_api_proto_rawDescData
}

var file_secrets_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_secrets_v1_api_proto_goTypes = []any{
	(*Secret)(nil),               // 0: secrets.v1.Secret
	(*SecretMetadata)(nil),       // 1: secrets.v1.SecretMetadata
	(*CreateSecretRequest)(nil),  // 2: secrets.v1.CreateSecretRequest
	(*CreateSecretResponse)(nil), // 3: secrets.v1.CreateSecretResponse
	(*GetSecretRequest)(nil),     // 4: secrets.v1.GetSecretRequest
	(*GetSecretResponse)(nil),    // 5: secrets.v1.GetSecretResponse
	(*UpdateSecretRequest)(nil),  // 6: secrets.v1.UpdateSecretRequest
	(*UpdateSecretResponse)(nil), // 7: secrets.v1.UpdateSecretResponse
	(*DeleteSecretRequest)(nil),  // 8: secrets.v1.DeleteSecretRequest
	(*DeleteSecretResponse)(nil), // 9: secrets.v1.DeleteSecretResponse
	(*RevertSecretRequest)(nil),  // 10: secrets.v1.RevertSecretRequest
	(*RevertSecretResponse)(nil), // 11: secrets.v1.RevertSecretResponse
	(*ListSecretsRequest)(nil),   // 12: secrets.v1.ListSecretsRequest
	(*ListSecretsResponse)(nil),  // 13: secrets.v1.ListSecretsResponse
	(*BackupRequest)(nil),        // 14: secrets.v1.BackupRequest
	(*BackupResponse)(nil),       // 15: secrets.v1.BackupResponse
}
var file_secrets_v1_api_proto_depIdxs = []int32{
	1,  // 0: secrets.v1.CreateSecretResponse.secret:type_name -> secrets.v1.SecretMetadata
	0,  // 1: secrets.v1.GetSecretResponse.secret:type_name -> secrets.v1.Secret
	1,  // 2: secrets.v1.ListSecretsResponse.secrets:type_name -> secrets.v1.SecretMetadata
	2,  // 3: secrets.v1.Kv2Service.CreateSecret:input_type -> secrets.v1.CreateSecretRequest
	4,  // 4: secrets.v1.Kv2Service.GetSecret:input_type -> secrets.v1.GetSecretRequest
	6,  // 5: secrets.v1.Kv2Service.UpdateSecret:input_type -> secrets.v1.UpdateSecretRequest
	8,  // 6: secrets.v1.Kv2Service.DeleteSecret:input_type -> secrets.v1.DeleteSecretRequest
	10, // 7: secrets.v1.Kv2Service.RevertSecret:input_type -> secrets.v1.RevertSecretRequest
	12, // 8: secrets.v1.Kv2Service.ListSecrets:input_type -> secrets.v1.ListSecretsRequest
	14, // 9: secrets.v1.Kv2Service.Backup:input_type -> secrets.v1.BackupRequest
	3,  // 10: secrets.v1.Kv2Service.CreateSecret:output_type -> secrets.v1.CreateSecretResponse
	5,  // 11: secrets.v1.Kv2Service.GetSecret:output_type -> secrets.v1.GetSecretResponse
	7,  // 12: secrets.v1.Kv2Service.UpdateSecret:output_type -> secrets.v1.UpdateSecretResponse
	9,  // 13: secrets.v1.Kv2Service.DeleteSecret:output_type -> secrets.v1.DeleteSecretResponse
	11, // 14: secrets.v1.Kv2Service.RevertSecret:output_type -> secrets.v1.RevertSecretResponse
	13, // 15: secrets.v1.Kv2Service.ListSecrets:output_type -> secrets.v1.ListSecretsResponse
	15, // 16: secrets.v1.Kv2Service.Backup:output_type -> secrets.v1.BackupResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_secrets_v1_api_proto_init() }
func file_secrets_v1_api_proto_init() {
	if File_secrets_v1_api_proto != nil {
		return
	}
	file_secrets_v1_api_proto_msgTypes[14].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_secrets_v1_api_proto_rawDesc), len(file_secrets_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_secrets_v1_api_proto_goTypes,
		DependencyIndexes: file_secrets_v1_api_proto_depIdxs,
		MessageInfos:      file_secrets_v1_api_proto_msgTypes,
	}.Build()
	File_secrets_v1_api_proto = out.File
	file_secrets_v1_api_proto_goTypes = nil
	file_secrets_v1_api_proto_depIdxs = nil
}
