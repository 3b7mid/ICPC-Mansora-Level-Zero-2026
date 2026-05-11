package main

import "fmt"

func main() {
	var t int64
	fmt.Scan(&t)
	for t > 0 {
		var n, k int64
		fmt.Scan(&n, &k)

		if n%2 == k%2 && k*k <= n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}
