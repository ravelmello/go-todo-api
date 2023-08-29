package database

import (
	"database/sql"
	"fmt"
	"main/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (db *sql.DB, err error) {
	conf := configs.GetDB()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode='disable'",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	return conn, err

}
