package Day139

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	separator string
	null      string
}

func Constructor() Codec {
	return Codec{
		separator: "#",
		null:      "n",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ret := make([]string, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			ret = append(ret, this.null)
			return
		}
		ret = append(ret, strconv.Itoa(root.Val))
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(ret, this.separator)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, this.separator)
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		if arr[0] == this.null {
			arr = arr[1:]
			return nil
		}
		val, _ := strconv.Atoi(arr[0])
		node := &TreeNode{
			Val: val,
		}
		arr = arr[1:]
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}
