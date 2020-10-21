package Week_10

import (
	"testing"
)

func TestDropEgg(t *testing.T) {
	level := 100
	cnt := 2
	ret := dropEggs(level, cnt)

	t.Logf("level %v, eggs %v, expect %v, yours %v", level, cnt, 14, ret)
}
