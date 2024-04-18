package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

func printLinkedList(node *Node) {
	currentNode := node

	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
		// same as below because selectors automatically dereference pointers to structs
		// fmt.Println((*currentNode).data)
		// currentNode = (*currentNode).next
	}
}

func main() {

	nodeA := Node{data: 1}
	nodeB := Node{data: 2}
	nodeA.next = &nodeB

	fmt.Println(nodeA)       // {1 0x140000b0000}
	fmt.Println(*nodeA.next) // {2 <nil>}
	printLinkedList(&nodeA)
}
