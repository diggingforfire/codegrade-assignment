package routes

import (
	"time"

	"codegrade.com/take-home/db"
	"github.com/gofiber/fiber/v2"
)

func Posts(app *fiber.App) {
	app.Get("/posts/", func(ctx *fiber.Ctx) error {
		// TODO: Only show non expired posts.
		posts, err := db.Client.Post.Query().All(ctx.Context())
		if err != nil {
			return err
		}
		return ctx.JSON(posts)
	})

	app.Post("/post", func(ctx *fiber.Ctx) error {
		payload := struct {
			Content string `json:"content"`
			Title   string `json:"title"`
			// You can change this to be in any type!
			ExpiresIn time.Time `json:"expiresIn"`
		}{}
		if err := ctx.BodyParser(&payload); err != nil {
			return err
		}

		// TODO: Use expires in!
		post, err := db.Client.Post.Create().
			SetContent(payload.Content).
			SetTitle(payload.Title).
			Save(ctx.Context())
		if err != nil {
			return err
		}
		return ctx.JSON(post)
	})
}
