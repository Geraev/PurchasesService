package driver

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"log"
	"os"
)

var db *sqlx.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", os.Getenv("MYSQL_URL"))
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}

func GetUUID() string {
	return fmt.Sprint(uuid.Must(uuid.NewV4()))
}
