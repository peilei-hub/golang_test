package main

// https://leetcode.cn/problems/next-permutation/

// https://leetcode.cn/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/

// 从后往前找到升降的转折点的升点，比如3 5 4，然后从后往前找到第一个比升点大的点 4，然后3和4替换变成4 5 3，然后从转折点开始往后的点reverse 4 3 5
// 注意一些特殊判断

func nextPermutation(nums []int) {
	var l, m int
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			l = i - 1
			m = i
			break
		}
	}
	if m == 0 { // 没有转折点，说明递降的，那翻转全部
		reverseNums(nums, 0, len(nums)-1)
		return
	}

	// 找到转折点
	// 最后两个点，那将两个点交换即可
	if m == len(nums)-1 {
		swapNums(nums, l, m)
		return
	}

	// 找第一个比转折点的升点大的点
	var r int
	for i := len(nums) - 1; i >= m; i-- {
		if nums[i] > nums[l] {
			r = i
			break
		}
	}

	swapNums(nums, l, r)                // 交换
	reverseNums(nums, l+1, len(nums)-1) // 将l+1开始后面的反转
}

func swapNums(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func reverseNums(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	nextPermutation([]int{1, 3, 2})
}
