package fib

func Fib(n int) []int {
	rs := []int{}
	for i, j := 0, 1; len(rs) < n; i, j = i+j, i {
		rs = append(rs, i)
	}
	return rs
}

func FibMemoized(n int) int {
	cache := map[int]int{}

	var fib func(n int) int

	impl := func(n int) int {
		if rs, ok := cache[n]; ok {
			return rs
		}
		if n == 1 || n == 2 {
			return 1
		}
		return fib(n-1) + fib(n-2)
	}
	fib = impl

	return fib(n)
}
