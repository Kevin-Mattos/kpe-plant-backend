package controllers

import (
	"net/http"
	"plant_api/datasource"
	"plant_api/entities"
	"strconv"

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
	strId := c.Param("id")

	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail, err := repo.db.GetDetail(id)

	//	todo verify 404
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, detail)
}

func (repo *PlantDetailsControllerImpl) GetDetails(c *gin.Context) {

	details, err := repo.db.GetDetails()

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
	_, err := repo.db.CreateDetails(&detail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (repo *PlantDetailsControllerImpl) DeleteDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = repo.db.DeleteDetails(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
