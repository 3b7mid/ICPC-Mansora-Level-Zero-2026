package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(n int) int { 
	return (n * (n + 1) * (2 * n + 1)) / 6
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int 
	fmt.Fscan(in, &n)
	fmt.Fprintln(out, sum(n))
}
