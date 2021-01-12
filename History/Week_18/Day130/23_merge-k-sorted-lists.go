package Day130

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	for len(lists) > 1 {
		tmp := make([]*ListNode, 0)
		for i := 0; i < len(lists); i += 2 {
			if i == len(lists)-1 {
				tmp = append(tmp, lists[i])
			} else {
				tmp = append(tmp, merge2Lists(lists[i], lists[i+1]))
			}
		}
		lists = tmp
	}
	if len(lists) == 0 {
		return nil
	}
	return lists[0]
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return dummy.Next
}
