package service

import (
	"errors"
	"os"
)

func Port() (string, error) {
	args := os.Args[1:]
	var port string
	if len(args) > 1 {
		return "", errors.New("input port number like go run . [PORT] or go run . ")
	} else if len(args) == 1 {
		port = args[0]
	} else {
		port = "8989"
	}
	return port, nil
}
