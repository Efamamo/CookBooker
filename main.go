package main

import (
	"example/htmx/domain"
	"example/htmx/repository"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func main() {

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.Static("/styles", "./styles")

	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/recipes", func(c *gin.Context) {

		recipes, err := repository.FindAll()
		if err != nil {
			c.IndentedJSON(500, err.Error())
			return
		}
		c.HTML(http.StatusOK, "recipies.html", recipes)
	})

	r.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_form.html", nil)
	})

	r.GET("/recipes/:id", func(c *gin.Context) {
		id := c.Param("id")
		recipe, err := repository.FindOne(id)

		if err != nil {
			c.IndentedJSON(404, err.Error())
			return
		}

		c.HTML(http.StatusOK, "indiv_recipie.html", recipe)

	})

	r.GET("/recipes/edit/:id", func(c *gin.Context) {
		id := c.Param("id")

		recipe, e := repository.FindOne(id)
		if e != nil {
			if e == mongo.ErrNoDocuments {
				c.IndentedJSON(404, gin.H{"error": "recipe not found"})
				return
			}
			c.IndentedJSON(500, gin.H{"error": "server error"})
		}

		c.HTML(http.StatusOK, "edit_recipie.html", recipe)

	})

	r.PUT("/recipes/:id", func(c *gin.Context) {
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

		e := repository.UpdateOne(id, Title, Ingredients, Instructions, ImageURL)

		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		recipe, e := repository.FindOne(id)

		if e != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "indiv_recipie.html", recipe)

	})

	r.POST("/recipes", func(c *gin.Context) {

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

		rec, err := repository.Save(newRecipe)

		if err != nil {
			c.IndentedJSON(500, gin.H{"error": "server error"})
			return
		}
		c.HTML(http.StatusOK, "recipie.html", rec)

	})

	r.DELETE("/recipes/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := repository.DeleteOne(id)

		if err != nil {
			c.IndentedJSON(500, gin.H{"error": "server error"})
			return
		}
		c.Redirect(http.StatusFound, "/recipies")

	})

	r.Run()
}
