package main

// 维护一个长度m+n的临时数组
func merge88_1(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, m+n)

	i, j, idx := 0, 0, 0

	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			result[idx] = nums1[i]
			idx++
			i++
		} else {
			result[idx] = nums2[j]
			idx++
			j++
		}
	}

	for i < m {
		result[idx] = nums1[i]
		idx++
		i++
	}

	for j < n {
		result[idx] = nums2[j]
		idx++
		j++
	}

	copy(nums1, result)
}

// 倒着遍历
func merge88_2(nums1 []int, m int, nums2 []int, n int) {
	idx := m + n - 1

	i, j := m-1, n-1

	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[idx] = nums1[i]
			idx--
			i--
		} else {
			nums1[idx] = nums2[j]
			idx--
			j--
		}
	}

	for j >= 0 {
		nums1[idx] = nums2[j]
		idx--
		j--
	}
}
