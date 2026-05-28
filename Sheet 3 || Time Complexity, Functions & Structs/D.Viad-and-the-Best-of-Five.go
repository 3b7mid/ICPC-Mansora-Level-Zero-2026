package main

import (
	"bufio"
	"fmt"
	"os"
)

func mostFrequent(str string) string { 
	mp := map[rune]int{}

	for _, ch := range str {
		mp[ch]++
	}
	if mp['A'] > mp['B'] {
		return "A"
	} 
	return "B"
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int 
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var str string 
		fmt.Fscan(in, &str)
		fmt.Fprintln(out, mostFrequent(str)) 
	}
}
