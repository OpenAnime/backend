package routers

import (
	fiber "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func authRequired() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
			return nil
		},
		SigningKey: []byte("secret"),
	})
}
func Handle(app *fiber.App) {
	api := app.Group("api")
	anime := api.Group("anime")
	playlist := api.Group("playlist")
	user := api.Group("user")

	// User
	user.Post("/signup", SignUp)
	user.Post("/login", Login)
	user.Get("/me", Me)

	anime.Get("/get/:Id", GetAnime)
	anime.Get("/get/:Id/home", GetHomeAnime)
	anime.Get("/", GetAllAnime)
	anime.Get("/home", GetAllHomeAnime)
	anime.Post("/create", CreateAnime)

	playlist.Get("/", GetAllPlayList)
	playlist.Post("/create", authRequired(), createPlayList)
}
