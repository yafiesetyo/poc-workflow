package config

import (
	"log"

	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	Port string    `mapstructure:"port"`
	DB   DBConfig  `mapstructure:"db"`
	GCP  GCPConfig `mapstructure:"gcp"`
}

type (
	// db config
	DBConfig struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbName"`
		Port     string `mapstructure:"port"`
	}
	// gcp infra config
	GCPConfig struct {
		ProjectID     string `mapstructure:"projectId"`
		PubsubTopicID string `mapstructure:"pubsubTopicId"`
	}
)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to initialize config, error: %v \n", err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal config, error: %v \n", err)
	}
}
