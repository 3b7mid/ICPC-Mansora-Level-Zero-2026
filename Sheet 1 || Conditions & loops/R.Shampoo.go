package main

import "fmt"

func main() {
	var v, a, b, c int32
	fmt.Scan(&v, &a, &b, &c)
	for {
		if v < a {
			fmt.Println("F")
			break
		}
		v -= a
		if v < b {
			fmt.Println("M")
			break
		}
		v -= b
		if v < c {
			fmt.Println("T")
			break
		}
		v -= c
	}
}
