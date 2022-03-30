package infra

import (
	"database/sql"
	"plant_api/controllers"
	"plant_api/datasource"

	"github.com/gin-gonic/gin"
)

func Dispatch(db *sql.DB) {

	plantController := controllers.CreatePlantController(datasource.CreatePlantDatabase(db))
	detailsController := controllers.CreatePlantDetailsController(datasource.CreatePlantDetailsDatabase(db))

	router := gin.Default()

	setPlantRoutes(router, plantController)
	setDetailsRoutes(router, detailsController)

	router.Run()
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
