package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var n int
		var str1, str2 string
		fmt.Scan(&n, &str1, &str2)
		is_match := true
		for i := 0; i < n; i++ {
			if (str1[i] == 'R' && str2[i] != 'R') || (str1[i] != 'R' && str2[i] == 'R') {
				is_match = false
				break
			}
		}
		if is_match {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}
