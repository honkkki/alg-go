package main

func reverseList(head *ListNode) *ListNode {
	var arr = make([]int, 0, 10)
	if head != nil {
		for {
			val := head.Val
			arr = append(arr, val)
			head = head.Next
			if head == nil {
				break
			}
		}
	} else {
		return nil
	}

	p := new(ListNode)
	h := p
	for i := len(arr) - 1; i >= 0; i-- {
		p.Val = arr[i]
		if i == 0 {
			break
		}
		p.Next = new(ListNode)
		p = p.Next
	}

	return h
}

// 双指针 链表反转
func reverseList2(head *ListNode) *ListNode {
	cur := head
	pre := new(ListNode)
	pre = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

func main() {

}
