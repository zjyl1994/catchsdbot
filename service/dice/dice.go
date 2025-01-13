package dice

import "math/rand/v2"

var (
	DiceMessage = []string{"深渊迷航", "一波三折", "风平浪静", "一帆风顺", "好运相随"}
	DiceBuff    = []float64{-.3, -.2, -.15, -.05, 0, .05, .15, .2, .3}
)

const DiceFace = 100

func Roll() int {
	return rand.IntN(DiceFace) + 1
}

func GetByDiceResult[T any](result int, arr []T) T {
	percent := float64(result-1) / float64(DiceFace)
	return arr[int(percent*float64(len(arr)))]
}

func GetDiceBuff(result int) float64 {
	return GetByDiceResult(result, DiceBuff)
}

func GetDiceMessage(result int) string {
	return GetByDiceResult(result, DiceMessage)
}
