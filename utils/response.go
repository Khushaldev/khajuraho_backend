package utils

import (
	"khajuraho/backend/dto"

	"github.com/gofiber/fiber/v2"
)

const (
	SuccessMessage = "Success"
	ErrorMessage   = "Error"

	ServerSideErrorMessage = "Something went wrong on the server side. Please try after sometime."
)

func Success(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(
		dto.APIResponse{
			Success: true,
			Message: message,
			Data:    data,
			Errors:  nil,
		},
	)
}

func Created(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusCreated).JSON(
		dto.APIResponse{
			Success: true,
			Message: message,
			Data:    nil,
			Errors:  nil,
		},
	)
}

func BadRequest(c *fiber.Ctx, message string, errs []string) error {
	return c.Status(fiber.StatusBadRequest).JSON(
		dto.APIResponse{
			Success: false,
			Message: message,
			Data:    nil,
			Errors:  errs,
		},
	)
}

func Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(
		dto.APIResponse{
			Success: false,
			Message: message,
			Data:    nil,
			Errors:  nil,
		},
	)
}

func Forbidden(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(
		dto.APIResponse{
			Success: false,
			Message: message,
			Data:    nil,
			Errors:  nil,
		},
	)
}

func NotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(
		dto.APIResponse{
			Success: false,
			Message: message,
			Data:    nil,
			Errors:  nil,
		},
	)
}

func ServerError(c *fiber.Ctx, message string, errs []string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(
		dto.APIResponse{
			Success: false,
			Message: message,
			Data:    nil,
			Errors:  errs,
		},
	)
}
