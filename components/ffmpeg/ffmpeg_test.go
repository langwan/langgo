package ffmpeg

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"reflect"
	"testing"
	"time"
)

func init() {
	core.EnvName = core.Development
	langgo.Run(&Instance{
		FFMpeg:         "/opt/homebrew/bin/ffmpeg",
		FFProbe:        "/opt/homebrew/bin/ffprobe",
		CommandTimeout: 30 * time.Second,
	})
}

func Test_Thumbnail(t *testing.T) {

	type args struct {
		src       string
		dst       string
		duration  time.Duration
		overwrite bool
	}
	tests := []struct {
		name string

		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				src:       "../../testdata/samples/sample.mp4",
				dst:       "../../testdata/samples/sample.mp4.thumbnail.jpg",
				duration:  10 * time.Second,
				overwrite: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Get().Thumbnail(tt.args.src, tt.args.dst, tt.args.duration, tt.args.overwrite); (err != nil) != tt.wantErr {
				t.Errorf("Thumbnail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Duration(t *testing.T) {
	td, _ := time.ParseDuration("14.667s")
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				src: "../../testdata/samples/sample.mp4",
			},
			want:    td,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get().Duration(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Duration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Duration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBind_Transcoding(t *testing.T) {

	type args struct {
		src       string
		dst       string
		overwrite bool
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []byte
		wantErr    bool
	}{
		{
			name: "1",
			args: args{
				src:       "../../testdata/samples/sample.mov",
				dst:       "../../testdata/samples/sample.mov.mp4",
				overwrite: true,
			},
			wantOutput: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := Get().Transcoding(tt.args.src, tt.args.dst, tt.args.overwrite)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transcoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Transcoding() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
