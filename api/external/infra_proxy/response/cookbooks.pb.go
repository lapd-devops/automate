// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/response/cookbooks.proto

package response

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Cookbooks struct {
	Cookbooks            []*CookbookVersion `protobuf:"bytes,1,rep,name=cookbooks,proto3" json:"cookbooks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Cookbooks) Reset()         { *m = Cookbooks{} }
func (m *Cookbooks) String() string { return proto.CompactTextString(m) }
func (*Cookbooks) ProtoMessage()    {}
func (*Cookbooks) Descriptor() ([]byte, []int) {
	return fileDescriptor_e929c46acd19ab86, []int{0}
}

func (m *Cookbooks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cookbooks.Unmarshal(m, b)
}
func (m *Cookbooks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cookbooks.Marshal(b, m, deterministic)
}
func (m *Cookbooks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookbooks.Merge(m, src)
}
func (m *Cookbooks) XXX_Size() int {
	return xxx_messageInfo_Cookbooks.Size(m)
}
func (m *Cookbooks) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookbooks.DiscardUnknown(m)
}

var xxx_messageInfo_Cookbooks proto.InternalMessageInfo

func (m *Cookbooks) GetCookbooks() []*CookbookVersion {
	if m != nil {
		return m.Cookbooks
	}
	return nil
}

type Cookbook struct {
	CookbookName         string          `protobuf:"bytes,1,opt,name=cookbook_name,json=cookbookName,proto3" json:"cookbook_name,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string          `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	ChefType             string          `protobuf:"bytes,4,opt,name=chef_type,json=chefType,proto3" json:"chef_type,omitempty"`
	Frozen               bool            `protobuf:"varint,5,opt,name=frozen,proto3" json:"frozen,omitempty"`
	JsonClass            string          `protobuf:"bytes,6,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty"`
	Files                []*CookbookItem `protobuf:"bytes,7,rep,name=files,proto3" json:"files,omitempty"`
	Templates            []*CookbookItem `protobuf:"bytes,8,rep,name=templates,proto3" json:"templates,omitempty"`
	Attributes           []*CookbookItem `protobuf:"bytes,9,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Recipes              []*CookbookItem `protobuf:"bytes,10,rep,name=recipes,proto3" json:"recipes,omitempty"`
	Definitions          []*CookbookItem `protobuf:"bytes,11,rep,name=definitions,proto3" json:"definitions,omitempty"`
	Libraries            []*CookbookItem `protobuf:"bytes,12,rep,name=libraries,proto3" json:"libraries,omitempty"`
	Providers            []*CookbookItem `protobuf:"bytes,13,rep,name=providers,proto3" json:"providers,omitempty"`
	Resources            []*CookbookItem `protobuf:"bytes,14,rep,name=resources,proto3" json:"resources,omitempty"`
	RootFiles            []*CookbookItem `protobuf:"bytes,15,rep,name=root_files,json=rootFiles,proto3" json:"root_files,omitempty"`
	Metadata             *CookbookMeta   `protobuf:"bytes,16,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Access               *CookbookAccess `protobuf:"bytes,17,opt,name=access,proto3" json:"access,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Cookbook) Reset()         { *m = Cookbook{} }
func (m *Cookbook) String() string { return proto.CompactTextString(m) }
func (*Cookbook) ProtoMessage()    {}
func (*Cookbook) Descriptor() ([]byte, []int) {
	return fileDescriptor_e929c46acd19ab86, []int{1}
}

func (m *Cookbook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cookbook.Unmarshal(m, b)
}
func (m *Cookbook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cookbook.Marshal(b, m, deterministic)
}
func (m *Cookbook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookbook.Merge(m, src)
}
func (m *Cookbook) XXX_Size() int {
	return xxx_messageInfo_Cookbook.Size(m)
}
func (m *Cookbook) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookbook.DiscardUnknown(m)
}

var xxx_messageInfo_Cookbook proto.InternalMessageInfo

func (m *Cookbook) GetCookbookName() string {
	if m != nil {
		return m.CookbookName
	}
	return ""
}

func (m *Cookbook) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cookbook) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Cookbook) GetChefType() string {
	if m != nil {
		return m.ChefType
	}
	return ""
}

func (m *Cookbook) GetFrozen() bool {
	if m != nil {
		return m.Frozen
	}
	return false
}

