package main

import "fmt"

func main() {
	var n int 
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i] <= 10 {
			fmt.Printf("A[%d] = %d\n", i, arr[i])
		}
	}
}