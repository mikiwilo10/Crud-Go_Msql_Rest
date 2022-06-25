package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

var dbConn = &DB{}

func ConexionBD() (*DB, error) {
	Driver := "mysql"
	Usuario := "root"
	Clave := "cuenca"
	Base := "go"

	conexion, err := sql.Open(Driver, Usuario+":"+Clave+"@tcp(127.0.0.1:3306)/"+Base)
	//defer db.Close()

	if err != nil {
		panic(err.Error())
	}
	//defer conexion.Close()
	dbConn.SQL = conexion
	fmt.Println("Success!")
	return dbConn, err
}
