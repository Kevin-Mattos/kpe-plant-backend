package datasource

import (
	"plant_api/entities"

	"github.com/jmoiron/sqlx"
)

const detailsTable = "details"

type PlantDetailsDatabase interface {
	GetDetail(id int) (*entities.Detail, error)
	GetDetails() (*entities.Details, error)
	CreateDetails(detail *entities.Detail) (*entities.Detail, error)
	DeleteDetails(id int) error
}

//TODO GENERICS

type PlantDetailsDataBaseImpl struct {
	db *sqlx.DB
}

func CreatePlantDetailsDatabase(db *sqlx.DB) PlantDetailsDatabase {
	database := PlantDetailsDataBaseImpl{}
	database.db = db
	return &database
}

func (database *PlantDetailsDataBaseImpl) GetDetail(id int) (*entities.Detail, error) {
	return GetById[entities.Detail](database.db, detailsTable, id)
}

func (database *PlantDetailsDataBaseImpl) GetDetails() (*entities.Details, error) {
	return GetAll[entities.Details](database.db, detailsTable)
}

func (database *PlantDetailsDataBaseImpl) CreateDetails(detail *entities.Detail) (*entities.Detail, error) {
	return Create(database.db, detailsTable, detail)
}

func (database *PlantDetailsDataBaseImpl) DeleteDetails(id int) error {
	return Delete[entities.Detail](database.db, detailsTable, id)
}

//umidade e temp externa
//umidade interna
//luminosidade
