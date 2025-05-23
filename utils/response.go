package utils

import (
	"khajuraho/backend/config"

	"github.com/gofiber/fiber/v2"
)

const (
	SuccessMessage          = "Success"
	InternalServerError     = "Internal server error"
	ResourceCreatedMessage  = "Resource created"
	BadRequestMessage       = "Bad Request"
	UnauthorizedMessage     = "Unauthorized"
	ForbiddenMessage        = "Forbidden"
	NotFoundMessage         = "Resource not found"
	ValidationFailedMessage = "Validation failed"
	ConflictMessage         = "Conflict"
	TooManyRequestsMessage  = "Too many requests"

	PermissionDeniedCode = "PERMISSION_DENIED"
	AccessDeniedCode     = "ACCESS_DENIED"
	TooManyRequestsCode  = "TOO_MANY_REQUESTS"
)

type ErrorType string

const (
	ValidationError ErrorType = "VALIDATION_ERROR"
	GeneralError    ErrorType = "GENERAL_ERROR"
	InternalError   ErrorType = "INTERNAL_ERROR"
)

type ErrorDetail struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}

type ErrorCode struct {
	Code string `json:"code"`
}

type ErrorPayload struct {
	Type    ErrorType     `json:"type"`
	Code    string        `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Errors  []ErrorDetail `json:"errors,omitempty"`
	Stack   string        `json:"stack,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    any           `json:"data,omitempty"`
	Error   *ErrorPayload `json:"error,omitempty"`
}

func sendResponse(c *fiber.Ctx, status int, message string, payload any) error {
	success := status >= 200 && status < 300

	res := Response{
		Status:  status,
		Success: success,
		Message: message,
	}

	if success && payload != nil {
		res.Data = payload
	} else if !success && payload != nil {
		switch e := payload.(type) {
		case []ErrorDetail:
			res.Error = &ErrorPayload{
				Type:   ValidationError,
				Errors: e,
			}
		case error:
			stack := ""
			if config.AppConfig.Env != "production" {
				stack = e.Error()
			}
			res.Error = &ErrorPayload{
				Type:    InternalError,
				Message: "Something went wrong on the server side.",
				Stack:   stack,
			}
		case ErrorCode:
			res.Error = &ErrorPayload{
				Type: GeneralError,
				Code: e.Code,
			}
		default:
			res.Error = &ErrorPayload{
				Type: GeneralError,
				Code: "Unknown_ERROR",
			}
		}
	}

	return c.Status(status).JSON(res)
}

func OK(c *fiber.Ctx, message string, data any) error {
	return sendResponse(c, fiber.StatusOK, message, data)
}

func Created(c *fiber.Ctx, message string, data any) error {
	return sendResponse(c, fiber.StatusCreated, message, data)
}

func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

func BadRequest(c *fiber.Ctx, message string, err any) error {
	return sendResponse(c, fiber.StatusBadRequest, message, err)
}

func Unauthorized(c *fiber.Ctx, message string, err any) error {
	return sendResponse(c, fiber.StatusUnauthorized, message, err)
}

func Forbidden(c *fiber.Ctx, message string, err any) error {
	return sendResponse(c, fiber.StatusForbidden, message, err)
}

func NotFound(c *fiber.Ctx, message string, err any) error {
	return sendResponse(c, fiber.StatusNotFound, message, err)
}

func Conflict(c *fiber.Ctx, message string, err any) error {
	return sendResponse(c, fiber.StatusConflict, message, err)
}

func Unprocessable(c *fiber.Ctx, message string, errs []ErrorDetail) error {
	return sendResponse(c, fiber.StatusUnprocessableEntity, message, errs)
}

func ServerError(c *fiber.Ctx, message string, err error) error {
	return sendResponse(c, fiber.StatusInternalServerError, message, err)
}

func TooManyRequests(c *fiber.Ctx, message string, err ErrorCode) error {
	return sendResponse(c, fiber.StatusTooManyRequests, message, err)
}
