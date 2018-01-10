// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rbac/v1alpha1/rbac.proto

/*
Package istio_rbac_v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	rbac/v1alpha1/rbac.proto

It has these top-level messages:
	ServiceRole
	AccessRule
	ServiceRoleBinding
	Subject
	RoleRef
*/
package istio_rbac_v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ServiceRole specification contains a list of access rules (permissions).
// This represent the "Spec" part of the ServiceRole object. The name and namespace
// of the ServiceRole is specified in "metadata" section of the ServiceRole object.
type ServiceRole struct {
	// Required. The set of access rules (permissions) that the role has.
	Rules []*AccessRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *ServiceRole) Reset()                    { *m = ServiceRole{} }
func (m *ServiceRole) String() string            { return proto.CompactTextString(m) }
func (*ServiceRole) ProtoMessage()               {}
func (*ServiceRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceRole) GetRules() []*AccessRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// AccessRule defines a permission to access a list of services.
type AccessRule struct {
	// Required. A list of service names.
	// Exact match, prefix match, and suffix match are supported for service names.
	// For example, the service name "bookstore.mtv.cluster.local" matches
	// "bookstore.mtv.cluster.local" (exact match), or "bookstore*" (prefix match),
	// or "*.mtv.cluster.local" (suffix match).
	// If set to ["*"], it refers to all services in the namespace.
	Services []string `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
	// Optional. A list of HTTP paths.
	// Exact match, prefix match, and suffix match are supported for paths.
	// For example, the path "/books/review" matches
	// "/books/review" (exact match), or "/books/*" (prefix match),
	// or "*/review" (suffix match).
	// If not specified, it applies to any path.
	Paths []string `protobuf:"bytes,2,rep,name=paths" json:"paths,omitempty"`
	// Required. A list of HTTP methods (e.g., "GET", "POST") or gRPC methods.
	// gRPC methods must be presented as fully-qualified name in the form of
	// packageName.serviceName/methodName.
	// If set to ["*"], it applies to any method.
	Methods []string `protobuf:"bytes,3,rep,name=methods" json:"methods,omitempty"`
	// Optional. Extra constraints in the ServiceRole specification.
	// The above ServiceRole examples shows an example of constraint "version".
	Constraints []*AccessRule_Constraint `protobuf:"bytes,4,rep,name=constraints" json:"constraints,omitempty"`
}

func (m *AccessRule) Reset()                    { *m = AccessRule{} }
func (m *AccessRule) String() string            { return proto.CompactTextString(m) }
func (*AccessRule) ProtoMessage()               {}
func (*AccessRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AccessRule) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *AccessRule) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *AccessRule) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *AccessRule) GetConstraints() []*AccessRule_Constraint {
	if m != nil {
		return m.Constraints
	}
	return nil
}

// Definition of a custom constraint.
type AccessRule_Constraint struct {
	// The name of the constraint.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The list of valid values for the constraint.
	// Exact match, prefix match, and suffix match are supported for constraint values.
	// For example, the value "v1alpha2" matches
	// "v1alpha2" (exact match), or "v1*" (prefix match),
	// or "*alpha2" (suffix match).
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *AccessRule_Constraint) Reset()                    { *m = AccessRule_Constraint{} }
func (m *AccessRule_Constraint) String() string            { return proto.CompactTextString(m) }
func (*AccessRule_Constraint) ProtoMessage()               {}
func (*AccessRule_Constraint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *AccessRule_Constraint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccessRule_Constraint) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// ServiceRoleBinding assigns a ServiceRole to a list of subjects.
// This represent the "Spec" part of the ServiceRoleBinding object. The name and namespace
// of the ServiceRoleBinding is specified in "metadata" section of the ServiceRoleBinding
// object.
type ServiceRoleBinding struct {
	// Required. List of subjects that are assigned the ServiceRole object.
	Subjects []*Subject `protobuf:"bytes,1,rep,name=subjects" json:"subjects,omitempty"`
	// Required. Reference to the ServiceRole object.
	RoleRef *RoleRef `protobuf:"bytes,2,opt,name=roleRef" json:"roleRef,omitempty"`
}

func (m *ServiceRoleBinding) Reset()                    { *m = ServiceRoleBinding{} }
func (m *ServiceRoleBinding) String() string            { return proto.CompactTextString(m) }
func (*ServiceRoleBinding) ProtoMessage()               {}
func (*ServiceRoleBinding) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServiceRoleBinding) GetSubjects() []*Subject {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *ServiceRoleBinding) GetRoleRef() *RoleRef {
	if m != nil {
		return m.RoleRef
	}
	return nil
}

// Subject defines an identity or a group of identities.
type Subject struct {
	// Optional. The user name/ID that the subject represents.
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	// Optional. The group that the subject belongs to.
	Group string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	// Optional. The set of properties that identify the subject.
	// In the above ServiceRoleBinding example, the second subject has two properties:
	//     service: "reviews"
	//     namespace: "abc"
	Properties map[string]string `protobuf:"bytes,3,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Subject) Reset()                    { *m = Subject{} }
func (m *Subject) String() string            { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()               {}
func (*Subject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Subject) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Subject) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Subject) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

