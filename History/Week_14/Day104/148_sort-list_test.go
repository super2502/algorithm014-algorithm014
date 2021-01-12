package Day104

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	ln := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"t1", args{ln}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
