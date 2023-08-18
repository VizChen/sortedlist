package sortedlist

import (
	"sort"
)

const (
	defaultLoadFactor = 1000
)

// cmp.Ordered interface is supported in 1.21.0.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type SortedList[K Ordered] struct {
	lists [][]K
	maxes []K
	len int
}

func New[K Ordered]() *SortedList[K] {
	return &SortedList[K]{
		lists: [][]K{},
		maxes: []K{},
		len: 0,
	}
}

func (s *SortedList[K]) Add(v K) {
	s.len += 1
	if len(s.maxes) == 0 {
		s.maxes = append(s.maxes, v)
		s.lists = append(s.lists, []K{v})
		return
	}
    pos := sort.Search(len(s.maxes), func(i int)bool{
			return s.maxes[i] >= v
		})
	if pos == len(s.maxes) {
		pos -= 1
		s.lists[pos] = append(s.lists[pos], v)
		s.maxes[pos] = v
	} else {
		idx := sort.Search(len(s.lists[pos]), func(i int) bool {
				return s.lists[pos][i] >= v
			})
		if idx != len(s.lists[pos]) {
			s.lists[pos] = append(s.lists[pos][:idx+1], s.lists[pos][idx:]...)
            s.lists[pos][idx] = v
		} else {
			s.lists[pos] = append(s.lists[pos], v)
		}
	}
	s.expand(pos)
}

func (s *SortedList[K]) expand(pos int) {
	if len(s.lists[pos]) > (defaultLoadFactor << 1) {
        half := append([]K{}, s.lists[pos][defaultLoadFactor:]...)
		s.lists[pos] = s.lists[pos][:defaultLoadFactor]
		s.maxes[pos] = s.lists[pos][len(s.lists[pos])-1]

		if pos == len(s.lists) - 1 {
			s.lists = append(s.lists, half)
			s.maxes = append(s.maxes, half[len(half)-1])
		} else {
			s.lists = append(s.lists[:pos+1], s.lists[pos:]...)
            s.lists[pos+1] = half
			s.maxes = append(s.maxes[:pos+1], s.maxes[pos:]...)
            s.maxes[pos+1] = half[len(half)-1]
		}
	}
}

func (s SortedList[K]) List() []K {
	if len(s.lists) == 0 {
		return []K{}
	} else if len(s.lists) == 1 {
		return s.lists[0]
	}

	ret := []K{}
	for _, l := range s.lists {
		ret = append(ret, l...)
	}
	return ret
}
