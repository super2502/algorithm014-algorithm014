package Week_08

import (
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	ret := solveNQueens(4)
	for _, board := range ret {
		t.Logf("============")
		for _, row := range board {
			t.Logf("%s", row)
		}
	}
}
