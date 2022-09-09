package langgo

import (
	"github.com/joho/godotenv"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	"os"
	"path"
)

func Init() {
	core.EnvName = os.Getenv("langgo_env")

	if core.EnvName == "" {
		core.EnvName = core.Development
	}
	if core.WorkerDir == "" {
		core.WorkerDir = os.Getenv("langgo_worker_dir")
	}

	if core.WorkerDir == "" {
		core.WorkerDir, _ = os.Getwd()
		os.Setenv("langgo_worker_dir", core.WorkerDir)
	}

	err := godotenv.Load(path.Join(core.WorkerDir, ".env."+core.EnvName+".yml"))
	if err != nil {
		panic(err)
	}
	l := log.Instance{}
	confName := os.Getenv("langgo_configuration_name")

	confPath := path.Join(core.WorkerDir, confName+".yml")
	err = core.LoadConfigurationFile(confPath)
	if err != nil {
		if core.EnvName == core.Development {
			log.Logger("langgo", "init").Warn().Msg("load app config failed.")
		}
	}

	l.Load()
}

//func init() {
//	Init()
//
//}

func Run(instances ...core.Component) {
	Init()
	core.AddComponents(instances...)
	core.LoadComponents()
	core.SignalNotify()
}