func (m *Cookbook) GetJsonClass() string {
	if m != nil {
		return m.JsonClass
	}
	return ""
}

func (m *Cookbook) GetFiles() []*CookbookItem {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Cookbook) GetTemplates() []*CookbookItem {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *Cookbook) GetAttributes() []*CookbookItem {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Cookbook) GetRecipes() []*CookbookItem {
	if m != nil {
		return m.Recipes
	}
	return nil
}

func (m *Cookbook) GetDefinitions() []*CookbookItem {
	if m != nil {
		return m.Definitions
	}
	return nil
}

func (m *Cookbook) GetLibraries() []*CookbookItem {
	if m != nil {
		return m.Libraries
	}
	return nil
}

func (m *Cookbook) GetProviders() []*CookbookItem {
	if m != nil {
		return m.Providers
	}
	return nil
}

func (m *Cookbook) GetResources() []*CookbookItem {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Cookbook) GetRootFiles() []*CookbookItem {
	if m != nil {
		return m.RootFiles
	}
	return nil
}

func (m *Cookbook) GetMetadata() *CookbookMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Cookbook) GetAccess() *CookbookAccess {
	if m != nil {
		return m.Access
	}
	return nil
}

type CookbookItem struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Checksum             string   `protobuf:"bytes,4,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Specificity          string   `protobuf:"bytes,5,opt,name=specificity,proto3" json:"specificity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookbookItem) Reset()         { *m = CookbookItem{} }
func (m *CookbookItem) String() string { return proto.CompactTextString(m) }
func (*CookbookItem) ProtoMessage()    {}
func (*CookbookItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e929c46acd19ab86, []int{2}
}

func (m *CookbookItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookItem.Unmarshal(m, b)
}
func (m *CookbookItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookItem.Marshal(b, m, deterministic)
}
func (m *CookbookItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookItem.Merge(m, src)
}
func (m *CookbookItem) XXX_Size() int {
	return xxx_messageInfo_CookbookItem.Size(m)
}
func (m *CookbookItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookItem.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookItem proto.InternalMessageInfo

func (m *CookbookItem) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CookbookItem) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CookbookItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CookbookItem) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *CookbookItem) GetSpecificity() string {
	if m != nil {
		return m.Specificity
	}
	return ""
}

type CookbookMeta struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	LongDescription      string   `protobuf:"bytes,4,opt,name=long_description,json=longDescription,proto3" json:"long_description,omitempty"`
	Maintainer           string   `protobuf:"bytes,5,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
	MaintainerEmail      string   `protobuf:"bytes,6,opt,name=maintainer_email,json=maintainerEmail,proto3" json:"maintainer_email,omitempty"`
	License              string   `protobuf:"bytes,7,opt,name=license,proto3" json:"license,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookbookMeta) Reset()         { *m = CookbookMeta{} }
func (m *CookbookMeta) String() string { return proto.CompactTextString(m) }
func (*CookbookMeta) ProtoMessage()    {}
func (*CookbookMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_e929c46acd19ab86, []int{3}
}

func (m *CookbookMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookMeta.Unmarshal(m, b)
}
func (m *CookbookMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookMeta.Marshal(b, m, deterministic)
}
func (m *CookbookMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookMeta.Merge(m, src)
}
func (m *CookbookMeta) XXX_Size() int {
	return xxx_messageInfo_CookbookMeta.Size(m)
}
func (m *CookbookMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookMeta.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookMeta proto.InternalMessageInfo

func (m *CookbookMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CookbookMeta) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CookbookMeta) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CookbookMeta) GetLongDescription() string {
	if m != nil {
		return m.LongDescription
	}
	return ""
}

func (m *CookbookMeta) GetMaintainer() string {
	if m != nil {
		return m.Maintainer
	}
	return ""
}

func (m *CookbookMeta) GetMaintainerEmail() string {
	if m != nil {
		return m.MaintainerEmail
	}
	return ""
}

func (m *CookbookMeta) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

