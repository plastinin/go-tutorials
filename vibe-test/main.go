package main

import (
	"cmp"
	"fmt"
)

func main() {
	bt := &BinaryTree[int]{}
	bt.Add(5)
	bt.Add(3)
	bt.Add(7)
	bt.Add(2)
	bt.Add(4)
	bt.Add(6)
	bt.Add(8)

	fmt.Println("Integer Tree Searches:")
	fmt.Println("Search for 5:", bt.Search(5))
	fmt.Println("Search for 10:", bt.Search(10))
	fmt.Println("Search for 3:", bt.Search(3))
	fmt.Println("Search for 7:", bt.Search(7))
	fmt.Println("Search for 2:", bt.Search(2))
	fmt.Println("Search for 4:", bt.Search(4))
	fmt.Println("Search for 6:", bt.Search(6))
	fmt.Println("Search for 8:", bt.Search(8))

	fmt.Println("\nFloat64 Tree Examples:")
	btFloat := &BinaryTree[float64]{}
	btFloat.Add(3.14)
	btFloat.Add(1.618)
	btFloat.Add(2.718)
	btFloat.Add(0.577)
	btFloat.Add(1.414)
	fmt.Println("Search for 3.14:", btFloat.Search(3.14))
	fmt.Println("Search for 1.618:", btFloat.Search(1.618))
	fmt.Println("Search for 4.0:", btFloat.Search(4.0))
	fmt.Println("Search for 0.577:", btFloat.Search(0.577))

	fmt.Println("\nString Tree Examples:")
	btString := &BinaryTree[string]{}
	btString.Add("apple")
	btString.Add("banana")
	btString.Add("cherry")
	btString.Add("date")
	btString.Add("grape")
	fmt.Println("Search for 'apple':", btString.Search("apple"))
	fmt.Println("Search for 'banana':", btString.Search("banana"))
	fmt.Println("Search for 'orange':", btString.Search("orange"))
	fmt.Println("Search for 'date':", btString.Search("date"))
}


type Ordered interface {
	cmp.Ordered
}

// Node represents a single node in the binary search tree.
type Node[T Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// BinaryTree represents the binary search tree itself.
type BinaryTree[T Ordered] struct {
	Root *Node[T]
}

// Add inserts a new value into the binary search tree.
func (bt *BinaryTree[T]) Add(value T) {
	newNode := &Node[T]{Value: value}
	if bt.Root == nil {
		bt.Root = newNode
		return
	}

	current := bt.Root
	for {
		if value < current.Value {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left
		} else if value > current.Value {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		} else { // Value already exists, do nothing
			return
		}
	}
}

// Search checks if a value exists in the binary search tree.
func (bt *BinaryTree[T]) Search(value T) bool {
	current := bt.Root
	for current != nil {
		if value == current.Value {
			return true
		} else if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

// Delete removes a value from the binary search tree.
func (bt *BinaryTree[T]) Delete(value T) {
	bt.Root = deleteNode(bt.Root, value)
}

// deleteNode is a recursive helper function to delete a node from a subtree.
// It returns the new root of the subtree after deletion.
func deleteNode[T Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil // Value not found
	}

	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else { // Node to be deleted is found
		// Case 1: No child or one child
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Case 2: Two children
		// Find the inorder successor (smallest in the right subtree)
		successor := findMin(node.Right)
		node.Value = successor.Value // Replace current node's value with successor's value
		// Delete the inorder successor from the right subtree
		node.Right = deleteNode(node.Right, successor.Value)
	}
	return node
}

// findMin finds the node with the minimum value in a given subtree.
func findMin[T Ordered](node *Node[T]) *Node[T] {
	current := node
	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}
