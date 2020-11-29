package Day111

import "testing"

func Test_reversePairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 3, 2, 3, 1}}, 2},
		//{"t2", args{[]int{2,4,3,5,1}},3},
		//{"t3", args{[]int{5,4,3,2,1}},4},
		//{"t2", args{[]int{-5, -5}},1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.nums); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
