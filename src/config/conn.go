package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "reservasi"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}