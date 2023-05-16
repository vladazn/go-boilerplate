package db

import "fmt"

type Connection struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	Connection Connection
}

func (c Config) Address() string {
	return fmt.Sprintf("%s:%v", c.Connection.Host, c.Connection.Port)
}
