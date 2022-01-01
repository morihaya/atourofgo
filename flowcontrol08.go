package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)

		//fmt.Println("z=", z)
	}
	return z
}

func main() {
	x := 2.0
	fmt.Println(x, " = ", Sqrt(x))
	x = 4.0
	fmt.Println(x, " = ", Sqrt(x))
	x = 10.0
	fmt.Println(x, " = ", Sqrt(x))
	x = 100
	fmt.Println(x, " = ", Sqrt(x))
}
