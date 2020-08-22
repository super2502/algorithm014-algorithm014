package Week_01

import (
	"testing"
)

func TestLrucache(t *testing.T) {

	lc := Constructor146(2)
	lc.Put(1, 1)
	lc.Put(2, 2)
	t.Logf("lc.Get(1), %v", lc.Get(1))
	lc.Put(3, 3)
	t.Logf("lc.Get(2), %v", lc.Get(2))
	lc.Put(4, 4)
	t.Logf("lc.Get(1), %v", lc.Get(1))
	t.Logf("lc.Get(3), %v", lc.Get(3))
	t.Logf("lc.Get(4), %v", lc.Get(4))

}
