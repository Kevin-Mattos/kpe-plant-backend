package datasource

import (
	"fmt"
	"plant_api/entities"

	"github.com/jmoiron/sqlx"
)

const detailsTable = "detail"

type PlantDetailsDatabase interface {
	GetDetail(plantId int64, detailId int64) (*entities.Detail, error)
	GetDetails(plantId int64) (*[]*entities.Detail, error)
	CreateDetails(detail *entities.Detail) (*entities.Detail, error)
	DeleteDetails(plantId int64, id int64) error
}

const (
	getDetails    = "SELECT * FROM detail where id_plant = $1"
	getDetailById = "SELECT * FROM detail where id_plant = $1 AND id_detail = $2"
	createDetail  = "INSERT into detail(id_plant, time, internal_humidity, external_humidity, temp, luminosity) VALUES(:id_plant, :time, :internal_humidity, :external_humidity, :temp, :luminosity)"
	deleteDetail  = "DELETE FROM detail WHERE id_plant = $1 AND id_detail = $2"
)

type PlantDetailsDataBaseImpl struct {
	db *sqlx.DB
}

func CreatePlantDetailsDatabase(db *sqlx.DB) PlantDetailsDatabase {
	database := PlantDetailsDataBaseImpl{}
	database.db = db
	return &database
}

func (database *PlantDetailsDataBaseImpl) GetDetail(plantId int64, detailId int64) (*entities.Detail, error) {
	var obj entities.Detail

	err := database.db.Get(&obj, getDetailById, plantId, detailId)
	if err != nil {
		return nil, err
	}

	return &obj, nil

}

func (database *PlantDetailsDataBaseImpl) GetDetails(plantId int64) (*[]*entities.Detail, error) {

	obj := make([]*entities.Detail, 0)

	err := database.db.Select(&obj, getDetails, plantId)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (database *PlantDetailsDataBaseImpl) CreateDetails(detail *entities.Detail) (*entities.Detail, error) {
	tx := database.db.MustBegin()

	_, err := tx.NamedExec(createDetail, detail)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return detail, nil
}

func (database *PlantDetailsDataBaseImpl) DeleteDetails(plantId int64, id int64) error {
	result, err := database.db.Exec(deleteDetail, plantId, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// todo 404
	if affected == 0 {
		return fmt.Errorf("Not Affected")
	}
	return err
}
