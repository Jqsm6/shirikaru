package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Host         string        `yaml:"Host"`
		Port         string        `yaml:"Port"`
		LoggingLevel string        `yaml:"LoggingLevel"`
		ReadTimeout  time.Duration `yaml:"ReadTimeout"`
		WriteTimeout time.Duration `yaml:"WriteTimeout"`
	} `yaml:"server"`
	Postgres struct {
		PostgresqlHost     string `yaml:"PostgresqlHost"`
		PostgresqlPort     string `yaml:"PostgresqlPort"`
		PostgresqlUser     string `yaml:"PostgresqlUser"`
		PostgresqlPassword string `yaml:"PostgresqlPassword"`
		PostgresqlDbname   string `yaml:"PostgresqlDbname"`
		PostgresqlSSLMode  bool   `yaml:"PostgresqlSSLMode"`
		PgDriver           string `yaml:"PgDriver"`
	} `yaml:"postgres"`
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}

		err := cleanenv.ReadConfig("config.yml", config)
		if err != nil {
			help, _ := cleanenv.GetDescription(config, nil)
			println(help)
			panic(err)
		}
	})

	return config
}
