package Day100

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{"t1", args{[]int{2,3,-2,4}}, 6},
		//{"t2", args{[]int{-2,0,-1}}, 0},
		{"t3", args{[]int{-4, -3, -2}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
