package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n % 10 == 0 {
		fmt.Println("YES")
		return
	}
	first_digit := n / 10
	second_digit := n % 10
	if first_digit%second_digit == 0 || second_digit%first_digit == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
