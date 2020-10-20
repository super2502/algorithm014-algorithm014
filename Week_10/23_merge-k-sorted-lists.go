package Week_10

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := (len(lists) - 1) / 2
	l1 := merge(lists[:mid+1])
	l2 := merge(lists[mid+1:])

	return mergeTwo(l1, l2)
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwo(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwo(l1, l2.Next)
		return l2
	}
}
