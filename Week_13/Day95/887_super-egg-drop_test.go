package Day95

import "testing"

func Test_superEggDrop(t *testing.T) {
	type args struct {
		K int
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{2, 100}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop(tt.args.K, tt.args.N); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
