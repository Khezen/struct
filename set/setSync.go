package set

import (
	"sync"

	"github.com/khezen/struct/collection"
)

// setSync defines a thread safe set data structure.
type setSync struct {
	set
	l sync.RWMutex // we name it because we don't want to expose it
}

// NewSync creates and initialize a new thread safe set. It's accept a variable number of
// arguments to populate the initial set. If nothing passed a set with zero
// size is created.
func NewSync(items ...interface{}) Interface {
	return &setSync{
		*New(items...).(*set),
		sync.RWMutex{},
	}
}

// Add includes the specified items (one or more) to the set. The underlying
// set s is modified. If passed nothing it silently returns.
func (s *setSync) Add(items ...interface{}) {
	if len(items) > 0 {
		s.l.Lock()
		defer s.l.Unlock()
		s.set.Add(items...)
	}
}

// Remove deletes the specified items from the set.  The underlying set s is
// modified. If passed nothing it silently returns.
func (s *setSync) Remove(items ...interface{}) {
	if len(items) > 0 {
		s.l.Lock()
		defer s.l.Unlock()
		s.set.Remove(items...)
	}
}

// Pop  deletes and return an item from the set. The underlying set s is
// modified. If set is empty, nil is returned.
func (s *setSync) Pop() interface{} {
	s.l.Lock()
	defer s.l.Unlock()
	return s.set.Pop()
}

// Has looks for the existence of items passed. It returns false if nothing is
// passed. For multiple items it returns true only if all of  the items exist.
func (s *setSync) Has(items ...interface{}) bool {
	if len(items) > 0 {
		s.l.RLock()
		defer s.l.RUnlock()
		return s.set.Has(items...)
	}
	return true
}

func (s *setSync) Replace(item, substitute interface{}) {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Replace(item, substitute)
}

// Len returns the number of items in a set.
func (s *setSync) Len() int {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.set.Len()
}

// Clear removes all items from the set.
func (s *setSync) Clear() {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Clear()
}

// IsEqual test whether s and t are the same in size and have the same items.
func (s *setSync) IsEqual(t collection.Interface) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.set.IsEqual(t)
}

// IsSubset tests whether t is a subset of s.
func (s *setSync) IsSubset(t Interface) (subset bool) {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.set.IsSubset(t)
}

// Each traverses the items in the set, calling the provided function for each
// set member. Traversal will continue until all items in the set have been
// visited, or if the closure returns false.
func (s *setSync) Each(f func(item interface{}) bool) {
	s.l.RLock()
	defer s.l.RUnlock()
	s.set.Each(f)
}

// Slice returns a slice of all items. There is also StringSlice() and
// IntSlice() methods for returning slices of type string or int.
func (s *setSync) Slice() []interface{} {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.set.Slice()
}

// Copy returns a new set with a copy of s.
func (s *setSync) CopySet() Interface {
	s.l.RLock()
	defer s.l.RUnlock()
	u := NewSync()
	for item := range s.m {
		u.Add(item)
	}
	return u
}

// Merge is like Union, however it modifies the current set it's applied on
// with the given t set.
func (s *setSync) Merge(t collection.Interface) {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Merge(t)
}

// Retain removes the set items not containing in t from set s.
func (s *setSync) Retain(t collection.Interface) {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Retain(t)
}

func (s *setSync) CopyCollection() collection.Interface {
	return s.CopySet()
}
