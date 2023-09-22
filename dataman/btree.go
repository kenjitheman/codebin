package dataman

import (
	"fmt"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(data int) {
	if bt.Root == nil {
		bt.Root = &TreeNode{Data: data, Left: nil, Right: nil}
	} else {
		bt.insertRecursive(bt.Root, data)
	}
}

func (bt *BinaryTree) insertRecursive(node *TreeNode, data int) {
	if data <= node.Data {
		if node.Left == nil {
			node.Left = &TreeNode{Data: data, Left: nil, Right: nil}
		} else {
			bt.insertRecursive(node.Left, data)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Data: data, Left: nil, Right: nil}
		} else {
			bt.insertRecursive(node.Right, data)
		}
	}
}

func (bt *BinaryTree) Search(data int) bool {
	return bt.searchRecursive(bt.Root, data)
}

func (bt *BinaryTree) searchRecursive(node *TreeNode, data int) bool {
	if node == nil {
		return false
	}
	if data == node.Data {
		return true
	} else if data < node.Data {
		return bt.searchRecursive(node.Left, data)
	} else {
		return bt.searchRecursive(node.Right, data)
	}
}

func (bt *BinaryTree) InOrderTraversal(node *TreeNode) {
	if node != nil {
		bt.InOrderTraversal(node.Left)
		fmt.Printf("%d ", node.Data)
		bt.InOrderTraversal(node.Right)
	}
}

func TryBTree() {
	binaryTree := NewBinaryTree()
	binaryTree.Insert(5)
	binaryTree.Insert(3)
	binaryTree.Insert(7)
	binaryTree.Insert(2)
	binaryTree.Insert(4)
	binaryTree.Insert(6)
	binaryTree.Insert(8)

	fmt.Println("In-Order Traversal:")
	binaryTree.InOrderTraversal(binaryTree.Root)

	fmt.Println("\nSearching for 4:", binaryTree.Search(4))
	fmt.Println("Searching for 9:", binaryTree.Search(9))
}
