package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/0xb10c/memo/memod/config"
	"github.com/0xb10c/memo/memod/logger"
)

var Database *sql.DB

// Setup the database connection
func Setup() (*sql.DB, error) {

	dbUser := config.GetString("database.user")
	dbPasswd := config.GetString("database.passwd")
	dbHost := config.GetString("database.host")
	dbName := config.GetString("database.name")
	dbConnection := config.GetString("database.connection")
	connectionString := dbUser + ":" + dbPasswd + "@" + dbConnection + "(" + dbHost + ")" + "/" + dbName

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Ping the database once since Open() doesn't open a connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	logger.Info.Println("Setup database connection")

	Database = db
	return Database, nil
}