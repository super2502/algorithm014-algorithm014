package Week_10

import (
	"testing"
)

func TestDropEgg(t *testing.T) {
	level := 5000
	cnt := 7
	ret := dropEggs(level, cnt)

	t.Logf("level %v, eggs %v, expect %v, yours %v", level, cnt, 14, ret)
}
