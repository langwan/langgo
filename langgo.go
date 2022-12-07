package langgo

import (
	"github.com/joho/godotenv"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	"github.com/langwan/langgo/helpers/os"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
)

func Init() {
	core.EnvName = os.Getenv("langgo_env")

	if core.EnvName == "" {
		core.EnvName = core.Development
	}

	if core.WorkDir == "" {
		core.WorkDir = os.Getenv("langgo_worker_dir")
	}

	if core.WorkDir == "" {
		core.WorkDir, _ = os.Getwd()
		os.Setenv("langgo_worker_dir", core.WorkDir)
	}

	envPath := filepath.Join(core.WorkDir, ".env."+core.EnvName+".yml")
	confName := "app"

	if helper_os.FileExists(envPath) {
		err := godotenv.Load(envPath)
		if err != nil {
			log.Logger("langgo", "run").Warn().Err(err).Msg("load env file")
		}
		confName = os.Getenv("langgo_configuration_name")
	}

	l := log.Instance{}

	confPath := filepath.Join(core.WorkDir, confName+".yml")
	core.LoadConfigurationFile(confPath)
	//if err != nil {
	//	log.Logger("langgo", "run").Warn().Str("path", confPath).Msg("load app config failed.")
	//}

	l.Load()
}

func Run(instances ...core.Component) {
	Init()
	core.LoadComponents(instances...)
}

func LoadComponents(instances ...core.Component) {
	core.LoadComponents(instances...)
}

func Logger(name string, tag string) *zerolog.Logger {
	return log.Logger(name, tag)
}

func SetWorkDir(p string) {
	core.WorkDir = p
}
func GetWorkDir() string {
	return core.WorkDir
}
