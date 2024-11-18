// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: json_report_spec.proto

package jsonreportspec

import (
	models "github.com/safedep/vet/gen/models"
	violations "github.com/safedep/vet/gen/violations"
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

type RemediationAdviceType int32

const (
	RemediationAdviceType_UnknownAdviceType       RemediationAdviceType = 0
	RemediationAdviceType_UpgradePackage          RemediationAdviceType = 1
	RemediationAdviceType_AlternatePopularPackage RemediationAdviceType = 2
	RemediationAdviceType_AlternateSecurePackage  RemediationAdviceType = 3
)

// Enum value maps for RemediationAdviceType.
var (
	RemediationAdviceType_name = map[int32]string{
		0: "UnknownAdviceType",
		1: "UpgradePackage",
		2: "AlternatePopularPackage",
		3: "AlternateSecurePackage",
	}
	RemediationAdviceType_value = map[string]int32{
		"UnknownAdviceType":       0,
		"UpgradePackage":          1,
		"AlternatePopularPackage": 2,
		"AlternateSecurePackage":  3,
	}
)

func (x RemediationAdviceType) Enum() *RemediationAdviceType {
	p := new(RemediationAdviceType)
	*p = x
	return p
}

func (x RemediationAdviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemediationAdviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_json_report_spec_proto_enumTypes[0].Descriptor()
}

func (RemediationAdviceType) Type() protoreflect.EnumType {
	return &file_json_report_spec_proto_enumTypes[0]
}

func (x RemediationAdviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RemediationAdviceType.Descriptor instead.
func (RemediationAdviceType) EnumDescriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{0}
}

type ReportThreat_Confidence int32

const (
	ReportThreat_UnknownConfidence ReportThreat_Confidence = 0
	ReportThreat_High              ReportThreat_Confidence = 1
	ReportThreat_Medium            ReportThreat_Confidence = 2
	ReportThreat_Low               ReportThreat_Confidence = 3
)

// Enum value maps for ReportThreat_Confidence.
var (
	ReportThreat_Confidence_name = map[int32]string{
		0: "UnknownConfidence",
		1: "High",
		2: "Medium",
		3: "Low",
	}
	ReportThreat_Confidence_value = map[string]int32{
		"UnknownConfidence": 0,
		"High":              1,
		"Medium":            2,
		"Low":               3,
	}
)

func (x ReportThreat_Confidence) Enum() *ReportThreat_Confidence {
	p := new(ReportThreat_Confidence)
	*p = x
	return p
}

func (x ReportThreat_Confidence) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportThreat_Confidence) Descriptor() protoreflect.EnumDescriptor {
	return file_json_report_spec_proto_enumTypes[1].Descriptor()
}

func (ReportThreat_Confidence) Type() protoreflect.EnumType {
	return &file_json_report_spec_proto_enumTypes[1]
}

func (x ReportThreat_Confidence) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportThreat_Confidence.Descriptor instead.
func (ReportThreat_Confidence) EnumDescriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{1, 0}
}

type ReportThreat_Source int32

const (
	ReportThreat_UnknownSource ReportThreat_Source = 0
	ReportThreat_CWE           ReportThreat_Source = 1
)

// Enum value maps for ReportThreat_Source.
var (
	ReportThreat_Source_name = map[int32]string{
		0: "UnknownSource",
		1: "CWE",
	}
	ReportThreat_Source_value = map[string]int32{
		"UnknownSource": 0,
		"CWE":           1,
	}
)

func (x ReportThreat_Source) Enum() *ReportThreat_Source {
	p := new(ReportThreat_Source)
	*p = x
	return p
}

func (x ReportThreat_Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportThreat_Source) Descriptor() protoreflect.EnumDescriptor {
	return file_json_report_spec_proto_enumTypes[2].Descriptor()
}

func (ReportThreat_Source) Type() protoreflect.EnumType {
	return &file_json_report_spec_proto_enumTypes[2]
}

func (x ReportThreat_Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportThreat_Source.Descriptor instead.
func (ReportThreat_Source) EnumDescriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{1, 1}
}

type ReportThreat_SubjectType int32

const (
	ReportThreat_UnknownSubject ReportThreat_SubjectType = 0
	ReportThreat_Package        ReportThreat_SubjectType = 1
	ReportThreat_Manifest       ReportThreat_SubjectType = 2
)

