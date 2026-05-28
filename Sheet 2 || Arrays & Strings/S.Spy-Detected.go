package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		for i := 1; i < n; i++ {
			if a[i] != a[i-1] {
				if i == n-1 {
					fmt.Println(i + 1)
				} else if a[i] != a[i+1] {
					fmt.Println(i + 1)
				} else {
					fmt.Println(i)
				}
				break
			}
		}
		t--
	}
}