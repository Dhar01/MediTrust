package database

import (
	"context"
	"medicine-app/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDB(rDbms config.DBConfig) (*Queries, error) {
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

func GetDSN(rDbms config.DBConfig) string {
	host := "host=" + rDbms.Host
	port := " port=" + rDbms.Port
	username := " user=" + rDbms.User
	database := " dbname=" + rDbms.DbName
	password := " password=" + rDbms.Pass

	dsn := host + port + username + database + password

	sslMode := " sslmode=" + rDbms.SslMode

	dsn += sslMode

	return dsn
}
