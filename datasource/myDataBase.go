package datasource

import (
	"fmt"
	"plant_api/secrets"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func CreateDatabase() *sqlx.DB {
	Connect()
	return db
}

func Connect() {

	var DataSourceName = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", secrets.User, secrets.Password, secrets.Host, secrets.Port, secrets.Dbname)
	fmt.Println(DataSourceName)

	var err error
	db, err = sqlx.Connect("postgres", DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}
}

func Close() {
	db.Close()
}
