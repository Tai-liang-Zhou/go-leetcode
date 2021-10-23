package main

import "reflect"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dictS := make(map[rune]int)

	for _, value := range s {
		dictS[value]++
	}

	dictT := make(map[rune]int)

	for _, value := range t {
		dictT[value]++
	}

	res := reflect.DeepEqual(dictS, dictT)

	return res
}

func main() {

}