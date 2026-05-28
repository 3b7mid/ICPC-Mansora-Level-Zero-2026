package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(a int, s1, s2 string) string {
	if a == 1 { 
		return s1 + s2
	} 
	return s2 + s1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	var s1, s2 string
	fmt.Fscan(in, &s1, &s2)
	result := f(
		1,
		f(
			1,
			f(0, s1, s2)+s1,
			s2,
		),
		s1+f(
			0,
			s1+s2,
			f(1, s2, s1),
		),
	)
	fmt.Fprintln(out, result)
}
