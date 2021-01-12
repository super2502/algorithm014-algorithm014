package Day142

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"/home//foo/"}, "/home/foo"},
		{"t2", args{"/../"}, "/"},
		{"t3", args{"/home/"}, "/home"},
		{"t4", args{"/a/./b/../../c/"}, "/c"},
		{"t5", args{"/a/../../b/../c//.//"}, "/c"},
		{"t6", args{"/a//b////c/d//././/.."}, "/a/b/c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
