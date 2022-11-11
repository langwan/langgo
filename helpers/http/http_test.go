package helper_http

import (
	"context"
	"github.com/panjf2000/ants"
	"testing"
)

func TestDownloadFilePart(t *testing.T) {
	type args struct {
		ctx      context.Context
		url      string
		dst      string
		gp       *ants.PoolWithFunc
		partSize int64
	}

	gp, _ := ants.NewPoolWithFunc(5, func(i interface{}) {
		gpa := i.(*DownloadFilePartTask)
		DownloadPartToFile(gpa.Ctx, gpa.Completed, gpa.Failed, gpa.Url, gpa.Buf, gpa.Dst, gpa.Part)
	})
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "one",
			args: args{
				ctx:      context.Background(),
				url:      "https://file-examples.com/storage/fe8c7eef0c6364f6c9504cc/2017/04/file_example_MP4_640_3MG.mp4",
				dst:      "/Users/langwan/Documents/data/langwan/langgo.git/testdata/dl.mp4",
				gp:       gp,
				partSize: 1024 * 1024,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DownloadFileMultiPart(tt.args.ctx, tt.args.url, tt.args.dst, tt.args.gp, tt.args.partSize); (err != nil) != tt.wantErr {
				t.Errorf("DownloadFilePart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
