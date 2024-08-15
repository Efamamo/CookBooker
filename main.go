package main

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Ingredients  string   `json:"ingredients"`
	Instructions []string `json:"instructions"`
	ImageURL     string   `json:"image"`
}

var recipes = []Recipe{
	{
		ID:          1,
		Title:       "Spaghetti Bolognese",
		Ingredients: "Spaghetti, ground beef, tomato sauce, onion, garlic, olive oil, salt, pepper",
		Instructions: []string{
			"Cook the spaghetti.",
			"Sauté the onion and garlic in olive oil.",
			"Add the ground beef and cook until browned.",
			"Stir in the tomato sauce and let simmer.",
			"Serve the sauce over the spaghetti."},
		ImageURL: "../assets/spaghetti.avif",
	},
	{
		ID:          2,
		Title:       "Chicken Curry",
		Ingredients: "Chicken breast, curry powder, coconut milk, onion, garlic, ginger, salt, pepper",
		Instructions: []string{
			"Sauté the onion, garlic, and ginger in oil.",
			"Add the chicken and cook until browned.",
			"Stir in the curry powder and coconut milk.",
			"Let simmer until the chicken is cooked through.",
			"Serve with rice."},
		ImageURL: "../assets/chicken.jpg",
	},
	{
		ID:          3,
		Title:       "Vegetable Stir-Fry",
		Ingredients: "Bell peppers, broccoli, carrots, soy sauce, garlic, ginger, olive oil, salt, pepper",
		Instructions: []string{
			"Sauté the garlic and ginger in olive oil.",
			"Add the vegetables and stir-fry until tender. ",
			"Stir in the soy sauce.",
			"Serve over rice or noodles."},
		ImageURL: "../assets/vegitable.webp",
	},
	{
		ID:          4,
		Title:       "Pancakes",
		Ingredients: "Flour, eggs, milk, sugar, baking powder, salt, butter",
		Instructions: []string{
			"Mix the dry ingredients.",
			"Add the wet ingredients and stir until combined.",
			"Heat a pan and cook the pancakes until golden brown on each side.",
			"Serve with syrup or your favorite toppings."},
		ImageURL: "../assets/pancakes.webp",
	},
	{
		ID:          5,
		Title:       "Caesar Salad",
		Ingredients: "Romaine lettuce, Caesar dressing, croutons, Parmesan cheese, lemon juice, olive oil, salt, pepper",
		Instructions: []string{
			"Chop the lettuce and place in a bowl.",
			"Toss with Caesar dressing.",
			"Add croutons and Parmesan cheese.",
			"Drizzle with lemon juice and olive oil. 5. Season with salt and pepper to taste."},
		ImageURL: "../assets/Caesar.jpg",
	},
	{
		ID:          6,
		Title:       "Beef Tacos",
		Ingredients: "Ground beef, taco seasoning, taco shells, lettuce, tomato, cheese, sour cream",
		Instructions: []string{
			"Cook the ground beef with taco seasoning.",
			"Fill taco shells with the beef mixture.",
			"Top with lettuce, tomato, cheese, and sour cream.",
			"Serve immediately.",
		},
		ImageURL: "../assets/tacos.jpg",
	},
	{
		ID:          7,
		Title:       "Greek Yogurt Parfait",
		Ingredients: "Greek yogurt, honey, granola, mixed berries",
		Instructions: []string{
			"Layer Greek yogurt in a glass.",
			"Drizzle with honey.",
			"Add a layer of granola.",
			"Top with mixed berries.",
			"Repeat layers as desired.",
		},
		ImageURL: "../assets/Yogurt.avif",
	},
	{
		ID:          8,
		Title:       "Margherita Pizza",
		Ingredients: "Pizza dough, tomato sauce, mozzarella cheese, fresh basil, olive oil, salt",
		Instructions: []string{
			"Preheat oven to 475°F (245°C).",
			"Roll out pizza dough on a floured surface.",
			"Spread tomato sauce evenly over the dough.",
			"Sprinkle mozzarella cheese on top.",
			"Bake in the oven for 10-12 minutes.",
			"Garnish with fresh basil and a drizzle of olive oil.",
		},
		ImageURL: "../assets/pizza.jpg",
	},
	{
		ID:          9,
		Title:       "Lemonade",
		Ingredients: "Lemon juice, sugar, water, ice",
		Instructions: []string{
			"In a pitcher, combine lemon juice and sugar.",
			"Stir until the sugar is dissolved.",
			"Add water and mix well.",
			"Serve over ice.",
		},
		ImageURL: "../assets/lemonade.jpg",
	},
}

func GetRecipeById(id int) *Recipe {
	for i, r := range recipes {
		if r.ID == id {
			return &recipes[i]
		}
	}
	return nil
}

func main() {

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.Static("/styles", "./styles")

	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/recipies", func(c *gin.Context) {
		c.HTML(http.StatusOK, "recipies.html", recipes)
	})

	r.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_form.html", nil)
	})

	r.GET("/recipie/:id", func(c *gin.Context) {
		id := c.Param("id")
		Id, err := strconv.Atoi(id)
		if err == nil {
			recipe := GetRecipeById(Id)
			c.HTML(http.StatusOK, "indiv_recipie.html", *recipe)

		}

	})

	r.GET("/recipes/edit/:id", func(c *gin.Context) {
		id := c.Param("id")
		Id, err := strconv.Atoi(id)
		if err == nil {
			recipe := GetRecipeById(Id)
			c.HTML(http.StatusOK, "edit_recipie.html", recipe)
		}
	})

	r.PUT("/recipes/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := c.Request.ParseForm(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		recipe := GetRecipeById(id)
		if recipe != nil {
			recipe.ID = id
			recipe.Title = c.PostForm("title")
			recipe.Ingredients = c.PostForm("ingredients")

			var instructions []string
			for i := 0; ; i++ {
				instruction := c.PostForm("instructions[" + strconv.Itoa(i) + "]")
				if instruction == "" {
					break
				}
				instructions = append(instructions, instruction)
			}
			recipe.Instructions = instructions

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

			recipe.ImageURL = "/" + filepath

			c.HTML(http.StatusOK, "indiv_recipie.html", *recipe)
		}
	})

	r.POST("/recipes", func(c *gin.Context) {

		if err := c.Request.ParseForm(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var newRecipe Recipe

		newRecipe.ID = len(recipes) + 1
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

		recipes = append(recipes, newRecipe)

		c.HTML(http.StatusOK, "recipie.html", newRecipe)

	})

	r.DELETE("/recipes/:id", func(c *gin.Context) {
		id := c.Param("id")
		Id, err := strconv.Atoi(id)
		if err == nil {
			idx := -1
			for i, val := range recipes {
				if val.ID == Id {
					idx = i
				}
			}
			if idx != -1 {
				recipes = append(recipes[:idx], recipes[idx+1:]...)
				c.Redirect(http.StatusFound, "/recipies")
			}

		}

	})

	r.Run()
}
