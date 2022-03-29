package controllers

import (
	"database/sql"
	"plant_api/datasource"
	"plant_api/entities"
)

type PlantDetailsController interface {
	GetDetail(id int) (*entities.Detail, error)
	GetDetails() (entities.Details, error)
	CreateDetails(detail *entities.Detail) (*entities.Detail, error)
	DeleteDetails(id int) error
}

type PlantDetailsControllerImpl struct {
	db datasource.PlantDetailsDatabase
}

func CreatePlantDetailsController(db *sql.DB) PlantDetailsController {
	repos := &PlantDetailsControllerImpl{}

	repos.db = datasource.CreatePlantDetailsDatabase(db)
	return repos
}

func (repo *PlantDetailsControllerImpl) GetDetail(id int) (*entities.Detail, error) {

	detail, err := repo.db.GetDetail(id)
	return detail, err
}

func (repo *PlantDetailsControllerImpl) GetDetails() (entities.Details, error) {
	return repo.db.GetDetails()
}

func (repo *PlantDetailsControllerImpl) CreateDetails(detail *entities.Detail) (*entities.Detail, error) {
	return repo.db.CreateDetails(detail)
}

func (repo *PlantDetailsControllerImpl) DeleteDetails(id int) error {
	return repo.db.DeleteDetails(id)
}
