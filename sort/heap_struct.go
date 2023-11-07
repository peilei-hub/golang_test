package main

import "fmt"

func main() {
	heap := NewHeap[int]([]int{1, 3, 4, 2, 6, 7}, func(child, parent int) bool {
		return child > parent
	})
	heap.Replace(5)
	fmt.Println(heap.Slice())
	fmt.Println(heap.OrderedSlice())
}

func NewHeap[T comparable](nums []T, cmp func(child, parent T) bool) Heap[T] {
	h := Heap[T]{
		cap: len(nums),
		cmp: cmp,
	}

	temp := make([]T, len(nums)+1)
	for i, num := range nums {
		temp[i+1] = num
	}

	h.nums = temp

	for i := h.cap / 2; i >= 1; i-- {
		h.adjustHeap(i, h.cap)
	}

	return h
}

type Heap[T comparable] struct {
	nums []T
	cap  int
	cmp  func(T, T) bool
}

func (h *Heap[T]) Cap() int {
	return h.cap
}

func (h *Heap[T]) Add(t T) {
	h.nums = append(h.nums, t)
	h.cap++
	for i := h.cap; i/2 >= 1; i = i / 2 {
		if h.cmp(h.nums[i], h.nums[i/2]) {
			h.nums[i], h.nums[i/2] = h.nums[i/2], h.nums[i]
		}
	}
}

func (h *Heap[T]) Replace(t T) {
	if h.cmp(h.nums[1], t) {
		h.nums[1] = t
		h.adjustHeap(1, h.cap)
	}
}

func (h *Heap[T]) OrderedSlice() []T {
	for i := h.cap; i > 1; i-- {
		h.nums[1], h.nums[i] = h.nums[i], h.nums[1]
		h.adjustHeap(1, i-1)
	}
	return h.Slice()
}

func (h *Heap[T]) Slice() []T {
	return h.nums[1:]
}

func (h *Heap[T]) adjustHeap(i int, cap int) {
	par := i

	for {
		pos := par
		if 2*par <= cap && h.cmp(h.nums[2*par], h.nums[pos]) {
			pos = 2 * par
		}
		if 2*par+1 <= cap && h.cmp(h.nums[2*par+1], h.nums[pos]) {
			pos = 2*par + 1
		}
		if pos == par {
			break
		}

		h.nums[pos], h.nums[par] = h.nums[par], h.nums[pos]
		par = pos
	}
}