type CookbookAccess struct {
	Read                 bool     `protobuf:"varint,1,opt,name=read,proto3" json:"read,omitempty"`
	Create               bool     `protobuf:"varint,2,opt,name=create,proto3" json:"create,omitempty"`
	Grant                bool     `protobuf:"varint,3,opt,name=grant,proto3" json:"grant,omitempty"`
	Update               bool     `protobuf:"varint,4,opt,name=update,proto3" json:"update,omitempty"`
	Delete               bool     `protobuf:"varint,5,opt,name=delete,proto3" json:"delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookbookAccess) Reset()         { *m = CookbookAccess{} }
func (m *CookbookAccess) String() string { return proto.CompactTextString(m) }
func (*CookbookAccess) ProtoMessage()    {}
func (*CookbookAccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_e929c46acd19ab86, []int{4}
}

func (m *CookbookAccess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookAccess.Unmarshal(m, b)
}
func (m *CookbookAccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookAccess.Marshal(b, m, deterministic)
}
func (m *CookbookAccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookAccess.Merge(m, src)
}
func (m *CookbookAccess) XXX_Size() int {
	return xxx_messageInfo_CookbookAccess.Size(m)
}
func (m *CookbookAccess) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookAccess.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookAccess proto.InternalMessageInfo

func (m *CookbookAccess) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func (m *CookbookAccess) GetCreate() bool {
	if m != nil {
		return m.Create
	}
	return false
}

func (m *CookbookAccess) GetGrant() bool {
	if m != nil {
		return m.Grant
	}
	return false
}

func (m *CookbookAccess) GetUpdate() bool {
	if m != nil {
		return m.Update
	}
	return false
}

func (m *CookbookAccess) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

type CookbooksAvailableVersions struct {
	Cookbooks            []*CookbookAllVersion `protobuf:"bytes,1,rep,name=cookbooks,proto3" json:"cookbooks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CookbooksAvailableVersions) Reset()         { *m = CookbooksAvailableVersions{} }
func (m *CookbooksAvailableVersions) String() string { return proto.CompactTextString(m) }
func (*CookbooksAvailableVersions) ProtoMessage()    {}
func (*CookbooksAvailableVersions) Descriptor() ([]byte, []int) {
	return fileDescriptor_e929c46acd19ab86, []int{5}
}

func (m *CookbooksAvailableVersions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbooksAvailableVersions.Unmarshal(m, b)
}
func (m *CookbooksAvailableVersions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbooksAvailableVersions.Marshal(b, m, deterministic)
}
func (m *CookbooksAvailableVersions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbooksAvailableVersions.Merge(m, src)
}
func (m *CookbooksAvailableVersions) XXX_Size() int {
	return xxx_messageInfo_CookbooksAvailableVersions.Size(m)
}
func (m *CookbooksAvailableVersions) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbooksAvailableVersions.DiscardUnknown(m)
}

var xxx_messageInfo_CookbooksAvailableVersions proto.InternalMessageInfo

func (m *CookbooksAvailableVersions) GetCookbooks() []*CookbookAllVersion {
	if m != nil {
		return m.Cookbooks
	}
	return nil
}

type CookbookVersion struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookbookVersion) Reset()         { *m = CookbookVersion{} }
func (m *CookbookVersion) String() string { return proto.CompactTextString(m) }
func (*CookbookVersion) ProtoMessage()    {}
func (*CookbookVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_e929c46acd19ab86, []int{6}
}

