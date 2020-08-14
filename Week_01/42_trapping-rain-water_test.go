package Week_01

import (
	"testing"
)

func TestWater(t *testing.T) {

	var heights = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	x := trap(heights)

	t.Logf("x is (%v)", x)

}
