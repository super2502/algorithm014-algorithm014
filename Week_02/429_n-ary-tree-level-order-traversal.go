package Week_02

type Node2 struct {
	Val      int
	Children []*Node2
}

func levelOrder(root *Node2) [][]int {
	ret := make([][]int, 0)
	q := make(queue, 0)
	q.offer(&LvNode{
		Node:  root,
		Level: 0,
	})
	for !q.isEmpty() {
		n := q.poll()
		if n.Node == nil {
			continue
		}
		if len(ret) < n.Level+1 {
			ret = append(ret, make([]int, 0))
		}
		ret[n.Level] = append(ret[n.Level], n.Node.Val)
		for _, child := range n.Node.Children {
			q.offer(&LvNode{
				Node:  child,
				Level: n.Level + 1,
			})
		}
	}

	return ret
}

type LvNode struct {
	Node  *Node2
	Level int
}

type queue []*LvNode

func (q *queue) len() int {
	return len(*q)
}
func (q *queue) isEmpty() bool {
	return q.len() == 0
}

func (q *queue) offer(n *LvNode) {
	*q = append(*q, n)
}
func (q *queue) poll() *LvNode {
	n := (*q)[0]
	*q = (*q)[1:q.len()]
	return n
}
