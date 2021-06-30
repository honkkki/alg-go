package main

import "fmt"

func reverse7(x int) int {
	var res int
	for x != 0 {
		num := x % 10
		res = num + res*10
		x = x / 10
	}
	if res > 1<<31-1 || res < -(1<<31) {
		return 0
	}
	return res
}


// 123 321
func main()  {
	//fmt.Println(123 % 10)
	//fmt.Println(123/10%10)
	//fmt.Println(123/100%10)
	//fmt.Println(123/1000)

	fmt.Println(reverse7(123))
	fmt.Println(reverse7(160))
}
