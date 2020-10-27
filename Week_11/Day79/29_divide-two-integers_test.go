package Day79

import (
	"math"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{10, 3}, 10 / 3},
		{"t1", args{-55, 10}, -55 / 10},
		{"t2", args{math.MaxInt32, 1}, math.MaxInt32},
		{"t3", args{-math.MaxInt32, -1}, math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
