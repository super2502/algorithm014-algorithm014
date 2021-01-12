package Day107

import "testing"

func Test_findMinMinutes(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{4, 7}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinMinutes(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findMinMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
