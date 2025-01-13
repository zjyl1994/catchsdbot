package debug

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/zjyl1994/catchsdbot/infra/vars"
	"github.com/zjyl1994/catchsdbot/service/stamina"
)

func Interceptor() func(c *fiber.Ctx) error {
	conf := basicauth.ConfigDefault
	conf.Authorizer = func(user, pass string) bool {
		return vars.AdminUser != "" && vars.AdminPass != "" &&
			vars.AdminUser == user && vars.AdminPass == pass
	}
	conf.Next = func(c *fiber.Ctx) bool {
		return vars.DebugMode
	}
	return basicauth.New(conf)
}

func AddSp(c *fiber.Ctx) error {
	userId := c.QueryInt("id", 1)
	addSp := c.QueryInt("sp", 1000)
	err := stamina.AddStaminPoint(int64(userId), int64(addSp))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString("OK")
}
