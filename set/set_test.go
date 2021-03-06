package set

import (
	"testing"

	"github.com/khezen/struct/collection"
)

func testErr(err error, expectErr bool, t *testing.T) {
	if (expectErr && err == nil) || (!expectErr && err != nil) {
		t.Errorf(" Error expected? %v. Got: %v.", expectErr, err)
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		set, toBeAdded, expected Interface
	}{
		{New(1, 4, -8), New(42, -1), New(1, 4, -8, 42, -1)},
		{New(), New(42, -1), New(42, -1)},
		{New(), New(), New()},
		{New(), New(nil), New(nil)},
		{New(), New(42, -1), New(42, -1)},
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, -8, 42, -1)},
		{NewSync(), New(42, -1), NewSync(42, -1)},
		{NewSync(), New(), NewSync()},
		{NewSync(), New(nil), NewSync(nil)},
		{NewSync(), New(42, -1), NewSync(42, -1)},
	}
	for _, c := range cases {
		c.set.Add(c.toBeAdded.Slice()...)
		if !c.set.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.set.Slice())
		}
	}
}

func TestPop(t *testing.T) {
	cases := []struct {
		s Interface
	}{
		{New(1, 2, "fatih")},
		{NewSync(1, 2, "fatih")},
	}
	for _, c := range cases {
		length := c.s.Len()
		for i := 0; i < length; i++ {
			item := c.s.Pop()
			if c.s.Len() != length-1-i {
				t.Errorf("Expected %v. Got %v", length-i, c.s.Len())
			}
			if c.s.Has(item) {
				t.Error("Pop: returned item should not exist")
			}
		}
		item := c.s.Pop()
		if item != nil {
			t.Error("Pop: should return nil because set is empty")
		}
	}
}

func TestIsSubset(t *testing.T) {
	cases := []struct {
		s, sub Interface
		isSub  bool
	}{
		{New("1", "2", "3", "4"), New("1", "2", "3"), true},
		{New("1", "2", "3"), New("1", "2", "3", "4"), false},
		{NewSync("1", "2", "3", "4"), NewSync("1", "2", "3"), true},
		{NewSync("1", "2", "3"), NewSync("1", "2", "3", "4"), false},
	}
	for _, c := range cases {
		ok := c.s.IsSubset(c.sub)
		if ok != c.isSub {
			t.Errorf("Expected %v. Got %v", c.isSub, ok)
		}
	}
}

func TestIsSuperset(t *testing.T) {
	cases := []struct {
		s, sub Interface
		isSub  bool
	}{
		{New("1", "2", "3", "4"), New("1", "2", "3"), false},
		{New("1", "2", "3"), New("1", "2", "3", "4"), true},
		{NewSync("1", "2", "3", "4"), NewSync("1", "2", "3"), false},
		{NewSync("1", "2", "3"), NewSync("1", "2", "3", "4"), true},
	}
	for _, c := range cases {
		ok := c.s.IsSuperset(c.sub)
		if ok != c.isSub {
			t.Errorf("Expected %v. Got %v", c.isSub, ok)
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		set, toBeRemoved, expected Interface
	}{
		{New(1, 4, -8), New(42, -1), New(1, 4, -8)},
		{New(1, 4, -8), New(1, -8), New(4)},
		{New(), New(42, -1), New()},
		{New(), New(), New()},
		{New(), New(nil), New()},
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, -8)},
		{NewSync(1, 4, -8), New(1, -8), NewSync(4)},
		{NewSync(), New(42, -1), NewSync()},
		{NewSync(), New(), NewSync()},
		{NewSync(), New(nil), NewSync()},
	}
	for _, c := range cases {
		c.set.Remove(c.toBeRemoved.Slice()...)
		if !c.set.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.set.Slice())
		}
	}
}

func TestReplace(t *testing.T) {
	cases := []struct {
		set, expected    Interface
		item, substitute interface{}
	}{
		{New(1, 4, -8), New(42, 4, -8), 1, 42},
		{New(1, 4, -8), New(1, 4, -8), 1000, 42},
		{NewSync(1, 4, -8), NewSync(42, 4, -8), 1, 42},
		{NewSync(1, 4, -8), NewSync(1, 4, -8), 1000, 42},
	}
	for _, c := range cases {
		c.set.Replace(c.item, c.substitute)
		if !c.set.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.set.Slice())
		}
	}
}

