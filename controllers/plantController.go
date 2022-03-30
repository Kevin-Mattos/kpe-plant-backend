package controllers

import (
	"net/http"
	"plant_api/datasource"
	"plant_api/entities"
	"strconv"

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

func CreatePlantController(db datasource.PlantDatabase) PlantController {
	return &PlantControllerImpl{
		db: db,
	}
}

func (repo *PlantControllerImpl) GetPlant(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	plant, err := repo.db.GetPlant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plant)
}

func (repo *PlantControllerImpl) GetPlants(c *gin.Context) {
	plants, err := repo.db.GetPlants()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plants)
}

func (repo *PlantControllerImpl) CreatePlant(c *gin.Context) {

	var plant entities.Plant
	if err := c.ShouldBind(&plant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := repo.db.CreatePlant(&plant)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (repo *PlantControllerImpl) DeletePlant(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repo.db.DeletePlant(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
