package main

import "fmt"

// 两数之和
func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int)

	for k, v := range nums {
		value, exist := m[target - v]
		if exist {
			res = append(res, value)
			res = append(res, k)
		}
		m[v] = k
	}

	return res
}


func main()  {
	nums := []int{1,2,3,4,5}
	res := twoSum(nums, 9)
	fmt.Println(res)
}
