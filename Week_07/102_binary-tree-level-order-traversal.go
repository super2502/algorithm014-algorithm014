package Week_07

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	dq := make(deque, 0)
	if root == nil {
		return ret
	}
	dq.Offer(root)
	for !dq.IsEmpty() {
		cns := make([]*TreeNode, 0)
		for !dq.IsEmpty() {
			cns = append(cns, dq.Poll())
		}
		row := make([]int, 0)
		for _, cn := range cns {
			row = append(row, cn.Val)
			if cn.Left != nil {
				dq.Offer(cn.Left)
			}
			if cn.Right != nil {
				dq.Offer(cn.Right)
			}
		}
		if len(row) > 0 {
			ret = append(ret, row)
		}
	}
	return ret
}

type deque []*TreeNode

func (dq *deque) Len() int {
	return len(*dq)
}
func (dq *deque) IsEmpty() bool {
	return dq.Len() == 0
}
func (dq *deque) Offer(x *TreeNode) {
	*dq = append(*dq, x)
}
func (dq *deque) Peek() *TreeNode {
	return (*dq)[0]
}
func (dq *deque) Poll() *TreeNode {
	x := (*dq)[0]
	*dq = (*dq)[1:dq.Len()]
	return x
}
