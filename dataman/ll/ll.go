package ll

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{Val: data, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
