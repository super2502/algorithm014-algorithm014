package Day107

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		jobs          int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{2, [][]int{{0, 1}, {1, 2}, {2, 1}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.jobs, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
