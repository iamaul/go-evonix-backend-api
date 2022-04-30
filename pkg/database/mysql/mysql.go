package mysql

import (
	"fmt"
	"time"

	"github.com/iamaul/go-evonix-backend-api/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

var conn *sqlx.DB

func NewMysqlDB(c *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.Mysql.MysqlUser,
		c.Mysql.MysqlPassword,
		c.Mysql.MysqlHost,
		c.Mysql.MysqlPort,
		c.Mysql.MysqlDbname,
	)

	db, err := sqlx.Connect(c.Mysql.MysqlDriver, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}
	SetConnection(db)

	return db, nil
}

func GetConnection() *sqlx.DB {
	return conn
}

func SetConnection(connection *sqlx.DB) {
	conn = connection
}
