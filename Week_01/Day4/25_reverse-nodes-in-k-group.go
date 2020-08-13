package Day4

func reverseKGroup(head *ListNode, k int) *ListNode {

	newHead := head
	for i := 0; i < k; i++ {
		if newHead == nil {
			return head
		}
		newHead = newHead.Next
	}
	pre := reverse(head, k)
	head.Next = reverseKGroup(newHead, k)
	return pre
}

func reverse(head *ListNode, k int) *ListNode {
	var pre *ListNode
	for i := 0; i < k; i++ {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
