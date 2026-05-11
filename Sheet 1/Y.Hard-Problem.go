package main

import "fmt"

func main() {
	var a, b, c int64
	fmt.Scan(&a, &b, &c)
	sum := a + b + c 
	mx := max(a, b, c)
	mn := min(a, b, c)
	fmt.Println(mn, sum - mx - mn, mx)
}