package config

import (
	"time"
)

// Entire structs of this file is copied from [pilinux/gorest](https://github.com/pilinux/gorest)
// Licensed under the MIT License

// DatabaseConfig - all database variables
type DatabaseConfig struct {
	RDbms RDBMS
	Redis REDIS
}

// rDbms - relational database variables
type RDBMS struct {
	Activate string
	Env      struct {
		Host     string
		Port     string
		TimeZone string
	}
	Access struct {
		DbName string
		User   string
		Pass   string
	}
	Ssl struct {
		SslMode    string
		MinTLS     string
		RootCA     string
		ServerCert string
		ClientCert string
		ClientKey  string
	}
	Conn struct {
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime time.Duration
	}
	Log struct {
		LogLevel int
	}
}

// REDIS - redis database variables
type REDIS struct {
	Activate string
	Env      struct {
		Host string
		Port string
	}
	Conn struct {
		PoolSize int
		ConnTTL  int
	}
}

// Get configuration of relational database - postgres
func databaseRDbms() (rDbms RDBMS, err error) {
	host, err := getEnvOrErr("DB_HOST")
	if err != nil {
		return
	}

	port, err := getEnvOrErr("DB_PORT")
	if err != nil {
		return
	}

	timezone, err := getEnvOrErr("DB_TIMEZONE")
	if err != nil {
		return
	}

	dbName, err := getEnvOrErr("DB_NAME")
	if err != nil {
		return
	}

	username, err := getEnvOrErr("DB_USER")
	if err != nil {
		return
	}

	password, err := getEnvOrErr("DB_PASS")
	if err != nil {
		return
	}

	// ENV
	rDbms.Env.Host = host
	rDbms.Env.Port = port
	rDbms.Env.TimeZone = timezone

	// ACCESS
	rDbms.Access.DbName = dbName
	rDbms.Access.User = username
	rDbms.Access.Pass = password

	// SSL
	// rDbms.Ssl.SslMode = mustGetEnv("DB_SSL_MODE")

	// CONN will be implemented later
	// Log will be implemented later

	return
}

// Get configuration of REDIS database
func databaseRedis() (redis REDIS, err error) {
	poolSize, err := getEnvNumber("POOL_SIZE")
	if err != nil {
		return
	}

	connTTL, err := getEnvNumber("CONN_TTL")
	if err != nil {
		return
	}

	host, err := getEnvOrErr("REDIS_HOST")
	if err != nil {
		return
	}

	port, err := getEnvOrErr("REDIS_PORT")
	if err != nil {
		return
	}

	redis.Env.Host = host
	redis.Env.Port = port
	redis.Conn.PoolSize = poolSize
	redis.Conn.ConnTTL = connTTL

	return
}
