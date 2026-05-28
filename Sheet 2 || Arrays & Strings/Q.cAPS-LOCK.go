package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	ok := true
	for i := 1; i < len(str); i++ {
		if 'a' <= str[i] && str[i] <= 'z' {
			ok = false
			break
		}
	}

	if ok {
		res := ""
		for i := range str {
			ch := str[i]

			if 'a' <= ch && ch <= 'z' {
				ch -= 32
			} else {
				ch += 32
			}
			res += string(ch)
		}
		fmt.Println(res)
	} else {
		fmt.Println(str)
	}
}
