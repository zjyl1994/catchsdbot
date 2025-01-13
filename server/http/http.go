package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/server/http/debug"
)

func Run(listen string) error {
	app := fiber.New()

	debugG := app.Group("/debug", debug.Interceptor())
	debugG.Get("/addsp", debug.AddSp)

	logrus.Infoln("Http server started.")
	return app.Listen(listen)
}
