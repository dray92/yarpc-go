// Code generated by thriftrw v1.9.0. DO NOT EDIT.
// @generated

package meta

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Meta_VersionInfo_Args represents the arguments for the Meta.versionInfo function.
//
// The arguments for versionInfo are sent and received over the wire as this struct.
type Meta_VersionInfo_Args struct {
}

// ToWire translates a Meta_VersionInfo_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Meta_VersionInfo_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Meta_VersionInfo_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Meta_VersionInfo_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Meta_VersionInfo_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Meta_VersionInfo_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a Meta_VersionInfo_Args
// struct.
func (v *Meta_VersionInfo_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("Meta_VersionInfo_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Meta_VersionInfo_Args match the
// provided Meta_VersionInfo_Args.
//
// This function performs a deep comparison.
func (v *Meta_VersionInfo_Args) Equals(rhs *Meta_VersionInfo_Args) bool {

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "versionInfo" for this struct.
func (v *Meta_VersionInfo_Args) MethodName() string {
	return "versionInfo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Meta_VersionInfo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Meta_VersionInfo_Helper provides functions that aid in handling the
// parameters and return values of the Meta.versionInfo
// function.
var Meta_VersionInfo_Helper = struct {
	// Args accepts the parameters of versionInfo in-order and returns
	// the arguments struct for the function.
	Args func() *Meta_VersionInfo_Args

	// IsException returns true if the given error can be thrown
	// by versionInfo.
	//
	// An error can be thrown by versionInfo only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for versionInfo
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// versionInfo into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by versionInfo
	//
	//   value, err := versionInfo(args)
	//   result, err := Meta_VersionInfo_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from versionInfo: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*VersionInfo, error) (*Meta_VersionInfo_Result, error)

	// UnwrapResponse takes the result struct for versionInfo
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if versionInfo threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Meta_VersionInfo_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Meta_VersionInfo_Result) (*VersionInfo, error)
}{}

func init() {
	Meta_VersionInfo_Helper.Args = func() *Meta_VersionInfo_Args {
		return &Meta_VersionInfo_Args{}
	}

	Meta_VersionInfo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Meta_VersionInfo_Helper.WrapResponse = func(success *VersionInfo, err error) (*Meta_VersionInfo_Result, error) {
		if err == nil {
			return &Meta_VersionInfo_Result{Success: success}, nil
		}

		return nil, err
	}
	Meta_VersionInfo_Helper.UnwrapResponse = func(result *Meta_VersionInfo_Result) (success *VersionInfo, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Meta_VersionInfo_Result represents the result of a Meta.versionInfo function call.
//
// The result of a versionInfo execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Meta_VersionInfo_Result struct {
	// Value returned by versionInfo after a successful execution.
	Success *VersionInfo `json:"success,omitempty"`
}

// ToWire translates a Meta_VersionInfo_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Meta_VersionInfo_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Meta_VersionInfo_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _VersionInfo_Read(w wire.Value) (*VersionInfo, error) {
	var v VersionInfo
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Meta_VersionInfo_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Meta_VersionInfo_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Meta_VersionInfo_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Meta_VersionInfo_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _VersionInfo_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Meta_VersionInfo_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Meta_VersionInfo_Result
// struct.
func (v *Meta_VersionInfo_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Meta_VersionInfo_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Meta_VersionInfo_Result match the
// provided Meta_VersionInfo_Result.
//
// This function performs a deep comparison.
func (v *Meta_VersionInfo_Result) Equals(rhs *Meta_VersionInfo_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "versionInfo" for this struct.
func (v *Meta_VersionInfo_Result) MethodName() string {
	return "versionInfo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Meta_VersionInfo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}