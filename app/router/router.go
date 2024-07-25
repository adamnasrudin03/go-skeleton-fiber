package router

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Routes struct {
	HttpServer *fiber.App
}

func NewRoutes() *Routes {
	r := &Routes{
		HttpServer: fiber.New(),
	}

	r.HttpServer.Use(logger.New())
	r.HttpServer.Use(cors.New())
	r.HttpServer.Use(recover.New())

	r.HttpServer.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(response_mapper.RenderStruct(http.StatusOK, response_mapper.MultiLanguages{
			ID: "selamat datang di server ini",
			EN: "Welcome this server",
		}))
	})
	r.HttpServer.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(response_mapper.RenderStruct(http.StatusNotFound, response_mapper.ErrRouteNotFound()))
	})

	return r
}

func (r *Routes) Run(addr string) error {
	return r.HttpServer.Listen(addr)
}
