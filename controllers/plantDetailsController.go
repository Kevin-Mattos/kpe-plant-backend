package controllers

import (
	"github.com/gin-gonic/gin"
	"plant_api/datasource"
	"plant_api/entities"
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

func CreatePlantDetailsController(db datasource.PlantDetailsDatabase) PlantDetailsController {
	return &PlantDetailsControllerImpl{
		db: db,
	}
}

func (repo *PlantDetailsControllerImpl) GetDetail(c *gin.Context) {
	Get[entities.Detail](c, repo.db.GetDetail)
}

func (repo *PlantDetailsControllerImpl) GetDetails(c *gin.Context) {
	GetAll[entities.Detail](c, repo.db.GetDetails)
}

func (repo *PlantDetailsControllerImpl) CreateDetails(c *gin.Context) {
	Create[entities.Detail](c, repo.db.CreateDetails)
}

func (repo *PlantDetailsControllerImpl) DeleteDetails(c *gin.Context) {
	Delete(c, repo.db.DeleteDetails)
}
