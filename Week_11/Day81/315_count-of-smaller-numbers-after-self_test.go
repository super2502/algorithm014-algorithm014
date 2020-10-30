package Day81

import (
	"reflect"
	"testing"
)

func Test_countSmaller(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{4, 2, 6, 3, 8, 1, 9, 7}}, []int{3, 1, 2, 1, 2, 0, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmaller(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
