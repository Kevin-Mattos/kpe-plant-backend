package main

import (
	"plant_api/datasource"
	"plant_api/infra"

	_ "github.com/lib/pq"
)

func main() {

	db := datasource.CreateDatabase()
	defer datasource.Close()
	infra.Dispatch(db)
	// repo := controllers.CreatePlantController(db)
	// detailsRepo := controllers.CreatePlantDetailsController(db)

	// // createDetails(detailsRepo)
	// a := getDetails(detailsRepo)

	// for _, detail := range a {
	// 	fmt.Printf("%d %s\n", detail.ID, detail.Name)
	// }

	// plants := getPlants(repo)
	// for _, plant := range plants {
	// 	fmt.Printf("%d %s %d\n", plant.ID, plant.Nome, plant.Idade)
	// }

}
