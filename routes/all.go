package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/leulshawel/go-blogger-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(client *mongo.Client) (*fiber.App, error) {
	app := fiber.New()

	app.Get("/api", func(ctx *fiber.Ctx) error {
		ctx.SendString("Blog server is up and running with air")
		return nil
	})

	app.Post("/api/post", func(ctx *fiber.Ctx) error {
		var post = new(models.Post)
		ctx.BodyParser(post)
		fmt.Println(post)
		//actions.Actions_.OnDelete(post)
		return nil
	})

	app.Get("/api/post/:id", func(ctx *fiber.Ctx) error {
		fmt.Println("Get request comming to /api/post/:id")
		//parse parameters
		id := ctx.Params("id")
		ctx.SendString(id)
		objId, err := primitive.ObjectIDFromHex(id)

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

	app.Delete("/api/post/:id", func(cts *fiber.Ctx) error {
		return nil
	})

	return app, nil
}
