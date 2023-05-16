package config

import "github.com/vladazn/go-boilerplate/app/server"

func NewGrpcServerConfigs() *server.GrpcServerConfigs {
	return &server.GrpcServerConfigs{
		Port: 8088,
	}
}
