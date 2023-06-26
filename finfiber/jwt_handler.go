package finfiber

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

type JwtHandler interface {
	Method() string
	SignInKey() interface{}
	GetErrorHandler() fiber.ErrorHandler
	GetSuccessHandler() fiber.Handler
	GenerateToken(claims jwt.Claims) (string, error)
}

func GetFiberJwtHandler(handler JwtHandler) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     handler.SignInKey(),
		SigningMethod:  handler.Method(),
		ErrorHandler:   handler.GetErrorHandler(),
		SuccessHandler: handler.GetSuccessHandler(),
	})
}
