package Day4

import (
	"path/filepath"
	"strings"
	"testing"
	"fmt"

)

var _ = fmt.Errorf("")

type N struct {
	Val int
	Next *N
}
func TestParsePath(t *testing.T) {

	for _, s :=  range []string{
		"/1/2/3/4",
		"///a//b",
		"/../a/b/c/",
		"/a/b/../c/./d/",
		"/a/../b/c/./d/",
	} {
		r := makePath(s)
		r1, _ := filepath.Abs(s)

		t.Logf("(%v)(%v)(%v)", s, r, r1)
	}


}

func makePath(s string) string {
	arr := strings.Split(s, "/")

	//for _, i := range arr {
	//	fmt.Printf("=== (%v)\n", i)
	//}
	stack := make([]string, 0)

	for i:=1; i <len(arr);i++ {

		if i == len(arr) - 1 && arr[i] == "" {
			break
		}
		stack = append(stack, "/")
		switch arr[i] {
		case "/":
		case "..":
			if len(stack) < 3 {
				stack = stack[:len(stack)-1]
			} else {
				stack = stack[:len(stack)-3]
			}
		case "",".":
			stack = stack[:len(stack) -1]
		default:
			stack = append(stack, arr[i])
		}
	}

	r := ""
	for _, i := range stack {
		r += i
	}
	return r
}


func TestRevK(t *testing.T) {
	a := []int{1,2,3,4,5,6,7,8,9}
	preHead := &N{}
	pre := preHead
	for i:=0;i<len(a);i++ {
		pre.Next = &N{Val: a[i]}
		pre = pre.Next
	}

	//PrintN(preHead.Next, t)

	revN := RevK(preHead.Next, 4)

	//revN := revert(preHead.Next, 9)
	var _ = revN
	PrintN(revN, t)

}

func PrintN(head *N, t *testing.T) {

	for head != nil {
		t.Logf("linked list (%v) ", head.Val)
		head = head.Next
	}


}

func RevK(head *N, k int ) *N{

	if head == nil {
		return nil
	}
	nextHead := head
	for i := 0; i< k; i++ {
		if nextHead == nil {
			k = i
			break
		}
		nextHead = nextHead.Next
	}

	newHead := revert(head, k)

	head.Next = RevK(nextHead, k)
	return newHead
}

func revert(head *N, k int) (*N){

	var pre *N
	cur := head

	for i :=0 ; i<k ;i++ {
		//fmt.Printf("1.pre (%v) cur (%v)\n", pre.Val, cur.Val)
		node := cur.Next
		//fmt.Printf("2.pre (%v) cur (%v) node (%v)\n", pre.Val, cur.Val, node.Next.Val)
		cur.Next = pre
		//fmt.Printf("3.pre (%v) cur (%v)\n", pre.Val, cur.Val)
		pre = cur
		//fmt.Printf("4.pre (%v) cur (%v)\n", pre.Val, cur.Val)
		cur = node
		//fmt.Printf("5.pre (%v) cur (%v)\n", pre.Val, cur.Val)
	}

	return pre

}