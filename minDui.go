package main

import "container/heap"

func main() {

}

type SeatManager struct {
	available *Heap
}

func Constructor(n int) SeatManager {
	h := &Heap{}
	heap.Init(h)
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
	}
	return SeatManager{available: h}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.available).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.available, seatNumber)
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
