package main

import (
	"fmt"
	"strconv"
)

// func hammingDistance(x int, y int) int {
// 	count := 0
// 	binaryX := paddingHelp(x)
// 	binaryY := paddingHelp(y)
// 	for i := range binaryX {
// 		if binaryX[i] != binaryY[i] {
// 			count += 1
// 		}
// 	}
// 	return count

// }

// func paddingHelp(x int) string {

// 	BinaryInt := strconv.FormatInt(int64(x), 2)
// 	stringTemp := ""
// 	for i := 0; i < 31; i++ {
// 		if i+len(BinaryInt) < 31 {
// 			stringTemp = stringTemp + "0"
// 		} else {
// 			stringTemp = stringTemp + BinaryInt
// 			break
// 		}

// 	}
// 	return stringTemp
// }

func hammingDistance(x int, y int) int {
	z := int64(x ^ y)
	fmt.Println(z)
	zStr := strconv.FormatInt(z, 2)

	c := 0
	for _, r := range zStr {
		if r == '1' {
			c += 1
		}
	}
	return c
}

// func main() {

// 	fmt.Println(hammingDistance(1, 4))
// }
