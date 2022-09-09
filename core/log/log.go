package log

import (
	"fmt"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/helpers/io"
	"github.com/langwan/langgo/helpers/reopen"
	"github.com/rs/zerolog"
	sysio "io"
	"log"
	"os"
	"path"
	"syscall"
	"time"
)

type Instance struct {
	ReopenSignal syscall.Signal `yaml:"reopen_signal"`
}

const name = "log"

var instance *Instance

type item struct {
	logger zerolog.Logger
	writer reopen.Writer
}

var loggers = make(map[string]item)

func (i *Instance) Load() {
	instance = i
	core.GetComponentConfiguration(name, i)
	if core.EnvName == core.Development {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	if i.ReopenSignal > 0 {
		core.SignalHandle(&core.SignalHandler{
			Sig: i.ReopenSignal,
			F: func() {
				for _, it := range loggers {
					it.writer.Reopen()
				}
			},
		})
	}
}

func (i *Instance) GetName() string {
	return name
}

func Logger(name string, tag string) *zerolog.Logger {
	if len(loggers) == 0 {
		l := zerolog.New(os.Stdout)
		return &l
	}
	rp := path.Join(core.WorkerDir, "logs")
	io.CreateFolder(rp, true)
	if _, ok := loggers[name]; !ok {
		p := path.Join(rp, fmt.Sprintf("%s.log", name))
		rf, err := reopen.NewFileWriter(p)
		if err != nil {
			log.Fatalf("%s create %s log file %s : %v", "langgo", name, p, err)
		}
		if core.EnvName == core.Development {
			mf := sysio.MultiWriter(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}, zerolog.ConsoleWriter{Out: rf, TimeFormat: time.RFC3339, NoColor: true})
			loggers[name] = item{
				logger: zerolog.New(mf),
				writer: rf,
			}
		} else {
			loggers[name] = item{
				logger: zerolog.New(rf),
				writer: rf,
			}
		}
	}
	it := loggers[name]
	return &it.logger
}
