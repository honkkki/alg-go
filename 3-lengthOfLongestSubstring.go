package main

import "fmt"

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var freq [256]int
	res, left, right := 0, 0, -1

	for right < len(s) && left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	s := "abcdacd"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
