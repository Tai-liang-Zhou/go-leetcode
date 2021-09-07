package main

import "fmt"

func RomanToInt() {
	fmt.Println(romanToInt("IV"))
}

func romanToInt(s string) int {
	romanNunber := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	result := 0

	for i := 0; i < len(s); {
		if i < len(s)-1 && romanNunber[s[i:i+2]] != 0 {
			result += romanNunber[s[i:i+2]]
			i += 2
		} else {
			result += romanNunber[s[i:i+1]]
			i += 1
		}
	}

	return result
}
