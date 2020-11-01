package Day84

import (
	"reflect"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1", args{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}}, []string{"cats and dog", "cat sand dog"}},
		{"t2", args{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}}, []string{}},
		{"t3", args{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Tr(t *testing.T) {
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	tr := &trie{}
	for _, word := range wordDict {
		tr.insert(word)
	}
	t.Logf("search dog (%v)", tr.search("dog"))
	t.Logf("search cats (%v)", tr.search("cats"))
	t.Logf("search mouse (%v)", tr.search("mouse"))
}
