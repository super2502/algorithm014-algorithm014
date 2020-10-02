package Day54

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	for i := 0; i < k; i++ {
		if head == nil {
			return start
		}
		head = head.Next
	}
	newHead := reserveK(start, k)
	start.Next = reverseKGroup(head, k)
	return newHead
}
func reserveK(head *ListNode, k int) *ListNode {
	var pre *ListNode
	for i := 0; i < k; i++ {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
