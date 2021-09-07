package main

// func main() {
// 	strs := []string{"a"}
// 	result := longestCommonPrefix(strs)
// 	fmt.Println(result)
// }

func longestCommonPrefix(strs []string) string {
	// var S_slice []string
	// temp := strs[0][0]
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if (i+1 > len(strs[j])) || strs[j][i] != c {
				return res
			}

		}
		res += string(c)
	}

	return ""
}
