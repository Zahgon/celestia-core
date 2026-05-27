package json

import (
	"reflect"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

var (
	// typeRegistry contains globally registered types for JSON encoding/decoding.
	typeRegistry = newTypes()
)

// RegisterType registers a type for Amino-compatible interface encoding in the global type
// registry. These types will be encoded with a type wrapper `{"type":"<type>","value":<value>}`
// regardless of which interface they are wrapped in (if any). If the type is a pointer, it will
// still be valid both for value and pointer types, but decoding into an interface will generate
// the a value or pointer based on the registered type.
//
// Should only be called in init() functions, as it panics on error.
func RegisterType(_type interface{}, name string) { _ = "STUB: not implemented"; return }

// typeInfo contains type information.
type typeInfo struct {
	name      string
	rt        reflect.Type
	returnPtr bool
}

// types is a type registry. It is safe for concurrent use.
type types struct {
	cmtsync.RWMutex
	byType map[reflect.Type]*typeInfo
	byName map[string]*typeInfo
}

// newTypes creates a new type registry.
func newTypes() types { _ = "STUB: not implemented"; return *new(types) }

// registers the given type with the given name. The name and type must not be registered already.
func (t *types) register(name string, rt reflect.Type) error { _ = "STUB: not implemented"; return nil }

// If this is a pointer type, we recursively resolve until we get a bare type, but register that
// we should return pointers.

// lookup looks up a type from a name, or nil if not registered.
func (t *types) lookup(name string) (reflect.Type, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), false
}

// name looks up the name of a type, or empty if not registered. Unwraps pointers as necessary.
func (t *types) name(rt reflect.Type) string { _ = "STUB: not implemented"; return "" }
