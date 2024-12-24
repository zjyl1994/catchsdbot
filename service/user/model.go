package user

type User struct {
	ID         int64  `gorm:"primaryKey"`
	TgUserId   int64  `gorm:"unique,column:tg_user_id"`
	TgUserName string `gorm:"column:tg_user_name"`
}
