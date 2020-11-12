package Day93

import "testing"

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{[]int{1, 2, 3}}},
		{"t2", args{[]int{3, 1, 1, 2, 1}}},
		{"t3", args{[]int{3, 2, 1}}},
		{"t4", args{[]int{1, 2, 4, 3, 5, 6}}},
		{"t5", args{[]int{1, 3, 6, 4, 2}}},
		{"t6", args{[]int{1, 3, 6, 6}}},
		{"t7", args{[]int{3, 6, 6}}},
		{"t8", args{[]int{3, 3, 6}}},
		{"t9", args{[]int{6, 6, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
		})
	}
}
