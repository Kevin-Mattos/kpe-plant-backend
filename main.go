package main

import (
	"fmt"
	"plant_api/datasource"
	plantDetailsDS "plant_api/datasource/details"
	plantDS "plant_api/datasource/plant"
	"plant_api/entities"

	_ "github.com/lib/pq"
)

func main() {

	db := datasource.CreateDatabase()
	defer datasource.Close()
	repo := plantDS.CreateRepository(db)
	detailsRepo := plantDetailsDS.CreatePlantDetailsRepository(db)

	// createDetails(detailsRepo)
	a := getDetails(detailsRepo)

	for _, detail := range a {
		fmt.Printf("%d %s\n", detail.ID, detail.Name)
	}

	plants := getPlants(repo)
	for _, plant := range plants {
		fmt.Printf("%d %s %d\n", plant.ID, plant.Nome, plant.Idade)
	}

}

func getPlant(repo plantDS.PlantRepository, id int) *entities.Plant {
	plant, _ := repo.GetPlant(id)
	return plant
}

func getPlants(repo plantDS.PlantRepository) entities.Plants {
	plants, _ := repo.GetPlants()
	return plants
}

func createPlant(repo plantDS.PlantRepository) *entities.Plant {
	plant := entities.Plant{
		Nome:  "nominho",
		Idade: 123,
	}

	repo.CreatePlant(&plant)
	return &plant
}

func deletePlant(repo plantDS.PlantRepository) *entities.Plant {
	plant := entities.Plant{
		ID: 3,
	}

	repo.DeletePlant(plant.ID)
	return &plant
}

func getDetails(repo plantDetailsDS.PlantDetailsRepository) entities.Details {
	details, _ := repo.GetDetails()
	return details
}

func createDetails(repo plantDetailsDS.PlantDetailsRepository) *entities.Detail {
	detail := entities.Detail{
		Name: "nominho",
	}

	_, err := repo.CreateDetails(&detail)

	if err != nil {
		panic(err)
	}

	return &detail
}
