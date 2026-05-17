package main

import "fmt"

func main() {
	var num string
	fmt.Scan(&num)
	var sum int64 = 0
	for i := range num {
		dig := num[i] - '0'
		sum += int64(dig)
	}
	if sum%3 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
