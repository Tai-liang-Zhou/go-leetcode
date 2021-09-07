package main

import (
	"fmt"
)

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	l := len(s) - 1
	for i := 0; i <= l/2; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}

	fmt.Println(string(s))
	return
}

// func main() {
// 	s := "hello"
// 	teststring := []byte(s)
// 	reverseString(teststring)
// 	fmt.Println(s)
// }
