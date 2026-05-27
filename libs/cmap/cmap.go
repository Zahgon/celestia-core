package cmap

import (
	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

// CMap is a goroutine-safe map
type CMap struct {
	m map[string]interface{}
	l cmtsync.Mutex
}

func NewCMap() *CMap { _ = "STUB: not implemented"; return nil }

func (cm *CMap) Set(key string, value interface{}) { _ = "STUB: not implemented"; return }

func (cm *CMap) Get(key string) interface{} { _ = "STUB: not implemented"; return nil }

func (cm *CMap) Has(key string) bool { _ = "STUB: not implemented"; return false }

func (cm *CMap) Delete(key string) { _ = "STUB: not implemented"; return }

func (cm *CMap) Size() int { _ = "STUB: not implemented"; return 0 }

func (cm *CMap) Clear() { _ = "STUB: not implemented"; return }

func (cm *CMap) Keys() []string { _ = "STUB: not implemented"; return nil }

func (cm *CMap) Values() []interface{} { _ = "STUB: not implemented"; return nil }
