package user

import (
	"errors"
	"time"

	"github.com/zjyl1994/catchsdbot/infra/vars"
	"gorm.io/gorm"
)

func GetOrCreateByTgUser(tgUserId int64, tgUserName string) (*User, error) {
	user := &User{}
	if err := vars.Database.Where("tg_user_id = ?", tgUserId).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.TgUserId = tgUserId
			user.TgUserName = tgUserName
			user.LastTick = time.Now().Unix()
			if err := vars.Database.Create(user).Error; err != nil {
				return nil, err
			}
			return user, nil
		}
		return nil, err
	}
	if user.TgUserName != tgUserName {
		user.TgUserName = tgUserName
		if err := vars.Database.Save(user).Error; err != nil {
			return nil, err
		}
	}
	return user, nil
}
