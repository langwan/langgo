package langgo

import (
	"github.com/joho/godotenv"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	helper_os "github.com/langwan/langgo/helpers/os"
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

	envPath := path.Join(core.WorkerDir, ".env."+core.EnvName+".yml")
	confName := "app"

	if helper_os.FileExists(envPath) {
		err := godotenv.Load(envPath)
		if err != nil {
			log.Logger("langgo", "run").Warn().Err(err).Msg("load env file")
		}
		confName = os.Getenv("langgo_configuration_name")
	} else {
		log.Logger("langgo", "run").Warn().Msg("env file not find")
	}

	l := log.Instance{}

	confPath := path.Join(core.WorkerDir, confName+".yml")
	err := core.LoadConfigurationFile(confPath)
	if err != nil {
		log.Logger("langgo", "run").Warn().Str("path", confPath).Msg("load app config failed.")
	}

	l.Load()
}

func Run(instances ...core.Component) {
	Init()
	core.AddComponents(instances...)
	core.LoadComponents()
	core.SignalNotify()
}

// Deprecated
// RunComponent 方法已经作废，完全可以使用 Run 方法替代
// 当 Run 方法找不到配置的时候，会保留默认传入的默认值
func RunComponent(instances ...core.Component) {
	core.RunComponents(instances...)
}
