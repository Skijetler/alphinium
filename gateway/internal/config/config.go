package config

import "sync"

type Config struct {
	Rest struct {
		Addr   string `yaml:"addr"`
		Logger struct {
			Dir  string `yaml:"dir"`
			File string `yaml:"file"`
		} `yaml:"logger"`
	} `yaml:"rest"`
	Auth struct {
		Addr string `yaml:"addr"`
	} `yaml:"auth"`
}

var instance *Config
var once sync.Once

func GetConfig(confFile string) *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig(confFile, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logrus.Info("read application configuration")
			logrus.Info(help)
			logrus.Fatal(err)
		}
	})
	return instance
}
