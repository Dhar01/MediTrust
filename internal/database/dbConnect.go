package database

import (
	"context"
	"medicine-app/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDB(rDbms config.RDBMS) (*Queries, error) {
	dsn := GetDSN(rDbms)

	conn, err := ConnectDB(dsn)
	if err != nil {
		return nil, err
	}

	return New(conn), nil
}

func ConnectDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func GetDSN(rDbms config.RDBMS) string {
	host := "host=" + rDbms.Env.Host
	port := " port=" + rDbms.Env.Port
	username := " user=" + rDbms.Access.User
	database := " dbname=" + rDbms.Access.DbName
	password := " password=" + rDbms.Access.Pass
	timezone := " TimeZone=" + rDbms.Env.TimeZone

	dsn := host + port + username + database + password + timezone

	// ! currently working for the development
	if rDbms.Ssl.SslMode == "" {
		rDbms.Ssl.SslMode = "disable"
	}

	sslMode := " sslmode=" + rDbms.Ssl.SslMode

	dsn += sslMode

	return dsn
}
