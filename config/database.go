package config

// Get configuration of relational database - postgres
func databaseRDbms() (dbConfig DBConfig, err error) {
	host, err := getEnvOrErr("DB_HOST")
	if err != nil {
		return
	}

	port, err := getEnvOrErr("DB_PORT")
	if err != nil {
		return
	}

	// timezone, err := getEnvOrErr("DB_TIMEZONE")
	// if err != nil {
	// 	return
	// }

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

	sslMode, err := getEnvOrErr("DB_SSL_MODE")
	if err != nil {
		return
	}

	// ENV
	dbConfig.Host = host
	dbConfig.Port = port
	dbConfig.DbName = dbName
	dbConfig.User = username
	dbConfig.Pass = password
	dbConfig.SslMode = sslMode

	// CONN will be implemented later
	// Log will be implemented later

	return
}
