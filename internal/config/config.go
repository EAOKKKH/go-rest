package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel string `yaml:"loglevel" env-default:"debug"`
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port string
}
type PostgresConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDbname   string
	PostgresSSLMode  string
	PgDriver         string
}

func LoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH isn't set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file doesn't exist :%s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Can't read config :%s", err)
	}

	return &cfg
}

func (c *Config) GetDbDsn() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		c.Postgres.PostgresHost, c.Postgres.PostgresPort, c.Postgres.PostgresUser, c.Postgres.PostgresDbname,
		c.Postgres.PostgresSSLMode, c.Postgres.PostgresPassword,
	)
	return dsn
}
