package main

import (
	"errors"
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x >= 0 {

		val := 1.0
		for i := 0; i < 10; i++ {
			val = val - ((val*val - x) / (2 * val))
		}
		return val, nil
	} else {
		return 0, errors.New("Bro will not go")
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
