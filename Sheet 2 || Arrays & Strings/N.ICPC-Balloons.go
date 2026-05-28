package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var n int
		var s string
		fmt.Scan(&n, &s)
		vis := make([]bool, 26)
		cnt := 0
		for i := range s {
			if vis[s[i]-'A'] == false {
				cnt += 2
				vis[s[i]-'A'] = true
			} else {
				cnt++
			}
		}
		fmt.Println(cnt)
		t--
	}
}
