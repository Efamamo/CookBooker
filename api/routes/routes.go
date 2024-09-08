package routes

import (
	"example/htmx/api/controller"
	"os" // To read the environment variables

	"github.com/gin-gonic/gin"
)

func StartRoute(controller controller.RecipeController) {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.Static("/styles", "./styles")

	r.LoadHTMLGlob("views/*")

	r.GET("/", controller.GetHome)
	r.GET("/recipes", controller.GetRecipes)
	r.GET("/form", controller.GetForm)
	r.GET("/recipes/:id", controller.GetSingleRecipe)
	r.GET("/recipes/edit/:id", controller.GetUpdate)

	r.PUT("/recipes/:id", controller.UpdateRecipe)
	r.POST("/recipes", controller.AddRecipe)
	r.DELETE("/recipes/:id", controller.DeleteRecipe)

	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Fallback to 5000 if no port is provided
	}

	// Use the dynamically assigned port from the environment variable
	r.Run(":" + port)
}
