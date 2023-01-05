package LinkList

// func swapPairs(head *ListNode) *ListNode {
// 	dummy := &ListNode{
// 		Val:  -1,
// 		Next: head,
// 	}

// 	pre := dummy

// 	for head != nil && head.Next != nil {
// 		pre.Next = head.Next
// 		next := head.Next.Next
// 		head.Next.Next = head
// 		head.Next = next

// 		pre = head

// 		head = next

// 	}
// 	return dummy.Next
// }

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head

	return next
}
