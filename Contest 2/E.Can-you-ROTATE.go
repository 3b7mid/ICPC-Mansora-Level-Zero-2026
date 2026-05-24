package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in  := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l, r := k, n-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	
	for _, v := range a {
		fmt.Fprintf(out, "%d ", v)
	}
}