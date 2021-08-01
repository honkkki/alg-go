package main

import "fmt"

// 中心扩散法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	res := ""
	for i := range s {
		// 比较奇数个与偶数个
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}

	return res
}

// res用来比较
func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	// 返回长的子串
	if len(res) < len(sub) {
		return sub
	}
	return res
}

func main() {
	res := longestPalindrome("abbaaba")
	fmt.Println(res)
}
