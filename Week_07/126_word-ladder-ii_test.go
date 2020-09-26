package Week_07

import "testing"

func Test_findLadders(t *testing.T) {
	begin := "hit"
	end := "cog"
	words := []string{"hot","dot","dog","lot","log","cog"}

	ret := findLadders(begin, end, words)
	for _, list := range ret {
		t.Logf("%+v", list)
	}


}
