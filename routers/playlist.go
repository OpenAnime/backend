package routers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"constani.me/constani"
	"constani.me/reference"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate = validator.New()
var playList *mongo.Collection = reference.GetCollection(reference.DB, "playlist")

func GetAllPlayList(c *fiber.Ctx) error {
	var results []constani.PlayList

	findOptions := options.Find()
	findOptions.SetLimit(100)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cur, err := playList.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		fmt.Println("Hata var ", err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result constani.PlayList
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	return c.JSON(results)
}

func createPlayList(c *fiber.Ctx) error {
	var playlist constani.PlayList

	if err := c.BodyParser(&playlist); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Status": http.StatusBadRequest, "Message": "error", "data": err.Error()})
	}

	newPlayList := constani.PlayList{
		Id:          playlist.Id,
		Name:        playlist.Name,
		Description: playlist.Description,
		Image:       playlist.Image,
		Items:       playlist.Items,
		Slot:        len(playlist.Items),
		CreatedAt:   playlist.CreatedAt,
	}
	result, err := playList.InsertOne(context.Background(), newPlayList)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Get Anime By id", "status": http.StatusBadRequest, "data": err})
	}
	return c.Status(http.StatusAccepted).JSON(fiber.Map{"message": "Get Anime By id", "status": http.StatusAccepted, "data": result})
}
