package top_frequent_words

import "container/heap"

type Word struct {
	val   string
	count int
	index int
}

var _ heap.Interface = &wordHeap{}

type wordHeap []Word

func (h *wordHeap) Push(w interface{}) {
	i := len(*h)
	word := w.(Word)
	word.index = i
	*h = append(*h, word)
}

func (h *wordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

func (h wordHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h wordHeap) Less(i int, j int) bool {
	return h[i].count > h[j].count
}

func (h wordHeap) Len() int {
	return len(h)
}

func TopKFrequent(words []string, k int) []string {
	out := make([]string, k)
	wcMap := make(map[string]int)
	h := new(wordHeap)

	for _, v := range words {
		if _, ok := wcMap[v]; ok {
			wcMap[v]++
			continue
		}
		wcMap[v] = 1
	}

	for w, c := range wcMap {
		heap.Push(h, Word{val: w, count: c})
	}

	for i := 0; i < k; i++ {
		out[i] = heap.Pop(h).(Word).val
	}

	return out
}
