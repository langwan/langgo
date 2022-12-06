package langgo

import (
	"github.com/joho/godotenv"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	"github.com/langwan/langgo/helpers/os"
	"os"
	"path/filepath"
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

	envPath := filepath.Join(core.WorkerDir, ".env."+core.EnvName+".yml")
	confName := "app"

	if helper_os.FileExists(envPath) {
		err := godotenv.Load(envPath)
		if err != nil {
			log.Logger("langgo", "run").Warn().Err(err).Msg("load env file")
		}
		confName = os.Getenv("langgo_configuration_name")
	}

	l := log.Instance{}

	confPath := filepath.Join(core.WorkerDir, confName+".yml")
	err := core.LoadConfigurationFile(confPath)
	if err != nil {
		log.Logger("langgo", "run").Warn().Str("path", confPath).Msg("load app config failed.")
	}

	l.Load()
}

func Run(instances ...core.Component) {
	Init()
	core.LoadComponents(instances...)
}
