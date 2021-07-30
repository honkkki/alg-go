package main

import "fmt"

func moveZeroes(nums []int) {
	var j int
	for i, num := range nums {
		if num != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}

			j++
		}
	}

	fmt.Println(nums)
}

func moveZeroes2(nums []int) {
	var j int
	for i, num := range nums {
		if num != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	fmt.Println(nums)
}

func main() {
	s := []int{0, 1, 0, 3, 12}
	moveZeroes(s)

	s2 := []int{0, 1, 0, 3, 12}
	moveZeroes2(s2)

}
