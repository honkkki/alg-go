package main

import "fmt"

var m = make(map[int]int)

func fib(n int) int {
	if n < 2 {
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

func main() {
	res := fib(100)
	fmt.Println(res)
}