// Enum value maps for ReportThreat_SubjectType.
var (
	ReportThreat_SubjectType_name = map[int32]string{
		0: "UnknownSubject",
		1: "Package",
		2: "Manifest",
	}
	ReportThreat_SubjectType_value = map[string]int32{
		"UnknownSubject": 0,
		"Package":        1,
		"Manifest":       2,
	}
)

func (x ReportThreat_SubjectType) Enum() *ReportThreat_SubjectType {
	p := new(ReportThreat_SubjectType)
	*p = x
	return p
}

func (x ReportThreat_SubjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportThreat_SubjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_json_report_spec_proto_enumTypes[3].Descriptor()
}

func (ReportThreat_SubjectType) Type() protoreflect.EnumType {
	return &file_json_report_spec_proto_enumTypes[3]
}

func (x ReportThreat_SubjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportThreat_SubjectType.Descriptor instead.
func (ReportThreat_SubjectType) EnumDescriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{1, 2}
}

type ReportThreat_ReportThreatId int32

const (
	ReportThreat_UnknownReportThreatId ReportThreat_ReportThreatId = 0
	ReportThreat_LockfilePoisoning     ReportThreat_ReportThreatId = 1
)

// Enum value maps for ReportThreat_ReportThreatId.
var (
	ReportThreat_ReportThreatId_name = map[int32]string{
		0: "UnknownReportThreatId",
		1: "LockfilePoisoning",
	}
	ReportThreat_ReportThreatId_value = map[string]int32{
		"UnknownReportThreatId": 0,
		"LockfilePoisoning":     1,
	}
)

func (x ReportThreat_ReportThreatId) Enum() *ReportThreat_ReportThreatId {
	p := new(ReportThreat_ReportThreatId)
	*p = x
	return p
}

func (x ReportThreat_ReportThreatId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportThreat_ReportThreatId) Descriptor() protoreflect.EnumDescriptor {
	return file_json_report_spec_proto_enumTypes[4].Descriptor()
}

func (ReportThreat_ReportThreatId) Type() protoreflect.EnumType {
	return &file_json_report_spec_proto_enumTypes[4]
}

func (x ReportThreat_ReportThreatId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportThreat_ReportThreatId.Descriptor instead.
func (ReportThreat_ReportThreatId) EnumDescriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{1, 3}
}

type RemediationAdvice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                          RemediationAdviceType `protobuf:"varint,1,opt,name=type,proto3,enum=RemediationAdviceType" json:"type,omitempty"`
	Package                       *models.Package       `protobuf:"bytes,2,opt,name=package,proto3" json:"package,omitempty"`
	TargetPackageName             string                `protobuf:"bytes,3,opt,name=target_package_name,json=targetPackageName,proto3" json:"target_package_name,omitempty"`
	TargetPackageVersion          string                `protobuf:"bytes,4,opt,name=target_package_version,json=targetPackageVersion,proto3" json:"target_package_version,omitempty"`
	TargetAlternatePackageName    string                `protobuf:"bytes,5,opt,name=target_alternate_package_name,json=targetAlternatePackageName,proto3" json:"target_alternate_package_name,omitempty"`
	TargetAlternatePackageVersion string                `protobuf:"bytes,6,opt,name=target_alternate_package_version,json=targetAlternatePackageVersion,proto3" json:"target_alternate_package_version,omitempty"`
}

func (x *RemediationAdvice) Reset() {
	*x = RemediationAdvice{}
	mi := &file_json_report_spec_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemediationAdvice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemediationAdvice) ProtoMessage() {}

func (x *RemediationAdvice) ProtoReflect() protoreflect.Message {
	mi := &file_json_report_spec_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemediationAdvice.ProtoReflect.Descriptor instead.
func (*RemediationAdvice) Descriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{0}
}

func (x *RemediationAdvice) GetType() RemediationAdviceType {
	if x != nil {
		return x.Type
	}
	return RemediationAdviceType_UnknownAdviceType
}

func (x *RemediationAdvice) GetPackage() *models.Package {
	if x != nil {
		return x.Package
	}
	return nil
}

func (x *RemediationAdvice) GetTargetPackageName() string {
	if x != nil {
		return x.TargetPackageName
	}
	return ""
}

func (x *RemediationAdvice) GetTargetPackageVersion() string {
	if x != nil {
		return x.TargetPackageVersion
	}
	return ""
}

func (x *RemediationAdvice) GetTargetAlternatePackageName() string {
	if x != nil {
		return x.TargetAlternatePackageName
	}
	return ""
}

