package main

import "fmt"

func main() {
	var n, a, b, c int
	fmt.Scan(&n) 
	cnt := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b, &c)
		if a + b + c >= 2 {
			cnt++
		}
	}
	fmt.Println(cnt)
}