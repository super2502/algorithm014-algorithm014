package Day60

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {
	_, num, nextS := parse(S)
	_, root := helper(nextS, num, 0)
	return root
}

func helper(S string, num int, level int) (string, *TreeNode) {

	node := &TreeNode{
		Val: num,
	}
	if S == "" {
		return "", node
	}
	cnt, nextNum, nextS := parse(S)
	//fmt.Printf("level(%v, %v) will parse (%v), get (%v)(%v)(%v)\n", level,num, S, cnt, nextNum, nextS)
	if cnt <= level { // 同层或父层，返回
		return S, node
	}
	nextS, node.Left = helper(nextS, nextNum, level+1)
	//fmt.Printf("%v .left is %v (%v)\n", node.Val, node.Left.Val, nextS)
	if nextS == "" {
		return "", node
	}
	// 看看能不能做右子树
	nextSTmp := nextS
	cnt, nextNum, nextS = parse(nextS)
	if cnt <= level {
		return nextSTmp, node
	}
	nextS, node.Right = helper(nextS, nextNum, level+1)
	fmt.Printf("%v .right is %v\n", node.Val, node.Right.Val)

	return nextS, node

}

func parse(S string) (int, int, string) {
	cnt := 0
	num := ""
	dashOver := false
	i := 0
	for ; i < len(S); i++ {
		if S[i] >= '0' && S[i] <= '9' {
			dashOver = true
			num += string(S[i])
		}
		if S[i] == '-' {
			if dashOver {
				break
			} else {
				cnt++
			}
		}
	}
	n, _ := strconv.Atoi(num)
	return cnt, n, S[i:]
}
