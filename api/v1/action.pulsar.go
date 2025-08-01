// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package orbiterv1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Action            protoreflect.MessageDescriptor
	fd_Action_id         protoreflect.FieldDescriptor
	fd_Action_attributes protoreflect.FieldDescriptor
)

func init() {
	file_noble_orbiter_v1_action_proto_init()
	md_Action = File_noble_orbiter_v1_action_proto.Messages().ByName("Action")
	fd_Action_id = md_Action.Fields().ByName("id")
	fd_Action_attributes = md_Action.Fields().ByName("attributes")
}

var _ protoreflect.Message = (*fastReflection_Action)(nil)

type fastReflection_Action Action

func (x *Action) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Action)(x)
}

func (x *Action) slowProtoReflect() protoreflect.Message {
	mi := &file_noble_orbiter_v1_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Action_messageType fastReflection_Action_messageType
var _ protoreflect.MessageType = fastReflection_Action_messageType{}

type fastReflection_Action_messageType struct{}

func (x fastReflection_Action_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Action)(nil)
}
func (x fastReflection_Action_messageType) New() protoreflect.Message {
	return new(fastReflection_Action)
}
func (x fastReflection_Action_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Action
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Action) Descriptor() protoreflect.MessageDescriptor {
	return md_Action
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Action) Type() protoreflect.MessageType {
	return _fastReflection_Action_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Action) New() protoreflect.Message {
	return new(fastReflection_Action)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Action) Interface() protoreflect.ProtoMessage {
	return (*Action)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Action) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Id))
		if !f(fd_Action_id, value) {
			return
		}
	}
	if x.Attributes != nil {
		value := protoreflect.ValueOfMessage(x.Attributes.ProtoReflect())
		if !f(fd_Action_attributes, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Action) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "noble.orbiter.v1.Action.id":
		return x.Id != 0
	case "noble.orbiter.v1.Action.attributes":
		return x.Attributes != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.orbiter.v1.Action"))
		}
		panic(fmt.Errorf("message noble.orbiter.v1.Action does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Action) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "noble.orbiter.v1.Action.id":
		x.Id = 0
	case "noble.orbiter.v1.Action.attributes":
		x.Attributes = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.orbiter.v1.Action"))
		}
		panic(fmt.Errorf("message noble.orbiter.v1.Action does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Action) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "noble.orbiter.v1.Action.id":
		value := x.Id
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "noble.orbiter.v1.Action.attributes":
		value := x.Attributes
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.orbiter.v1.Action"))
		}
		panic(fmt.Errorf("message noble.orbiter.v1.Action does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Action) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "noble.orbiter.v1.Action.id":
		x.Id = (ActionID)(value.Enum())
	case "noble.orbiter.v1.Action.attributes":
		x.Attributes = value.Message().Interface().(*anypb.Any)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.orbiter.v1.Action"))
		}
		panic(fmt.Errorf("message noble.orbiter.v1.Action does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Action) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "noble.orbiter.v1.Action.attributes":
		if x.Attributes == nil {
			x.Attributes = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.Attributes.ProtoReflect())
	case "noble.orbiter.v1.Action.id":
		panic(fmt.Errorf("field id of message noble.orbiter.v1.Action is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.orbiter.v1.Action"))
		}
		panic(fmt.Errorf("message noble.orbiter.v1.Action does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Action) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "noble.orbiter.v1.Action.id":
		return protoreflect.ValueOfEnum(0)
	case "noble.orbiter.v1.Action.attributes":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.orbiter.v1.Action"))
		}
		panic(fmt.Errorf("message noble.orbiter.v1.Action does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Action) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in noble.orbiter.v1.Action", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Action) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Action) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Action) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Action) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Action)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		if x.Attributes != nil {
			l = options.Size(x.Attributes)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Action)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Attributes != nil {
			encoded, err := options.Marshal(x.Attributes)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Action)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Action: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Action: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= ActionID(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Attributes == nil {
					x.Attributes = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Attributes); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: noble/orbiter/v1/action.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// buf:lint:ignore ENUM_VALUE_PREFIX
// ActionID identifies the type of action to be performed on a transfer.
type ActionID int32

const (
	// ACTION_UNSUPPORTED represents an unknown or unsupported action type.
	// This is the default zero value and should not be used in production.
	ActionID_ACTION_UNSUPPORTED ActionID = 0
	// ACTION_FEE represents a fee collection action.
	ActionID_ACTION_FEE ActionID = 1
	// ACTION_SWAP represents a token swap action.
	ActionID_ACTION_SWAP ActionID = 2
)

// Enum value maps for ActionID.
var (
	ActionID_name = map[int32]string{
		0: "ACTION_UNSUPPORTED",
		1: "ACTION_FEE",
		2: "ACTION_SWAP",
	}
	ActionID_value = map[string]int32{
		"ACTION_UNSUPPORTED": 0,
		"ACTION_FEE":         1,
		"ACTION_SWAP":        2,
	}
)

func (x ActionID) Enum() *ActionID {
	p := new(ActionID)
	*p = x
	return p
}

func (x ActionID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionID) Descriptor() protoreflect.EnumDescriptor {
	return file_noble_orbiter_v1_action_proto_enumTypes[0].Descriptor()
}

func (ActionID) Type() protoreflect.EnumType {
	return &file_noble_orbiter_v1_action_proto_enumTypes[0]
}

func (x ActionID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionID.Descriptor instead.
func (ActionID) EnumDescriptor() ([]byte, []int) {
	return file_noble_orbiter_v1_action_proto_rawDescGZIP(), []int{0}
}

// Action represents a pre-processing step to be executed on a transfer
// before routing to the destination counterparty.
//
// Actions contain an ID that specifies the type of action and attributes
// that provide action-specific configuration. The attributes field uses
// the Any type to allow for different action implementations while
// maintaining type safety through interface constraints.
type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id specifies the type of action to be performed.
	// This determines how the attributes field should be interpreted.
	Id ActionID `protobuf:"varint,1,opt,name=id,proto3,enum=noble.orbiter.v1.ActionID" json:"id,omitempty"`
	// attributes contains the action-specific configuration data.
	// The actual type depends on the action ID and must implement
	// the `ActionAttributes` interface.
	Attributes *anypb.Any `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noble_orbiter_v1_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_noble_orbiter_v1_action_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetId() ActionID {
	if x != nil {
		return x.Id
	}
	return ActionID_ACTION_UNSUPPORTED
}

func (x *Action) GetAttributes() *anypb.Any {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_noble_orbiter_v1_action_proto protoreflect.FileDescriptor

var file_noble_orbiter_v1_action_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x2e, 0x6f, 0x72, 0x62,
	0x69, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x5b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x25,
	0xca, 0xb4, 0x2d, 0x21, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x2a, 0x49, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x45, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x57, 0x41, 0x50, 0x10, 0x02, 0x1a, 0x04, 0x88, 0xa3,
	0x1e, 0x00, 0x42, 0xb1, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x62, 0x6c, 0x65,
	0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x6f, 0x72, 0x62, 0x69,
	0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x62, 0x6c,
	0x65, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x72, 0x62,
	0x69, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x4f, 0x58, 0xaa, 0x02, 0x10, 0x4e,
	0x6f, 0x62, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x10, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x5c, 0x4f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1c, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x5c, 0x4f, 0x72, 0x62, 0x69, 0x74,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x12, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x3a, 0x3a, 0x4f, 0x72, 0x62, 0x69, 0x74,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_noble_orbiter_v1_action_proto_rawDescOnce sync.Once
	file_noble_orbiter_v1_action_proto_rawDescData = file_noble_orbiter_v1_action_proto_rawDesc
)

func file_noble_orbiter_v1_action_proto_rawDescGZIP() []byte {
	file_noble_orbiter_v1_action_proto_rawDescOnce.Do(func() {
		file_noble_orbiter_v1_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_noble_orbiter_v1_action_proto_rawDescData)
	})
	return file_noble_orbiter_v1_action_proto_rawDescData
}

var file_noble_orbiter_v1_action_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_noble_orbiter_v1_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_noble_orbiter_v1_action_proto_goTypes = []interface{}{
	(ActionID)(0),     // 0: noble.orbiter.v1.ActionID
	(*Action)(nil),    // 1: noble.orbiter.v1.Action
	(*anypb.Any)(nil), // 2: google.protobuf.Any
}
var file_noble_orbiter_v1_action_proto_depIdxs = []int32{
	0, // 0: noble.orbiter.v1.Action.id:type_name -> noble.orbiter.v1.ActionID
	2, // 1: noble.orbiter.v1.Action.attributes:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_noble_orbiter_v1_action_proto_init() }
func file_noble_orbiter_v1_action_proto_init() {
	if File_noble_orbiter_v1_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_noble_orbiter_v1_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
			RawDescriptor: file_noble_orbiter_v1_action_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_noble_orbiter_v1_action_proto_goTypes,
		DependencyIndexes: file_noble_orbiter_v1_action_proto_depIdxs,
		EnumInfos:         file_noble_orbiter_v1_action_proto_enumTypes,
		MessageInfos:      file_noble_orbiter_v1_action_proto_msgTypes,
	}.Build()
	File_noble_orbiter_v1_action_proto = out.File
	file_noble_orbiter_v1_action_proto_rawDesc = nil
	file_noble_orbiter_v1_action_proto_goTypes = nil
	file_noble_orbiter_v1_action_proto_depIdxs = nil
}
