package main

import (
	"bufio"
	"fmt"
	"os"
)

func threeNumbersSumCounter(k, s int) int {
	counter := 0
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			z := s - x - y
			if 0 <= z && z <= k {
				counter++
			}
		}
	}
	return counter
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var k, s int 
	fmt.Fscan(in, &k, &s)
	fmt.Fprint(out, threeNumbersSumCounter(k, s))
}
