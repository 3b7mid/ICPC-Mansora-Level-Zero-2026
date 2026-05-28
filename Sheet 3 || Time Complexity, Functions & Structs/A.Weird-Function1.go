package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(x int) int {
	return 2 * x + 3
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	var x int 
	fmt.Fscan(in, &x)
	fmt.Fprintln(out, f(f(f(x))) + (2 * f(x * f(x))))
}
