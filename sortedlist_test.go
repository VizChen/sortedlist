package sortedlist

import (
	"testing"
)

func TestAddLessThanLoad(t *testing.T) {
	sl := New[int]()
	for i := 100; i >= 0; i-- {
		sl.Add(i)
	}
	if sl.len != 101 {
		t.Fatalf("Length of SortedList is %d, expected 101.", sl.len)
	}
	if len(sl.lists) != 1 || len(sl.lists[0]) != 101 {
		t.Fatalf("Lists of SortedList is %v, expected [[0...100]].", sl.lists)
	}
	if len(sl.maxes) != 1 || sl.maxes[0] != 100 {
		t.Fatalf("Maxes of SortedList is %v, expected [100].", sl.maxes)
	}
}

func TestAddLargerThanLoad(t *testing.T) {
	sl := New[int]()
	for i := 2*defaultLoadFactor; i >= 0; i-- {
		sl.Add(i)
	}
	if sl.len != 2*defaultLoadFactor+1 {
		t.Fatalf("Length of SortedList is %d, expected %d.", sl.len, 2*defaultLoadFactor+1)
	}
	if len(sl.lists) != 2 || len(sl.lists[0]) != defaultLoadFactor || len(sl.lists[1]) != defaultLoadFactor+1 {
		t.Fatalf("Lists of SortedList is %v, expected [[0...%d], [%d...%d+1]].", sl.lists, defaultLoadFactor-1, defaultLoadFactor, 2*defaultLoadFactor)
	}
	if len(sl.maxes) != 2 || sl.maxes[0] != defaultLoadFactor-1 || sl.maxes[1] != 2*defaultLoadFactor{
		t.Fatalf("Maxes of SortedList is %v, expected [%d, %d].", sl.maxes, defaultLoadFactor-1, 2*defaultLoadFactor)
	}
}
