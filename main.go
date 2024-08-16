package main

import (
	"context"
	"example/htmx/api/controller"
	"example/htmx/api/routes"
	"example/htmx/repository"
	"example/htmx/usecase"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	var client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepo(client)
	usecase := usecase.RecipeUsecase{RecipeRepository: repo}
	c := controller.RecipeController{RecipeUsecase: usecase}

	routes.StartRoute(c)

}
