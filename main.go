package main

import (
	"fmt"
	"time"
)

type ErrNegativeSqrt struct {
	when time.Time
	what string
}

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error: %v at %v", err.what, err.when)
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {

		val := 1.0
		for i := 0; i < 10; i++ {
			val = val - ((val*val - x) / (2 * val))
		}
		return val, nil
	} else {
		return 0, ErrNegativeSqrt{when: time.Now(),
			what: "Not gonna work"}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
