package download

import (
	"context"
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	helper_progress "github.com/langwan/langgo/helpers/progress"
	"testing"
	"time"
)

type Listener struct {
}

func (l Listener) ProgressChanged(event *helper_progress.ProgressEvent) {
	fmt.Println(event)
}

func TestDownload(t *testing.T) {
	dl := Instance{
		Workers:  5,
		partSize: 1024 * 1024 * 1,

		bufSize: 1024 * 200,
	}
	dl.Run()
	dl.Download(context.Background(), "https://file-examples.com/storage/fe8c7eef0c6364f6c9504cc/2017/04/file_example_MP4_640_3MG.mp4", "./example.mp4", &Listener{})
}

func TestTimeout(t *testing.T) {
	dl := Instance{
		Workers:  5,
		partSize: 1024 * 1024 * 5,
		bufSize:  1024 * 200,
	}
	dl.Run()
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	dl.Download(ctx, "https://file-examples.com/storage/fe8c7eef0c6364f6c9504cc/2017/04/file_example_MP4_640_3MG.mp4", "./example.mp4", &Listener{})
}

func TestDownloadFromConfiguration(t *testing.T) {
	core.EnvName = core.Development
	core.LoadConfigurationFile("../../testdata/configuration_test.app.yml")
	langgo.Run(&Instance{})
	Get().Tune(10)
	Get().Download(context.Background(), "https://file-examples.com/storage/fe8c7eef0c6364f6c9504cc/2017/04/file_example_MP4_640_3MG.mp4", "./example.mp4", &Listener{})

}
