package fibnumdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	// "github.com/thanhftu/api-multi/utils"
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
	dbSource = "postgresql://root:secret@postgres:5432/root?sslmode=disable"
)

var (
	Client *sql.DB
)

func init() {
	var err error
	// config, err := utils.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config", err)
	// }

	Client, err = sql.Open(dbDriver, dbSource)
	// Client, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
