package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp", a)

	a[4] = 10
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
}
