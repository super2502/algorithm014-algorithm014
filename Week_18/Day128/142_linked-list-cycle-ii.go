package Day128

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	fast, slow := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for {
				if slow == fast {
					return slow
				} else {
					slow = slow.Next
					fast = fast.Next
				}
			}
		}
	}

	return nil
}
