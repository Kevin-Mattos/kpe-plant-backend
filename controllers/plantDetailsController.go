package controllers

import (
	"net/http"
	"plant_api/datasource"
	"plant_api/entities"
	"strconv"
	"time"

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

func CreatePlantDetailsController(db datasource.PlantDetailsDatabase) PlantDetailsController {
	return &PlantDetailsControllerImpl{
		db: db,
	}
}

func (repo *PlantDetailsControllerImpl) GetDetail(c *gin.Context) {
	plantIdStr := c.Param("id_plant")

	plantId, err := strconv.ParseInt(plantIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detailIdStr := c.Param("id")

	detailId, err := strconv.ParseInt(detailIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail, err := repo.db.GetDetail(plantId, detailId)

	//	todo verify 404
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, detail)
}

func (repo *PlantDetailsControllerImpl) GetDetails(c *gin.Context) {
	plantIdStr := c.Param("id_plant")

	plantId, err := strconv.ParseInt(plantIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	details, err := repo.db.GetDetails(plantId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}

func (repo *PlantDetailsControllerImpl) CreateDetails(c *gin.Context) {
	var detail entities.Detail
	if err := c.ShouldBind(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plantIdStr := c.Param("id_plant")

	plantId, err := strconv.ParseInt(plantIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail.PlantId = plantId
	detail.Time = time.Now()
	//todo get Time
	_, err = repo.db.CreateDetails(&detail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (repo *PlantDetailsControllerImpl) DeleteDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plantIdStr := c.Param("id_plant")

	plantId, err := strconv.ParseInt(plantIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repo.db.DeleteDetails(plantId, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
