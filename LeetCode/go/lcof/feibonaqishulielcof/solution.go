package feibonaqishulielcof

func fib(n int) int {
	if n < 2 {
		return n
	}
	p2, p1 := 0, 1
	for i := 2; i <= n; i++ {
		p2, p1 = p1, (p1+p2)%(1e9+7)
	}
	return p1
}
