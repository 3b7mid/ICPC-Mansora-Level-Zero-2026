package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	x := make([]int, n)
	y := make([]int, n)
	z := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i])
		fmt.Fscan(in, &y[i])
		fmt.Fscan(in, &z[i])
	}

	for i := 0; i < n; i++ {
		cx, cy, cz := 0, 0, 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if x[i] == x[j] {
				cx++
			}
			if y[i] == y[j] {
				cy++
			}
			if z[i] == z[j] {
				cz++
			}
		}
		fmt.Fprint(out, cx, " ", cy, " ", cz, "\n")
	}
}
