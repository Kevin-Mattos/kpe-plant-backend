package datasource

import (
	"plant_api/entities"

	"github.com/jmoiron/sqlx"
)

type PlantDatabase interface {
	GetPlant(id int) (*entities.Plant, error)
	GetPlants() (entities.Plants, error)
	CreatePlant(plant *entities.Plant) (*entities.Plant, error)
	DeletePlant(id int) error
}

const plantsTable = "plant"

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

	return GetById[entities.Plant](database.db, plantsTable, id)
}

func (database *PlantDataBaseImpl) GetPlants() (entities.Plants, error) {
	return GetAll[entities.Plant](database.db, plantsTable)
}

func (database *PlantDataBaseImpl) CreatePlant(plant *entities.Plant) (*entities.Plant, error) {
	return Create(database.db, plantsTable, plant)
}

func (database *PlantDataBaseImpl) DeletePlant(id int) error {
	return Delete[entities.Plant](database.db, plantsTable, id)
}
