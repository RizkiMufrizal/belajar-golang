package config

import (
	"belajar-golang/exception"
	"github.com/joho/godotenv"
	"os"
)

type ConfigEnv interface {
	Get(key string) string
}

type ConfigEnvImpl struct {
}

func (config *ConfigEnvImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) ConfigEnv {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &ConfigEnvImpl{}
}
