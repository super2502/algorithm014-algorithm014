package Day153

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	ret := make([]int, 0)
	var findK func(root *TreeNode, dist int)
	findK = func(root *TreeNode, dist int) {
		if root == nil {
			return
		}
		if dist == 0 {
			ret = append(ret, root.Val)
			return
		}
		findK(root.Left, dist-1)
		findK(root.Right, dist-1)
	}
	onStack := make(map[*TreeNode]int) // 构建map保存所有target节点的父节点，以及其到target的距离dist
	// 之后遍历所有父节点，找到与其距离为K-target的子节点即可，期间子节点已经在map中的分支不必再找
	var dfs func(root *TreeNode) (bool, int)
	dfs = func(root *TreeNode) (bool, int) {
		if root == nil {
			return false, 0
		}
		if root == target {
			onStack[root] = 0
			return true, 0
		}
		found, dist := dfs(root.Left)
		if found {
			onStack[root] = dist + 1
			return true, dist + 1
		}
		found, dist = dfs(root.Right)
		if found {
			onStack[root] = dist + 1
			return true, dist + 1
		}
		return false, 0
	}
	dfs(root)
	for node, dist := range onStack {
		if dist > K {
			continue
		}
		if dist == K {
			ret = append(ret, node.Val)
			continue
		}
		if dist == 0 {
			findK(node, K)
			continue
		}
		_, ok := onStack[node.Left]
		if ok {
			findK(node.Right, K-dist-1)
		}
		_, ok = onStack[node.Right]
		if ok {
			findK(node.Left, K-dist-1)
		}
	}
	return ret
}
