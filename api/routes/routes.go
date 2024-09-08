package routes

import (
	"example/htmx/api/controller"

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

	r.Run()
}
