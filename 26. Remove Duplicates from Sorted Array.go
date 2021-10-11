package main

import "fmt"

func removeDuplicates(nums []int) int {

	ln := len(nums)
    if ln <= 1 {
        return ln
    }
    
    j := 0 // points to  the index of last filled position
    for i := 1; i < ln; i++ {
        if nums[j] != nums[i] {
            j++
            nums[j] = nums[i]
        }
    }

	fmt.Println(nums)
    
    return j + 1
    
}
func main(){
	input :=  []int{0,0,1,1,1,2,2,3,3,4}

	output := removeDuplicates(input)

	fmt.Println(output)
}