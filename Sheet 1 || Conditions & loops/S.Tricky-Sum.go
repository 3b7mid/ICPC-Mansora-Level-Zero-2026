package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int64
		fmt.Scan(&n)
		var allSum int64 = n * (n + 1) / 2
		var powerSum int64 = 0
		var pow int64 = 1
		for n > 0 {
			powerSum += pow
			pow *= 2
			n /= 2
		}
		fmt.Println(allSum - powerSum - powerSum)
	}
}
