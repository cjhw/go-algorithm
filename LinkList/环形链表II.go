package LinkList

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}

	return nil
}
