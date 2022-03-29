package datasource

import (
	"database/sql"
	"fmt"
	"plant_api/secrets"
)

var db *sql.DB

var schema = `
	CREATE TABLE IF NOT EXISTS DETAILS(id serial, name text)
`

func CreateDatabase() *sql.DB {
	Connect()
	return db
}

func Connect() {
	var DataSourceName = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", secrets.User, secrets.Password, secrets.Host, secrets.Port, secrets.Dbname)
	fmt.Println(DataSourceName)

	var err error
	db, err = sql.Open("postgres", DataSourceName)
	db.Exec(schema)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}
}

func Close() {
	db.Close()
}
