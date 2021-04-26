package fibnumdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "root"
// 	password = "mysecretpassword"
// 	dbname   = "root"
// )
const (
	dbDriver = "postgres"
	dbSource = "postgres://root:mysecretpassword@localhost:5432/root?sslmode=disable"
)

var (
	Client *sql.DB
)

func init() {
	// datasourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	Client, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
