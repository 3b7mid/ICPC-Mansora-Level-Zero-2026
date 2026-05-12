package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)

	for _, v := range arr {
		fmt.Print(v, " ")
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Printf("\n")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}