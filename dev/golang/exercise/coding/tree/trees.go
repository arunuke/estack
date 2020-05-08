// implementing binary search trees and performing operations

// insert a node, delete a node
// in-order, pre-order and post-order traversal

package main

import (
	"fmt"
)
type Node struct {
	Value 	int
	Right 	*Node
	Left 	*Node

}

type Tree struct {

	Root 	*Node
	Level	int

}


func (t *Tree) InOrder() {

	
	if t.Root != nil {
		InOrderTraversal(t.Root)
	} 

	

}

func InOrderTraversal(n *Node) {

	if n != nil {
		InOrderTraversal(n.Left)
		fmt.Println("Val: ",n.Value)
		InOrderTraversal(n.Right)
	}

}
func (t *Tree) Insert(v int) {

	newNode := &Node{v,nil,nil}
	if t.Root == nil {
		t.Root = newNode 
		return

	}
	rootNode := t.Root
	AddNodeRec(rootNode, newNode)	
}

func AddNodeRec(rt *Node, in *Node) {

	if rt.Value < in.Value {
		if rt.Right == nil {
			rt.Right = in
			return
		} else {
			AddNodeRec(rt.Right,in) 
		}

	} else {
		if rt.Left == nil {
			rt.Left = in
			return
		} else {
			AddNodeRec(rt.Left,in)	
		}
	}
}



func main() {

	ptestTree := &Tree{}
	/*
	testTree := Tree{}
	fmt.Println("Test Tree values", testTree)
	fmt.Println("P Test Tree values", ptestTree)
	fmt.Printf("P Test Tree values, with %p\n", ptestTree)
	fmt.Println("addr Test Tree values", &testTree)
	fmt.Printf("addr Test Tree values with %p\n", &testTree)
	fmt.Println("val P Test Tree values", *ptestTree)
	*/

	ptestTree.Insert(8)
	ptestTree.Insert(4)
	ptestTree.Insert(10)
	ptestTree.Insert(2)
	ptestTree.Insert(6)
	ptestTree.Insert(1)
	ptestTree.Insert(3)
	ptestTree.Insert(5)
	ptestTree.Insert(7)
	ptestTree.Insert(9)
	ptestTree.Insert(11)

	ptestTree.InOrder()

}



