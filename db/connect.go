package db

import (
	"database/sql"
	"fmt"
	"os"

	// _ "github.com/alexbrainman/odbc"
	_ "bitbucket.org/miquella/mgodbc"

	u "github.com/nombiezinja/ignite-cursor-example/utils"
)

func Connect() *sql.DB {
	Setenv()

	var db *sql.DB
	var err error

	dbstring := fmt.Sprintf("Driver=Apache Ignite;ADDRESS=%s;Cache=%s",
		os.Getenv("address"), os.Getenv("cache"))

	db, err = sql.Open("mgodbc", dbstring)
	u.FailOnError(err, "Failed to open ODBC connection")

	err = db.Ping()
	u.FailOnError(err, "Failed to ping ODBC")

	fmt.Println("Apache Ignite connection successfully established")

	return db
}
