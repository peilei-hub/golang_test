package main

import "fmt"

func main() {

	//nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 110, 23, 12, 432423, 543, 543, 54, 55, 6543, 24, 753, 1, 4, 64, 2, 4, 2, 54, 3, 2, 45}
	nums := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 9, 10, 11, 12}

	array1 := minK(nums, 10)
	fmt.Println(array1)

	array2 := maxK(nums, 10)
	fmt.Println(array2)
}

func minK(nums []int, k int) []int {
	if k == 0 {
		return nil
	}

	heap := make([]int, k+1) // 下标从 1 开始。此时 parent(i) = n/2 。 leftChild(i) = 2*i, rightChild(i) = 2*i+1
	for i := 0; i < k; i++ {
		heap[i+1] = nums[i]
	}

	// 先构造堆
	for i := k / 2; i >= 1; i-- { // k/2 为最后一个非叶子节点
		adjustMaxHeap(heap, i, k)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] < heap[1] { // 当前值比大顶堆的顶小
			heap[1] = nums[i] // 将小的值放到堆里

			adjustMaxHeap(heap, 1, k)
		}
	}
	// 此时堆里已经是最小的k个数

	// 这里是递增排序的解法
	for i := k; i >= 1; i-- {
		heap[1], heap[i] = heap[i], heap[1] // 每次将堆顶(因为堆顶最大)的数据移到末尾
		adjustMaxHeap(heap, 1, i-1)
	}

	return heap[1:]
}

// 调整最大堆
func adjustMaxHeap(heap []int, i, n int) { // 从节点 i 开始调整容量为 n 的堆
	par := i

	for {
		// par 为 parent
		maxPos := par // maxPos 定为当前位置
		// 找到左右子节点中最大的元素，与当前节点交换
		if par*2 <= n && heap[maxPos] < heap[par*2] {
			maxPos = par * 2
		}
		if par*2+1 <= n && heap[maxPos] < heap[par*2+1] {
			maxPos = par*2 + 1
		}

		// 退出条件 maxPos 为 par，说明子节点没有比当前节点大，当前节点已经最大了
		if par == maxPos {
			break
		}

		heap[par], heap[maxPos] = heap[maxPos], heap[par]
		par = maxPos // par改为最大子节点，准备下一轮循环
	}
}

func maxK(nums []int, k int) []int {
	if k == 0 {
		return nil
	}

	heap := make([]int, k+1)

	for i := 0; i < k; i++ {
		heap[i+1] = nums[i]
	}
	// 构造小顶堆
	for i := k / 2; i >= 1; i-- { // k/2 为最后一个非叶子节点
		adjustMinHeap(heap, i, k)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > heap[1] {
			heap[1] = nums[i]
			adjustMinHeap(heap, 1, k)
		}
	}
	// 此时堆里已经是最大的 k 个数

	// 这里是排序递减的解法
	for i := k; i >= 1; i-- {
		heap[i], heap[1] = heap[1], heap[i] // 将第一位与最后一位替换
		adjustMinHeap(heap, 1, i-1)
	}

	return heap[1:]
}

func adjustMinHeap(heap []int, i, n int) { // 从节点 i 开始调整容量为 n 的堆
	par := i

	for {
		minPos := par
		if par*2 <= n && heap[par*2] < heap[minPos] {
			minPos = par * 2
		}
		if par*2+1 <= n && heap[par*2+1] < heap[minPos] {
			minPos = par*2 + 1
		}
		if minPos == par { // minPos 为 par，说明子节点没有比当前节点小的
			break
		}

		heap[minPos], heap[par] = heap[par], heap[minPos]
		par = minPos
	}
}
