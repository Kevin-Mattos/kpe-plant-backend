package datasource

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"plant_api/secrets"
	"reflect"
	"strings"
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

func GetAll[T any](database *sqlx.DB, table string) (*[]*T, error) {
	query := fmt.Sprintf("SELECT * FROM %s", table)

	obj := make([]*T, 0)

	err := database.Select(&obj, query)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func GetById[T any](database *sqlx.DB, table string, id string) (*T, error) {
	var obj T
	query := fmt.Sprintf("SELECT * FROM %s where %s = $1", table, GetDatabaseIdTag(&obj))

	if err := database.Get(&obj, query, id); err != nil {
		return nil, err
	}
	return &obj, nil
}

func Create[T any](database *sqlx.DB, table string, obj *T) (*T, error) {

	fields := GetDatabaseTagsWithoutId(obj)
	columns := strings.Join(fields, ",")
	values := fmt.Sprintf(":%s", strings.Join(fields, ",:"))

	query := fmt.Sprintf("INSERT into %s(%s) VALUES(%s)", table, columns, values)

	tx := database.MustBegin()

	_, err := tx.NamedExec(query, obj)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func Delete[T any](database *sqlx.DB, table string, id int) error {
	var t T
	query := fmt.Sprintf("DELETE FROM %s where %s = $1", table, GetDatabaseIdTag(&t))
	result, err := database.Exec(query, id)
	affected, _ := result.RowsAffected()

	// todo 404
	if affected == 0 {
		return fmt.Errorf("Not Affected")
	}
	return err
}

func GetDatabaseTagsWithoutId[T any](obj *T) []string {
	fields := reflect.TypeOf(obj).Elem().NumField()
	var fieldArr []string

	for i := 1; i < fields; i++ {
		field := reflect.TypeOf(obj).Elem().Field(i).Tag.Get("db")
		fieldArr = append(fieldArr, field)

	}

	return fieldArr
}

func GetDatabaseIdTag[T any](obj *T) string {
	field := reflect.TypeOf(obj).Elem().Field(0).Tag.Get("db")
	return field
}
