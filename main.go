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

	for _, user := range a {
		fmt.Printf("%d %s\n", user.ID, user.Name)
	}

	users := getPlants(repo)
	for _, user := range users {
		fmt.Printf("%d %s %d\n", user.ID, user.Nome, user.Idade)
	}

}

func getPlant(repo plantDS.PlantRepository, id int) *entities.Plant {
	user, _ := repo.GetPlant(id)
	return user
}

func getPlants(repo plantDS.PlantRepository) []*entities.Plant {
	users, _ := repo.GetPlants()
	return users
}

func createPlant(repo plantDS.PlantRepository) *entities.Plant {
	user := entities.Plant{
		Nome:  "nominho",
		Idade: 123,
	}

	repo.CreatePlant(&user)
	return &user
}

func deletePlant(repo plantDS.PlantRepository) *entities.Plant {
	user := entities.Plant{
		ID: 3,
	}

	repo.DeletePlant(user.ID)
	return &user
}

func getDetails(repo plantDetailsDS.PlantDetailsRepository) []*entities.Details {
	users, _ := repo.GetDetails()
	return users
}

func createDetails(repo plantDetailsDS.PlantDetailsRepository) *entities.Details {
	user := entities.Details{
		Name: "nominho",
	}

	_, err := repo.CreateDetails(&user)

	if err != nil {
		panic(err)
	}

	return &user
}

// func Teste[T any](value T) any {
// 	var t T
// 	fmt.Printf("%T", t)

// 	var plant entities.Plant
// 	return {&plant.ID, &plant.Nome, &plant.Idade}
// }
