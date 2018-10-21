package kira

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Lafriakh/config"
	"github.com/Lafriakh/log"
)

// WithMysql - open an mysql connection.
func (a *App) WithMysql() {
	var err error

	options := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s",
		config.GetString("DB_USERNAME"),
		config.GetString("DB_PASSWORD"),
		config.GetDefault("DB_PROTOCOL", "tcp").(string),
		config.GetString("DB_HOST"),
		config.GetInt("DB_PORT"),
		config.GetString("DB_DATABASE"),
		config.GetDefault("DB_PARAMS", "charset=utf8&parseTime=true").(string),
	)

	// log.Debug(options)

	a.DB, err = sql.Open("mysql", options)
	if err != nil {
		log.Panic(err.Error())
	}

	// Open doesn't open a connection. Validate DSN data:
	err = a.DB.Ping()
	if err != nil {
		log.Panic(err)
	}
}

// CloseMysql - for close mysql connection
func (a *App) CloseMysql() {
	a.DB.Close()
}
