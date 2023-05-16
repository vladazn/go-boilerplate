package config

import "github.com/vladazn/go-boilerplate/gateway"

func NewGatewayServerConfigs(config *Config) *gateway.GatewayServerConfig {
	return &gateway.GatewayServerConfig{
		Api: struct {
			Host string
			Port int
		}{
			Port: 8000,
		},
		Services: gateway.GatewayServices{
			AuthService: config.AuthService.Host,
			UserService: config.UserService.Host,
		},
	}
}
