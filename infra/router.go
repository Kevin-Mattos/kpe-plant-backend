package infra

import (
	"plant_api/controllers"
	"plant_api/datasource"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Dispatch(db *sqlx.DB) {

	plantController := controllers.CreatePlantController(datasource.CreatePlantDatabase(db))
	detailsController := controllers.CreatePlantDetailsController(datasource.CreatePlantDetailsDatabase(db))

	router := gin.Default()

	setPlantRoutes(router, plantController)
	setDetailsRoutes(router, detailsController)

	err := router.Run()
	if err != nil {
		panic(err.Error())
	}
}

func setPlantRoutes(router *gin.Engine, controller controllers.PlantController) {
	router.GET("/plant", controller.GetPlants)
	router.POST("/plant", controller.CreatePlant)
	router.GET("/plant/:id", controller.GetPlant)
	router.DELETE("/plant/:id", controller.DeletePlant)
}

func setDetailsRoutes(router *gin.Engine, controller controllers.PlantDetailsController) {
	router.GET("/detail", controller.GetDetails)
	router.POST("/detail", controller.CreateDetails)
	router.GET("/detail/:id", controller.GetDetail)
	router.DELETE("/detail/:id", controller.DeleteDetails)
}
