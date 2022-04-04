package controllers

import (
	"github.com/gin-gonic/gin"
	"plant_api/datasource"
	"plant_api/entities"
)

type PlantController interface {
	GetPlant(c *gin.Context)
	GetPlants(c *gin.Context)
	CreatePlant(c *gin.Context)
	DeletePlant(c *gin.Context)
}

type PlantControllerImpl struct {
	db datasource.PlantDatabase
}

func CreatePlantController(db datasource.PlantDatabase) PlantController {
	return &PlantControllerImpl{
		db: db,
	}
}

func (repo *PlantControllerImpl) GetPlant(c *gin.Context) {
	Get[entities.Plant](c, repo.db.GetPlant)
}

func (repo *PlantControllerImpl) GetPlants(c *gin.Context) {
	GetAll[entities.Plant](c, repo.db.GetPlants)
}

func (repo *PlantControllerImpl) CreatePlant(c *gin.Context) {
	Create[entities.Plant](c, repo.db.CreatePlant)
}

func (repo *PlantControllerImpl) DeletePlant(c *gin.Context) {
	Delete(c, repo.db.DeletePlant)
}
