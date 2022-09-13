package main

import (
	"fmt"
	"github.com/finfree-backend/finfree-fiber/finfiber"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestDefaultJwt(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: finfiber.DefaultErrorHandler,
	})
	publicKey := os.Getenv("ES256_KEY")
	secretKey := os.Getenv("HS256_KEY")
	hs256Mw := finfiber.NewDefaultHS256([]byte(secretKey))
	es256Mw := finfiber.NewDefaultES256([]byte(publicKey))
	router := app.Use("", finfiber.GetFiberJwtHandler(hs256Mw), finfiber.GetFiberJwtHandler(es256Mw))

	router.Get("", func(ctx *fiber.Ctx) error {

		for _, key := range finfiber.DEFAULT_JWT_KEYS {
			fmt.Println(key, ":", ctx.Locals(key))

		}
		m := map[string]string{}

		resp := finfiber.NewSuccessResponseWithNextUrl(m, ctx.Request().URI(), 24)
		ctx.JSON(resp)

		return nil
	})
	app.Listen(":8080")

}

func TestDefaultHS256(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: finfiber.DefaultErrorHandler,
	})
	secretKey := os.Getenv("HS256_KEY")
	hs256Mw := finfiber.NewDefaultHS256([]byte(secretKey))
	router := app.Use("", finfiber.GetFiberJwtHandler(hs256Mw))
	router.Get("", func(ctx *fiber.Ctx) error {

		for _, key := range finfiber.DEFAULT_JWT_KEYS {
			fmt.Println(key, ":", ctx.Locals(key))

		}
		m := map[string]string{}

		resp := finfiber.NewSuccessResponseWithNextUrl(m, ctx.Request().URI(), 24)
		ctx.JSON(resp)

		return nil
	})
	app.Listen(":8080")
}

func TestDefaultES256(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: finfiber.DefaultErrorHandler,
	})

	// Get & validate finfree public key file
	// 'FINFREE_PUBLIC_KEY_FILE' os variable required
	es256Key, err := ioutil.ReadFile(os.Getenv("FINFREE_PUBLIC_KEY_FILE"))
	if err != nil {
		log.Println("Error reading finfree public key file")
		t.FailNow()
	}

	es256Mw := finfiber.NewDefaultES256(es256Key)

	router := app.Use("", finfiber.GetFiberJwtHandler(es256Mw))
	router.Get("", func(ctx *fiber.Ctx) error {

		for _, key := range finfiber.DEFAULT_JWT_KEYS {
			fmt.Println(key, ":", ctx.Locals(key))

		}
		m := map[string]string{}

		resp := finfiber.NewSuccessResponseWithNextUrl(m, ctx.Request().URI(), 24)
		ctx.JSON(resp)

		return nil
	})
	app.Listen(":8080")
}
