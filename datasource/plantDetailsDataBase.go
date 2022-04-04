package datasource

import (
	"plant_api/entities"

	"github.com/jmoiron/sqlx"
)

const detailsTable = "detail"

type PlantDetailsDatabase interface {
	GetDetail(id int64, filters *map[string]any) (*entities.Detail, error)
	GetDetails(filters *map[string]any) (*[]*entities.Detail, error)
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

func (database *PlantDetailsDataBaseImpl) GetDetail(id int64, filters *map[string]any) (*entities.Detail, error) {

	if filters == nil {
		return GetById[entities.Detail](database.db, detailsTable, id)
	} else {
		return FilterOne[entities.Detail](database.db, detailsTable, filters)
	}

}

func (database *PlantDetailsDataBaseImpl) GetDetails(filters *map[string]any) (*[]*entities.Detail, error) {
	if filters == nil {
		return GetAll[entities.Detail](database.db, detailsTable)
	} else {
		return Filter[entities.Detail](database.db, detailsTable, filters)
	}
}

func (database *PlantDetailsDataBaseImpl) CreateDetails(detail *entities.Detail) (*entities.Detail, error) {
	return Create(database.db, detailsTable, detail)
}

func (database *PlantDetailsDataBaseImpl) DeleteDetails(id int) error {
	return Delete[entities.Detail](database.db, detailsTable, id)
}
