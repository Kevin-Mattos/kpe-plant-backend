package datasource

import (
	"fmt"
	"plant_api/entities"

	"github.com/jmoiron/sqlx"
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
	db *sqlx.DB
}

func CreatePlantDatabase(db *sqlx.DB) PlantDatabase {
	database := PlantDataBaseImpl{}
	database.db = db
	return &database
}

func (database *PlantDataBaseImpl) GetPlant(id int) (*entities.Plant, error) {
	query := fmt.Sprintf("SELECT * FROM %s where id = $1", plantsTable)

	var plant entities.Plant

	if err := database.db.Get(&plant, query, id); err != nil {
		return nil, err
	}

	return &plant, nil
}

func (database *PlantDataBaseImpl) GetPlants() (entities.Plants, error) {
	test := Teste{
		db: database.db,
	}

	return GetById[entities.Plants](test, 3, plantsTable)
	// query := fmt.Sprintf("SELECT * FROM %s", plantsTable)

	// var plants entities.Plants

	// err := database.db.Select(&plants, query)
	// if err != nil {
	// 	return nil, err
	// }

	// return plants, nil
}

func (database *PlantDataBaseImpl) CreatePlant(plant *entities.Plant) (*entities.Plant, error) {
	query := fmt.Sprintf("INSERT into %s(nome, idade) VALUES(:nome, :idade)", plantsTable)

	tx := database.db.MustBegin()

	tx.NamedExec(query, plant)
	err := tx.Commit()

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
