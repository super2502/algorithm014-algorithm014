package Day17

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	ret := make([]int, 0)
	helper06(head, &ret)
	return ret
}

func helper06(head *ListNode, ret *[]int) {
	if head == nil {
		return
	}
	helper06(head.Next, ret)
	*ret = append(*ret, head.Val)
}
