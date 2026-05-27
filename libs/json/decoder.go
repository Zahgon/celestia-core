package json

import (
	"encoding/json"
	"reflect"
)

// Unmarshal unmarshals JSON into the given value, using Amino-compatible JSON encoding (strings
// for 64-bit numbers, and type wrappers for registered types).
func Unmarshal(bz []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

func decode(bz []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

// If this is a registered type, defer to interface decoder regardless of whether the input is
// an interface or a bare value. This retains Amino's behavior, but is inconsistent with
// behavior in structs where an interface field will get the type wrapper while a bare value
// field will not.

func decodeReflect(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Handle null for slices, interfaces, and pointers

// Dereference-and-construct pointers, to handle nested pointers.

// Times must be UTC and end with Z

// If value implements json.Umarshaler, call it.

// Decode complex types recursively.

// For 64-bit integers, unwrap expected string and defer to stdlib for integer decoding.

// Anything else we defer to the stdlib.

func decodeReflectList(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Decode base64-encoded bytes using stdlib decoder, via byte slice for arrays.

// Decode anything else into a raw JSON slice, and decode values recursively.

// arrays of wrong size

// Replace empty slices with nil slices, for Amino compatibility

func decodeReflectMap(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Decode into a raw JSON map, using string keys.

// Recursively decode values.

func decodeReflectStruct(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Decode raw JSON values into a string-keyed map.

func decodeReflectInterface(bz []byte, rv reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// Decode the interface wrapper.

// Dereference-and-construct pointers, to handle nested pointers.

// Look up the interface type, and construct a concrete value.

// This makes sure interface implementations with pointer receivers (e.g. func (c *Car)) are
// constructed as pointers behind the interface. The types must be registered as pointers with
// RegisterType().

func decodeStdlib(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Make sure we are unmarshaling into a pointer.

type interfaceWrapper struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}
