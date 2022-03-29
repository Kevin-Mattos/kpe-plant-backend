package plantDetailsDS

import (
	"database/sql"
	"fmt"
	"plant_api/entities"
)

const table = "details"

type PlantDetailsDatabase interface {
	GetDetail(userId int) (*entities.Details, error)
	GetDetails() ([]*entities.Details, error)
	CreateDetails(user *entities.Details) (*entities.Details, error)
	DeleteDetails(id int) error
}

//TODO GENERICS

type PlantDetailsDataBaseImpl struct {
	db *sql.DB
}

func CreatePlantDatabase(db *sql.DB) PlantDetailsDatabase {
	database := PlantDetailsDataBaseImpl{}
	database.db = db
	return &database
}

func (database *PlantDetailsDataBaseImpl) GetDetail(userId int) (*entities.Details, error) {
	userSql := fmt.Sprintf("SELECT id, name FROM %s where id = $1", table)

	var user entities.Details

	if err := database.db.QueryRow(userSql, userId).Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *PlantDetailsDataBaseImpl) GetDetails() ([]*entities.Details, error) {
	userSql := fmt.Sprintf("SELECT id, name FROM %s", table)

	// An album slice to hold data from returned rows.
	var users []*entities.Details

	rows, err := repo.db.Query(userSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {

		user := &entities.Details{}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (database *PlantDetailsDataBaseImpl) CreateDetails(user *entities.Details) (*entities.Details, error) {
	userSql := fmt.Sprintf("INSERT into %s(name) VALUES($1)", table)

	_, err := database.db.Exec(userSql, user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (database *PlantDetailsDataBaseImpl) DeleteDetails(id int) error {
	userSql := fmt.Sprintf("DELETE FROM %s where id = $1", table)

	_, err := database.db.Exec(userSql, id)

	return err
}
