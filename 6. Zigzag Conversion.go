package main

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	N:= len(s)
	lines := mae([][]byte, numRows)
	for i := range lines {
		lines[i] = []byte{}
	}

row, step := 0, -1
	for i := 0; i < N;i++ {
		lines[row] = append(lins[row], s[i])
		if row == 0 || row == numRows-1 {
			step *= -1
		}
		rw += step
	}
	rturn string(bytes.Join(lines, []byte{}))
}

/ func main() {
// 	// var testSring string = "PAYPALISHIRING"
// 	// results := convert(testString, 3)
// 	// fmt.Println(results)
// 	// 回文或頂真的  for loop 寫法
// 	row, step := 0, -1
// 	for i := 0; i < 6;i++ {
// 		if row == 0 || row == (-1) {
// 			step *= -1
// 		}
// 		rw += step
// 		fmt.Printlnrow)
// 	}
// }
