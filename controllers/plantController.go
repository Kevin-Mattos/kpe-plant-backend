package controllers

import (
	"database/sql"
	"plant_api/datasource"
	"plant_api/entities"

	"github.com/gin-gonic/gin"
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

func CreatePlantController(db *sql.DB) PlantController {
	repos := &PlantControllerImpl{}

	repos.db = datasource.CreatePlantDatabase(db)

	return repos
}

func (repo *PlantControllerImpl) GetPlant(c *gin.Context) {

	id := 1
	repo.db.GetPlant(id)

}

func (repo *PlantControllerImpl) GetPlants(c *gin.Context) {
	repo.db.GetPlants()
}

func (repo *PlantControllerImpl) CreatePlant(c *gin.Context) {
	plant := &entities.Plant{}
	repo.db.CreatePlant(plant)
}

func (repo *PlantControllerImpl) DeletePlant(c *gin.Context) {
	id := 1
	repo.db.DeletePlant(id)
}
