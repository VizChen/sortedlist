package sortedlist

import "sort"

const (
	defaultLoadFactor = 1000
)

type SortedList struct {
	lists [][]int
	maxes []int
	len   int
}

func New() *SortedList {
	return &SortedList{
		lists: [][]int{},
		maxes: []int{},
		len: 0,
	}
}

func (s *SortedList) Add(v int) {
	s.len += 1
	if len(s.maxes) == 0 {
		s.maxes = append(s.maxes, v)
		s.lists = append(s.lists, []int{v})
		return
	}
    pos := sort.SearchInts(s.maxes, v)
	if pos == len(s.maxes) {
		pos -= 1
		s.lists[pos] = append(s.lists[pos], v)
		s.maxes[pos] = v
	} else {
		idx := sort.SearchInts(s.lists[pos], v)
		if idx != len(s.lists[pos]) {
			s.lists[pos] = append(s.lists[pos][:idx+1], s.lists[pos][idx:]...)
            s.lists[pos][idx] = v
		} else {
			s.lists[pos] = append(s.lists[pos], v)
		}
	}
	s.expand(pos)
}

func (s *SortedList) expand(pos int) {
	if len(s.lists[pos]) > (defaultLoadFactor << 1) {
        half := append([]int{}, s.lists[pos][defaultLoadFactor:]...)
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

func (s SortedList) List() []int {
	if len(s.lists) == 0 {
		return []int{}
	} else if len(s.lists) == 1 {
		return s.lists[0]
	}

	ret := []int{}
	for _, l := range s.lists {
		ret = append(ret, l...)
	}
	return ret
}
