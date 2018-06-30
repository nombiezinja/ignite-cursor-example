package main

import (
	"log"

	_ "github.com/alexbrainman/odbc"

	"github.com/nombiezinja/ignite-cursor-example/db"
	u "github.com/nombiezinja/ignite-cursor-example/utils"
)

var DB = db.Connect()

// Create a table and some hard coded rows to query with
func init() {

	_, err := DB.Exec("DROP TABLE IF EXISTS FOOS;")
	u.FailOnError(err, "Failed to delete table")

	_, err = DB.Exec("CREATE TABLE FOOS (ID VARCHAR(255) PRIMARY KEY, VALUE NUMERIC );")
	u.FailOnError(err, "Failed to create table")

	_, err = DB.Exec("INSERT INTO foos VALUES('9661eef1-8cbd-4cc9-9eac-5786e137f800', 1.0)")
	u.FailOnError(err, "Failed to insert")

	_, err = DB.Exec("INSERT INTO foos VALUES('9661eef1-8cbd-4cc9-9eac-5786e137f801', 1.1)")
	u.FailOnError(err, "Failed to insert")

	_, err = DB.Exec("INSERT INTO foos VALUES('9661eef1-8cbd-4cc9-9eac-5786e137f802', 1.2)")
	u.FailOnError(err, "Failed to insert")

	_, err = DB.Exec("INSERT INTO foos VALUES('9661eef1-8cbd-4cc9-9eac-5786e137f803', 1.3)")
	u.FailOnError(err, "Failed to insert")

}

func main() {

	for i := 0; i <= 300; i++ {
		getEntry()
	}

}

func getEntry() {
	statement := `SELECT ID, VALUE FROM FOOS ORDER BY VALUE DESC LIMIT 1`

	var b Bar

	rows, err := DB.Query(statement)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&b.ID, &b.Value)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Retrieved entry:", b)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

}
