package heap

type comparable interface {
	~int | ~float32 | ~int64 | ~float64 | ~int16 | ~int8
}

type Heap[T comparable] struct {
	items []T
	size  int
}

func New[T comparable](capacity int) *Heap[T] {
	return &Heap[T]{
		items: make([]T, capacity),
		size:  0,
	}
}

func (h *Heap[T]) LeftIx(ix int) int {
	return 2*ix + 1
}

func (h *Heap[T]) RightIx(ix int) int {
	return 2*ix + 2
}

func (h *Heap[T]) ParentIx(ix int) int {
	return (ix - 1) / 2
}

func (h *Heap[T]) HasLeft(ix int) bool {
	return h.LeftIx(ix) < h.size
}

func (h *Heap[T]) HasRight(ix int) bool {
	return h.LeftIx(ix) < h.size
}

func (h *Heap[T]) HasParent(ix int) bool {
	return h.LeftIx(ix) >= 0
}

func (h *Heap[T]) Parent(ix int) (val T, ok bool) {
	if h.HasParent(ix) {
		return h.items[h.ParentIx(ix)], true
	}
	return val, false
}

func (h *Heap[T]) Left(ix int) (val T, ok bool) {
	if h.HasLeft(ix) {
		return h.items[h.LeftIx(ix)], true
	}
	return val, false
}

func (h *Heap[T]) Right(ix int) (val T, ok bool) {
	if h.HasRight(ix) {
		return h.items[h.RightIx(ix)], true
	}
	return val, false
}

func (h *Heap[T]) Swap(ix, jx int) bool {
	if ix <= h.size && jx <= h.size && ix >= 0 && jx >= 0 {
		h.items[ix], h.items[jx] = h.items[jx], h.items[ix]
		return true
	}
	return false
}

func (h *Heap[T]) ensureCapacity() {
	if h.size < cap(h.items) {
		return
	}
	items := make([]T, h.size*2)
	_ = copy(items, h.items)
	h.items = items
}

func (h *Heap[T]) Peek() (val T, ok bool) {
	if h.size == 0 {
		return
	}
	return h.items[0], true
}

func (h *Heap[T]) Poll() (val T, ok bool) {
	if h.size == 0 {
		return
	}
	item := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--
	h.heapifyDown()
	return item, true
}

func (h *Heap[T]) Add(val T) {
	h.ensureCapacity()
	h.items[h.size] = val
	h.size++
	h.heapifyUp()
}

func (h *Heap[T]) heapifyDown() {
	ix := 0
	for h.HasLeft(ix) {
		smallerIx := h.LeftIx(ix)
		left, _ := h.Left(ix)
		right, ok := h.Right(ix)
		if ok && left > right {
			smallerIx = h.RightIx(ix)
		}
		if h.items[ix] < h.items[smallerIx] {
			break
		}
		h.Swap(ix, smallerIx)
		ix = smallerIx
	}
}

func (h *Heap[T]) heapifyUp() {
	ix := h.size - 1
	parent, ok := h.Parent(ix)
	for ok && parent > h.items[ix] {
		h.Swap(h.ParentIx(ix), ix)
		ix = h.ParentIx(ix)
		parent, ok = h.Parent(ix)
	}
}

func (h *Heap[T]) Items() []T {
	dup := make([]T, h.size)
	copy(dup, h.items)
	return dup
}