func (x *RemediationAdvice) GetTargetAlternatePackageVersion() string {
	if x != nil {
		return x.TargetAlternatePackageVersion
	}
	return ""
}

// We are introducing the concept of Threat as a reporting entity so
// that we can report threats like lockfile poisoning using a standard schema.
// But why do we need threats? Why not just use vet's paradigm of policy over
// enriched packages? The reason is, there are threats that are applicable in
// an environment, against a manifest or other entities or even group of entities.
// Hence it is required to introduce a threat as a reporting entity so that external
// tools can consume vet's reports and take actions based on the threats.
type ReportThreat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          ReportThreat_ReportThreatId `protobuf:"varint,1,opt,name=id,proto3,enum=ReportThreat_ReportThreatId" json:"id,omitempty"`
	InstanceId  string                      `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId,omitempty"` // Unique threat instance ID per (ID, SubjectType, Subject) tuple
	Message     string                      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	SubjectType ReportThreat_SubjectType    `protobuf:"varint,4,opt,name=subject_type,json=subjectType,proto3,enum=ReportThreat_SubjectType" json:"subject_type,omitempty"`
	Subject     string                      `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Confidence  ReportThreat_Confidence     `protobuf:"varint,6,opt,name=confidence,proto3,enum=ReportThreat_Confidence" json:"confidence,omitempty"`
	Source      ReportThreat_Source         `protobuf:"varint,7,opt,name=source,proto3,enum=ReportThreat_Source" json:"source,omitempty"`
	SourceId    string                      `protobuf:"bytes,8,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
}

func (x *ReportThreat) Reset() {
	*x = ReportThreat{}
	mi := &file_json_report_spec_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportThreat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportThreat) ProtoMessage() {}

func (x *ReportThreat) ProtoReflect() protoreflect.Message {
	mi := &file_json_report_spec_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportThreat.ProtoReflect.Descriptor instead.
func (*ReportThreat) Descriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{1}
}

func (x *ReportThreat) GetId() ReportThreat_ReportThreatId {
	if x != nil {
		return x.Id
	}
	return ReportThreat_UnknownReportThreatId
}

func (x *ReportThreat) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *ReportThreat) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ReportThreat) GetSubjectType() ReportThreat_SubjectType {
	if x != nil {
		return x.SubjectType
	}
	return ReportThreat_UnknownSubject
}

func (x *ReportThreat) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ReportThreat) GetConfidence() ReportThreat_Confidence {
	if x != nil {
		return x.Confidence
	}
	return ReportThreat_UnknownConfidence
}

func (x *ReportThreat) GetSource() ReportThreat_Source {
	if x != nil {
		return x.Source
	}
	return ReportThreat_UnknownSource
}

func (x *ReportThreat) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

type PackageManifestReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ecosystem   models.Ecosystem `protobuf:"varint,2,opt,name=ecosystem,proto3,enum=Ecosystem" json:"ecosystem,omitempty"`
	Path        string           `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Threats     []*ReportThreat  `protobuf:"bytes,4,rep,name=threats,proto3" json:"threats,omitempty"`
	DisplayPath string           `protobuf:"bytes,5,opt,name=display_path,json=displayPath,proto3" json:"display_path,omitempty"`
	SourceType  string           `protobuf:"bytes,6,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	Namespace   string           `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *PackageManifestReport) Reset() {
	*x = PackageManifestReport{}
	mi := &file_json_report_spec_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PackageManifestReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageManifestReport) ProtoMessage() {}

func (x *PackageManifestReport) ProtoReflect() protoreflect.Message {
	mi := &file_json_report_spec_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageManifestReport.ProtoReflect.Descriptor instead.
func (*PackageManifestReport) Descriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{2}
}

func (x *PackageManifestReport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PackageManifestReport) GetEcosystem() models.Ecosystem {
	if x != nil {
		return x.Ecosystem
	}
	return models.Ecosystem(0)
}

func (x *PackageManifestReport) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PackageManifestReport) GetThreats() []*ReportThreat {
	if x != nil {
		return x.Threats
	}
	return nil
}

func (x *PackageManifestReport) GetDisplayPath() string {
	if x != nil {
		return x.DisplayPath
	}
	return ""
}

