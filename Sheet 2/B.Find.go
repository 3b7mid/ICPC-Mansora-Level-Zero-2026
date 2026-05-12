package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i, v := range arr {
		if v == x {
			fmt.Println(i)
			return 
		}
	}
	
	fmt.Println("Not Found")
}
