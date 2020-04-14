// All linked list based tests


/* 

1. create singly linked list - add nodes to front or end
2. create doubly linked list - add nodes to front or end
3. implement stack in a doubly linked list - first in first out (push, dequeue)
4. implmenet queue in the same doubly linked list - last in first out (push, pop)
5. find 'kth' element from the end in singly linked list
5. reverse singly linked list
6. add two numbers when each node is a digit with head having 1s place
7. add two numbers when each node is a dight with head having largest place

*/

package main

import (
	"fmt"
//	"math/rand"
)



type Node struct {

	// Node struct has next and prev. But single list will
	// use only next while double list will use both
	Value int
	Next *Node
	Prev *Node

}

type List struct {
	Head *Node
	Tail *Node
	IsDouble bool // Single or Doubly linked list
}

/*

push - push into a stack.
pop - pop out of a stack.

enqueue - add to a queue
dequeue - remove from a queue

*/

func (l *List) Push(v int) {

	// Push adds to the end of list. 
	// For both stacks and queues

	// Create the new value
	n := new(Node)
	n.Value = v
	n.Next = nil

	// If doubly linked list, add information 
	// on previous
	if l.IsDouble != false {
		//Doubly linked list 
		n.Prev = l.Tail
	}
	// If new list, set both head and tail to point
	// to this node and return
	if l.Head == nil && l.Tail == nil{
		//New list starting
		l.Head = n
		l.Tail = n	
	} else {
		// If previously existing list, move tail to new node
		l.Tail.Next = n
		l.Tail = n
	}

}
func (l *List) Pop() int {

	// Pop removes from top of stack
	var v = -1
	if l.IsDouble != false {
		fmt.Println("Doubly link list. Can easily pop")	
		v = l.Tail.Value
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
	} else {
		fmt.Println("Singly linked list. Need some things done")
	}
	return v

}

func (l *List) Dequeue() int {

	// Dequeue removes from front of queue
	v := l.Head.Value
	l.Head = l.Head.Next
	l.Head.Prev = nil
	return v

}

func (l *List) Enqueue(v int)  {

	// Enqueue adds to front of queue
	n := new(Node)
	n.Value = v
	n.Prev = nil

	if l.Head != nil {
		// Some list exists
		if l.IsDouble != false {
			l.Head.Prev = n
		}
		n.Next = l.Head
		l.Head = n
	} else {
		//
		l.Head = n
		l.Tail = n

	}
}

func (l* List) Display() {

	fmt.Printf("Displaying List: ")
	if l.Head == nil || l.Tail == nil {
		fmt.Println("\nHead %x Tail %x \n",l.Head, l.Tail)
	} else {
		fmt.Printf("Head ")	
		for p := l.Head; p != nil; p = p.Next {
			fmt.Printf(" --> %d ",p.Value)
		}
		fmt.Printf("--> Tail \n")	
	}

}

func (l* List) SumList() int {

// Sum the digits in the list. List is built with head pointing to 1st
// place, followed by 10s, 100s and so on 
	var sum, mult  = 0,1
	for p := l.Head ; p != nil; p = p.Next {

		sum += mult * p.Value
		mult *= 10
	}
	return sum

}

func (l* List) MidList() *Node {

	//Use two pointers, first walks one at a time, second two at a time
	//By the time second reaches end, first hits middle 
	var mid, end *Node

	mid = l.Head
	end = mid.Next

	//for _ ; end != nil; mid = mid.Next, end = (end.Next).Next {
	for ; end != nil; mid, end = mid.Next, (end.Next).Next {
		fmt.Println("Walking", mid.Value, end.Value)
	}
	return mid



}

func (l* List) Reverse() {
// Reverse a link list in-place. If not in-place, second list can be created
// by walking through the list and creating a new list by adding at the head
// Or in other words, Enqueueing

// Need three pointers. start, curr and fin

	var start *Node
	cur := l.Head
	fin := cur.Next
	start = nil

	//Flip  Head and Tail of 
	tmp := l.Head
	l.Head = l.Tail
	l.Tail = tmp

	for ; cur != nil; {	
		cur.Next = start
		start = cur
		cur = fin
		if fin != nil {
			fin = fin.Next
		}
	}


}
//func (l* List) CheckRemoveLoop() 


func main() {

	//Sll - Single Linked List, Dll - Double Linked List
	var Sll, Dll, Revll List
	var max = 10

	// Problem 1: Singly linked list as a stack
	for i := 0; i <= max; i++ {
		Sll.Push(i)
	}
	Sll.Display()

	// Problem 2: Doubly linked list and insertion at the head
	for i := 0; i <= max; i++ {
		Dll.Enqueue(i)
	}
	Dll.Display()

	// Problem 3: Adding two linked lists
	/*
	var N1, N2,N3 List
	N1.Enqueue(3)
	N1.Enqueue(4)
	N1.Enqueue(2)

	N2.Enqueue(4)
	N2.Enqueue(6)
	N2.Enqueue(5)

	N1.Display()
	N2.Display()

	fmt.Println("N1 is",N1.SumList())
	fmt.Println("N2 is",N2.SumList())
	fmt.Println("Sum of N1 and N2", N1.SumList() + N2.SumList())	
	for tot := N1.SumList() + N2.SumList(); tot > 0; tot /= 10 {
		N3.Push(tot % 10)
	}
	N3.Display()



	// Problem 4: Finding mid point of a linked list
	pMid := Sll.MidList()
	fmt.Println("Mid point",pMid.Value)
	*/

	// Problem 5: Create a link list, reverse it and print
	for i := 0; i <= max * 2 ; i++ {
		Revll.Push(i)
	}
	Revll.Display()
	Revll.Reverse()
	Revll.Display()

}