func (x *PackageManifestReport) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *PackageManifestReport) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// PackageReport represents the first class entity for which we have different type
// of reporting information
type PackageReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *models.Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	// The manifests identified by IDs where this package belongs to
	Manifests  []string                `protobuf:"bytes,2,rep,name=manifests,proto3" json:"manifests,omitempty"`
	Violations []*violations.Violation `protobuf:"bytes,3,rep,name=violations,proto3" json:"violations,omitempty"`
	Advices    []*RemediationAdvice    `protobuf:"bytes,4,rep,name=advices,proto3" json:"advices,omitempty"`
	// Insights data
	Vulnerabilities []*models.InsightVulnerability `protobuf:"bytes,5,rep,name=vulnerabilities,proto3" json:"vulnerabilities,omitempty"`
	Licenses        []*models.InsightLicenseInfo   `protobuf:"bytes,6,rep,name=licenses,proto3" json:"licenses,omitempty"`
	// Threats
	Threats []*ReportThreat `protobuf:"bytes,7,rep,name=threats,proto3" json:"threats,omitempty"`
}

func (x *PackageReport) Reset() {
	*x = PackageReport{}
	mi := &file_json_report_spec_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PackageReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageReport) ProtoMessage() {}

func (x *PackageReport) ProtoReflect() protoreflect.Message {
	mi := &file_json_report_spec_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageReport.ProtoReflect.Descriptor instead.
func (*PackageReport) Descriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{3}
}

func (x *PackageReport) GetPackage() *models.Package {
	if x != nil {
		return x.Package
	}
	return nil
}

func (x *PackageReport) GetManifests() []string {
	if x != nil {
		return x.Manifests
	}
	return nil
}

func (x *PackageReport) GetViolations() []*violations.Violation {
	if x != nil {
		return x.Violations
	}
	return nil
}

func (x *PackageReport) GetAdvices() []*RemediationAdvice {
	if x != nil {
		return x.Advices
	}
	return nil
}

func (x *PackageReport) GetVulnerabilities() []*models.InsightVulnerability {
	if x != nil {
		return x.Vulnerabilities
	}
	return nil
}

func (x *PackageReport) GetLicenses() []*models.InsightLicenseInfo {
	if x != nil {
		return x.Licenses
	}
	return nil
}

func (x *PackageReport) GetThreats() []*ReportThreat {
	if x != nil {
		return x.Threats
	}
	return nil
}

type ReportMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToolName    string `protobuf:"bytes,1,opt,name=tool_name,json=toolName,proto3" json:"tool_name,omitempty"`
	ToolVersion string `protobuf:"bytes,2,opt,name=tool_version,json=toolVersion,proto3" json:"tool_version,omitempty"`
	CreatedAt   string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ReportMeta) Reset() {
	*x = ReportMeta{}
	mi := &file_json_report_spec_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportMeta) ProtoMessage() {}

func (x *ReportMeta) ProtoReflect() protoreflect.Message {
	mi := &file_json_report_spec_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportMeta.ProtoReflect.Descriptor instead.
func (*ReportMeta) Descriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{4}
}

func (x *ReportMeta) GetToolName() string {
	if x != nil {
		return x.ToolName
	}
	return ""
}

func (x *ReportMeta) GetToolVersion() string {
	if x != nil {
		return x.ToolVersion
	}
	return ""
}

func (x *ReportMeta) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *ReportMeta              `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Manifests []*PackageManifestReport `protobuf:"bytes,2,rep,name=manifests,proto3" json:"manifests,omitempty"`
	Packages  []*PackageReport         `protobuf:"bytes,3,rep,name=packages,proto3" json:"packages,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	mi := &file_json_report_spec_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_json_report_spec_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_json_report_spec_proto_rawDescGZIP(), []int{5}
}

func (x *Report) GetMeta() *ReportMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Report) GetManifests() []*PackageManifestReport {
	if x != nil {
		return x.Manifests
	}
	return nil
}

func (x *Report) GetPackages() []*PackageReport {
	if x != nil {
		return x.Packages
	}
	return nil
}

var File_json_report_spec_proto protoreflect.FileDescriptor

