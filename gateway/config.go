package gateway

import "fmt"

type GatewayServerConfig struct {
	Api struct {
		Host string
		Port int
	}
	Services GatewayServices
}

type GatewayServices struct {
	AuthService string
	UserService string
}

func (c *GatewayServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Api.Host, c.Api.Port)
}
