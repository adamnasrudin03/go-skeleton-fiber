package middlewares

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

// BasicAuth function basic auth
func BasicAuth(basicUsername, basicPassword string) func(c *fiber.Ctx) error {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			basicUsername: basicPassword,
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == basicUsername && pass == basicPassword {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			err := response_mapper.NewError(response_mapper.ErrUnauthorized, response_mapper.NewResponseMultiLang(
				response_mapper.MultiLanguages{
					ID: "Token tidak valid",
					EN: "Invalid token",
				},
			))
			return c.Status(http.StatusUnauthorized).JSON(response_mapper.RenderStruct(http.StatusUnauthorized, err))
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	})
}
