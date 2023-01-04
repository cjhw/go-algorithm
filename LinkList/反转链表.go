package LinkList

// func reverseList(head *ListNode) *ListNode {

// 	var pre *ListNode
// 	cur := head

// 	for cur != nil {
// 		next := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = next
// 	}

// 	return pre

// }

func reverseList(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}

	next := head.Next
	head.Next = pre

	return help(head, next)
}
