package main

import (
	"fmt"
	"plant_api/controllers"
	"plant_api/datasource"
	"plant_api/entities"

	_ "github.com/lib/pq"
)

func main() {

	db := datasource.CreateDatabase()
	defer datasource.Close()
	repo := controllers.CreateRepository(db)
	detailsRepo := controllers.CreatePlantDetailsRepository(db)

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

func getPlant(repo controllers.PlantRepository, id int) *entities.Plant {
	plant, _ := repo.GetPlant(id)
	return plant
}

func getPlants(repo controllers.PlantRepository) entities.Plants {
	plants, _ := repo.GetPlants()
	return plants
}

func createPlant(repo controllers.PlantRepository) *entities.Plant {
	plant := entities.Plant{
		Nome:  "nominho",
		Idade: 123,
	}

	repo.CreatePlant(&plant)
	return &plant
}

func deletePlant(repo controllers.PlantRepository) *entities.Plant {
	plant := entities.Plant{
		ID: 3,
	}

	repo.DeletePlant(plant.ID)
	return &plant
}

func getDetails(repo controllers.PlantDetailsRepository) entities.Details {
	details, _ := repo.GetDetails()
	return details
}

func createDetails(repo controllers.PlantDetailsRepository) *entities.Detail {
	detail := entities.Detail{
		Name: "nominho",
	}

	_, err := repo.CreateDetails(&detail)

	if err != nil {
		panic(err)
	}

	return &detail
}
