package btree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(value int) {
	if bt.Root == nil {
		bt.Root = &TreeNode{Value: value, Left: nil, Right: nil}
	} else {
		bt.insertRecursive(bt.Root, value)
	}
}

func (bt *BinaryTree) insertRecursive(node *TreeNode, value int) {
	if value <= node.Value {
		if node.Left == nil {
			node.Left = &TreeNode{Value: value, Left: nil, Right: nil}
		} else {
			bt.insertRecursive(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Value: value, Left: nil, Right: nil}
		} else {
			bt.insertRecursive(node.Right, value)
		}
	}
}

func (bt *BinaryTree) Search(value int) bool {
	return bt.searchRecursive(bt.Root, value)
}

func (bt *BinaryTree) searchRecursive(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return bt.searchRecursive(node.Left, value)
	} else {
		return bt.searchRecursive(node.Right, value)
	}
}

func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Value)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	current := root

	for len(stack) > 0 || current != nil {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, current.Value)
			current = current.Right
		}
	}

	return result
}

func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	prevVisited := (*TreeNode)(nil)

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			top := stack[len(stack)-1]

			if top.Right == nil || top.Right == prevVisited {
				result = append(result, top.Value)
				prevVisited = top
				stack = stack[:len(stack)-1]
			} else {
				root = top.Right
			}
		}
	}

	return result
}
