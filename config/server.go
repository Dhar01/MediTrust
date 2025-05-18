package config

import "strings"

// the struct is copied from [pilinux/gorest](https://github.com/pilinux/gorest)
// Licensed under the MIT License
type ServerConfig struct {
	ServerHost string
	ServerPort string // public port of server
	ServerEnv  string
}

// server - host, port and env
func server() (serverConfig ServerConfig, err error) {
	host, err := getEnvOrErr("SERVER_HOST")
	if err != nil {
		return
	}

	port, err := getEnvOrErr("SERVER_PORT")
	if err != nil {
		return
	}

	environment, err := getEnvOrErr("SERVER_ENV")
	if err != nil {
		return
	}

	serverConfig.ServerHost = host
	serverConfig.ServerPort = port
	serverConfig.ServerEnv = strings.ToLower(environment)

	return
}
