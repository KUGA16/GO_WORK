package temp

// // 大文字関数は外部に公開
func DivideCall(beDivNum int, divNum float64) float64 {
	var number float64 = dividing(beDivNum, divNum)
	return number
}

func dividing(beDivNum int, divNum float64) float64 {
	var answer float64 = float64(beDivNum)
	return answer / divNum
}
