package Day3

import (
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {

	var x []int
	//x := maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3)
	//
	//t.Logf("%v\n", x)

	x = maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)

	t.Logf("%v\n", x)

}
