package domain

type Recipe struct {
	ID           string   `json:"id" bson:"_id"`
	Title        string   `json:"title"`
	Ingredients  string   `json:"ingredients"`
	Instructions []string `json:"instructions"`
	ImageURL     string   `json:"image"`
}
