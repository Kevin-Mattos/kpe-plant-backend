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

	filter := []datasource.DBFilter{}

	filter = append(filter, datasource.DBFilter{Key: "id_plant", Opr: datasource.EQUALS, Value: plantId})
	filter = append(filter, datasource.DBFilter{Key: "id_detail", Opr: datasource.EQUALS, Value: detailId})

	Get[entities.Detail](c, func(id int64) (*entities.Detail, error) {
		return repo.db.GetDetail(id, &filter) //&map[string]any{"id_plant": plantId, "id_detail": detailId})
	})
}

func (repo *PlantDetailsControllerImpl) GetDetails(c *gin.Context) {
	plantIdStr := c.Param("id_plant")

	plantId, err := strconv.ParseInt(plantIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := []datasource.DBFilter{}

	filter = append(filter, datasource.DBFilter{Key: "id_plant", Opr: datasource.EQUALS, Value: plantId})

	GetAll[entities.Detail](c, func() (*[]*entities.Detail, error) {
		return repo.db.GetDetails(&filter)
	})
}

func (repo *PlantDetailsControllerImpl) CreateDetails(c *gin.Context) {
	Create[entities.Detail](c, repo.db.CreateDetails)
}

func (repo *PlantDetailsControllerImpl) DeleteDetails(c *gin.Context) {
	Delete(c, repo.db.DeleteDetails)
}
