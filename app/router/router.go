package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	return r
}

func (r *Routes) Run(addr string) error {
	return r.HttpServer.Listen(addr)
}
