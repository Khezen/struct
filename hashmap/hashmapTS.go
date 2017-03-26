package hashmap

import (
	"sync"
)

type hashmapTS struct {
	hashmap
	l sync.RWMutex
}

// NewTS creates a new thread safe hashmap
func NewTS(pairs ...interface{}) Interface {
	return &hashmapTS{
		*New(pairs...).(*hashmap),
		sync.RWMutex{},
	}
}

func (h *hashmapTS) Get(k interface{}) (interface{}, error) {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Get(k)
}

func (h *hashmapTS) Put(k, v interface{}) {
	h.l.Lock()
	defer h.l.Unlock()
	h.hashmap.Put(k, v)
}

func (h *hashmapTS) Remove(keys ...interface{}) {
	h.l.Lock()
	defer h.l.Unlock()
	h.hashmap.Remove(keys...)
}

func (h *hashmapTS) Has(keys ...interface{}) bool {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Has(keys...)
}

func (h *hashmapTS) HasValue(values ...interface{}) bool {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.HasValue(values...)
}

func (h *hashmapTS) KeyOf(value interface{}) (interface{}, error) {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.KeyOf(value)
}

func (h *hashmapTS) Each(f func(k, v interface{}) bool) {
	h.l.RLock()
	defer h.l.RUnlock()
	h.hashmap.Each(f)
}

func (h *hashmapTS) Len() int {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Len()
}

func (h *hashmapTS) Clear() {
	h.l.Lock()
	defer h.l.Unlock()
	h.hashmap.Clear()
}

func (h *hashmapTS) IsEmpty() bool {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.IsEmpty()
}

func (h *hashmapTS) IsEqual(t Interface) bool {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.IsEqual(t)
}

func (h *hashmapTS) String() string {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.String()
}

func (h *hashmapTS) Keys() []interface{} {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Keys()
}

func (h *hashmapTS) Values() []interface{} {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Values()
}

func (h *hashmapTS) Copy() Interface {
	h.l.RLock()
	defer h.l.RUnlock()
	cpy := NewTS()
	h.Each(func(k, v interface{}) bool {
		cpy.Put(k, v)
		return true
	})
	return cpy
}
