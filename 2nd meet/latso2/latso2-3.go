package main

import "fmt"

func main() {
	var x int
	var y bool
	fmt.Print("Tahun: ")
	fmt.Scan(&x)
	y = x%4 == 0
	fmt.Print("Kabisat: ", y)
}
