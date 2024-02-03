package stack_ll

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(val int) {
	newNode := &Node{Val: val, Next: s.Top}
	s.Top = newNode
}

func (s *Stack) Pop() int {
	if s.Top == nil {
		return -1
	}
	val := s.Top.Val
	s.Top = s.Top.Next
	return val
}
