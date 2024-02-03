package dll

import "fmt"

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (dll *DoublyLinkedList) Append(data int) {
	newNode := &Node{Data: data, Prev: nil, Next: nil}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}
}

func (dll *DoublyLinkedList) Display() {
	current := dll.Head
	for current != nil {
		fmt.Printf("%d <-> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}
