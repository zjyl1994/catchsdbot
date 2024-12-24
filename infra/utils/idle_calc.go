package utils

func IdleCalcWithMax(lastTick, currentTick, tickGenNum, currentNum, maxNum int64) int64 {
	calcauteNum := IdleCalcNoLimit(lastTick, currentTick, tickGenNum, currentNum)
	if maxNum <= 0 { // 无限制，返回计算结果
		return calcauteNum
	} else {
		if currentNum >= maxNum { // 当前超过自然恢复上限，可能是额外奖励，直接返回当前值
			return currentNum
		} else {
			if calcauteNum > maxNum { // 算出来的超出上限，受上限管控
				return maxNum
			} else {
				return calcauteNum
			}
		}
	}
}

func IdleCalcNoLimit(lastTick, currentTick, tickGenNum, currentNum int64) int64 {
	tickElapsed := currentTick - lastTick
	if tickElapsed < 0 { // 时间倒退无法计算
		return currentNum
	}
	// 计算当前tick区间生成的数量
	generated := tickElapsed * tickGenNum
	return currentNum + generated
}
