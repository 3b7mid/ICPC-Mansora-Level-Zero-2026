package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	frq := make([]int, 1e5)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if 1 <= a[i] && a[i] <= m {
			frq[a[i]]++
		}
	}
	for i := 1; i <= m; i++ {
		fmt.Println(frq[i])
	}
}