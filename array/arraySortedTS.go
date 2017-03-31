package array

import (
	"sort"
)

// Sorted is the interface for sortable arrays
type arraySortTS struct {
	*arrayTS
	less func(slice []interface{}, i, j int) bool
}

// NewSortedTS creates a thread safe array that expose Sort method
func NewSortedTS(less func(slice []interface{}, i, j int) bool, items ...interface{}) Sorted {
	return &arraySortTS{
		(NewTS(items...).(*arrayTS)),
		less,
	}
}

func (a *arraySortTS) Less(i, j int) bool {
	a.arrayTS.l.RLock()
	defer a.arrayTS.l.RUnlock()
	return a.less(a.arrayTS.s, i, j)
}

func (a *arraySortTS) Sort() {
	sort.Sort(a)
}
