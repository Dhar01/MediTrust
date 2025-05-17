package config

import (
	"strconv"
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

func databaseRDbms() (rDbms RDBMS, err error) {
	// ENV
	rDbms.Env.Host = mustGetEnv("DB_HOST")
	rDbms.Env.Port = mustGetEnv("DB_PORT")
	rDbms.Env.TimeZone = mustGetEnv("DB_TIMEZONE")

	// ACCESS
	rDbms.Access.DbName = mustGetEnv("DB_NAME")
	rDbms.Access.User = mustGetEnv("DB_USER")
	rDbms.Access.Pass = mustGetEnv("DB_PASS")

	// SSL
	// rDbms.Ssl.SslMode = mustGetEnv("DB_SSL_MODE")

	// CONN will be implemented later
	// Log will be implemented later

	return
}

func databaseRedis() (redis REDIS, err error) {
	poolSize, err := strconv.Atoi(mustGetEnv("POOL_SIZE"))
	if err != nil {
		return
	}

	connTTL, err := strconv.Atoi(mustGetEnv("CONN_TTL"))
	if err != nil {
		return
	}

	redis.Env.Host = mustGetEnv("REDIS_HOST")
	redis.Env.Port = mustGetEnv("REDIS_PORT")
	redis.Conn.PoolSize = poolSize
	redis.Conn.ConnTTL = connTTL
	return
}
