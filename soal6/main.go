package main

import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any          { old := *h; v := old[len(old)-1]; *h = old[:len(old)-1]; return v }

func GetMinimumPenalty(quantity []int, m int) int64 {
	h := &minHeap{}
	for _, q := range quantity {
		if q > 0 {
			*h = append(*h, q)
		}
	}
	heap.Init(h)
	var total int64
	for m > 0 && h.Len() > 0 {
		x := heap.Pop(h).(int)
		total += int64(x)
		x--
		if x > 0 {
			heap.Push(h, x)
		}
		m--
	}
	// If m still > 0 (not enough stock), penalties of zero (selling nothing) → stays as it is.
	return total
}
