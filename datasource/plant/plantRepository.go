package plantDS

import (
	"database/sql"
	"plant_api/entities"
)

type PlantRepository interface {
	GetPlant(id int) (*entities.Plant, error)
	GetPlants() ([]*entities.Plant, error)
	CreatePlant(plant *entities.Plant) (*entities.Plant, error)
	DeletePlant(id int) error
}

type PlantRepositoryImpl struct {
	db PlantDatabase
}

func CreateRepository(db *sql.DB) PlantRepository {
	repos := &PlantRepositoryImpl{}

	repos.db = CreatePlantDatabase(db)
	return repos
}

func (repo *PlantRepositoryImpl) GetPlant(id int) (*entities.Plant, error) {

	plant, err := repo.db.GetPlant(id)
	return plant, err
}

func (repo *PlantRepositoryImpl) GetPlants() ([]*entities.Plant, error) {
	return repo.db.GetPlants()
}

func (repo *PlantRepositoryImpl) CreatePlant(plant *entities.Plant) (*entities.Plant, error) {
	return repo.db.CreatePlant(plant)
}

func (repo *PlantRepositoryImpl) DeletePlant(id int) error {
	return repo.db.DeletePlant(id)
}
