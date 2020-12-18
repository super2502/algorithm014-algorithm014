package Day131

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	dummy := &ListNode{}
	pre := dummy
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			val := head.Val
			for head != nil && head.Val == val {
				head = head.Next
			}
		} else {
			pre.Next = head
			head = head.Next
			pre = pre.Next
		}
	}
	pre.Next = head
	return dummy.Next
}
