package cargo

import (
	"errors"

	"github.com/zjyl1994/catchsdbot/infra/vars"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetCargo(userId int64) (map[int32]int64, error) {
	var m []CargoItem
	err := vars.Database.Where(CargoItem{UserId: userId}).Find(&m).Error
	if err != nil {
		return nil, err
	}
	result := make(map[int32]int64)
	for _, item := range m {
		result[item.ItemId] = item.Amount
	}
	return result, nil
}

func GetCargoItem(db *gorm.DB, userId int64, itemId int32) (int64, error) {
	var m CargoItem
	err := db.Where(CargoItem{UserId: userId, ItemId: itemId}).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 不存在的库存项视为 0 不报错
			return 0, nil
		}
		return 0, err
	}
	return m.Amount, nil
}

func SetCargoItem(db *gorm.DB, userId int64, itemId int32, amount int64) error {
	m := CargoItem{
		UserId: userId,
		ItemId: itemId,
		Amount: amount,
	}
	return db.Clauses(clause.OnConflict{ // 当唯一key冲突时，更新数量
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "item_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"amount"}),
	}).Create(&m).Error
}
