package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var config *GConfig

// GConfig global config
type GConfig struct {
	Mod      string `yaml:"mod"`
	LogLevel int    `yaml:"logLevel"`
}

func LoadConfig() *GConfig {
	vconfig := viper.New()

	vconfig.AutomaticEnv()
	vconfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	vconfig.SetConfigName("config")
	vconfig.AddConfigPath(".")
	vconfig.SetConfigType("yaml")

	err := vconfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := vconfig.Unmarshal(&config); err != nil {
		log.Panicln("\"unmarshal cng file fail " + err.Error())
	}
	return config
}
