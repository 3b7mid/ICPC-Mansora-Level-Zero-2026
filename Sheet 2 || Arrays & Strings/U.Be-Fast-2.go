package main

import "fmt"

func main() {
	var l, r, n int32
	fmt.Scan(&l, &r, &n)
	numbers := r - l + 1
	numOfDivisible := r/n - (l-1)/n
	fmt.Println(numbers - numOfDivisible)
}
