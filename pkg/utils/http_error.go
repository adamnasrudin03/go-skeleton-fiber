package utils

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/gofiber/fiber/v2"
)

func HttpError(c *fiber.Ctx, err error) error {
	var statusCode int
	if e, ok := err.(*response_mapper.ResponseError); ok {
		statusCode = response_mapper.StatusErrorMapping(e.Code)
	}
	if statusCode == 0 || http.StatusText(statusCode) == "" {
		statusCode = http.StatusInternalServerError
	}

	return c.Status(statusCode).JSON(response_mapper.RenderStruct(statusCode, err))
}
