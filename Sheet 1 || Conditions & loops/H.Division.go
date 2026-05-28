package main

import "fmt"

func main() {
	var t, rate int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&rate)
		if rate <= 1399 {
			fmt.Println("Division 4")
		} else if 1400 <= rate && rate <= 1599 {
			fmt.Println("Division 3")
		} else if 1600 <= rate && rate <= 1899 {
			fmt.Println("Division 2")
		} else {
			fmt.Println("Division 1")
		}
	}
}
