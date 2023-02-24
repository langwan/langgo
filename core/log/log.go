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

var loggerWriters = make(map[string]helper_reopen.Writer)
var lock sync.Mutex

func SetCuttingSignal(sig os.Signal) {
	core.SignalHandlers(func(sig os.Signal) {
		for _, it := range loggerWriters {
			it.Reopen()
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
	rp := filepath.Join(core.WorkDir, "logs")

	if _, ok := loggerWriters[name]; !ok {
		func(loggerWriters map[string]helper_reopen.Writer) {
			lock.Lock()
			if _, ok = loggerWriters[name]; ok {
				return
			}
			defer lock.Unlock()
			helper_os.CreateFolder(rp, true)
			p := filepath.Join(rp, fmt.Sprintf("%s.log", name))
			rf, err := helper_reopen.NewFileWriter(p)
			if err != nil {
				log.Fatalf("%s create %s log file %s : %v", "langgo", name, p, err)
			}
			loggerWriters[name] = rf

		}(loggerWriters)
	}

	if core.EnvName == core.Development {
		mf := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen, NoColor: false}, loggerWriters[name])
		l := zerolog.New(mf).With().Str("tag", tag).Timestamp().Caller().Logger()
		return &l
	} else {
		l := zerolog.New(loggerWriters[name]).With().Str("tag", tag).Timestamp().Logger()
		return &l
	}

}
