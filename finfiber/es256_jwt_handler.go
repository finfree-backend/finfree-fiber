package finfiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type ES256Middleware struct {
	Key            []byte
	ErrorHandler   fiber.ErrorHandler
	SuccessHandler fiber.Handler
}

func NewES256Middleware(key []byte, errorHandler fiber.ErrorHandler, successHandler fiber.Handler) *ES256Middleware {
	return &ES256Middleware{Key: key, ErrorHandler: errorHandler, SuccessHandler: successHandler}
}

func NewDefaultES256(key []byte, payloadKeys ...string) *ES256Middleware {
	var errorHandler fiber.ErrorHandler = func(ctx *fiber.Ctx, err error) error {

		if ctx.Locals(IS_AUTHORIZED_LOCAL_KEY).(bool) {
			return ctx.Next()
		}
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
	return &ES256Middleware{Key: key, ErrorHandler: errorHandler, SuccessHandler: successHandler}
}

func (e *ES256Middleware) Method() string {
	return "ES256"
}

func (e *ES256Middleware) SignInKey() interface{} {
	globalPublicKey, err := jwt.ParseECPublicKeyFromPEM(e.Key)
	if err != nil {
		panic("Error on transforming ES256 public key. Err:" + err.Error())
	}
	return globalPublicKey
}

func (e *ES256Middleware) GetErrorHandler() fiber.ErrorHandler {
	return e.ErrorHandler
}

func (e *ES256Middleware) GetSuccessHandler() fiber.Handler {
	return e.SuccessHandler
}
