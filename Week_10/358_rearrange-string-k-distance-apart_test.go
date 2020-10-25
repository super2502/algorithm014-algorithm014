package Week_10

import "testing"

func Test_rearrangeString(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{"t1", args{"aabbcc", 3}, "abcabc" },
		//{"t2", args{"aaabc", 3}, "" },
		//{"t3", args{"aaadbbcc", 2}, "abacabcd" },
		//{"t4", args{"aabbcc", 0}, "aabbcc" },
		//{"t4", args{"abb", 2}, "bab" },
		{"t5", args{"bbabcaccaaabababbaaaaccbbcbacbacacccbbcaabcbcacaaccbabbbbbcacccaccbabaccbacabcabcacb", 2}, "cacabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeString(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("rearrangeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
