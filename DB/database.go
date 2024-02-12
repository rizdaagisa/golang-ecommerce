package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

const (
	host     = "sql.bsite.net\\MSSQL2016"
	username = "wiryawww_BikeStore"
	password = "Batch13"
	dbName   = "wiryawww_BikeStore"
	port     = ""
)

func Connect() *sql.DB {
	dbAuth := fmt.Sprintf("odbc:server=%s;user id=%s;password=%s;database=%s", host, username, password, dbName)
	open, err := sql.Open("sqlserver", dbAuth)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println("Connection succsess")

	return open
}
