package main

import (
	"bufio"
	"fmt"
	"os"
)

func divisible_numbers(l, r, n int) int {
	numbers := r - l + 1
	numOfDivisible := r/n - (l-1)/n
	return numbers - numOfDivisible
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var l, r, n int
	fmt.Fscan(in, &l, &r, &n)
	fmt.Fprintln(out, divisible_numbers(l, r, n))
}
