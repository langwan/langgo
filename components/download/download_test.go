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
	httpReader := HttpReader{Url: "https://file-examples.com/storage/fe4bd0f32f63701d79a6df5/2017/04/file_example_MP4_640_3MG.mp4"}
	dl.Download(context.Background(), "./example.mp4", &httpReader, &Listener{})
}

func TestTimeout(t *testing.T) {
	dl := Instance{
		Workers:  5,
		partSize: 1024 * 1024 * 5,
		bufSize:  1024 * 200,
	}
	dl.Run()
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	httpReader := HttpReader{Url: "https://file-examples.com/storage/fe8c7eef0c6364f6c9504cc/2017/04/file_example_MP4_640_3MG.mp4"}
	dl.Download(ctx, "./example.mp4", &httpReader, &Listener{})
}

func TestDownloadFromConfiguration(t *testing.T) {
	core.EnvName = core.Development
	core.LoadConfigurationFile("../../testdata/configuration_test.app.yml")
	langgo.Run(&Instance{})
	httpReader := HttpReader{Url: "https://file-examples.com/storage/fe4bd0f32f63701d79a6df5/2017/04/file_example_MP4_640_3MG.mp4"}

	Get().Download(context.Background(), "./example.mp4", &httpReader, &Listener{})

}
