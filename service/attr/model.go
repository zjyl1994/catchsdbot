package attr

type Attr struct {
	ID        int64 `gorm:"primaryKey"`
	UserId    int64 `gorm:"unique,column:user_id"`
	AttrId    int32 `gorm:"unique,column:attr_id"`
	AttrType  int32 `gorm:"column:attr_type"`
	AttrValue string
}
