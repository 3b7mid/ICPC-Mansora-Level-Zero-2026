package main

import "fmt"

func main() {
	var a, b, c int 
	fmt.Scan(&a, &b, &c)
	if a + b > c && a + c > b && b + c > a {
		fmt.Println(0)
	} else {
		sum := a + b + c 
		mx := max(a, b, c)
		sum -= mx
		fmt.Println(mx - sum + 1)
	}
}
