package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zjyl1994/catchsdbot/infra/vars"
)

func Run() error {
	app := fiber.New()
	return app.Listen(vars.ListenAddr)
}
