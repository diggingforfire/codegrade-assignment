package main

import (
	"context"
	"errors"

	"codegrade.com/take-home/db"
	"codegrade.com/take-home/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			ctx.Context().Logger().Printf("Request failed: %v", err)

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// In case the SendFile fails
			return ctx.Status(code).JSON(map[string]string{
				"msg": err.Error(),
			})
		},
	})
	ctx := context.Background()
	db.Connect(ctx)
	routes.Posts(app)

	app.Listen(":9001")
}
