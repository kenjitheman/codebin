package dataman

type Node__ struct {
	Value int
	Next *Node__
}

type Stack_ struct {
	Top *Node__
}

func (s *Stack_) Push(val int) {
	newNode := &Node__{Value: val, Next: s.Top}
	s.Top = newNode
}

func (s *Stack_) Pop() int {
	if s.Top == nil {
		return -1
	}
	val := s.Top.Value
	s.Top = s.Top.Next
	return val
}
