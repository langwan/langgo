package langgo

import (
	"github.com/joho/godotenv"
	"github.com/langwan/langgo/core"
	"os"
	"path"
)

func Init() {
	env := os.Getenv("LANGGO_ENV")
	if env == "" {
		env = "development"
	}
	workerDir := os.Getenv("langgo_worker_dir")
	if workerDir == "" {
		workerDir, _ = os.Getwd()
		os.Setenv("langgo_worker_dir", workerDir)
	}

	err := godotenv.Load(path.Join(workerDir, ".env."+env+".yml"))
	if err != nil {
		panic(err)
	}
	confName := os.Getenv("langgo_configuration_name")

	confPath := path.Join(workerDir, confName+".yml")
	err = core.LoadConfigurationFile(confPath)
	if err != nil {
		panic(err)
	}
}

func init() {
	Init()

}
