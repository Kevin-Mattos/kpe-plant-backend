package plantDS

import (
	"database/sql"
	"plant_api/entities"
)

type PlantRepository interface {
	GetPlant(userId int) (*entities.Plant, error)
	GetPlants() ([]*entities.Plant, error)
	CreatePlant(user *entities.Plant) (*entities.Plant, error)
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

func (repo *PlantRepositoryImpl) GetPlant(userId int) (*entities.Plant, error) {

	user, err := repo.db.GetPlant(userId)
	return user, err
}

func (repo *PlantRepositoryImpl) GetPlants() ([]*entities.Plant, error) {
	return repo.db.GetPlants()
}

func (repo *PlantRepositoryImpl) CreatePlant(user *entities.Plant) (*entities.Plant, error) {
	return repo.db.CreatePlant(user)
}

func (repo *PlantRepositoryImpl) DeletePlant(id int) error {
	return repo.db.DeletePlant(id)
}
