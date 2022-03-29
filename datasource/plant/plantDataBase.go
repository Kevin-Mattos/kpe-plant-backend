package plantDS

import (
	"database/sql"
	"plant_api/entities"
)

type PlantDatabase interface {
	GetPlant(userId int) (*entities.Plant, error)
	GetPlants() ([]*entities.Plant, error)
	CreatePlant(user *entities.Plant) (*entities.Plant, error)
	DeletePlant(id int) error
}

//TODO GENERICS

type PlantDataBaseImpl struct {
	db *sql.DB
}

func CreatePlantDatabase(db *sql.DB) PlantDatabase {
	database := PlantDataBaseImpl{}
	database.db = db
	return &database
}

func (database *PlantDataBaseImpl) GetPlant(userId int) (*entities.Plant, error) {
	userSql := "SELECT id, nome, idade FROM teste where id = $1"

	var user entities.Plant

	if err := database.db.QueryRow(userSql, userId).Scan(&user.ID, &user.Nome, &user.Idade); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *PlantDataBaseImpl) GetPlants() ([]*entities.Plant, error) {
	userSql := "SELECT id, nome, idade FROM teste"
	// An album slice to hold data from returned rows.
	var users []*entities.Plant

	rows, err := repo.db.Query(userSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {

		user := &entities.Plant{}
		if err := rows.Scan(&user.ID, &user.Nome, &user.Idade); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (database *PlantDataBaseImpl) CreatePlant(user *entities.Plant) (*entities.Plant, error) {
	userSql := "INSERT into TESTE(nome, idade) VALUES($1, $2)"

	_, err := database.db.Exec(userSql, user.Nome, user.Idade)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (database *PlantDataBaseImpl) DeletePlant(id int) error {
	userSql := "DELETE FROM TESTE where id = $1"

	_, err := database.db.Exec(userSql, id)

	return err
}
