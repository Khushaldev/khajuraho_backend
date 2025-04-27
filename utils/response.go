package utils

import (
	"khajuraho/backend/config"
	"khajuraho/backend/dto"

	"github.com/gofiber/fiber/v2"
)

const (
	SuccessMessage         = "Success"
	ErrorMessage           = "Error"
	TooManyRequestsMessage = "Too many requests"
	GoogleAuthErrorMessage = "Google authentication failed."
	RateLimitErrorMessage  = "Rate limit exceeded. Please wait a moment and try again."
	ServerSideErrorMessage = "Something went wrong on the server side. Please try after sometime."
)

func getDevStackTrace(err error) string {
	if config.AppConfig.Env == "development" && err != nil {
		return err.Error()
	}
	return ""
}

func OK(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(
		dto.SendSuccess{
			Success: true,
			Message: message,
			Data:    data,
		},
	)
}

func Created(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusCreated).JSON(
		dto.SendSuccess{
			Success: true,
			Message: message,
			Data:    data,
		},
	)
}

func BadRequest(c *fiber.Ctx, message string, errs []string, internalError error) error {
	return c.Status(fiber.StatusBadRequest).JSON(
		dto.SendError{
			Success:    false,
			Message:    message,
			Errors:     errs,
			StackTrace: getDevStackTrace(internalError),
		},
	)
}

func Unauthorized(c *fiber.Ctx, message string, errs []string, internalError error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(
		dto.SendError{
			Success:    false,
			Message:    message,
			Errors:     errs,
			StackTrace: getDevStackTrace(internalError),
		},
	)
}

func Forbidden(c *fiber.Ctx, message string, errs []string, internalError error) error {
	return c.Status(fiber.StatusForbidden).JSON(
		dto.SendError{
			Success:    false,
			Message:    message,
			Errors:     errs,
			StackTrace: getDevStackTrace(internalError),
		},
	)
}

func NotFound(c *fiber.Ctx, message string, errs []string, internalError error) error {
	return c.Status(fiber.StatusNotFound).JSON(
		dto.SendError{
			Success:    false,
			Message:    message,
			Errors:     errs,
			StackTrace: getDevStackTrace(internalError),
		},
	)
}

func TooManyRequests(c *fiber.Ctx, message string, errs []string, internalError error) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(
		dto.SendError{
			Success:    false,
			Message:    message,
			Errors:     errs,
			StackTrace: getDevStackTrace(internalError),
		},
	)
}

func ServerError(c *fiber.Ctx, message string, errs []string, internalError error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(
		dto.SendError{
			Success:    false,
			Message:    message,
			Errors:     errs,
			StackTrace: getDevStackTrace(internalError),
		},
	)
}
