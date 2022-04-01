package datasource

import (
	"fmt"
	"plant_api/entities"

	"github.com/jmoiron/sqlx"
)

const detailsTable = "details"

type PlantDetailsDatabase interface {
	GetDetail(id int) (*entities.Detail, error)
	GetDetails() (entities.Details, error)
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
	query := fmt.Sprintf("SELECT * FROM %s where id = $1", detailsTable)

	var detail entities.Detail

	if err := database.db.Get(&detail, query, id); err != nil {
		return nil, err
	}

	return &detail, nil
}

func (database *PlantDetailsDataBaseImpl) GetDetails() (entities.Details, error) {
	test := Teste{
		db: database.db,
	}

	return GetById[entities.Details](test, 3, detailsTable)

	// query := fmt.Sprintf("SELECT * FROM %s", detailsTable)

	// var details entities.Details

	// err := database.db.Select(&details, query)
	// if err != nil {
	// 	return nil, err
	// }

	// return details, nil
}

func (database *PlantDetailsDataBaseImpl) CreateDetails(detail *entities.Detail) (*entities.Detail, error) {
	query := fmt.Sprintf("INSERT into %s(name) VALUES(:name)", detailsTable)

	tx := database.db.MustBegin()

	tx.NamedExec(query, detail)
	err := tx.Commit()
	if err != nil {
		return nil, err
	}

	return detail, nil
}

func (database *PlantDetailsDataBaseImpl) DeleteDetails(id int) error {
	query := fmt.Sprintf("DELETE FROM %s where id = $1", detailsTable)

	_, err := database.db.Exec(query, id)

	return err
}
