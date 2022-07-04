package main

import (
	"fmt"
	"github.com/finfree-backend/finfree-fiber/finfiber"
	"github.com/gofiber/fiber/v2"
	"os"
	"testing"
)

func TestWithPayloadKeysUsage(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: finfiber.DefaultErrorHandler,
	})
	publicKey := os.Getenv("ES256_KEY")
	secretKey := os.Getenv("HS256_KEY")
	// Custom payload keys can be set to request context
	// To be able to set custom keys, after NewDefaultHS256 or NewDefaultES256 methods, required keys must be given as string
	// Remember that if you give payload keys by yourself, default keys no longer be used by your NewDefaultHS256 and NewDefaultES256 methods' SuccessHandler
	// 'username' and 'locale' parameters given in this example
	hs256Mw := finfiber.NewDefaultHS256([]byte(secretKey), "username", "locale")
	es256Mw := finfiber.NewDefaultES256([]byte(publicKey), "username", "locale")
	router := app.Use("", finfiber.GetFiberJwtHandler(hs256Mw), finfiber.GetFiberJwtHandler(es256Mw))

	router.Get("", func(ctx *fiber.Ctx) error {
		fmt.Println("Username:", ctx.Locals("username"))
		fmt.Println("Locale:", ctx.Locals("locale"))
		m := map[string]string{}

		resp := finfiber.NewSuccessResponseWithNextUrl(m, ctx.Request().URI(), 24)
		ctx.JSON(resp)

		return nil
	})
	app.Listen(":8080")

}
