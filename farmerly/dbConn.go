package farmerly

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func databaseConn() (db *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePassword := ""
	databaseName := "farmerly"

	// mysql, root:Password@/farmerly
	db, err := sql.Open(databaseDriver, databaseUser+":"+databasePassword+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
