package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Pg          PgConfig    `yaml:"pg" env-prefix:"PG_"`
	JWT         JWTConfig   `yaml:"jwt" env-prefix:"JWT_"`
	AuthService AuthService `yaml:"auth_service" env-prefix:"AUTH_SERVICE_"`
	UserService UserService `yaml:"user_service" env-prefix:"USER_SERVICE_"`
}

type AuthService struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost:8088"`
}

type UserService struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost:8088"`
}

type JWTConfig struct {
	Secret string `yaml:"secret" env:"SECRET"`
}

type PgConfig struct {
	Host string `yaml:"host" env:"HOST"`
	Port int    `yaml:"port" env:"PORT"`
	User string `yaml:"user" env:"USER"`
	Pass string `yaml:"pass" env:"PASS"`
	Db   string `yaml:"db" env:"DB"`
}

type RedisConfig struct {
	Host string `yaml:"host" env:"HOST"`
	Port int    `yaml:"port" env:"PORT"`
	Db   int    `yaml:"db" env:"DB"`
}

type RmqConfigs struct {
	User string `yaml:"user" env:"USER"`
	Pass string `yaml:"pass" env:"PASS"`
	Host string `yaml:"host" env:"HOST"`
}

func NewConfig() (*Config, error) {
	path := "config.yaml"

	conf := Config{}
	err := cleanenv.ReadConfig(path, &conf)
	if err != nil {
		return &conf, fmt.Errorf("error unmarshalling config: %w", err)
	}

	err = validateConfigs(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func validateConfigs(config *Config) error {
	if err := validateJWTConfigs(config.JWT); err != nil {
		return err
	}

	return nil
}

func validateJWTConfigs(config JWTConfig) error {
	if config.Secret == "" {
		return errors.New("missing config: jwt.secret")
	}

	return nil
}
