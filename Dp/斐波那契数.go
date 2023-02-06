package Dp

// func fib(n int) int {
// 	var fibRunner func(a, b, n int) int

// 	fibRunner = func(a, b, n int) int {
// 		if n == 0 {
// 			return a
// 		}

// 		return fibRunner(b, a+b, n-1)
// 	}

// 	return fibRunner(0, 1, n)

// }

func fib(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 1, 0

	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}

	return c
}
