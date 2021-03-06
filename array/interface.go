// Package array provides both threadsafe and non-threadsafe implementations of
// a generic dynamic array. In the threadsafe array, safety encompasses all
// operations on one array. Operations on multiple arrays are consistent in that
// the elements of each array used was valid at exactly one point in time
// between the start and the end of the operation.
package array

import (
	"errors"

	"github.com/khezen/struct/collection"
)

// Interface is describing a Set. Sets are an unordered, unique list of values.
type Interface interface {
	collection.Interface
	Get(i int) interface{}
	Insert(i int, item ...interface{})
	RemoveAt(i int) interface{}
	ReplaceAt(i int, substitute interface{}) interface{}
	IndexOf(interface{}) (int, error)
	Swap(i, j int)
	SubArray(i, j int) Interface
	CopyArr() Interface
}

var (
	// ErrIndexOutOfBounds - index is out of bounds
	ErrIndexOutOfBounds = errors.New("ErrIndexOutOfBounds")
	// ErrBadSubsetBoudaries - subset boudaries must be 0 <= i < j <= length
	ErrBadSubsetBoudaries = errors.New("ErrBadSubsetBoudaries -  subset boudaries must be 0 <= i < j <= length")
	// ErrNotFound - item not found
	ErrNotFound = errors.New("ErrNotFound")
)
