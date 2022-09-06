package log

import (
	"fmt"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/helpers/io"
	"github.com/langwan/langgo/helpers/reopen"
	"github.com/rs/zerolog"
	"log"
	"path"
)

type Instance struct {
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
}

func (i *Instance) GetName() string {
	return name
}

func Logger(name string, tag string) *zerolog.Logger {
	rp := path.Join(core.WorkerDir, "logs")
	io.CreateFolder(rp, true)
	if it, ok := loggers[name]; ok {

	} else {

		p := path.Join(rp, fmt.Sprintf("%s.log", logName))
		f, err := reopen.NewFileWriter(p)
		if err != nil {
			log.Fatalf("%s create %s log file %s : %v", "langgo", logName, p, err)
		}
		loggers[name] = item{
			logger: zerolog.New(),
			writer: nil,
		}
	}
}
