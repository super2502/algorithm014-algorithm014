package Day92

import (
	"reflect"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{"2-1-1"}, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffWaysToCompute(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
