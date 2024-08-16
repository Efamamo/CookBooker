package repository

import (
	"context"
	"errors"
	"example/htmx/domain"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	collection *mongo.Collection
}

func NewRepo(client *mongo.Client) Repo {
	collection := client.Database("recipe-hub").Collection("recipes")

	return Repo{
		collection: collection,
	}
}

func (r Repo) FindAll() (*[]domain.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, e := r.collection.Find(ctx, bson.D{})
	if e != nil {

		return nil, errors.New("server error")
	}

	recipes := make([]domain.Recipe, 0)
	for cur.Next(ctx) {
		var recipe domain.Recipe

		e := cur.Decode(&recipe)

		if e != nil {
			return nil, errors.New("server error")

		}
		recipes = append(recipes, recipe)
	}

	if cur.Err() != nil {
		return nil, errors.New("server error")
	}
	cur.Close(ctx)

	return &recipes, nil
}

func (r Repo) FindOne(id string) (*domain.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var recipe domain.Recipe
	e := r.collection.FindOne(ctx, filter).Decode(&recipe)
	if e != nil {
		if e == mongo.ErrNoDocuments {
			return nil, errors.New("recipe Not Found")
		}
		return nil, errors.New("server Error")
	}

	return &recipe, nil
}

func (r Repo) UpdateOne(id string, Title string, Ingredients string, Instructions []string, ImageURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title":        Title,
			"ingredients":  Ingredients,
			"instructions": Instructions,
			"imageurl":     ImageURL,
		},
	}
	fmt.Println(update)
	_, e := r.collection.UpdateOne(ctx, filter, update)

	if e != nil {
		return e
	}
	return nil

}

func (r Repo) Save(newRecipe domain.Recipe) (*domain.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id := primitive.NewObjectID()

	newRecipe.ID = id.Hex()

	_, e := r.collection.InsertOne(ctx, newRecipe)

	if e != nil {

		return nil, errors.New("server error")
	}

	return &newRecipe, nil
}

func (r Repo) DeleteOne(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	_, e := r.collection.DeleteOne(ctx, filter)

	if e != nil {
		return e
	}
	return nil
}
