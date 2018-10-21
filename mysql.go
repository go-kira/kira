package kira

import (
	"database/sql"
	"fmt"

	"github.com/Lafriakh/log"
	_ "github.com/go-sql-driver/mysql"
)

// WithMysql - open an mysql connection.
func (a *App) WithMysql() {
	var err error

	options := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s",
		a.Configs.GetString("DB_USERNAME"),
		a.Configs.GetString("DB_PASSWORD"),
		a.Configs.GetDefault("DB_PROTOCOL", "tcp").(string),
		a.Configs.GetString("DB_HOST"),
		a.Configs.GetInt("DB_PORT"),
		a.Configs.GetString("DB_DATABASE"),
		a.Configs.GetDefault("DB_PARAMS", "charset=utf8&parseTime=true").(string),
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
