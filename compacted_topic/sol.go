package compacted_topic

import (
	"sort"
)

func Compact(nums []int) []int {
	s := &SortedUniqueNums{
		indexes: map[int]int{},
	}

	for i, n := range nums {
		s.Add(n, i)
	}
	sort.Sort(s)
	return s.Set()
}

type SortedUniqueNums struct {
	indexes map[int]int
	set     []int
}

func (s *SortedUniqueNums) Add(n, i int) {
	_, ok := s.indexes[n]
	if !ok {
		s.set = append(s.set, n)
	}
	s.indexes[n] = i
}

func (s SortedUniqueNums) Set() []int {
	return s.set
}

func (s SortedUniqueNums) Len() int {
	return len(s.set)
}

func (s SortedUniqueNums) Less(i, j int) bool {
	return s.indexes[s.set[i]] < s.indexes[s.set[j]]
}

func (s *SortedUniqueNums) Swap(i, j int) {
	s.set[i], s.set[j] = s.set[j], s.set[i]
}
