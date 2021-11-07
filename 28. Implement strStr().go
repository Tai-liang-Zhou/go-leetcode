package main

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		// fmt.Println("i: ",i)
		for j := 0; j < len(needle) && (i+j) < len(haystack); j++ {
			// fmt.Println("j: ",j)
			if haystack[i+j] != needle[j] {
				break
			}

			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1

}

// func main()  {
// 	var string1 string = "mississippi"

// 	var string2 string = "issipi"

// 	output := strStr(string1, string2)

// 	fmt.Println(output)
<<<<<<< HEAD
// }
=======
// }
>>>>>>> 3e628df1926e5b6f9990151ecc6131225bb56a89
