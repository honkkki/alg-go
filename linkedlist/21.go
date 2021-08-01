package main

import "sort"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Arr := make([]int, 0, 10)
	l2Arr := make([]int, 0, 10)
	if l1 == nil && l2 == nil {
		return nil
	}
	for l1 != nil {
		l1Arr = append(l1Arr, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		l2Arr = append(l2Arr, l2.Val)
		l2 = l2.Next
	}

	arr := append(l1Arr, l2Arr...)
	sort.Ints(arr)
	list := new(ListNode)
	resList := list
	for k, val := range arr {
		list.Val = val
		if k == len(arr)-1 {
			list.Next = nil
		} else {
			list.Next = new(ListNode)
			list = list.Next
		}
	}

	return resList
}

// 递归解决两条有序链表组成一条有序链表
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l2.Next, l1)
		return l2
	}
}

