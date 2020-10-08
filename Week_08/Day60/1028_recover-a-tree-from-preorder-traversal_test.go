package Day60

import "testing"

func Test_parse(t *testing.T) {
	s := "1-2--3--4-5--6--7"
	s = "1-2--3---4-5--6---7"
	recoverFromPreorder(s)
}
