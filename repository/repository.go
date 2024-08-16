package repository

import (
	"context"
	"errors"
	"example/htmx/domain"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

func init() {
	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	var client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	userCollection = client.Database("recipe-hub").Collection("recipes")
}

func FindAll() (*[]domain.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, e := userCollection.Find(ctx, bson.D{})
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

func FindOne(id string) (*domain.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var recipe domain.Recipe
	e := userCollection.FindOne(ctx, filter).Decode(&recipe)
	if e != nil {
		if e == mongo.ErrNoDocuments {
			return nil, errors.New("recipe Not Found")
		}
		return nil, errors.New("server Error")
	}

	return &recipe, nil
}

func UpdateOne(id string, Title string, Ingredients string, Instructions []string, ImageURL string) error {
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
	_, e := userCollection.UpdateOne(ctx, filter, update)

	if e != nil {
		return e
	}
	return nil

}

func Save(newRecipe domain.Recipe) (*domain.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id := primitive.NewObjectID()

	newRecipe.ID = id.Hex()

	_, e := userCollection.InsertOne(ctx, newRecipe)

	if e != nil {

		return nil, errors.New("server error")
	}

	return &newRecipe, nil
}

func DeleteOne(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	_, e := userCollection.DeleteOne(ctx, filter)

	if e != nil {
		return e
	}
	return nil
}
