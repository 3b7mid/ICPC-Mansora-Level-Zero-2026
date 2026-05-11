package main

import "fmt"

func main() {
	const PI float32 = 3.14159
	var r float32
	fmt.Scan(&r)
	fmt.Printf("A=%.4f\n", r*r*PI)
}
