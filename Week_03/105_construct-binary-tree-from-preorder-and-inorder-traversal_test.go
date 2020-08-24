package Week_03

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}

	root := buildTree(pre, in)
	t.Logf("root %+v", root)

	pre1 := preTraversal(root)
	t.Logf("pre: %+v", pre1)
}
