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
