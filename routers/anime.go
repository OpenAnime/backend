package routers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"constani.me/constani"
	"constani.me/reference"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var animeCol *mongo.Collection = reference.GetCollection(reference.DB, "animes")

func GetAllAnime(c *fiber.Ctx) error {
	return c.Status(http.StatusAccepted).JSON(GetAllData())
}

func GetAnime(c *fiber.Ctx) error {
	data, err := GetAnimeById(c.Params("Id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Anime Not Found", "status": http.StatusBadGateway})
	}
	return c.Status(http.StatusAccepted).JSON(fiber.Map{"message": "Get Anime By id", "status": http.StatusAccepted, "data": data})
}

func CreateAnime(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var anime constani.Anime
	defer cancel()

	if err := c.BodyParser(&anime); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Status": http.StatusBadRequest, "Message": "error", "data": err.Error()})
	}
	newAnime := constani.Anime{
		Id:       anime.Id,
		Name:     anime.Name,
		Info:     anime.Info,
		Pictures: anime.Pictures,
		Episodes: anime.Episodes,
	}
	result, err := animeCol.InsertOne(ctx, newAnime)
	if err != nil {
		fmt.Println("Hata var", err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"Status": http.StatusCreated, "Message": "Created", "data": result})
}

func GetHomeAnime(c *fiber.Ctx) error {
	data, err := GetAnimeHomePage(c.Params("Id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Anime Not Found", "status": http.StatusBadGateway})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"Status": http.StatusCreated, "Message": "Created", "data": data})
}

func GetAllHomeAnime(c *fiber.Ctx) error {
	return c.Status(http.StatusAccepted).JSON(GetAllDataHome())

}