// RoleRef refers to a role object.
type RoleRef struct {
	// Required. The type of the role being referenced.
	// Currently, "ServiceRole" is the only supported value for "kind".
	Kind string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	// Required. The name of the ServiceRole object being referenced.
	// The ServiceRole object must be in the same namespace as the ServiceRoleBinding
	// object.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *RoleRef) Reset()                    { *m = RoleRef{} }
func (m *RoleRef) String() string            { return proto.CompactTextString(m) }
func (*RoleRef) ProtoMessage()               {}
func (*RoleRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RoleRef) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *RoleRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceRole)(nil), "istio.rbac.v1alpha1.ServiceRole")
	proto.RegisterType((*AccessRule)(nil), "istio.rbac.v1alpha1.AccessRule")
	proto.RegisterType((*AccessRule_Constraint)(nil), "istio.rbac.v1alpha1.AccessRule.Constraint")
	proto.RegisterType((*ServiceRoleBinding)(nil), "istio.rbac.v1alpha1.ServiceRoleBinding")
	proto.RegisterType((*Subject)(nil), "istio.rbac.v1alpha1.Subject")
	proto.RegisterType((*RoleRef)(nil), "istio.rbac.v1alpha1.RoleRef")
}

func init() { proto.RegisterFile("rbac/v1alpha1/rbac.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdd, 0xca, 0x13, 0x31,
	0x10, 0x25, 0xdb, 0x9f, 0x6d, 0x67, 0x2f, 0x94, 0x28, 0x12, 0x8a, 0x60, 0xd9, 0xab, 0x22, 0xb2,
	0xa5, 0x15, 0xa5, 0x08, 0x5e, 0xf8, 0x77, 0xd7, 0x0b, 0x49, 0x9f, 0x60, 0xbb, 0x1d, 0xdb, 0xd8,
	0xed, 0x66, 0x49, 0xb2, 0x85, 0xbe, 0x80, 0xcf, 0xe5, 0xa3, 0xf8, 0x28, 0x92, 0x64, 0xb3, 0x2d,
	0x52, 0xbf, 0xef, 0x6e, 0xce, 0xcc, 0x99, 0xe1, 0x9c, 0x93, 0x00, 0x53, 0xdb, 0xbc, 0x98, 0x9f,
	0x17, 0x79, 0x59, 0x1f, 0xf2, 0xc5, 0xdc, 0xa2, 0xac, 0x56, 0xd2, 0x48, 0xfa, 0x4c, 0x68, 0x23,
	0x64, 0xe6, 0x3a, 0x61, 0x9e, 0x7e, 0x85, 0x64, 0x83, 0xea, 0x2c, 0x0a, 0xe4, 0xb2, 0x44, 0xfa,
	0x0e, 0x06, 0xaa, 0x29, 0x51, 0x33, 0x32, 0xed, 0xcd, 0x92, 0xe5, 0xab, 0xec, 0xce, 0x4e, 0xf6,
	0xa9, 0x28, 0x50, 0x6b, 0xde, 0x94, 0xc8, 0x3d, 0x3b, 0xfd, 0x43, 0x00, 0xae, 0x5d, 0x3a, 0x81,
	0x91, 0xf6, 0x47, 0xfd, 0xa1, 0x31, 0xef, 0x30, 0x7d, 0x0e, 0x83, 0x3a, 0x37, 0x07, 0xcd, 0x22,
	0x37, 0xf0, 0x80, 0x32, 0x88, 0x4f, 0x68, 0x0e, 0x72, 0xa7, 0x59, 0xcf, 0xf5, 0x03, 0xa4, 0x6b,
	0x48, 0x0a, 0x59, 0x69, 0xa3, 0x72, 0x51, 0x19, 0xcd, 0xfa, 0x4e, 0xd7, 0xeb, 0x47, 0x74, 0x65,
	0x5f, 0xba, 0x15, 0x7e, 0xbb, 0x3e, 0x59, 0x01, 0x5c, 0x47, 0x94, 0x42, 0xbf, 0xca, 0x4f, 0xc8,
	0xc8, 0x94, 0xcc, 0xc6, 0xdc, 0xd5, 0xf4, 0x05, 0x0c, 0xcf, 0x79, 0xd9, 0x60, 0x10, 0xd8, 0xa2,
	0xf4, 0x17, 0x01, 0x7a, 0x93, 0xd4, 0x67, 0x51, 0xed, 0x44, 0xb5, 0xa7, 0x2b, 0x18, 0xe9, 0x66,
	0xfb, 0x13, 0x0b, 0x13, 0x32, 0x7b, 0x79, 0x57, 0xdb, 0xc6, 0x93, 0x78, 0xc7, 0xa6, 0xef, 0x21,
	0x56, 0xb2, 0x44, 0x8e, 0x3f, 0x58, 0x34, 0x25, 0xff, 0x5d, 0xe4, 0x9e, 0xc3, 0x03, 0x39, 0xfd,
	0x4d, 0x20, 0x6e, 0xaf, 0x59, 0x03, 0x8d, 0x46, 0x15, 0x0c, 0xd8, 0xda, 0x06, 0xbc, 0x57, 0xb2,
	0xa9, 0xdd, 0xd5, 0x31, 0xf7, 0x80, 0xae, 0x01, 0x6a, 0x25, 0x6b, 0x54, 0x46, 0xa0, 0xcf, 0x38,
	0x59, 0xbe, 0x79, 0x48, 0x69, 0xf6, 0xbd, 0xa3, 0x7f, 0xab, 0x8c, 0xba, 0xf0, 0x9b, 0xfd, 0xc9,
	0x47, 0x78, 0xf2, 0xcf, 0x98, 0x3e, 0x85, 0xde, 0x11, 0x2f, 0xad, 0x12, 0x5b, 0x5a, 0x21, 0x2e,
	0xbb, 0x20, 0xc4, 0x81, 0x0f, 0xd1, 0x8a, 0xa4, 0x0b, 0x88, 0x5b, 0x5b, 0xd6, 0xc1, 0x51, 0x54,
	0xbb, 0xe0, 0xc0, 0xd6, 0xdd, 0xb3, 0x44, 0xd7, 0x67, 0xd9, 0x0e, 0xdd, 0x1f, 0x7e, 0xfb, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0x1a, 0xbb, 0xce, 0xdf, 0x02, 0x00, 0x00,
}