package routes

import (
	"github.com/gofiber/fiber/v2"
	// validator "github.com/asaskevich/govalidator"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	"fmt"
	"errors"
)


func CreateApp() (*fiber.App, error) {
	fmt.Println("App started")
	
	app := fiber.New()

	app.Get("/api", func(ctx *fiber.Ctx ) error {
		return ctx.SendString("Blog server is up and running	")
	})

	app.Post("/api/post", func(ctx *fiber.Ctx) error {
		params := struct {Id string `params:"id"`}{}
		ctx.ParamsParser(&params)

		ctx.JSON(&params)
		return nil
	})

	app.Get("/api/post/:id", func(ctx *fiber.Ctx) error {
		//parse parameters
		params := struct {Id string `params:"id"`}{}
		ctx.ParamsParser(&params)

		if !primitive.IsValidObjectID(params.Id){
			ctx.Status(233)
			ctx.JSON(struct{error string}{})
			return nil
		}

		ctx.JSON(&params)

		return nil
	})


	return app
}