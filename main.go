package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (neg ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error: %v is not compatible", float64(neg))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		val := 1.0
		for i := 0; i < 10; i++ {
			val = val - ((val*val - x) / (2 * val))
		}
		return val, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