func TestHas(t *testing.T) {
	cases := []struct {
		set, items Interface
		expected   bool
	}{
		{New(1, 42, -8), New(1, 42, -8), true},
		{New(1, 42, -8), New(-8), true},
		{New(1, 42, -8), New(34), false},
		{New(1, 42, -8), New(nil), false},
		{New(1, 42, -8), New(), true},
		{NewSync(1, 42, -8), NewSync(1, 42, -8), true},
		{NewSync(1, 42, -8), NewSync(-8), true},
		{NewSync(1, 42, -8), NewSync(34), false},
		{NewSync(1, 42, -8), NewSync(nil), false},
		{NewSync(1, 42, -8), NewSync(), true},
	}
	for _, c := range cases {
		has := c.set.Has(c.items.Slice()...)
		if has != c.expected {
			t.Errorf("Expected %v. Got %v.", c.expected, has)
		}
	}
}

func TestEach(t *testing.T) {
	cases := []struct {
		set  Interface
		stop bool
	}{
		{New(1, -8, 42), true},
		{NewSync(1, -8, 42), false},
	}
	for _, c := range cases {
		count := 0
		callback := func(item interface{}) bool {
			count++
			return !c.stop
		}
		c.set.Each(callback)
		if c.stop && count != 1 {
			t.Errorf("Expected %v. Got %v.", 1, count)
		}
	}
}

func TestLen(t *testing.T) {
	cases := []struct {
		set Interface
		len int
	}{
		{New(), 0},
		{New(1), 1},
		{New(1, 42, -8), 3},
		{NewSync(), 0},
		{NewSync(1), 1},
		{NewSync(1, 42, -8), 3},
	}
	for _, c := range cases {
		if c.set.Len() != c.len {
			t.Errorf("Expected %v. Got %v.", c.len, c.set.Len())
		}
	}
}

func TestClear(t *testing.T) {
	cases := []struct {
		set Interface
	}{
		{New(1, 42, -8)},
		{New()},
		{NewSync(1, 42, -8)},
		{NewSync()},
	}
	for _, c := range cases {
		c.set.Clear()
		if !c.set.IsEmpty() {
			t.Error("Array should be empty")
		}
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		set     Interface
		isEmpty bool
	}{
		{New(), true},
		{New(1, 42, -8), false},
		{NewSync(), true},
		{NewSync(1, 42, -8), false},
	}
	for _, c := range cases {
		if c.set.IsEmpty() != c.isEmpty {
			t.Errorf("Expected %v. Got %v.", c.isEmpty, c.set.IsEmpty())
		}
	}
}

func TestIsEqual(t *testing.T) {
	cases := []struct {
		set, toBeCompared Interface
		isEqual           bool
	}{
		{New(), New(), true},
		{New(1, 42, -8), New(1, 42, -8), true},
		{New(1, 42, -8), New(1, "42", -8), false},
		{New(1, 42, -8), New(), false},
		{New(66, -1000), New(42, 1, 8), false},
		{NewSync(), NewSync(), true},
		{NewSync(1, 42, -8), NewSync(1, 42, -8), true},
		{NewSync(1, 42, -8), NewSync(1, "42", -8), false},
		{NewSync(1, 42, -8), NewSync(), false},
		{NewSync(66, -1000), NewSync(42, 1, 8), false},
	}
	for _, c := range cases {
		isEqual := c.set.IsEqual(c.toBeCompared)
		if isEqual != c.isEqual {
			t.Errorf("Expected %v to be equal to %v? %v. Got: %v", c.set.Slice(), c.toBeCompared.Slice(), c.isEqual, isEqual)
		}
	}
}

func TestMerge(t *testing.T) {
	cases := []struct {
		set, toBeMerged, expected Interface
	}{
		{New(1, 42), New(-8), New(1, 42, -8)},
		{New(1, 42), New(-8, nil), New(1, 42, -8, nil)},
		{New(1, 42), New(), New(1, 42)},
		{New(), New(), New()},
		{NewSync(1, 42), NewSync(-8), NewSync(1, 42, -8)},
		{NewSync(1, 42), NewSync(-8, nil), NewSync(1, 42, -8, nil)},
		{NewSync(1, 42), NewSync(), NewSync(1, 42)},
		{NewSync(), NewSync(), NewSync()},
	}
	for _, c := range cases {
		c.set.Merge(c.toBeMerged)
		if !c.set.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.set.Slice())
		}
	}
}

func TestSeparate(t *testing.T) {
	cases := []struct {
		set, toBeMerged, expected Interface
	}{
		{New(1, 42, -8), New(1, 42), New(-8)},
		{New(1, 42, -8), New(1, 42, nil), New(-8)},
		{New(1, 42, -8), New(), New(1, 42, -8)},
		{New(), New(), New()},
		{NewSync(1, 42, -8), NewSync(1, 42), NewSync(-8)},
		{NewSync(1, 42, -8), NewSync(1, 42, nil), NewSync(-8)},
		{NewSync(1, 42, -8), NewSync(), NewSync(1, 42, -8)},
		{NewSync(), NewSync(), NewSync()},
	}
	for _, c := range cases {
		c.set.Separate(c.toBeMerged)
		if !c.set.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.set.Slice())
		}
	}
}

