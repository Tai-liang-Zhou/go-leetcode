package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList = &ListNode{}
	var head = newList

	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			newList.Next = l1
			l1 = l1.Next
			newList =  newList.Next
		}else{
			newList.Next = l2
			l2 = l2.Next
			newList =  newList.Next
		}
	}

	if l1 != nil {
		newList.Next = l1
	} else if l2 != nil {
		newList.Next = l2
	}

	return head.Next

}

type ListNode struct {
	Next *ListNode
	Val int 
}

// listPrinter(list &ListNode){

// }
// func main(){

// 	l1 := &ListNode{nil,1}
// 	l1.Next = &ListNode{nil,2}
// 	l1.Next.Next = &ListNode{nil,4}

// 	l2 := &ListNode{nil,1}
// 	l2.Next = &ListNode{nil,3}
// 	l2.Next.Next = &ListNode{nil,4}
// 	mergeTwoLists(l1, l2)

// }