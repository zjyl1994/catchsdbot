package startup

import (
	"errors"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/infra/vars"
	"github.com/zjyl1994/catchsdbot/server"
)

func Startup() (err error) {
	vars.DebugMode, _ = strconv.ParseBool(os.Getenv("CATCHSD_DEBUG"))
	if vars.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}

	vars.ListenAddr = os.Getenv("CATCHSD_LISTEN")
	if vars.BotToken == "" {
		return errors.New("listen addr empty")
	}

	vars.BotToken = os.Getenv("CATCHSD_BOTTOKEN")
	if vars.BotToken == "" {
		return errors.New("telegram token not found")
	}
	vars.BotInstance, err = tgbotapi.NewBotAPI(vars.BotToken)
	if err != nil {
		return err
	}
	vars.BotInstance.Debug = vars.DebugMode
	server.StartBot()
	return server.Run()
}
