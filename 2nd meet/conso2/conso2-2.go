package main

import "fmt"

func main() {
	var alas, tinggi, luas float64
	fmt.Print("masukan alas: ")
	fmt.Scan(&alas)
	fmt.Print("masukan tinggi: ")
	fmt.Scan(&tinggi)
	luas = alas * tinggi / 2
	fmt.Print(luas)
}
