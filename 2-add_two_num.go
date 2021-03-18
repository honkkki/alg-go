package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	// 头节点
	res := head
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		head.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}

		tmp = tmp / 10
		head = head.Next
	}

	return res.Next
}

func main() {
	node1 := &ListNode{
		Val:  6,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  6,
		Next: nil,
	}
	node3 := &ListNode{
		Val:  6,
		Next: nil,
	}
	node1.Next = node2
	node2.Next = node3

	newNode1 := node1

	res := addTwoNumbers(node1, newNode1)

	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}
