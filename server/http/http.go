package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Run(listen string) error {
	app := fiber.New()
	logrus.Infoln("Http server started.")
	return app.Listen(listen)
}
