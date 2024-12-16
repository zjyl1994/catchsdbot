package startup

import (
	"errors"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/server"
	"github.com/zjyl1994/catchsdbot/vars"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Startup() (err error) {
	vars.DebugMode, _ = strconv.ParseBool(os.Getenv("CATCHSD_DEBUG"))
	if vars.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
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
	
	return server.Run()
}
