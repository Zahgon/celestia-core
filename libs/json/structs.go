package json

import (
	"reflect"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

var (
	// cache caches struct info.
	cache = newStructInfoCache()
)

// structCache is a cache of struct info.
type structInfoCache struct {
	cmtsync.RWMutex
	structInfos map[reflect.Type]*structInfo
}

func newStructInfoCache() *structInfoCache { _ = "STUB: not implemented"; return nil }

func (c *structInfoCache) get(rt reflect.Type) *structInfo { _ = "STUB: not implemented"; return nil }

func (c *structInfoCache) set(rt reflect.Type, sInfo *structInfo) {
	_ = "STUB: not implemented"
	return
}

// structInfo contains JSON info for a struct.
type structInfo struct {
	fields []*fieldInfo
}

// fieldInfo contains JSON info for a struct field.
type fieldInfo struct {
	jsonName  string
	omitEmpty bool
	hidden    bool
}

// makeStructInfo generates structInfo for a struct as a reflect.Value.
func makeStructInfo(rt reflect.Type) *structInfo { _ = "STUB: not implemented"; return nil }
