package Day83

import "testing"

func Test_distributeCoins(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCoins(tt.args.root); got != tt.want {
				t.Errorf("distributeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
