package infra

import (
	"database/sql"
	"plant_api/controllers"

	"github.com/gin-gonic/gin"
)

func Dispatch(db *sql.DB) {

	plantController := controllers.CreatePlantController(db)
	detailsController := controllers.CreatePlantDetailsController(db)

	router := gin.Default()

	setPlantRoutes(router, plantController)
	setDetailsRoutes(router, detailsController)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
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