func (m *CookbookVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookVersion.Unmarshal(m, b)
}
func (m *CookbookVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookVersion.Marshal(b, m, deterministic)
}
func (m *CookbookVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookVersion.Merge(m, src)
}
func (m *CookbookVersion) XXX_Size() int {
	return xxx_messageInfo_CookbookVersion.Size(m)
}
func (m *CookbookVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookVersion.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookVersion proto.InternalMessageInfo

func (m *CookbookVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CookbookVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type CookbookAllVersion struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CurrentVersion       string   `protobuf:"bytes,2,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty"`
	Versions             []string `protobuf:"bytes,3,rep,name=versions,proto3" json:"versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookbookAllVersion) Reset()         { *m = CookbookAllVersion{} }
func (m *CookbookAllVersion) String() string { return proto.CompactTextString(m) }
func (*CookbookAllVersion) ProtoMessage()    {}
func (*CookbookAllVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_e929c46acd19ab86, []int{7}
}

func (m *CookbookAllVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookAllVersion.Unmarshal(m, b)
}
func (m *CookbookAllVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookAllVersion.Marshal(b, m, deterministic)
}
func (m *CookbookAllVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookAllVersion.Merge(m, src)
}
func (m *CookbookAllVersion) XXX_Size() int {
	return xxx_messageInfo_CookbookAllVersion.Size(m)
}
func (m *CookbookAllVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookAllVersion.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookAllVersion proto.InternalMessageInfo

func (m *CookbookAllVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CookbookAllVersion) GetCurrentVersion() string {
	if m != nil {
		return m.CurrentVersion
	}
	return ""
}

func (m *CookbookAllVersion) GetVersions() []string {
	if m != nil {
		return m.Versions
	}
	return nil
}

type CookbookFileContent struct {
	// Cookbook data file content.
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookbookFileContent) Reset()         { *m = CookbookFileContent{} }
func (m *CookbookFileContent) String() string { return proto.CompactTextString(m) }
func (*CookbookFileContent) ProtoMessage()    {}
func (*CookbookFileContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e929c46acd19ab86, []int{8}
}

func (m *CookbookFileContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookFileContent.Unmarshal(m, b)
}
func (m *CookbookFileContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookFileContent.Marshal(b, m, deterministic)
}
func (m *CookbookFileContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookFileContent.Merge(m, src)
}
func (m *CookbookFileContent) XXX_Size() int {
	return xxx_messageInfo_CookbookFileContent.Size(m)
}
func (m *CookbookFileContent) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookFileContent.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookFileContent proto.InternalMessageInfo

func (m *CookbookFileContent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Cookbooks)(nil), "chef.automate.api.infra_proxy.response.Cookbooks")
	proto.RegisterType((*Cookbook)(nil), "chef.automate.api.infra_proxy.response.Cookbook")
	proto.RegisterType((*CookbookItem)(nil), "chef.automate.api.infra_proxy.response.CookbookItem")
	proto.RegisterType((*CookbookMeta)(nil), "chef.automate.api.infra_proxy.response.CookbookMeta")
	proto.RegisterType((*CookbookAccess)(nil), "chef.automate.api.infra_proxy.response.CookbookAccess")
	proto.RegisterType((*CookbooksAvailableVersions)(nil), "chef.automate.api.infra_proxy.response.CookbooksAvailableVersions")
	proto.RegisterType((*CookbookVersion)(nil), "chef.automate.api.infra_proxy.response.CookbookVersion")
	proto.RegisterType((*CookbookAllVersion)(nil), "chef.automate.api.infra_proxy.response.CookbookAllVersion")
	proto.RegisterType((*CookbookFileContent)(nil), "chef.automate.api.infra_proxy.response.CookbookFileContent")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/response/cookbooks.proto", fileDescriptor_e929c46acd19ab86)
}

var fileDescriptor_e929c46acd19ab86 = []byte{
	// 759 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xdf, 0x6b, 0x2c, 0x35,
	0x14, 0xc7, 0x99, 0xbb, 0xed, 0xee, 0xcc, 0xd9, 0xde, 0x6e, 0x8d, 0x22, 0xa1, 0xa2, 0x2c, 0x23,
	0x68, 0x7d, 0x99, 0xc1, 0x1f, 0x28, 0x5c, 0x04, 0xb9, 0x56, 0x05, 0x05, 0x8b, 0x8c, 0xd7, 0x22,
	0xbe, 0x0c, 0xd9, 0xec, 0xd9, 0x36, 0x76, 0x26, 0x19, 0x92, 0xcc, 0xd2, 0xf5, 0xd1, 0x07, 0xff,
	0x30, 0xff, 0x20, 0xff, 0x06, 0x49, 0xe6, 0xc7, 0x4e, 0x6d, 0x41, 0xd7, 0x79, 0xcb, 0xf9, 0x26,
	0xe7, 0x93, 0x93, 0xb3, 0xd9, 0x6f, 0x06, 0x3e, 0x64, 0x95, 0x48, 0xf1, 0xde, 0xa2, 0x96, 0xac,
	0x48, 0x85, 0xdc, 0x68, 0x96, 0x57, 0x5a, 0xdd, 0xef, 0x52, 0x8d, 0xa6, 0x52, 0xd2, 0x60, 0xca,
	0x95, 0xba, 0x5b, 0x29, 0x75, 0x67, 0x92, 0x4a, 0x2b, 0xab, 0xc8, 0x7b, 0xfc, 0x16, 0x37, 0x09,
	0xab, 0xad, 0x2a, 0x99, 0xc5, 0x84, 0x55, 0x22, 0x19, 0xe4, 0x25, 0x5d, 0x5e, 0xbc, 0x82, 0xe8,
	0xb2, 0x4b, 0x25, 0x3f, 0x41, 0xd4, 0x73, 0x68, 0xb0, 0x9c, 0x5c, 0xcc, 0x3f, 0xfa, 0x2c, 0xf9,
	0x6f, 0xa0, 0xa4, 0xa3, 0x5c, 0xa3, 0x36, 0x42, 0xc9, 0x6c, 0x4f, 0x8a, 0xff, 0x0c, 0x21, 0xec,
	0xa6, 0xc9, 0xbb, 0xf0, 0xbc, 0x9b, 0xc9, 0x25, 0x2b, 0x91, 0x06, 0xcb, 0xe0, 0x22, 0xca, 0x4e,
	0x3a, 0xf1, 0x8a, 0x95, 0x48, 0x08, 0x1c, 0xf9, 0xb9, 0x67, 0x7e, 0xce, 0x8f, 0x09, 0x85, 0xd9,
	0xb6, 0x61, 0xd3, 0x89, 0x97, 0xbb, 0x90, 0xbc, 0x05, 0x91, 0x2b, 0x32, 0xb7, 0xbb, 0x0a, 0xe9,
	0x91, 0x9f, 0x0b, 0x9d, 0xf0, 0x6a, 0x57, 0x21, 0x79, 0x13, 0xa6, 0x1b, 0xad, 0x7e, 0x43, 0x49,
	0x8f, 0x97, 0xc1, 0x45, 0x98, 0xb5, 0x11, 0x79, 0x1b, 0xe0, 0x57, 0xa3, 0x64, 0xce, 0x0b, 0x66,
	0x0c, 0x9d, 0xfa, 0xac, 0xc8, 0x29, 0x97, 0x4e, 0x20, 0xdf, 0xc1, 0xf1, 0x46, 0x14, 0x68, 0xe8,
	0xcc, 0xb7, 0xe1, 0x93, 0x43, 0xdb, 0xf0, 0xad, 0xc5, 0x32, 0x6b, 0x10, 0x24, 0x83, 0xc8, 0x62,
	0x59, 0x15, 0xcc, 0xa2, 0xa1, 0xe1, 0x08, 0xde, 0x1e, 0x43, 0x5e, 0x01, 0x30, 0x6b, 0xb5, 0x58,
	0xd5, 0x0e, 0x1a, 0x8d, 0x80, 0x0e, 0x38, 0xe4, 0x0a, 0x66, 0x1a, 0xb9, 0xa8, 0xd0, 0x50, 0x18,
	0x81, 0xec, 0x20, 0xe4, 0x1a, 0xe6, 0x6b, 0xdc, 0x08, 0x29, 0xac, 0x50, 0xd2, 0xd0, 0xf9, 0x08,
	0xe6, 0x10, 0xe4, 0x3a, 0x5a, 0x88, 0x95, 0x66, 0x5a, 0xa0, 0xa1, 0x27, 0x63, 0x3a, 0xda, 0x63,
	0x1c, 0xb3, 0xd2, 0x6a, 0x2b, 0xd6, 0xa8, 0x0d, 0x7d, 0x3e, 0x86, 0xd9, 0x63, 0x1c, 0x53, 0xa3,
	0x51, 0xb5, 0xe6, 0x68, 0xe8, 0xe9, 0x18, 0x66, 0x8f, 0x21, 0x3f, 0x02, 0x68, 0xa5, 0x6c, 0xde,
	0x5c, 0xcf, 0xc5, 0x28, 0xa8, 0x52, 0xf6, 0x1b, 0x7f, 0x45, 0x7f, 0x80, 0xb0, 0x44, 0xcb, 0xd6,
	0xcc, 0x32, 0x7a, 0xb6, 0x0c, 0xfe, 0x0f, 0xf2, 0x7b, 0xb4, 0x2c, 0xeb, 0x29, 0xe4, 0x0a, 0xa6,
	0x8c, 0x73, 0x34, 0x86, 0xbe, 0xe6, 0x79, 0x9f, 0x1e, 0xca, 0x7b, 0xe9, 0xb3, 0xb3, 0x96, 0x12,
	0xff, 0x11, 0xc0, 0xc9, 0xb0, 0x7a, 0x72, 0x06, 0x93, 0x5a, 0x17, 0xad, 0x7d, 0xb8, 0xa1, 0x73,
	0x8d, 0x8a, 0xd9, 0xdb, 0xce, 0x35, 0xdc, 0xb8, 0x77, 0x92, 0xc9, 0xc0, 0x49, 0xce, 0xc1, 0xd9,
	0x03, 0xbf, 0x33, 0x75, 0x39, 0xb0, 0x0b, 0x1f, 0x93, 0x25, 0xcc, 0x4d, 0x85, 0x5c, 0x6c, 0x04,
	0x17, 0x76, 0xe7, 0x3d, 0x23, 0xca, 0x86, 0x52, 0xfc, 0xd7, 0xa0, 0x10, 0x77, 0xe6, 0x7e, 0x8b,
	0xe0, 0x69, 0xb3, 0x7a, 0xf6, 0xd0, 0xac, 0x96, 0xee, 0x2f, 0x61, 0xb8, 0x16, 0x95, 0xdd, 0x5b,
	0xd9, 0x50, 0x22, 0x1f, 0xc0, 0x59, 0xa1, 0xe4, 0x4d, 0x3e, 0x5c, 0xd6, 0x94, 0xb9, 0x70, 0xfa,
	0x57, 0x83, 0xa5, 0xef, 0x00, 0x94, 0x4c, 0x48, 0xcb, 0x84, 0x44, 0xdd, 0x16, 0x3b, 0x50, 0x1c,
	0x6a, 0x1f, 0xe5, 0x58, 0x32, 0x51, 0xb4, 0x56, 0xb7, 0xd8, 0xeb, 0x5f, 0x3b, 0xd9, 0x55, 0x5c,
	0x08, 0x8e, 0xd2, 0x20, 0x9d, 0x35, 0x15, 0xb7, 0x61, 0xfc, 0x7b, 0x00, 0xa7, 0x0f, 0x7f, 0x14,
	0x77, 0x64, 0x8d, 0x6c, 0xed, 0x8f, 0x1c, 0x66, 0x7e, 0xec, 0x8c, 0x96, 0x6b, 0x64, 0xb6, 0x71,
	0xed, 0x30, 0x6b, 0x23, 0xf2, 0x06, 0x1c, 0xdf, 0x68, 0x26, 0xad, 0x3f, 0x6a, 0x98, 0x35, 0x81,
	0x5b, 0x5d, 0x57, 0x6b, 0xb7, 0xfa, 0xa8, 0x59, 0xdd, 0x44, 0x4e, 0x5f, 0x63, 0x81, 0x16, 0x3b,
	0xbb, 0x6e, 0xa2, 0x78, 0x0b, 0xe7, 0xfd, 0x3b, 0xf5, 0x72, 0xcb, 0x44, 0xc1, 0x56, 0x05, 0xb6,
	0x6f, 0x8d, 0x21, 0x3f, 0x3f, 0x7e, 0xb8, 0x5e, 0x1c, 0x7c, 0xdf, 0x8a, 0xe2, 0x89, 0xb7, 0xeb,
	0x0b, 0x58, 0xfc, 0xe3, 0x65, 0x3b, 0xec, 0xf7, 0x8e, 0x4b, 0x20, 0x8f, 0x77, 0x78, 0x92, 0xf1,
	0x3e, 0x2c, 0x78, 0xad, 0x35, 0x4a, 0x9b, 0x3f, 0x64, 0x9d, 0xb6, 0x72, 0x97, 0x7c, 0x0e, 0x61,
	0xbb, 0xc0, 0xd0, 0xc9, 0x72, 0xe2, 0xee, 0x6f, 0x17, 0xc7, 0x29, 0xbc, 0xde, 0x6d, 0xe7, 0xfe,
	0xd9, 0x97, 0x4a, 0x5a, 0x94, 0xd6, 0xd5, 0xc7, 0x9b, 0x61, 0xbb, 0x65, 0x17, 0x7e, 0xf9, 0xf9,
	0x2f, 0x2f, 0x6e, 0x84, 0xbd, 0xad, 0x57, 0x09, 0x57, 0x65, 0xea, 0x7a, 0x96, 0x76, 0x3d, 0x4b,
	0xff, 0xf5, 0xb3, 0x63, 0x35, 0xf5, 0x5f, 0x1b, 0x1f, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x02,
	0xa7, 0x15, 0x6f, 0xa2, 0x08, 0x00, 0x00,
}
