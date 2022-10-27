package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Config struct {
	Server struct {
		Http struct {
			Addr string `yaml:"addr"`
		} `yaml:"http"`
		Grpc struct {
			Addr string `yaml:"addr"`
		} `yaml:"grpc"`
	} `yaml:"server"`
	Logger struct {
		Dir  string `yaml:"dir"`
		File string `yaml:"file"`
	} `yaml:"logger"`
	Storage struct {
		Database struct {
			Driver string `yaml:"driver"`
			Source string `yaml:"source"`
		} `yaml:"database"`
		Redis struct {
			Addr     string `yaml:"addr"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
		} `yaml:"redis"`
	} `yaml:"storage"`
	Hasher struct {
		Salt       string `yaml:"salt"`
		Iterations uint32 `yaml:"iterations"`
		Memory     uint32 `yaml:"memory"`
		Threads    uint8  `yaml:"threads"`
		KeyLen     uint32 `yaml:"keyLen"`
	} `yaml:"hasher"`
	UUID struct {
		Iterations int `yaml:"iterations"`
	} `yaml:"uuid"`
	TokenMaker struct {
		AccessKey  string        `yaml:"accessKey"`
		AccessTtl  time.Duration `yaml:"accessTtl"`
		RefreshKey string        `yaml:"refreshKey"`
		RefreshTtl time.Duration `yaml:"refreshTtl"`
		Assert     string        `yaml:"assert"`
	} `yaml:"tokenMaker"`
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