var file_json_report_spec_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x76,
	0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd5, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x64, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x52, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x64, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x1d, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x1a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47,
	0x0a, 0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xbf, 0x04, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74,
	0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x74, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x42, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x6f, 0x77, 0x10,
	0x03, 0x22, 0x24, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x43, 0x57, 0x45, 0x10, 0x01, 0x22, 0x3c, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x10, 0x02, 0x22, 0x42, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x64,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x6f,
	0x69, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x22, 0xf0, 0x01, 0x0a, 0x15, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x74, 0x52, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xc6, 0x02, 0x0a,
	0x0d, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x22,
	0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x2a, 0x0a, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x07,
	0x61, 0x64, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x52, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x07, 0x61, 0x64, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x76, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x56, 0x75, 0x6c,
	0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x76, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x52, 0x07, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x74, 0x73, 0x22, 0x6b, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x34,
	0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73,
	0x2a, 0x7b, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x64, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x41, 0x64, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x10,
	0x02, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x10, 0x03, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x66, 0x65,
	0x64, 0x65, 0x70, 0x2f, 0x76, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_json_report_spec_proto_rawDescOnce sync.Once
	file_json_report_spec_proto_rawDescData = file_json_report_spec_proto_rawDesc
)

func file_json_report_spec_proto_rawDescGZIP() []byte {
	file_json_report_spec_proto_rawDescOnce.Do(func() {
		file_json_report_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_json_report_spec_proto_rawDescData)
	})
	return file_json_report_spec_proto_rawDescData
}

var file_json_report_spec_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_json_report_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_json_report_spec_proto_goTypes = []any{
	(RemediationAdviceType)(0),          // 0: RemediationAdviceType
	(ReportThreat_Confidence)(0),        // 1: ReportThreat.Confidence
	(ReportThreat_Source)(0),            // 2: ReportThreat.Source
	(ReportThreat_SubjectType)(0),       // 3: ReportThreat.SubjectType
	(ReportThreat_ReportThreatId)(0),    // 4: ReportThreat.ReportThreatId
	(*RemediationAdvice)(nil),           // 5: RemediationAdvice
	(*ReportThreat)(nil),                // 6: ReportThreat
	(*PackageManifestReport)(nil),       // 7: PackageManifestReport
	(*PackageReport)(nil),               // 8: PackageReport
	(*ReportMeta)(nil),                  // 9: ReportMeta
	(*Report)(nil),                      // 10: Report
	(*models.Package)(nil),              // 11: Package
	(models.Ecosystem)(0),               // 12: Ecosystem
	(*violations.Violation)(nil),        // 13: Violation
	(*models.InsightVulnerability)(nil), // 14: InsightVulnerability
	(*models.InsightLicenseInfo)(nil),   // 15: InsightLicenseInfo
}
var file_json_report_spec_proto_depIdxs = []int32{
	0,  // 0: RemediationAdvice.type:type_name -> RemediationAdviceType
	11, // 1: RemediationAdvice.package:type_name -> Package
	4,  // 2: ReportThreat.id:type_name -> ReportThreat.ReportThreatId
	3,  // 3: ReportThreat.subject_type:type_name -> ReportThreat.SubjectType
	1,  // 4: ReportThreat.confidence:type_name -> ReportThreat.Confidence
	2,  // 5: ReportThreat.source:type_name -> ReportThreat.Source
	12, // 6: PackageManifestReport.ecosystem:type_name -> Ecosystem
	6,  // 7: PackageManifestReport.threats:type_name -> ReportThreat
	11, // 8: PackageReport.package:type_name -> Package
	13, // 9: PackageReport.violations:type_name -> Violation
	5,  // 10: PackageReport.advices:type_name -> RemediationAdvice
	14, // 11: PackageReport.vulnerabilities:type_name -> InsightVulnerability
	15, // 12: PackageReport.licenses:type_name -> InsightLicenseInfo
	6,  // 13: PackageReport.threats:type_name -> ReportThreat
	9,  // 14: Report.meta:type_name -> ReportMeta
	7,  // 15: Report.manifests:type_name -> PackageManifestReport
	8,  // 16: Report.packages:type_name -> PackageReport
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_json_report_spec_proto_init() }
func file_json_report_spec_proto_init() {
	if File_json_report_spec_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_json_report_spec_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_json_report_spec_proto_goTypes,
		DependencyIndexes: file_json_report_spec_proto_depIdxs,
		EnumInfos:         file_json_report_spec_proto_enumTypes,
		MessageInfos:      file_json_report_spec_proto_msgTypes,
	}.Build()
	File_json_report_spec_proto = out.File
	file_json_report_spec_proto_rawDesc = nil
	file_json_report_spec_proto_goTypes = nil
	file_json_report_spec_proto_depIdxs = nil
}
