package main

import "fmt"

func main() {
	var c, r, f, k float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&c)
	f = (c * 9 / 5) + 32
	r = c * 4 / 5
	k = (f + 459.67) * 5 / 9
	fmt.Println("Derajat Reamur: ", r)
	fmt.Println("Derajat Fahrenheit: ", f)
	fmt.Println("Derajat Kelvin: ", k)
}
