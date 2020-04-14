// All linked list based tests


/* 

1. create singly linked list
2. create doubly linked list
3. implement stack in a doubly linked list
4. implmenet queue in the same doubly linked list
5. find 'kth' element from the end in singly linked list
5. reverse singly linked list
6. add two numbers when each node is a digit with head having 1s place
7. add two numbers when each node is a dight with head having largest place

*/

package main

import (
	"fmt"
	"math/rand"
)


type SNode struct {

	Value int
	Next *SNode

}

type DNode struct {

	Value int
	Next *DNode
	Prev *DNode

}

func StackPushSll(head **SNode,tail **SNode, v int) {
// Stack implementation Add Singly Linked List

	tgt := new(SNode)
	tgt.Value = v
	tgt.Next = nil
	fmt.Printf("\n StackPushSll: Head %p, Tail %p, Tgt %p, Value %d",head, tail, tgt,v)
	if (*head == nil) {
		//First node of Singly Linked List
		*head = tgt
	} else {
		//head at start of list, update current tail
		//to point to new node
		(*tail).Next = tgt
	}
	*tail = tgt
}

func PrintSll(h *SNode) {

	fmt.Println("PrintSll:",h)
	for i := h; i != nil; i = i.Next {
		fmt.Printf("Ptr %p Value %d\n",i, i.Value)
	}



}

func main() {

	var shead, stail *SNode = nil, nil
	//var dhead, dtail *DNode = nil, nil
	var listsize = 10

	for i := 0; i < listsize; i ++ {
		StackPushSll(&shead, &stail, rand.Int()) 
	}

	PrintSll(shead)




}

