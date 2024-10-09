package main

import "fmt"

func main() {
	var phi, v, l, r float64
	fmt.Scan(&r)
	phi = 3.1415926535
	v = 4 * r * r * r * phi / 3
	l = phi * (r * r) * 4
	fmt.Print("Bola dengan jejari ", r, " memiliki volume ", v, " dan luas kulit ", l)
}
