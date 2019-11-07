package persistence

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DB_HOST")
	port, _  = strconv.Atoi(os.Getenv("DB_PORT"))
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	Db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return Db
}
