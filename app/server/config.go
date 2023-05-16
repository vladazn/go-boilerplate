package server

import "fmt"

type GrpcServerConfigs struct {
	Host string
	Port int
}

func (c GrpcServerConfigs) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
