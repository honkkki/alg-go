package main

import "fmt"

func maxArea(height []int) int {
	var max int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			h := minNum(height[i], height[j])
			l := j - i
			area := l * h
			max = maxNum(max, area)
		}
	}

	return max
}

// O(n) 左右边界向中间移动 移动短的柱子
func maxArea2(height []int) int {
	var max int
	i := 0
	j := len(height) - 1
	for i < j {
		h := minNum(height[i], height[j])
		l := j - i
		area := l * h
		max = maxNum(max, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}


func minNum(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	max := maxArea2(height)
	fmt.Println(max)
}
