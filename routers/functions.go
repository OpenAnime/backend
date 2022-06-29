package routers

import (
	"context"
	"fmt"
	"log"
	"time"

	"constani.me/constani"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAnimeById(id string) (constani.Anime, error) {
	var result constani.Anime
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	err := animeCol.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return constani.Anime{}, err
	}
	return result, nil
}

func GetAllData() []constani.Anime {
	var results []constani.Anime

	findOptions := options.Find()
	findOptions.SetLimit(100)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cur, err := animeCol.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		fmt.Println("Hata var ", err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result constani.Anime
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	return results
}

func GetAnimeHomePage(id string) (*fiber.Map, error) {
	var result constani.Anime
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	err := animeCol.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return &fiber.Map{}, err
	}
	return &fiber.Map{"id": result.Id, "name": result.Name, "info": result.Info, "avatar": result.Pictures.Avatar, "banner": result.Pictures.Banner}, nil
}

func GetAllDataHome() []constani.AnimeHome {
	var results []constani.AnimeHome

	findOptions := options.Find()
	findOptions.SetLimit(100)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cur, err := animeCol.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		fmt.Println("Hata var ", err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result constani.AnimeHome
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	return results
}

func IsUserFound(email string) (bool, constani.User) {
	var user constani.User
	filter := bson.D{primitive.E{Key: "email", Value: email}}
	err := userCol.FindOne(context.TODO(), filter).Decode(&user)
	fmt.Println("sa")
	if err != nil {
		return false, constani.User{}
	}
	return true, user
}
