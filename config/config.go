package config

import (
	"fmt"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Server struct {
		Port         int           `yaml:"port", envconfig:"SERVER_PORT"`
		TimeoutRead  time.Duration `yaml:"TimeoutRead", envconfig:"TimeoutRead"`
		TimeoutWrite time.Duration `yaml:"TimeoutWrite", envconfig:"TimeoutWrite"`
		TimeoutIdle  time.Duration `yaml:"env:TimeoutIdle", envconfig:"TimeoutIdle"`
		Debug        bool          `yaml:"DEBUG", envconfig:"DEBUG"`
	} `yaml:"server"`
	Common struct {
		Xkskey string `envconfig:"XKSKEY", yaml:"XKSKEY"`
	} `yaml:"common"`
}

// type serverConf struct {
// 	Port         string        `yaml:"port", envconfig:"SERVER_PORT" "`
// 	TimeoutRead  time.Duration `yaml:"TimeoutRead", env:"TimeoutRead"`
// 	TimeoutWrite time.Duration `yaml:"TimeoutWrite", env:"TimeoutWrite"`
// 	TimeoutIdle  time.Duration `yaml:"env:TimeoutIdle", env:"TimeoutIdle"`
// 	SecretCode   string        `yaml:"SecretCode", env:"SecretCode"`
// }`yaml:"server"`

func AppConfig() *Conf {
	// var c Conf
	// if err := envdecode.StrictDecode(&c); err != nil {
	// 	log.Fatalf("Failed to decode: %s", err)
	// }
	var cfg Conf
	readFile(&cfg)
	readEnv(&cfg)

	return &cfg
}

func readFile(cfg *Conf) {
	f, err := os.Open("config.yaml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Conf) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
