package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func(e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negatice number: %d", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)	
	}
	z := float64(1)
	for i := 1; i <= 10; i++ {
		fmt.Println(z)
 		z -= (z * z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(-2))
}
