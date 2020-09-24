package Week_07

import "testing"

func Test_shortestPathBinaryMatrix(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 1},
		{0, 0, 0, 1},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
	}
	dist := shortestPathBinaryMatrix(grid)
	t.Logf("dist: %v", dist)
}
