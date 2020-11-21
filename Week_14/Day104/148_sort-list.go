package Day104

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	printList(head)
	dummy := &ListNode{}
	tmp := head
	cnt := 0
	for tmp != nil {
		cnt++
		tmp = tmp.Next
	}
	dummy.Next = head
	//dummyTmp := dummy
	step := 1
	for {
		pre := dummy
		head := pre.Next
		for i := 0; i < cnt; i += 2 * step {
			//fmt.Printf("step, i (%v)(%v)\n", step, i)
			first, second := cut(head, step)
			if second != nil {
				second, third := cut(second, step)
				merged := merge(first, second)
				pre.Next = merged
				for pre.Next != nil {
					pre = pre.Next
				}
				head = third
			} else {
				pre.Next = first
			}
		}
		step *= 2
		//fmt.Printf("round end\n")
		//printList(dummy.Next)
		if step >= cnt {
			break
		}
	}
	return dummy.Next
}

func merge(l1, l2 *ListNode) *ListNode {
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

func cut(head *ListNode, cnt int) (*ListNode, *ListNode) {
	//fmt.Printf("cut: (%v)", cnt)
	//printList(head)
	firstDummy := &ListNode{}
	first := firstDummy
	for head != nil && cnt > 0 {
		first.Next = head
		cnt--
		head = head.Next
		first = first.Next
	}
	last := first.Next
	first.Next = nil
	//fmt.Printf("get \n")
	//printList(firstDummy.Next)
	//printList(last)
	return firstDummy.Next, last
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Printf("nil\n")
}
