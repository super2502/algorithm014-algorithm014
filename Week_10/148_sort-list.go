package Week_10

import "fmt"

func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	step := 1
	for {
		//fmt.Printf("step %v\n", step)
		//printList(dummy.Next)
		head := dummy.Next
		newDummy := &ListNode{
			Next: head,
		}
		pre := newDummy
		for head != nil {
			first, others := cut(head, step)
			if others == nil {
				pre.Next = first
				break
			}
			second, others := cut(others, step)
			sorted := mergeList(first, second)
			pre.Next = sorted
			for pre.Next != nil {
				pre = pre.Next
			}
			head = others
		}
		if newDummy.Next == head {
			break
		}
		head = newDummy.Next
		dummy.Next = head
		step *= 2
	}

	return dummy.Next
}

func mergeList(l1, l2 *ListNode) *ListNode {
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
	} else {
		pre.Next = l2
	}
	return dummy.Next
}

func cut(head *ListNode, cnt int) (*ListNode, *ListNode) {
	pre := &ListNode{
		Next: head,
	}
	cur := head
	for i := 0; i < cnt; i++ {
		if cur == nil {
			return head, nil
		}
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = nil
	return head, cur
}
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("nil \n")
}
func initList(nums []int) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for i := 0; i < len(nums); i++ {
		pre.Next = &ListNode{
			Val: nums[i],
		}
		pre = pre.Next
	}
	return dummy.Next
}
