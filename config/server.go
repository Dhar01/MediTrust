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
func server() (serverConfig ServerConfig) {
	serverConfig.ServerHost = mustGetEnv("SERVER_HOST")
	serverConfig.ServerPort = mustGetEnv("SERVER_PORT")
	serverConfig.ServerEnv = strings.ToLower(mustGetEnv("SERVER_ENV"))

	return
}
