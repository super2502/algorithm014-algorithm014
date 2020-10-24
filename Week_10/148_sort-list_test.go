package Week_10

import (
	"testing"
)

func Test_printList(t *testing.T) {
	nums := []int{2, 4, 6, 3, 1, 7}
	head := initList(nums)
	printList(head)
}

func Test_cut(t *testing.T) {
	nums := []int{2, 4, 6, 3, 1, 7, 5}
	head := initList(nums)
	printList(head)
	left, right := cut(head, 2)
	printList(left)
	printList(right)
}

func Test_sortList(t *testing.T) {
	nums := []int{2, 4, 6, 3, 1, 7, 5}
	head := initList(nums)
	printList(head)
	newHead := sortList(head)
	printList(newHead)
}
