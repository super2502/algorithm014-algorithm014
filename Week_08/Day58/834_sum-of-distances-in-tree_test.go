package Day58

import (
	"testing"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	N := 6
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	N = 4
	edges = [][]int{{1, 2}, {2, 0}, {0, 3}}
	ret := sumOfDistancesInTree(N, edges)

	t.Logf("%+v", edges)
	t.Logf("ret: %+v", ret)
}
