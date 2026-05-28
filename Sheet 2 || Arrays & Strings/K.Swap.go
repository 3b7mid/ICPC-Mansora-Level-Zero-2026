package main

import "fmt"

func main() {
	var n int 
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	min_idx, max_idx := 0, 0
	for i, v := range a {
		if v < a[min_idx] {
			min_idx = i
		}
		if v > a[max_idx] {
			max_idx = i
		}
	}
	a[min_idx], a[max_idx] = a[max_idx], a[min_idx]
	for _, v := range a {
		fmt.Print(v, " ")
	}
}
