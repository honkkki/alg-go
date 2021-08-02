package main

import (
	"strings"
)

func simplifyPath(path string) string {
	pathArr := strings.Split(path, "/")
	resArr := make([]string, 0, 10)
	for _, v := range pathArr {
		if v == "." || v == "" {
			continue
		} else if v == ".." {
			if len(resArr) > 0 {
				resArr = resArr[:len(resArr)-1]
			}
		} else {
			resArr = append(resArr, v)
		}
	}

	res := "/" + strings.Join(resArr, "/")
	return res
}
