package misc

func Fibonacci(n int) int {

	if n == 1 || n == 2 {
		return n - 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciOne(n int) int {
	memo := map[int]int{}
	return fibonacciMemoization(n, memo)
}

func fibonacciMemoization(n int, memo map[int]int) int {
	if n == 1 || n == 2 {
		return n - 1
	}
	if v, ok := memo[n]; ok {
		return v
	}
	memo[n] = fibonacciMemoization(n-1, memo) + fibonacciMemoization(n-2, memo)
	return memo[n]
}
func FibonacciTwo(n int) int {
	if n == 1 || n == 2 {
		return n - 1
	}
	first, second := 0, 1
	var third int
	for i := 3; i <= n; i++ {
		third = first + second
		first = second
		second = third
	}
	return third
}
