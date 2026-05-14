package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	s = strings.ToLower(s)
	frq := map[rune]int{}
	for _, char := range s {
		frq[char]++
	}
	if len(frq) == 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
