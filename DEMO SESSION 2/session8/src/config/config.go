package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//ket noi db
/**/
func Connect() (db *sql.DB, err error) {
	dbName := "demo_session_8"
	dbUser := "root"
	dbPassword := ""
	dbDriver := "mysql"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	if err != nil {
		fmt.Println(err)
	} else {
		err2 := db.Ping()
		if err2 != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok")
		}
	}
	return
}
