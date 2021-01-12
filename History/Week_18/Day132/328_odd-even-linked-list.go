package Day132

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	isOdd := true
	odd := &ListNode{}
	even := &ListNode{}
	oddPre := odd
	evenPre := even
	for head != nil {
		if isOdd {
			oddPre.Next = head
			oddPre = oddPre.Next
		} else {
			evenPre.Next = head
			evenPre = evenPre.Next
		}
		head = head.Next
		isOdd = !isOdd
	}
	oddPre.Next = even.Next
	evenPre.Next = nil
	return odd.Next
}
