
package main

import (
	"fmt"
)

func main() {
	var beDivNum int = 20
	var divNum float64 = 0.002

	fmt.Println(dividing(beDivNum, divNum))
}

func dividing(beDivNum int, divNum float64) float64 {
	var answer float64 = float64(beDivNum)
	return answer / divNum
}

