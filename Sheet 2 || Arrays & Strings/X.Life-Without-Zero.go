package main

import "fmt"

func removeZero(n int) int {
	ret, ten := 0, 1
	for n > 0 {
		dig := n % 10
		n /= 10
		if dig == 0 {
			continue
		}
		ret += dig * ten
		ten *= 10
	}
	return ret
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sum := a + b
	a = removeZero(a)
	b = removeZero(b)
	sum = removeZero(sum)
	if a + b == sum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
