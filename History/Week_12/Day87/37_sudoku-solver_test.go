package Day87

import (
	"fmt"
	"testing"
)

var board = [][]byte{
	{'1', '2', '3', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '1', '2', '3', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '1', '2', '3'},
	{'4', '5', '6', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '4', '5', '6', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '4', '5', '6'},
	{'7', '8', '9', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '7', '8', '9', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '7', '8', '9'},
}

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{board}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.args.board)
		})
	}
}

func TestByte(t *testing.T) {
	a := '1'
	intA := int(a - '1')
	fmt.Printf("%v\n", intA)

}
