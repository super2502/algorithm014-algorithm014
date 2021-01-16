package Day159

import "testing"

func Test_shortestBridge(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestBridge(tt.args.A); got != tt.want {
				t.Errorf("shortestBridge() = %v, want %v", got, tt.want)
			}
		})
	}
}
