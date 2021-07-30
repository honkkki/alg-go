package main

import "fmt"

var m = make(map[int]int)

// 1 1 2 3 5 8
func fib(n int) int {
	if n <= 2 {
		return 1
	}

	if num, ok := m[n]; ok {
		return num
	} else {
		res := fib(n-1) + fib(n-2)
		m[n] = res
		return res
	}
}

func fib2(n int) int {
	if n <= 2 {
		return 1
	}

	f1 := 1
	f2 := 1
	for i := 3; i <= n; i++ {
		f3 := f1+f2
		f1 = f2
		f2 = f3
	}

	return f2
}

func main() {
	res := fib(6)
	fmt.Println(res)
	fmt.Println(fib2(6))
}
