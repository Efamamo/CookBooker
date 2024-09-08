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

var client *mongo.Client

// Handler function for Vercel
func main() {
	// Initialize MongoDB client if not already done
	if client == nil {
		var err error
		clientOptions := options.Client().ApplyURI("mongodb+srv://nest:efamamo@cluster0.avreuwg.mongodb.net/recipe-hub")
		client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Initialize repositories, usecases, and controllers
	repo := repository.NewRepo(client)
	usecase := usecase.RecipeUsecase{RecipeRepository: repo}
	c := controller.RecipeController{RecipeUsecase: usecase}

	// Setup routes using Gin
	routes.StartRoute(c)

}
