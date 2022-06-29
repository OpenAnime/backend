package routers

import (
	"context"
	"net/http"
	"time"

	"constani.me/constani"
	"constani.me/reference"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCol *mongo.Collection = reference.GetCollection(reference.DB, "users")

type Claims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func SignUp(c *fiber.Ctx) error {
	var user constani.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Status": http.StatusBadRequest, "Message": "error", "data": err.Error()})
	}

	if err := validate.Struct(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Status": http.StatusBadRequest, "Message": "error", "data": err.Error()})
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	newUser := constani.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  string(password),
		CreatedAt: user.CreatedAt,
		Avatar:    user.Avatar,
		Banner:    user.Banner,
		PlayList:  user.PlayList,
		AnimeList: user.AnimeList,
		IsAdmin:   user.IsAdmin,
		Presence:  user.Presence,
	}
	result, err := userCol.InsertOne(context.Background(), newUser)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Status": http.StatusBadRequest, "Message": "error", "data": err.Error()})
	}

	return c.Status(http.StatusAccepted).JSON(fiber.Map{"Status": http.StatusAccepted, "Message": "success", "data": result})
}

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user", "atlasch1903@gmail.com")
	// Throws Unauthorized error
	err, result := IsUserFound(user)

	if err == false {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Status": http.StatusUnauthorized, "Message": "error", "data": "User not found"})
	}

	claims := Claims{
		Id:    result.Id,
		Email: result.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: t,
	})

	return c.JSON(fiber.Map{"token": t, "user": fiber.Map{"id": result.Id, "name": result.Name, "email": result.Email}})
}

func Me(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Status": http.StatusUnauthorized, "Message": "error", "data": "Couild NOt LOgin"})
	}

	claims := token.Claims.(*Claims)

	var user constani.User

	filter := bson.D{primitive.E{Key: "id", Value: claims.Id}}
	err = userCol.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Status": http.StatusUnauthorized, "Message": "error", "data": " User Not Found"})
	}
	return c.JSON(fiber.Map{"message": "success", "data": user})
}
