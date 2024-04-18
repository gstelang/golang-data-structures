package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func insert(data int, root *Node) {
	currentNode := root

	for {
		if data < currentNode.data {
			if currentNode.left == nil {
				currentNode.left = &Node{data: data}
				break
			} else {
				currentNode = currentNode.left
			}
		}
		if data > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data: data}
				break
			} else {
				currentNode = currentNode.right
			}
		}
	}

}

// inorder traversal
func sort(node *Node) {

	if node == nil {
		return
	}
	if node.left != nil {
		sort(node.left)
	}

	fmt.Println(node.data)

	if node.right != nil {
		sort(node.right)
	}
}

func main() {

	root := Node{data: 9}
	insert(8, &root)
	insert(20, &root)
	insert(22, &root)

	fmt.Println("Printing inorder")
	sort(&root)
}
