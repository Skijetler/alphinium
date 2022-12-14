package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	Rest struct {
		Addr string `yaml:"addr"`
	} `yaml:"rest"`
	Logger struct {
		Dir  string `yaml:"dir"`
		File string `yaml:"file"`
	} `yaml:"logger"`
	Auth struct {
		Addr string `yaml:"addr"`
	} `yaml:"auth"`
}

var instance *Config
var once sync.Once

func GetConfig(confFile string) *Config {
	once.Do(func() {
		instance = &Config{}
		logrus.Info("read application configuration")
		if err := cleanenv.ReadConfig(confFile, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logrus.Info(help)
			logrus.Fatal(err)
		}
	})
	return instance
}
