package dataman

type Queue []int

func (q *Queue) Enqueue(val int) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		return -1
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

// func TryQueue() {
// 	queue := Queue{}
// 	queue.Enqueue(1)
// 	queue.Enqueue(2)
// 	queue.Enqueue(3)
//
// 	fmt.Println(queue.Dequeue())
// 	fmt.Println(queue.Dequeue())
// }
