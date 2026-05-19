package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	l, r, diff := 0, len(str)-1, 0
	for l <= r {
		if str[l] != str[r] {
			diff++
		}
		l++
		r--
	}
	if len(str)%2 == 1 && diff == 0 {
		diff++
	}
	if diff == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