func TestRetain(t *testing.T) {
	cases := []struct {
		set, toBeMerged, expected Interface
	}{
		{New(1, 42, -8), New(1, -8, 100), New(1, -8)},
		{New(1, 42, -8), New(1, -8, 100, nil), New(1, -8)},
		{New(1, 42, -8), New(), New()},
		{New(), New(), New()},
		{NewSync(1, 42, -8), NewSync(1, -8, 100), NewSync(1, -8)},
		{NewSync(1, 42, -8), NewSync(1, -8, 100, nil), NewSync(1, -8)},
		{NewSync(1, 42, -8), NewSync(), NewSync()},
		{NewSync(), NewSync(), NewSync()},
	}
	for _, c := range cases {
		c.set.Retain(c.toBeMerged)
		if !c.set.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.set.Slice())
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		set      Interface
		expected string
	}{
		{New(1), "[1]"},
		{New(), "[]"},
		{New(nil), "[<nil>]"},
		{NewSync(1), "[1]"},
		{NewSync(), "[]"},
		{NewSync(nil), "[<nil>]"},
	}

	for _, c := range cases {
		str := c.set.String()
		if str != c.expected {
			t.Errorf("Expected %v. Got %v", c.expected, str)
		}
	}

}

func TestSlice(t *testing.T) {
	cases := []struct {
		slice []interface{}
	}{
		{[]interface{}{1, 5, -76}},
	}
	for _, c := range cases {
		arr, arrSync := New(c.slice...), NewSync(c.slice...)
		s := arr.Slice()
		for i := range s {
			has := false
			for j := range c.slice {
				has = s[i] == c.slice[j]
				if has {
					break
				}
			}
			if !has {
				t.Errorf("Expected %v. Got %v.", c.slice, s)
			}
		}
		s = arrSync.Slice()
		for i := range s {
			has := false
			for j := range c.slice {
				has = s[i] == c.slice[j]
				if has {
					break
				}
			}
			if !has {
				t.Errorf("Expected %v. Got %v.", c.slice, s)
			}
		}
	}
}

func TestCopySet(t *testing.T) {
	cases := []struct {
		set Interface
	}{
		{New(1, 42, -8)},
		{New(-66, 1000, 32)},
		{NewSync(1, 42, -8)},
		{NewSync(-66, 1000, 32)},
	}
	for _, c := range cases {
		cpy := c.set.CopySet()
		if !cpy.IsEqual(c.set) {
			t.Errorf("Expected %v. Got %v.", c.set.Slice(), cpy.Slice())
		}
	}
}

func TestCopyCollection(t *testing.T) {
	cases := []struct {
		set Interface
	}{
		{New(1, 42, -8)},
		{New(-66, 1000, 32)},
		{NewSync(1, 42, -8)},
		{NewSync(-66, 1000, 32)},
	}
	for _, c := range cases {
		cpy := c.set.CopyCollection()
		if !cpy.IsEqual(c.set) {
			t.Errorf("Expected %v. Got %v.", c.set.Slice(), cpy.Slice())
		}
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		sets     []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(5, 42, 6), New(1, 42, -8, 7)}, New(1, 42, -8, 5, 6, 7)},
		{[]collection.Interface{New(1, 42, -8), New(5, 42, 6)}, New(1, 42, -8, 5, 6)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(5, 42, 6), NewSync(1, 42, -8, 7)}, NewSync(1, 42, -8, 5, 6, 7)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(5, 42, 6)}, NewSync(1, 42, -8, 5, 6)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Union(c.sets...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		sets     []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(-8, 6, 6), New(1, 7)}, New(42)},
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6)}, New(42)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 6, 6), NewSync(1, 7)}, NewSync(42)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6)}, NewSync(42)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Difference(c.sets...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		sets     []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6), New(1, 7)}, New(1)},
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6)}, New(1, -8)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6), NewSync(1, 7)}, NewSync(1)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6)}, NewSync(1, -8)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Intersection(c.sets...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestExclusion(t *testing.T) {
	cases := []struct {
		sets     []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6), New(1, 7)}, New(42, 6, 7)},
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6)}, New(42, 6)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6), NewSync(1, 7)}, NewSync(42, 6, 7)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6)}, NewSync(42, 6)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Exclusion(c.sets...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}
