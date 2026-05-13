package main

import "fmt"

func main()  {
	var n int 
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var min_idx int = 0
	for i, v := range a {
		if v < a[min_idx] {
			min_idx = i
		}
	}
	fmt.Println(a[min_idx], min_idx + 1)
}