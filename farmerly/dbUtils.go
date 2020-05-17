package farmerly

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connectToDatabase() (db *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePassword := ""
	databaseName := "farmerly"

	db, err := sql.Open(databaseDriver, databaseUser+":"+databasePassword+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	return db
}