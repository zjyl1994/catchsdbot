package cargo

const ( // 物品词典
	ITEM_AIR = iota
	ITEM_PENGUIN
	ITEM_FISH
)

var (
	ItemName = map[int32]string{
		ITEM_PENGUIN: "企鹅",
		ITEM_FISH:    "鱼",
	}
)
