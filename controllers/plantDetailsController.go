package controllers

import (
	"database/sql"
	"plant_api/datasource"
	"plant_api/entities"

	"github.com/gin-gonic/gin"
)

type PlantDetailsController interface {
	GetDetail(c *gin.Context)
	GetDetails(c *gin.Context)
	CreateDetails(c *gin.Context)
	DeleteDetails(c *gin.Context)
}

type PlantDetailsControllerImpl struct {
	db datasource.PlantDetailsDatabase
}

func CreatePlantDetailsController(db *sql.DB) PlantDetailsController {
	repos := &PlantDetailsControllerImpl{}

	repos.db = datasource.CreatePlantDetailsDatabase(db)
	return repos
}

func (repo *PlantDetailsControllerImpl) GetDetail(c *gin.Context) {
	id := 3
	repo.db.GetDetail(id)

}

func (repo *PlantDetailsControllerImpl) GetDetails(c *gin.Context) {
	repo.db.GetDetails()
}

func (repo *PlantDetailsControllerImpl) CreateDetails(c *gin.Context) {
	detail := &entities.Detail{}
	repo.db.CreateDetails(detail)
}

func (repo *PlantDetailsControllerImpl) DeleteDetails(c *gin.Context) {
	id := 3
	repo.db.DeleteDetails(id)
}
