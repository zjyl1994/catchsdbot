package cargo

type CargoItem struct {
	ID     int64 `gorm:"primaryKey"`
	UserId int64 `gorm:"unique,column:user_id"`
	ItemId int32 `gorm:"unique,column:item_id"`
	Amount int64
}
