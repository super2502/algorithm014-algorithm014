package Week_08

import "fmt"

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func mergeSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	tmp := slow
	head2 := slow.Next
	tmp.Next = nil // 打折

	head = mergeSortList(head)
	head2 = mergeSortList(head2)
	newHead := mergeList(head, head2)
	return newHead
}

func mergeSortListIteration(head *ListNode) *ListNode {

	length := 0
	tmp := head
	for tmp != nil {
		length++
		tmp = tmp.Next
	}

	//fmt.Printf("len of list %v\n", length)
	step := 1
	for {
		//fmt.Printf("step: %v\n", step)
		if step >= length {
			break
		}
		dummy := &ListNode{
			Next: head,
		}
		pre := dummy
		for i := 0; i < length; i += step * 2 {
			if i + step >= length {
				pre.Next = head
				break
			}
			//fmt.Printf("    into i (%v)\n", i)
			//printList(head)
			l2 := cutList(head, step)
			next := cutList(l2, step)
			//printList(head)
			//printList(l2)
			pre.Next = mergeListIteration(head, l2)
			for pre.Next != nil {
				pre = pre.Next
			}
			head = next
		}
		head = dummy.Next
		step = step * 2
	}

	return head
}

func mergeListIteration(l1, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	pre := preHead
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

	return preHead.Next
}

func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeList(l1, l2.Next)
		return l2
	}
}

func NewListNode(num int) *ListNode {

	return &ListNode{
		Val: num,
	}
}

func buildList(nums []int) *ListNode {

	preHead := &ListNode{}
	pre := preHead
	for _, num := range nums {
		pre.Next = NewListNode(num)
		pre = pre.Next
	}

	return preHead.Next
}

func printList(head *ListNode) {

	s := ""
	for head != nil {
		s += fmt.Sprintf("%v -> ", head.Val)
		head = head.Next
	}
	s += "nil"
	fmt.Printf("%v\n", s)
}

func cutList(head *ListNode, k int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	for i := 0; i < k ; i++ {
		if head == nil {
			break
		}

		pre = pre.Next
		head = head.Next
	}
	pre.Next = nil
	return head
}