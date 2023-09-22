package dataman

import "fmt"

type Node_ struct {
	Data int
	Prev *Node_
	Next *Node_
}

type DoublyLinkedList struct {
	Head *Node_
	Tail *Node_
}

func (dll *DoublyLinkedList) Append(data int) {
	newNode := &Node_{Data: data, Prev: nil, Next: nil}
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

// func TryDLL() {
// 	dll := DoublyLinkedList{}
// 	dll.Append(1)
// 	dll.Append(2)
// 	dll.Append(3)
//
// 	dll.Display()
// }
