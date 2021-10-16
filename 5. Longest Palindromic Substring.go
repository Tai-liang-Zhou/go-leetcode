package main

func longestPalindrome(s string)string{
	if len(s) < 2 {
		return s
	}

	var max string
	for i:=0;i<len(s);i++{
		// 如果是單數
		max = maxPaildrome(s, i, i, max)
		max = maxPaildrome(s, i, i+1,max)
		
	}
return max
}

func maxPaildrome(s string, left, right  int, max string)string{

	var sub string
	for left >= 0 && right < len(s) && s[left] == s[right]{
		sub = s[left:right+1]
		left--
		right++
	}
	if len(max) < len(sub){
		return sub
	}
	return max
}

// func main(){
// 	var test1 string = "babad"

// 	results := longestPalindrome(test1)
// 	fmt.Println(results)

// }