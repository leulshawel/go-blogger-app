package routes

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/leulshawel/go-blogger-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(client *mongo.Client) (*fiber.App, error) {
	app := fiber.New()
	dataValidator := validator.New()
	db := client.Database("posts")
	post_collection := db.Collection("posts")
	//user_collection := db.Collection("user")
	//data_collection := db.Collection("data")
	app.Get("/api", func(ctx *fiber.Ctx) error {
		ctx.SendString("Blog server is up and running with air")
		return nil
	})

	app.Get("/api/post/:id", func(ctx *fiber.Ctx) error {
		fmt.Println("Get request comming to /api/post/:id")
		//parse parameters
		id := ctx.Params("id")
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(Error{Code: 0, Message: "invalid post id"})
			return err
		}

		fmt.Println(objId)
		return nil
	})

	app.Post("/api/post", func(ctx *fiber.Ctx) error {
		body := new(models.Post)
		if err := ctx.BodyParser(body); err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(Error{Code: 1, Message: "invalid blog post data"})
		}

		if err := dataValidator.Struct(body); err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(Error{Code: 2, Message: "required fields are missing"})
		}

		res, error := post_collection.InsertOne(context.TODO(), body)

		if error != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(Error{Code: 3, Message: error.Error()})
		}
		ctx.Status(fiber.StatusOK).JSON(res)
		return nil
	})

	app.Delete("/api/post/:id", func(cts *fiber.Ctx) error {
		return nil
	})

	return app, nil
}
