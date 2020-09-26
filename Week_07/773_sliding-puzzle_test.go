package Week_07

import "testing"

func Test_slidingPuzzle(t *testing.T) {
	board := [][]int{{1, 2, 3}, {4, 0, 5}}
	board = [][]int{{4, 1, 2}, {5, 0, 3}}
	board = [][]int{{3, 2, 4}, {1, 5, 0}}
	ret := slidingPuzzle(board)
	t.Logf("ret: %v", ret)
}
