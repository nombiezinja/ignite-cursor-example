package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/amsokol/ignite-go-client/sql"
	u "github.com/nombiezinja/ignite-cursor-example/utils"
)

func Connect() *sql.DB {
	Setenv()

	var db *sql.DB
	var err error

	dbstring := fmt.Sprintf("Driver=Apache Ignite;ADDRESS=%s;Cache=%s",
		os.Getenv("address"), os.Getenv("cache"))

	db, err = sql.Open("odbc", dbstring)
	u.FailOnError(err, "Failed to open ODBC connection")

	err = db.Ping()
	u.FailOnError(err, "Failed to ping ODBC")

	fmt.Println("Apache Ignite connection successfully established")

	return db
}

func ClientConnect() *sql.DB {

	db, err := sql.Open("ignite", "tcp://localhost:10800/books")

	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	// ping
	err = db.Ping()
	u.FailOnError(err, "Failed to ping ODBC")

	return db
}
