package gamedata

var Data = GameData{
	Catch: []Catch{
		{
			Name:        "SilentDepth",
			ShortHand:   "sd",
			Emoji:       "🐧",
			FoodRequire: 1,
		},
	},
	Farm: []Farm{
		{
			Name: "Fish farm",
			Foods: []FarmFood{
				{
					Food: Food{
						Name:      "Fish",
						FoodValue: 1,
					},
					FarmingSpeed: 1,
				},
			},
		},
	},
}
