package Week_09

import (
	"reflect"
	"testing"
)

func Test_numIslands2(t *testing.T) {
	type args struct {
		m         int
		n         int
		positions [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{
			3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}, {1, 1}},
		}, []int{1, 1, 2, 3, 1}},
		{"t2", args{
			1, 2, [][]int{{0, 1}, {0, 0}},
		}, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands2(tt.args.m, tt.args.n, tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numIslands2() = %v, want %v", got, tt.want)
			}
		})
	}
}
