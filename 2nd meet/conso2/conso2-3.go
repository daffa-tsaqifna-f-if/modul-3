package main

import "fmt"

func main() {
	var IDR, USD float64
	fmt.Print("masukan nominal Rupiah: ")
	fmt.Scan(&IDR)
	USD = IDR / 15000
	fmt.Print("jadi ", IDR, "rupiah = ", USD, "dolar")
}
