package log

import (
	"fmt"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/helpers/os"
	"github.com/langwan/langgo/helpers/reopen"
	"github.com/rs/zerolog"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Instance struct {
}

const name = "log"

var instance *Instance

type item struct {
	logger zerolog.Logger
	writer helper_reopen.Writer
}

var loggers = make(map[string]item)
var lock sync.Mutex

func SetCuttingSignal(sig os.Signal) {
	core.SignalHandlers(func(sig os.Signal) {
		for _, it := range loggers {
			it.writer.Reopen()
		}
	}, sig)
}

func (i *Instance) Load() {
	instance = i
	core.GetComponentConfiguration(name, i)
	if core.EnvName == core.Development {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func (i *Instance) GetName() string {
	return name
}

func Logger(name string, tag string) *zerolog.Logger {
	rp := filepath.Join(core.WorkerDir, "logs")

	if _, ok := loggers[name]; !ok {
		func() {
			lock.Lock()
			if _, ok = loggers[name]; ok {
				fmt.Println("name", "exists")
				return
			}
			defer lock.Unlock()
			helper_os.CreateFolder(rp, true)
			p := filepath.Join(rp, fmt.Sprintf("%s.log", name))
			rf, err := helper_reopen.NewFileWriter(p)
			if err != nil {
				log.Fatalf("%s create %s log file %s : %v", "langgo", name, p, err)
			}
			if core.EnvName == core.Development {
				mf := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen, NoColor: false}, rf)
				l := zerolog.New(mf).With().Str("tag", tag).Timestamp().Logger()
				loggers[name] = item{
					logger: l,
					writer: rf,
				}
			} else {

				l := zerolog.New(rf).With().Str("tag", tag).Timestamp().Logger()
				loggers[name] = item{
					logger: l,
					writer: rf,
				}
			}
		}()
	}
	it := loggers[name]
	return &it.logger
}
