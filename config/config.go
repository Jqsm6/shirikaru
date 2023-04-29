package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Host         string        `yaml:"host"`
		Port         string        `yaml:"port"`
		LoggingLevel string        `yaml:"loggingLevel"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
	} `yaml:"server"`
	Postgres struct {
		PostgresqlHost     string `yaml:"postgresqlHost"`
		PostgresqlPort     string `yaml:"postgresqlPort"`
		PostgresqlUser     string `yaml:"postgresqlUser"`
		PostgresqlPassword string `yaml:"postgresqlPassword"`
		PostgresqlDbname   string `yaml:"postgresqlDbname"`
		PgDriver           string `yaml:"pgDriver"`
	} `yaml:"postgres"`
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		config = &Config{}

		err = cleanenv.ReadConfig("config.yml", config)
	})

	if err != nil {
		return nil, err
	}

	return config, nil
}
