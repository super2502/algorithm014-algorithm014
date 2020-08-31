package Week_04

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	deque := list.New()
	deque.PushBack(&levelNode{
		node:  root,
		level: 0,
	})

	for deque.Len() > 0 {
		cn := deque.Front().Value.(*levelNode)
		deque.Remove(deque.Front())
		if len(ret) < cn.level+1 {
			ret = append(ret, make([]int, 0))
		}
		ret[cn.level] = append(ret[cn.level], cn.node.Val)
		if cn.node.Left != nil {
			deque.PushBack(&levelNode{
				node:  cn.node.Left,
				level: cn.level + 1,
			})
		}
		if cn.node.Right != nil {
			deque.PushBack(&levelNode{
				node:  cn.node.Right,
				level: cn.level + 1,
			})
		}
	}

	return ret
}

type levelNode struct {
	node  *TreeNode
	level int
}
