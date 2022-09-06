package langgo

import (
	"github.com/joho/godotenv"
	"github.com/langwan/langgo/components/log"
	"github.com/langwan/langgo/core"
	"os"
	"path"
)

func Init() {
	core.EnvName = os.Getenv("langgo_env")

	if core.EnvName == "" {
		core.EnvName = core.Development
	}

	core.WorkerDir = os.Getenv("langgo_worker_dir")

	if core.WorkerDir == "" {
		core.WorkerDir, _ = os.Getwd()
		os.Setenv("langgo_worker_dir", core.WorkerDir)
	}

	err := godotenv.Load(path.Join(core.WorkerDir, ".env."+core.EnvName+".yml"))
	if err != nil {
		panic(err)
	}

	confName := os.Getenv("langgo_configuration_name")

	confPath := path.Join(core.WorkerDir, confName+".yml")
	err = core.LoadConfigurationFile(confPath)
	if err != nil {
		panic(err)
	}
	core.AddComponents(&log.Instance{})
}

func init() {
	Init()

}

func Run() {
	core.SignalNotify()
}
