package dataman

import (
	"fmt"
)

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func TryStackSlice() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
