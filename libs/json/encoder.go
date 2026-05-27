package json

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"time"
)

var (
	timeType            = reflect.TypeOf(time.Time{})
	jsonMarshalerType   = reflect.TypeOf(new(json.Marshaler)).Elem()
	jsonUnmarshalerType = reflect.TypeOf(new(json.Unmarshaler)).Elem()
)

// Marshal marshals the value as JSON, using Amino-compatible JSON encoding (strings for
// 64-bit numbers, and type wrappers for registered types).
func Marshal(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalIndent marshals the value as JSON, using the given prefix and indentation.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encode(w *bytes.Buffer, v any) error {
	_ = "STUB: not implemented"
	// Bare nil values can't be reflected, so we must handle them here.
	return nil
}

// If this is a registered type, defer to interface encoder regardless of whether the input is
// an interface or a bare value. This retains Amino's behavior, but is inconsistent with
// behavior in structs where an interface field will get the type wrapper while a bare value
// field will not.

func encodeReflect(w *bytes.Buffer, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Recursively dereference if pointer.

// Convert times to UTC.

// If the value implements json.Marshaler, defer to stdlib directly. Since we've already
// dereferenced, we try implementations with both value receiver and pointer receiver. We must
// do this after the time normalization above, and thus after dereferencing.

// Complex types must be recursively encoded.

// 64-bit integers are emitted as strings, to avoid precision problems with e.g.
// Javascript which uses 64-bit floats (having 53-bit precision).

// For everything else, defer to the stdlib encoding/json encoder

func encodeReflectList(w *bytes.Buffer, rv reflect.Value) error {
	_ = "STUB: not implemented"
	// Emit nil slices as null.
	return nil
}

// Encode byte slices as base64 with the stdlib encoder.

// Stdlib does not base64-encode byte arrays, only slices, so we copy to slice.

// Anything else we recursively encode ourselves.

func encodeReflectMap(w *bytes.Buffer, rv reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// nil maps are not emitted as nil, to retain Amino compatibility.

func encodeReflectStruct(w *bytes.Buffer, rv reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeReflectInterface(w *bytes.Buffer, rv reflect.Value) error {
	_ = "STUB: not implemented"
	// Get concrete value and dereference pointers.
	return nil
}

// Look up the name of the concrete type

// Write value wrapped in interface envelope

func encodeStdlib(w *bytes.Buffer, v any) error {
	_ = "STUB: not implemented"
	// Stream the output of the JSON marshaling directly into the buffer.
	// The stdlib encoder will write a newline, so we must truncate it,
	// which is why we pass in a bytes.Buffer throughout, not io.Writer.
	return nil
}

// Remove the last byte from the buffer

func writeStr(w io.Writer, s string) error { _ = "STUB: not implemented"; return nil }
