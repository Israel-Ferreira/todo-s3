package db

import (
	"database/sql"
	"fmt"

	"github.com/Israel-Ferreira/todo-s3/internal/config"
	_ "github.com/lib/pq"
)

func OpenDbConnection() (*sql.DB, error) {
	confg := config.ConfigVars

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		confg.DbHost,
		confg.DbPort,
		confg.DbUsername,
		confg.DbPass,
		confg.DbName,
	)

	connection, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	if err := connection.Ping(); err != nil {
		return nil, err
	}

	return connection, nil
}
