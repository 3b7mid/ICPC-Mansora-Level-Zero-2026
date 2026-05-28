package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	if n >= 13 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
