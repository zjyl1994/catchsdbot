package attr

import (
	"gorm.io/gorm"
)

func GetUserAttr(db *gorm.DB, userId int64) ([]Attr, error) {
	var attrs []Attr
	if err := db.Where("user_id = ?", userId).Find(&attrs).Error; err != nil {
		return nil, err
	}
	return attrs, nil
}

