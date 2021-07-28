package main

import "fmt"

var mapCache = make(map[int]int)

// 缓存法
func climbStairs(n int) int {

	if n <= 2 {
		return n
	}

	if num, ok := mapCache[n]; ok {
		return num
	} else {
		res := climbStairs(n-1) + climbStairs(n-2)
		mapCache[n] = res
		return res
	}
}

// 滚动变量
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}

	f1 := 1
	f2 := 2
	for i := 3; i <= n; i++ {
		f3 := f1 + f2
		f1 = f2
		f2 = f3
	}

	return f2
}

func main() {
	fmt.Println(climbStairs(100))
	fmt.Println(climbStairs2(100))
}
