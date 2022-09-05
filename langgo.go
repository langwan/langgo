package langgo

import (
	"github.com/joho/godotenv"
	"os"
	"path"
)

func init() {
	env := os.Getenv("LANGGO_ENV")
	if env == "" {
		env = "development"
	}
	err := godotenv.Load(".env." + env + ".yml")
	if err != nil {
		panic(err)
	}
	confName := os.Getenv("langgo_configuration_name")
	workerDir := os.Getenv("langgo_worker_dir")
	if workerDir == "" {
		workerDir, _ = os.Getwd()
	}
	confPath := path.Join(workerDir, confName+".yml")
	err = LoadConfigurationFile(confPath)
	if err != nil {
		panic(err)
	}
}
