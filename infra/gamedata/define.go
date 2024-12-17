package gamedata

// 生成自 https://zhwt.github.io/yaml-to-go/

type GameData struct {
	Catch Catch `yaml:"catch"`
	Food  Food  `yaml:"food"`
}
type CatchItem struct {
	Name        string `yaml:"name"`
	Emoji       string `yaml:"emoji"`
	ShortHand   string `yaml:"short_hand"`
	CatchDict   string `yaml:"catch_dict"`
	FoodRequire int    `yaml:"food_require"`
}
type Catch struct {
	List []CatchItem `yaml:"list"`
}
type FoodItem struct {
	Name      string `yaml:"name"`
	Emoji     string `yaml:"emoji"`
	FoodValue int    `yaml:"food_value"`
}
type Food struct {
	List []FoodItem `yaml:"list"`
}
