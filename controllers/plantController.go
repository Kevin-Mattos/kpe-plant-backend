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
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	detail, err := repo.db.GetPlant(id)
	//	todo verify 404
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, detail)
}

func (repo *PlantControllerImpl) GetPlants(c *gin.Context) {
	details, err := repo.db.GetPlants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}

func (repo *PlantControllerImpl) CreatePlant(c *gin.Context) {
	var detail entities.Plant
	if err := c.ShouldBind(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := repo.db.CreatePlant(&detail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
