package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/iamaul/evonix-backend-api/config"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	SQL *sql.DB
}

var dbConnection = &Database{}

func ConnectDatabase(c *config.Configuration) (*Database, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConnect, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConnect.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// defer dbConnect.Close()

	dbConnection.SQL = dbConnect

	return dbConnection, err
}
