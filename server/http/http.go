package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/service/stamina"
)

func Run(listen string) error {
	app := fiber.New()
	app.Get("/debug/addsp", debugAddSp)
	logrus.Infoln("Http server started.")
	return app.Listen(listen)
}

func debugAddSp(c *fiber.Ctx) error {
	userId := c.QueryInt("id", 1)
	addSp := c.QueryInt("sp", 1000)
	err := stamina.AddStaminPoint(int64(userId), int64(addSp))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString("OK")
}
