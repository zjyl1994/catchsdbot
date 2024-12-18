package gamedata

type GameData struct {
	Catch []Catch
	Farm  []Farm
}

type Catch struct {
	Name        string
	ShortHand   string
	Emoji       string
	FoodRequire float64
}

type Farm struct {
	Name  string
	Foods []FarmFood
}

type Food struct {
	Name      string
	FoodValue float64
}

type FarmFood struct {
	Food
	FarmingSpeed float64
}
