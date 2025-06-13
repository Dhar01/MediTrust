package config

import "strings"

// get configuration for server - host, port and env
func server() (appConfig AppConfig, err error) {
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

	appConfig.Host = host
	appConfig.Port = port
	appConfig.Env = strings.ToLower(environment)

	return
}
