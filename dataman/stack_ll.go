package dataman

import (
	"fmt"
)

type Node__ struct {
	Data int
	Next *Node__
}

type Stack_ struct {
	Top *Node__
}

func (s *Stack_) Push(val int) {
	newNode := &Node__{Data: val, Next: s.Top}
	s.Top = newNode
}

func (s *Stack_) Pop() int {
	if s.Top == nil {
		return -1
	}
	val := s.Top.Data
	s.Top = s.Top.Next
	return val
}

func TryStackLL() {
	stack := Stack_{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
