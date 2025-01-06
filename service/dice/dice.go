package dice

import "math/rand/v2"

var (
	DiceMessage = []string{"好运相随", "一帆风顺", "风平浪静", "一波三折", "厄运缠身"}
	DiceBuff    = []float64{.5, .2, 0, -.2, -.5}
)

func Roll() int {
	return rand.IntN(6) + 1
}

func getDicePos[T any](result int, arr []T) T {
	percent := float64(result) / float64(6)
	return arr[int(percent*float64(len(arr)))]
}

func GetDiceBuff(result int) float64 {
	return getDicePos(result, DiceBuff)
}

func GetDiceMessage(result int) string {
	return getDicePos(result, DiceMessage)
}
