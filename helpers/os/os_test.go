package helper_os

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Listener struct{}

func (l *Listener) ProgressChanged(event *ProgressEvent) {
	fmt.Printf("event = %d, %d / %d\n", event.EventType, event.ConsumedBytes, event.TotalBytes)
}

func TestCopyFileWatcher(t *testing.T) {
	src := "../../testdata/sample.jpg"
	stat, err := os.Stat(src)
	assert.NoError(t, err)
	type args struct {
		source   string
		dest     string
		buf      []byte
		listener IOProgressListener
	}
	tests := []struct {
		name        string
		args        args
		wantWritten int64
		wantErr     bool
	}{
		{
			name: "copy",
			args: args{
				source:   src,
				dest:     "../../testdata/sample2.jpg",
				buf:      nil,
				listener: &Listener{},
			},
			wantWritten: stat.Size(),
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWritten, err := CopyFileWatcher(tt.args.dest, tt.args.source, tt.args.buf, tt.args.listener)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyFileWatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWritten != tt.wantWritten {
				t.Errorf("CopyFileWatcher() gotWritten = %v, want %v", gotWritten, tt.wantWritten)
			}
		})
	}
}

func TestMoveFileWatcher(t *testing.T) {
	src := "../../testdata/sample.jpg"
	stat, err := os.Stat(src)
	assert.NoError(t, err)
	type args struct {
		source   string
		dest     string
		buf      []byte
		listener IOProgressListener
	}
	tests := []struct {
		name        string
		args        args
		wantWritten int64
		wantErr     bool
	}{
		{
			name: "move",
			args: args{
				source:   src,
				dest:     "../../testdata/sample2.jpg",
				buf:      make([]byte, 128*1024),
				listener: &Listener{},
			},
			wantWritten: stat.Size(),
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWritten, err := MoveFileWatcher(tt.args.dest, tt.args.source, tt.args.buf, tt.args.listener)
			if (err != nil) != tt.wantErr {
				t.Errorf("MoveFileWatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWritten != tt.wantWritten {
				t.Errorf("MoveFileWatcher() gotWritten = %v, want %v", gotWritten, tt.wantWritten)
			}
		})
	}
}

func TestGetFileInfo(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args

		wantErr bool
	}{
		{
			name: "info",
			args: args{
				src: "../../testdata/sample.jpg",
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFi, err := GetFileInfo(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotFi.Stat.Size())
			t.Log(gotFi.Mime)
		})
	}
}

func TestFileNameWithoutExt(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				filename: "langgo.mp4",
			},
			want: "langgo",
		}, {
			name: "two",
			args: args{
				filename: "/user/langwan/langgo.mp4",
			},
			want: "/user/langwan/langgo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FileNameWithoutExt(tt.args.filename), "FileNameWithoutExt(%v)", tt.args.filename)
		})
	}
}

func TestTouchFile(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "1",
			args: args{
				p: "/Users/langwan/Documents/a/b",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, TouchFile(tt.args.p, true, true), fmt.Sprintf("TouchFile(%v)", tt.args.p))
		})
	}
}

func TestNewFilename(t *testing.T) {
	type args struct {
		filename string
		tries    int
		rule     func(name string) string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "one",
			args: args{
				filename: "../../testdata/sample.jpg",
				tries:    10,
				rule:     nil,
			},
			want:    "../../testdata/sample1.jpg",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUniqueFilename(tt.args.filename, tt.args.tries, tt.args.rule)
			if !tt.wantErr(t, err, fmt.Sprintf("NewUniqueFilename(%v, %v)", tt.args.filename, tt.args.tries)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewUniqueFilename(%v, %v, %v)", tt.args.filename, tt.args.tries, tt.args.rule)
		})
	}
}

func TestReadFileAt(t *testing.T) {
	type args struct {
		name string
		off  int64
		size int
	}
	tests := []struct {
		name    string
		args    args
		wantB   []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			args: args{
				name: "",
				off:  0,
				size: 0,
			},
			wantB:   nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := ReadFileAt(tt.args.name, tt.args.off, tt.args.size)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadFileAt(%v, %v, %v)", tt.args.name, tt.args.off, tt.args.size)) {
				return
			}
			assert.Equalf(t, tt.wantB, gotB, "ReadFileAt(%v, %v, %v)", tt.args.name, tt.args.off, tt.args.size)
		})
	}
}
