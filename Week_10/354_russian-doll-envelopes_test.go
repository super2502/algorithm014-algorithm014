package Week_10

import "testing"

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}}, 3},
		{"t2", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {5, 5}, {6, 7}, {7, 8}}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
