package Day70

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"ADOBECODEBANC", "ABC"}, "BANC"},
		{"t2", args{"ABDEFGHIJKC", "ABC"}, "ABDEFGHIJKC"},
		{"t3", args{"ABDEFGHIJKCDEFG", "ABC"}, "ABDEFGHIJKC"},
		{"t4", args{"ABDBGACAH", "ABC"}, "BGAC"},
		{"t5", args{"a", "a"}, "a"},
		{"t5", args{"a", "aa"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
