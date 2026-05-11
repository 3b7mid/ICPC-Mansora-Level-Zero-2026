package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("MEDIA = %.1f\n", ((a*2.00)+(b*3.00)+(c*5.00))/10.00)
}
