package main

// 一個 linklist 一該會有以下 interfac 的實作方法
// 新增 1.加在前 2.加在後
// 移除 1.移除前 2.移除後

// type LinkedList interface{
//     Prepend(*Node)    // 將節點插入串列開頭
//     Append(*Node)     // 將節點插入串列結尾
//     Pop() *Node       // 將節點自尾部移除並返回節點
//     PopFirst() *Node  // 將節點自頭部移除並返回節點
//     Head() *Node      // 取點頭節點
//     Tail()            // 取得尾節點
//     Remove(*Node)     // 移除指定節點
// }

type Node struct {
	Val interface{}
	Perv *Node
	Next *Node
}

func NewNode(Val interface{}) *Node{
	return &Node{Val, nil,nil}
}

// Singly structure with length of the list and its head
type Single struct {
	lenth int
	Head *Node
}

func NewSingle() *Single {
	return &Single{}
}

// AddAtEnd adds a new snode with given value at the end of the list.

func (ll *Single)AddAtEnd(val int){
	n := NewNode(val)
	if ll.Head == nil {
		ll.Head = n
		ll.lenth ++
		return
	}

	cur := ll.Head
	for ; cur != nil;cur = cur.Next{
	}
	cur.Next = n
	ll.lenth ++
}

// DelAtBeg deletes the snode at the head(beginning) of the list and returns its value. Returns -1 if the list is empty.

func (ll *Single) DelAtGeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.lenth --

	return cur.Val
}

// DelAtEnd deletes the snode at the tail(end) of the list and returns its value. Returns -1 if the list is empty
func (ll *Single) DelAtEnd() interface{} {
	if ll.Head == nil {return -1}

	if ll.Head.Next == nil{

	}

	cur := ll.Head.Next

	for ;cur.Next.Next != nil;cur = cur.Next {	
	}

	reval := cur.Next.Val
	cur.Next = nil
	ll.lenth --
	return reval
}

