package LinkList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}

	fast, slow := dummy, dummy

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
