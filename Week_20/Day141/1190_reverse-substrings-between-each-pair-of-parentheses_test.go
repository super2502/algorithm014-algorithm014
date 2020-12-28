package Day141

import "testing"

func Test_reverseParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"(abcd)"}, "dcba"},
		{"t2", args{"(u(love)i)"}, "iloveu"},
		{"t3", args{"a(bcdefghijkl(mno)p)q"}, "apmnolkjihgfedcbq"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.args.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
