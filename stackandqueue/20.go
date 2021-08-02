package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	m := make(map[int32]int32)
	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'

	l := list.New()
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			l.PushBack(v)
		}
		if v == ')' || v == ']' || v == '}' {
			if l.Len() == 0 {
				return false
			}
			last := l.Back().Value
			lastV := last.(int32)
			if m[lastV] == v {
				l.Remove(l.Back())
			} else {
				return false
			}
		}
	}

	return l.Len() == 0
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
