package plantDetailsDS

import (
	"database/sql"
	"plant_api/entities"
)

type PlantDetailsRepository interface {
	GetDetail(userId int) (*entities.Details, error)
	GetDetails() ([]*entities.Details, error)
	CreateDetails(user *entities.Details) (*entities.Details, error)
	DeleteDetails(id int) error
}

type PlantDetailsRepositoryImpl struct {
	db PlantDetailsDatabase
}

func CreatePlantDetailsRepository(db *sql.DB) PlantDetailsRepository {
	repos := &PlantDetailsRepositoryImpl{}

	repos.db = CreatePlantDatabase(db)
	return repos
}

func (repo *PlantDetailsRepositoryImpl) GetDetail(userId int) (*entities.Details, error) {

	user, err := repo.db.GetDetail(userId)
	return user, err
}

func (repo *PlantDetailsRepositoryImpl) GetDetails() ([]*entities.Details, error) {
	return repo.db.GetDetails()
}

func (repo *PlantDetailsRepositoryImpl) CreateDetails(user *entities.Details) (*entities.Details, error) {
	return repo.db.CreateDetails(user)
}

func (repo *PlantDetailsRepositoryImpl) DeleteDetails(id int) error {
	return repo.db.DeleteDetails(id)
}
