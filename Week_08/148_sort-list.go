package Week_08

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return mergeSortList0(head)
}

func mergeSortList0(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	mid := findMid(head)
	right := mid.Next
	mid.Next = nil
	head = mergeSortList0(head)
	right = mergeSortList0(right)
	return mergeList0(head, right)
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList0(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = mergeList0(left.Next, right)
		return left
	} else {
		right.Next = mergeList0(left, right.Next)
		return right
	}
}
