package Week_02

import (
	"fmt"
	"strconv"
	"strings"
)

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
	q := make(queue297, 0)
	q.offer(root)
	s := ""
	for !q.isEmpty() {
		n := q.poll()
		if n == nil {
			s += this.separator + this.null
			continue
		}
		s += this.separator + strconv.Itoa(n.Val)
		q.offer(n.Left)
		q.offer(n.Right)
	}
	return strings.TrimPrefix(s, this.separator)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	fmt.Printf("%v\n", data)
	arr := strings.Split(data, this.separator)
	if arr[0] == this.null {
		return nil
	}
	rootVal, _ := strconv.Atoi(arr[0])
	root := &TreeNode{
		Val: rootVal,
	}
	if len(arr) == 1 {
		return root
	}
	q := make(queue297, 0)
	q.offer(root)
	i := 0
	for !q.isEmpty() && i+2 < len(arr) {
		n := q.poll()
		if arr[i+1] != this.null {
			left, _ := strconv.Atoi(arr[i+1])
			n.Left = &TreeNode{
				Val: left,
			}
			q.offer(n.Left)
		}
		if arr[i+2] != this.null {
			right, _ := strconv.Atoi(arr[i+2])
			n.Right = &TreeNode{
				Val: right,
			}
			q.offer(n.Right)
		}
		i = i + 2
	}
	return root
}

type queue297 []*TreeNode

func (q *queue297) len() int {
	return len(*q)
}
func (q *queue297) isEmpty() bool {
	return q.len() == 0
}

func (q *queue297) offer(n *TreeNode) {
	*q = append(*q, n)
}

func (q *queue297) poll() *TreeNode {
	n := (*q)[0]
	*q = (*q)[1:q.len()]
	return n
}
