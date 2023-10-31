package queue_ll

type Node struct {
	Val  int
	Next *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) AddToQueue(val int) {
	newNode := &Node{Val: val, Next: nil}
	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
		return
	}
	q.Tail.Next = newNode
	q.Tail = newNode
}

func (q *Queue) DeQueue() int {
	if q.Head == nil {
		return -1
	}
	val := q.Head.Val
	q.Head = q.Head.Next
	return val
}
