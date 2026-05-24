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

	var n, x int
	fmt.Fscan(in, &n, &x)

	less, greater := make([]int, 0, n), make([]int, 0, n)
	for i := 0; i < n; i++ {
		var v int 
		fmt.Fscan(in, &v)
		if v < x { 
			less = append(less, v)
		} else if v > x {
			greater = append(greater, v)
		}
	}
	
	for _, v := range less {
		fmt.Fprintf(out, "%d ", v)
	}
	for _, v := range greater {
		fmt.Fprintf(out, "%d ", v)
	}
}
