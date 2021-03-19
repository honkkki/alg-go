package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	res, left, right := 0, 0, 0
	for right < len(s) && left < len(s) {
		right++

	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main()  {
	s := "abcdacd"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
