package controller

import (
	"example/htmx/domain"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RecipeController struct {
	RecipeUsecase IRecipeUsecase
}

func (rc RecipeController) GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (rc RecipeController) GetRecipes(c *gin.Context) {
	recipes, err := rc.RecipeUsecase.GetRecipes()
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}
	c.HTML(http.StatusOK, "recipies.html", recipes)
}

func (rc RecipeController) GetForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add_form.html", nil)
}

func (rc RecipeController) GetSingleRecipe(c *gin.Context) {
	id := c.Param("id")
	recipe, err := rc.RecipeUsecase.GetRecipeById(id)

	if err != nil {
		c.IndentedJSON(404, err.Error())
		return
	}

	c.HTML(http.StatusOK, "indiv_recipie.html", recipe)
}

func (rc RecipeController) GetUpdate(c *gin.Context) {
	id := c.Param("id")

	recipe, e := rc.RecipeUsecase.GetRecipeById(id)
	if e != nil {
		c.IndentedJSON(404, gin.H{"error": "recipe not found"})
		return
	}

	c.HTML(http.StatusOK, "edit_recipie.html", recipe)

}

func (rc RecipeController) UpdateRecipe(c *gin.Context) {
	id := c.Param("id")
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Title := c.PostForm("title")
	Ingredients := c.PostForm("ingredients")

	var instructions []string
	for i := 0; ; i++ {
		instruction := c.PostForm("instructions[" + strconv.Itoa(i) + "]")
		if instruction == "" {
			break
		}
		instructions = append(instructions, instruction)
	}
	Instructions := instructions

	file, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(file.Filename)
	filepath := filepath.Join("assets", filename)

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ImageURL := "/" + filepath

	e := rc.RecipeUsecase.UpdateRecipe(id, Title, Ingredients, Instructions, ImageURL)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, e := rc.RecipeUsecase.GetRecipeById(id)

	if e != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "indiv_recipie.html", recipe)
}

func (rc RecipeController) AddRecipe(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newRecipe domain.Recipe

	newRecipe.Title = c.PostForm("title")
	newRecipe.Ingredients = c.PostForm("ingredients")

	var instructions []string
	for i := 0; ; i++ {
		instruction := c.PostForm("instructions[" + strconv.Itoa(i) + "]")
		if instruction == "" {
			break
		}
		instructions = append(instructions, instruction)
	}
	newRecipe.Instructions = instructions

	file, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(file.Filename)
	filepath := filepath.Join("assets", filename)

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newRecipe.ImageURL = "/" + filepath

	rec, err := rc.RecipeUsecase.AddRecipe(newRecipe)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "server error"})
		return
	}

	c.HTML(http.StatusOK, "recipie.html", rec)
}

func (rc RecipeController) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")

	err := rc.RecipeUsecase.DeleteRecipe(id)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "server error"})
		return
	}
	c.Redirect(http.StatusFound, "/recipies")
}
