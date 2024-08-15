package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	ID           int
	Title        string
	Ingredients  string
	Instructions []string
	ImageURL     string
}

var recipes = []Recipe{
	{
		ID:          1,
		Title:       "Spaghetti Bolognese",
		Ingredients: "Spaghetti, ground beef, tomato sauce, onion, garlic, olive oil, salt, pepper",
		Instructions: []string{
			"1. Cook the spaghetti.",
			"2. Sauté the onion and garlic in olive oil.",
			"3. Add the ground beef and cook until browned.",
			"4. Stir in the tomato sauce and let simmer.",
			" 5. Serve the sauce over the spaghetti."},
		ImageURL: "../assets/spaghetti.avif",
	},
	{
		ID:          2,
		Title:       "Chicken Curry",
		Ingredients: "Chicken breast, curry powder, coconut milk, onion, garlic, ginger, salt, pepper",
		Instructions: []string{
			"1. Sauté the onion, garlic, and ginger in oil.",
			"2. Add the chicken and cook until browned.",
			"3. Stir in the curry powder and coconut milk.",
			"4. Let simmer until the chicken is cooked through.",
			"5. Serve with rice."},
		ImageURL: "../assets/chicken.jpg",
	},
	{
		ID:          3,
		Title:       "Vegetable Stir-Fry",
		Ingredients: "Bell peppers, broccoli, carrots, soy sauce, garlic, ginger, olive oil, salt, pepper",
		Instructions: []string{
			"1. Sauté the garlic and ginger in olive oil.",
			"2. Add the vegetables and stir-fry until tender. ",
			"3. Stir in the soy sauce. 4. Serve over rice or noodles."},
		ImageURL: "../assets/vegitable.webp",
	},
	{
		ID:          4,
		Title:       "Pancakes",
		Ingredients: "Flour, eggs, milk, sugar, baking powder, salt, butter",
		Instructions: []string{
			"1. Mix the dry ingredients.",
			"2. Add the wet ingredients and stir until combined.",
			"3. Heat a pan and cook the pancakes until golden brown on each side.",
			"4. Serve with syrup or your favorite toppings."},
		ImageURL: "../assets/pancakes.webp",
	},
	{
		ID:          5,
		Title:       "Caesar Salad",
		Ingredients: "Romaine lettuce, Caesar dressing, croutons, Parmesan cheese, lemon juice, olive oil, salt, pepper",
		Instructions: []string{
			"1. Chop the lettuce and place in a bowl.",
			"2. Toss with Caesar dressing.",
			"3. Add croutons and Parmesan cheese.",
			"4. Drizzle with lemon juice and olive oil. 5. Season with salt and pepper to taste."},
		ImageURL: "../assets/Caesar.jpg",
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

	r.GET("/recipie/:id", func(c *gin.Context) {
		id := c.Param("id")
		Id, err := strconv.Atoi(id)
		if err == nil {
			recipe := GetRecipeById(Id)
			c.HTML(http.StatusOK, "recipie.html", *recipe)

		}

	})

	r.Run()
}
