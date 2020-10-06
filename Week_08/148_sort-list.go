package Week_08

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return mergeSortList(head)
}

func mergeSortList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	mid := findMid(head)
	right := mid.Next
	mid.Next = nil
	head = mergeSortList(head)
	right = mergeSortList(right)
	return mergeList(head, right)
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = mergeList(left.Next, right)
		return left
	} else {
		right.Next = mergeList(left, right.Next)
		return right
	}
}
