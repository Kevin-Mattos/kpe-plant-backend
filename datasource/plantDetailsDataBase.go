package datasource

import (
	"plant_api/entities"

	"github.com/jmoiron/sqlx"
)

const detailsTable = "detail"

type PlantDetailsDatabase interface {
	GetDetail(id int64, filters *[]DBFilter) (*entities.Detail, error)
	GetDetails(filters *[]DBFilter) (*[]*entities.Detail, error)
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

func (database *PlantDetailsDataBaseImpl) GetDetail(id int64, filters *[]DBFilter) (*entities.Detail, error) {

	if filters == nil {
		return GetById[entities.Detail](database.db, detailsTable, id)
	} else {

		query, values, err := Filter(detailsTable, filters)

		if err != nil {
			return nil, err
		}

		var obj entities.Detail
		err = database.db.Get(&obj, query, (*values)...)
		if err != nil {
			return nil, err
		}

		return &obj, nil
	}

}

func (database *PlantDetailsDataBaseImpl) GetDetails(filters *[]DBFilter) (*[]*entities.Detail, error) {
	if filters == nil {
		return GetAll[entities.Detail](database.db, detailsTable)
	} else {
		query, values, err := Filter(detailsTable, filters)

		if err != nil {
			return nil, err
		}

		obj := make([]*entities.Detail, 0)
		err = database.db.Select(&obj, query, (*values)...)
		if err != nil {
			return nil, err
		}

		return &obj, nil
	}
}

func (database *PlantDetailsDataBaseImpl) CreateDetails(detail *entities.Detail) (*entities.Detail, error) {
	return Create(database.db, detailsTable, detail)
}

func (database *PlantDetailsDataBaseImpl) DeleteDetails(id int) error {
	return Delete[entities.Detail](database.db, detailsTable, id)
}
