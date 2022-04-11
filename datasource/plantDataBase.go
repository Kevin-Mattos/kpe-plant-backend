package datasource

import (
	"fmt"
	"plant_api/entities"

	"github.com/jmoiron/sqlx"
)

type PlantDatabase interface {
	GetPlant(id int64) (*entities.Plant, error)
	GetPlants() (*[]*entities.Plant, error)
	CreatePlant(plant *entities.Plant) (*entities.Plant, error)
	DeletePlant(id int) error
}

const plantsTable = "plant"

const (
	getPlants    = "SELECT * FROM plant"
	getPlantById = "SELECT * FROM plant where id_plant = $1"
	createPlant  = "INSERT into plant(name, species) VALUES(:name, :species)"
	deletePlant  = "DELETE FROM plant WHERE id_plant = $1"
)

type PlantDataBaseImpl struct {
	db *sqlx.DB
}

func CreatePlantDatabase(db *sqlx.DB) PlantDatabase {
	database := PlantDataBaseImpl{}
	database.db = db
	return &database
}

func (database *PlantDataBaseImpl) GetPlant(id int64) (*entities.Plant, error) {
	var obj entities.Plant

	err := database.db.Get(&obj, getPlantById, id)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (database *PlantDataBaseImpl) GetPlants() (*[]*entities.Plant, error) {

	obj := make([]*entities.Plant, 0)

	err := database.db.Select(&obj, getPlants)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (database *PlantDataBaseImpl) CreatePlant(plant *entities.Plant) (*entities.Plant, error) {

	tx := database.db.MustBegin()

	_, err := tx.NamedExec(createPlant, plant)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return plant, nil
}

func (database *PlantDataBaseImpl) DeletePlant(id int) error {

	result, err := database.db.Exec(deletePlant, id)
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
