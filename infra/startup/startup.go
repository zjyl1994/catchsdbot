package startup

import (
	"errors"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/infra/gamedata"
	"github.com/zjyl1994/catchsdbot/infra/vars"
	"github.com/zjyl1994/catchsdbot/server"
)

func Startup() (err error) {
	// 加载游戏基本数据
	gd, err := gamedata.Load()
	if err != nil {
		return err
	}
	vars.GameData = gd
	// 初始化环境变量内容
	vars.DebugMode, _ = strconv.ParseBool(os.Getenv("CATCHSD_DEBUG"))
	if vars.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
	vars.DataDir = os.Getenv("CATCHSD_DATADIR")

	vars.ListenAddr = os.Getenv("CATCHSD_LISTEN")
	if vars.ListenAddr == "" {
		return errors.New("listen addr empty")
	}
	vars.BotToken = os.Getenv("CATCHSD_BOTTOKEN")
	if vars.BotToken == "" {
		return errors.New("telegram token not found")
	}
	// 启动bot实例
	vars.BotInstance, err = tgbotapi.NewBotAPI(vars.BotToken)
	if err != nil {
		return err
	}
	vars.BotInstance.Debug = vars.DebugMode
	server.StartBot()
	// 启动http服务
	return server.Run(vars.ListenAddr)
}
