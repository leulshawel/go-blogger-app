package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "github.com/leulshawel/go-blogger-app/model"
)

func New(client *mongo.Client) (*fiber.App, error) {
	app := fiber.New()

	app.Get("/api", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Blog server is up and running	")
	})

	app.Post("/api/post", func(ctx *fiber.Ctx) error {
		return nil
	})

	app.Get("/api/post/:id", func(ctx *fiber.Ctx) error {
		//parse parameters
		params := struct {
			Id string `params:"id"`
		}{}
		ctx.ParamsParser(&params)

		objId, err := primitive.ObjectIDFromHex(params.Id)

		if err != nil {
			ctx.Status(233)
			ctx.JSON(struct {
				status string
				code   uint8
			}{status: "", code: 1})
			return err
		}

		fmt.Println(objId)

		return nil
	})

	return app, nil
}
