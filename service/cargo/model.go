package cargo

type CargoItem struct {
	ID     int64 `gorm:"primaryKey"`
	UserId int64 `gorm:"uniqueIndex:idx_user_item,column:user_id"`
	ItemId int32 `gorm:"uniqueIndex:idx_user_item,column:item_id"`
	Amount int64
}
