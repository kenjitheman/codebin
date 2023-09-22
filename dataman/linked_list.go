package dataman

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}
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

// func TryLL() {
// 	linkedList := LinkedList{}
// 	linkedList.Append(1)
// 	linkedList.Append(2)
// 	linkedList.Append(3)
//
// 	current := linkedList.Head
// 	for current != nil {
// 		fmt.Println(current.Data)
// 		current = current.Next
// 	}
// }
