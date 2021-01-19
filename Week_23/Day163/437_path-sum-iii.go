package Day163

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	queue := make(map[*TreeNode][]int)
	if root == nil {
		return 0
	}
	queue[root] = []int{sum}
	ret := 0
	for len(queue) > 0 {
		tmp := make(map[*TreeNode][]int)
		for node, targets := range queue {
			nextTargets := []int{sum}
			for _, target := range targets {
				if target == node.Val {
					ret++
				}
				nextTargets = append(nextTargets, target-node.Val)
			}
			if node.Left != nil {
				tmp[node.Left] = nextTargets
			}
			if node.Right != nil {
				tmp[node.Right] = nextTargets
			}
		}
		queue = tmp
	}
	return ret
}
