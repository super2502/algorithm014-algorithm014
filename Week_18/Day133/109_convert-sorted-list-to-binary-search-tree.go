package Day133

type ListNode struct {
	Val  int
	Next *ListNode
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	pre := &ListNode{Next: head}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = pre.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	first, second := head, slow.Next
	pre.Next = nil
	tr := &TreeNode{
		Val: slow.Val,
	}
	tr.Left = sortedListToBST(first)
	tr.Right = sortedListToBST(second)
	return tr
}
