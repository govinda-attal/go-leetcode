package link_list

type comparable interface {
	~int | ~float32 | ~int64 | ~float64 | ~int16 | ~int8
}

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

func (ll *ListNode[T]) ToSlice() []T {
	items := []T{}
	cur := ll
	for cur != nil {
		items = append(items, cur.Val)
		cur = cur.Next
	}
	return items
}

func NewFromSlice[T comparable](items ...T) *ListNode[T] {
	if len(items) == 0 {
		return nil
	}
	head := &ListNode[T]{
		Val: items[0],
	}
	cur := head
	i := 1
	for ; i < len(items); i++ {
		cur.Next = &ListNode[T]{
			Val: items[i],
		}
		cur = cur.Next
	}
	return head
}
