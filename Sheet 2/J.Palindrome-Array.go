package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n/2; i++ {
		if a[i] != a[n - i - 1] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
