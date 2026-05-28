package main

import "fmt"

func main() {
	var s string
	var n, x int 
	fmt.Scan(&s, &n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		for j := 0; j < x; j++ {
			fmt.Print(s)
		}
		fmt.Printf("\n")
	}
}