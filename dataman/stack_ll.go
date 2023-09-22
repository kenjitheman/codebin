package dataman

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(val int) {
	newNode := &Node{Data: val, Next: s.Top}
	s.Top = newNode
}

func (s *Stack) Pop() int {
	if s.Top == nil {
		return -1
	}
	val := s.Top.Data
	s.Top = s.Top.Next
	return val
}

func TryStackLL() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
