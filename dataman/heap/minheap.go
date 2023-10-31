package heap

type MinHeap struct {
	data   []int
	length int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data:   make([]int, 0),
		length: 0,
	}
}

func (h *MinHeap) Parent(idx int) int {
	return (idx - 1) / 2
}

func (h *MinHeap) LeftChild(idx int) int {
	return idx*2 + 1
}

func (h *MinHeap) RightChild(idx int) int {
	return idx*2 + 2
}

func (h *MinHeap) HeapifyUp(idx int) {
	for h.data[idx] < h.data[h.Parent(idx)] {
		h.data[idx], h.data[h.Parent(idx)] = h.data[h.Parent(idx)], h.data[idx]
		idx = h.Parent(idx)
	}
}

func (h *MinHeap) HeapifyDown(idx int) {
    for h.LeftChild(idx) < h.length {
        left := h.LeftChild(idx)
        right := h.RightChild(idx)
        if right < h.length && h.data[right] < h.data[left] {
            left = right
        }
        if h.data[idx] < h.data[left] {
            break
        }
        h.data[idx], h.data[left] = h.data[left], h.data[idx]
        idx = left
    }
}

func (h *MinHeap) Insert(val int) {
    h.data = append(h.data, val)
    h.length++
    h.HeapifyUp(h.length - 1)
}
