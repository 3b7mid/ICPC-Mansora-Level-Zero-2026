package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var x, y, z int
		fmt.Scan(&x, &y, &z)
		if y == 0 || z == 0 || y%z == 1 || y%z == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(x / (y % z))
		}
	}
}
