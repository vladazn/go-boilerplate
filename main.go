package main

import (
	"github.com/vladazn/go-boilerplate/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
