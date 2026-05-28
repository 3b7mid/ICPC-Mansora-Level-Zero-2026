package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	sum := a + b + c
	mx := max(a, b, c)
	mn := min(a, b, c)
	se_mx := sum - (mx + mn)
	fmt.Println(mx + se_mx)
}
