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

package listpack

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

	fmt.Println("Displaying List")
	if l.Head == nil || l.Tail == nil {
		fmt.Println("\nHead %x Tail %x \n",l.Head, l.Tail)
	} else {
		
		for p := l.Head; p != nil; p = p.Next {
			fmt.Println("Val: ",p.Value)
		}
	}

}

func (l* List) SumList() int {

// Sum the digits in the list. List is built with head pointing to 1st
// place, followed by 10s, 100s and so on 
	var sum, mult  = 0,1
	for p := l.Head ; p != nil; p = p.Next {

		sum += mult * p.Value
		mult *= mult * 10
	}
	return sum

}

//func (l* list) ReverseList() 


