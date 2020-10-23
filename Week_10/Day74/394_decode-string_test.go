package Day74

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"3[a]2[bc]"}, "aaabcbc"},
		{"t2", args{"3[a2[c]]"}, "accaccacc"},
		{"t3", args{"2[abc]3[cd]ef"}, "abcabccdcdcdef"},
		{"t4", args{"2[a2[2[c]]d2[e]]"}, "accccdeeaccccdee"},
		{"t5", args{"10[a]"}, "aaaaaaaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
