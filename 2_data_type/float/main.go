package main

import "fmt"
import "math"

func main() {
	var fval1 float32

	fval1 = 12

	fmt.Println(fval1)

	fval2 := 12.0 // float64

	//fval1 = fval2 // error!
	fmt.Println(fval2)
	fval1 = float32(fval2)
}

func IsEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}
