package datasource

import (
	"database/sql"
	"fmt"
	"plant_api/entities"
)

type PlantDatabase interface {
	GetPlant(id int) (*entities.Plant, error)
	GetPlants() (entities.Plants, error)
	CreatePlant(plant *entities.Plant) (*entities.Plant, error)
	DeletePlant(id int) error
}

const plantsTable = "teste"

//TODO GENERICS

type PlantDataBaseImpl struct {
	db *sql.DB
}

func CreatePlantDatabase(db *sql.DB) PlantDatabase {
	database := PlantDataBaseImpl{}
	database.db = db
	return &database
}

func (database *PlantDataBaseImpl) GetPlant(id int) (*entities.Plant, error) {
	query := fmt.Sprintf("SELECT id, nome, idade FROM %s where id = $1", plantsTable)

	var plant entities.Plant

	if err := database.db.QueryRow(query, id).Scan(&plant.ID, &plant.Nome, &plant.Idade); err != nil {
		return nil, err
	}

	return &plant, nil
}

func (repo *PlantDataBaseImpl) GetPlants() (entities.Plants, error) {
	query := fmt.Sprintf("SELECT id, nome, idade FROM %s", plantsTable)

	var plants []*entities.Plant

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		plant := &entities.Plant{}
		if err := rows.Scan(&plant.ID, &plant.Nome, &plant.Idade); err != nil {
			return plants, err
		}
		plants = append(plants, plant)
	}
	if err = rows.Err(); err != nil {
		return plants, err
	}

	return plants, nil
}

func (database *PlantDataBaseImpl) CreatePlant(plant *entities.Plant) (*entities.Plant, error) {
	query := fmt.Sprintf("INSERT into %s(nome, idade) VALUES($1, $2)", plantsTable)

	_, err := database.db.Exec(query, plant.Nome, plant.Idade)
	if err != nil {
		return nil, err
	}

	return plant, nil
}

func (database *PlantDataBaseImpl) DeletePlant(id int) error {
	query := fmt.Sprintf("DELETE FROM %s where id = $1", plantsTable)

	_, err := database.db.Exec(query, id)

	return err
}
