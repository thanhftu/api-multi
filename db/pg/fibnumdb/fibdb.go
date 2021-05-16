package fibnumdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/thanhftu/api-multi/utils"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "root"
// 	password = "mysecretpassword"
// 	dbname   = "root"
// )

var (
	Client *sql.DB
)

func init() {
	var err error
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println("DBDriver", config.DBDriver)
	fmt.Println("DBSource", config.DBSource)

	Client, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
