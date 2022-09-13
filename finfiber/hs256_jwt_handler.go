package finfiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type HS256Middleware struct {
	Key            []byte
	ErrorHandler   fiber.ErrorHandler
	SuccessHandler fiber.Handler
}

func NewHS256Middleware(key []byte, errorHandler fiber.ErrorHandler, successHandler fiber.Handler) *HS256Middleware {
	return &HS256Middleware{Key: key, ErrorHandler: errorHandler, SuccessHandler: successHandler}
}

func NewDefaultHS256(key []byte, payloadKeys ...string) *HS256Middleware {
	var errorHandler fiber.ErrorHandler = func(ctx *fiber.Ctx, err error) error {
		return &fiber.Error{Code: 401, Message: "Invalid/Blank or expired auth token"}
	}

	var successHandler fiber.Handler = func(ctx *fiber.Ctx) error {
		ctx.Locals(IS_AUTHORIZED_LOCAL_KEY, true)
		payload := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
		if payloadKeys == nil {
			// Default JWT keys works
			payloadKeys = DEFAULT_JWT_KEYS
		}
		for _, payloadKey := range payloadKeys {
			ctx.Locals(payloadKey, payload[payloadKey])
		}
		return ctx.Next()
	}
	return &HS256Middleware{Key: key, ErrorHandler: errorHandler, SuccessHandler: successHandler}
}

func (s *HS256Middleware) Method() string {
	return "HS256"
}

func (s *HS256Middleware) SignInKey() interface{} {
	return s.Key
}

func (s *HS256Middleware) GetErrorHandler() fiber.ErrorHandler {
	return s.ErrorHandler
}

func (s *HS256Middleware) GetSuccessHandler() fiber.Handler {
	return s.SuccessHandler
}
