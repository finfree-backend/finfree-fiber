package finfiber

import "github.com/gofiber/fiber/v2"

var DefaultErrorHandler fiber.ErrorHandler = func(ctx *fiber.Ctx, err error) error {
	e, ok := err.(*fiber.Error)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(&fiber.Error{Code: 500, Message: "Unknown Error!"}))
	}

	resp := NewErrorResponse(e)
	if err = ctx.Status(resp.Code).JSON(&resp); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(&fiber.Error{Code: 500, Message: "Unknown Error!"}))
	}

	return nil
}
