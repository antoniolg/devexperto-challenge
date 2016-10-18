package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	n := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %v", n)
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	} else {
		z := 1.0
		for i := 1.0; i < x; i++ {
			z = z - (z * z - x) / (2 * z)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
