package heap

type MinHeap struct {
    data []int
    length int
}

func NewMinHeap() *MinHeap {
    return &MinHeap{
        data: make([]int, 0),
        length: 0,
    }
}

func (h *MinHeap) Parent(i int) int {
    return (i - 1) / 2
}
