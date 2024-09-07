package Handler

import (
	"context"
	"example/htmx/api/controller"
	"example/htmx/api/routes"
	"example/htmx/repository"
	"example/htmx/usecase"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Handler function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize MongoDB client if not already done
	if client == nil {
		var err error
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
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
	router := routes.StartRoute(c)

	// Serve the request using Gin's ServeHTTP
	router.ServeHTTP(w, r)
}

