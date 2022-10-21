package helper_platform

import "testing"

func TestOpenFileExplorer(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "open",
			args: args{
				path: "../../testdata/sample2.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenFileExplorer(tt.args.path)
		})
	}
}
