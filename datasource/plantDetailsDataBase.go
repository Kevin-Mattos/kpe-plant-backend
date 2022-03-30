package datasource

import (
	"database/sql"
	"fmt"
	"plant_api/entities"
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
	db *sql.DB
}

func CreatePlantDetailsDatabase(db *sql.DB) PlantDetailsDatabase {
	database := PlantDetailsDataBaseImpl{}
	database.db = db
	return &database
}

func (database *PlantDetailsDataBaseImpl) GetDetail(id int) (*entities.Detail, error) {
	query := fmt.Sprintf("SELECT id, name FROM %s where id = $1", detailsTable)

	var detail entities.Detail

	if err := database.db.QueryRow(query, id).Scan(&detail.ID, &detail.Name); err != nil {
		return nil, err
	}

	return &detail, nil
}

func (repo *PlantDetailsDataBaseImpl) GetDetails() (entities.Details, error) {
	query := fmt.Sprintf("SELECT id, name FROM %s", detailsTable)

	var details []*entities.Detail

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		detail := &entities.Detail{}
		if err := rows.Scan(&detail.ID, &detail.Name); err != nil {
			return details, err
		}
		details = append(details, detail)
	}
	if err = rows.Err(); err != nil {
		return details, err
	}

	return details, nil
}

func (database *PlantDetailsDataBaseImpl) CreateDetails(detail *entities.Detail) (*entities.Detail, error) {
	query := fmt.Sprintf("INSERT into %s(name) VALUES($1)", detailsTable)

	_, err := database.db.Exec(query, detail.Name)
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
