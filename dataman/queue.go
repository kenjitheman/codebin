package dataman

type Queue []int

func (q *Queue) AddToQueue(val int) {
	*q = append(*q, val)
}

func (q *Queue) DeleteFromQueue() int {
	if len(*q) == 0 {
		return -1
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}
