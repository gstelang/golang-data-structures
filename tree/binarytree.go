package main

import "fmt"

type TreeNode struct {
	data  interface{}
	left  *TreeNode
	right *TreeNode
}

// inorder traversal
func inorder(node *TreeNode) {

	if node == nil {
		return
	}
	if node.left != nil {
		inorder(node.left)
	}

	fmt.Print(node.data)

	if node.right != nil {
		inorder(node.right)
	}
}

// preorder traversal
func preorder(node *TreeNode) {

	if node == nil {
		return
	}
	fmt.Print(node.data)

	if node.left != nil {
		preorder(node.left)
	}

	if node.right != nil {
		preorder(node.right)
	}
}

func main() {

	nodeA := TreeNode{data: 1}
	nodeB := TreeNode{data: 2}
	nodeC := TreeNode{data: 3}
	nodeA.left = &nodeB
	nodeA.right = &nodeC
	fmt.Println("Printing inorder")
	inorder(&nodeA)
	fmt.Println("Printing preorder")
	preorder(&nodeA)
}
