package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	fmt.Println("\n")
	fmt.Println(math.Sincos(n))
	fmt.Println("\n")
	// fmt.Println(math.Sinh(n))
}
