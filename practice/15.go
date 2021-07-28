package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum > 0 {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else {
				var item []int
				item = append(item, nums[i], nums[l], nums[r])
				res = append(res, item)
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}

	return res
}

func main() {
	s := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(s))
	s1 := []int{0, 0, 0}
	fmt.Println(threeSum(s1))
}
