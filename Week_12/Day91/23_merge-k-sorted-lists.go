package Day91

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	nh := make(nodeHeap, 0)
	for _, head := range lists {
		if head != nil {
			heap.Push(&nh, head)
		}
	}
	for nh.Len() > 0 {
		node := heap.Pop(&nh).(*ListNode)
		if node.Next != nil {
			heap.Push(&nh, node.Next)
		}
		pre.Next = node
		pre = pre.Next
	}
	return dummy.Next
}

type nodeHeap []*ListNode

func (nh *nodeHeap) Len() int {
	return len(*nh)
}
func (nh *nodeHeap) Less(i, j int) bool {
	return (*nh)[i].Val < (*nh)[j].Val
}
func (nh *nodeHeap) Swap(i, j int) {
	(*nh)[i], (*nh)[j] = (*nh)[j], (*nh)[i]
}
func (nh *nodeHeap) Push(x interface{}) {
	*nh = append(*nh, x.(*ListNode))
}
func (nh *nodeHeap) Pop() interface{} {
	x := (*nh)[nh.Len()-1]
	*nh = (*nh)[:nh.Len()-1]
	return x
}

var _ heap.Interface = &nodeHeap{}

func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	left := mergeKLists1(lists[:mid])
	right := mergeKLists1(lists[mid:])
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
}
