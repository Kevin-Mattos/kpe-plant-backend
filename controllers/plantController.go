package controllers

import (
	"database/sql"
	"plant_api/datasource"
	"plant_api/entities"
)

type PlantController interface {
	GetPlant(id int) (*entities.Plant, error)
	GetPlants() (entities.Plants, error)
	CreatePlant(plant *entities.Plant) (*entities.Plant, error)
	DeletePlant(id int) error
}

type PlantControllerImpl struct {
	db datasource.PlantDatabase
}

func CreateController(db *sql.DB) PlantController {
	repos := &PlantControllerImpl{}

	repos.db = datasource.CreatePlantDatabase(db)
	return repos
}

func (repo *PlantControllerImpl) GetPlant(id int) (*entities.Plant, error) {

	plant, err := repo.db.GetPlant(id)
	return plant, err
}

func (repo *PlantControllerImpl) GetPlants() (entities.Plants, error) {
	return repo.db.GetPlants()
}

func (repo *PlantControllerImpl) CreatePlant(plant *entities.Plant) (*entities.Plant, error) {
	return repo.db.CreatePlant(plant)
}

func (repo *PlantControllerImpl) DeletePlant(id int) error {
	return repo.db.DeletePlant(id)
}
