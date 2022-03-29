package plantDetailsDS

import (
	"database/sql"
	"plant_api/entities"
)

type PlantDetailsRepository interface {
	GetDetail(id int) (*entities.Details, error)
	GetDetails() ([]*entities.Details, error)
	CreateDetails(detail *entities.Details) (*entities.Details, error)
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

func (repo *PlantDetailsRepositoryImpl) GetDetail(id int) (*entities.Details, error) {

	user, err := repo.db.GetDetail(id)
	return user, err
}

func (repo *PlantDetailsRepositoryImpl) GetDetails() ([]*entities.Details, error) {
	return repo.db.GetDetails()
}

func (repo *PlantDetailsRepositoryImpl) CreateDetails(detail *entities.Details) (*entities.Details, error) {
	return repo.db.CreateDetails(detail)
}

func (repo *PlantDetailsRepositoryImpl) DeleteDetails(id int) error {
	return repo.db.DeleteDetails(id)
}
