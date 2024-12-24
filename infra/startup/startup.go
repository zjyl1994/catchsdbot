package startup

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
	gorm_logrus "github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/infra/vars"
	bot_srv "github.com/zjyl1994/catchsdbot/server/bot"
	http_srv "github.com/zjyl1994/catchsdbot/server/http"
	"github.com/zjyl1994/catchsdbot/service/stamina"
	"github.com/zjyl1994/catchsdbot/service/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Startup() (err error) {
	// 初始化环境变量内容
	vars.DebugMode, _ = strconv.ParseBool(os.Getenv("CATCHSD_DEBUG"))
	if vars.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
	vars.DataDir = os.Getenv("CATCHSD_DATADIR")
	if vars.DataDir != "" {
		err = os.MkdirAll(vars.DataDir, 0755)
		if err != nil {
			return err
		}
	}

	vars.ListenAddr = os.Getenv("CATCHSD_LISTEN")
	if vars.ListenAddr == "" {
		return errors.New("listen addr empty")
	}
	vars.BotToken = os.Getenv("CATCHSD_BOTTOKEN")
	if vars.BotToken == "" {
		return errors.New("telegram token not found")
	}
	// 链接数据库
	databaseFilepath := filepath.Join(vars.DataDir, "catchsd.sqlite")
	vars.Database, err = gorm.Open(sqlite.Open(databaseFilepath), &gorm.Config{
		Logger: gorm_logrus.New(),
	})
	if err != nil {
		return err
	}
	// 初始化数据库 WAL 模式
	sqlDB, err := vars.Database.DB()
	if err != nil {
		return err
	}
	_, err = sqlDB.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		return err
	}
	err = vars.Database.AutoMigrate(&user.User{}, &stamina.Stamina{})
	if err != nil {
		return err
	}
	// 启动bot实例
	vars.BotInstance, err = tgbotapi.NewBotAPI(vars.BotToken)
	if err != nil {
		return err
	}
	vars.BotInstance.Debug = vars.DebugMode
	bot_srv.StartBot()
	// 启动http服务
	return http_srv.Run(vars.ListenAddr)
}
