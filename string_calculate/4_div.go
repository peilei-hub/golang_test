package main

import (
	"fmt"
	"strconv"
)

func div(nums1, nums2 string) string {
	result := "0"

	for len(nums1) > len(nums2) || (len(nums1) == len(nums2) && nums1 >= nums2) { // nums1 >= nums2
		nums3 := nums2
		count := "1"

		for i := 0; i < len(nums1)-len(nums2); i++ { // 往nums3后面添0，count也对应添0
			nums3 += "0"
			count += "0"
		}

		if len(nums1) < len(nums3) || (len(nums1) == len(nums3) && nums1 < nums3) { // 30>20，但30<200的情况，将200减去一个0
			nums3 = nums3[:len(nums3)-1]
			count = count[:len(count)-1]
		}

		for len(nums1) > len(nums3) || (len(nums1) == len(nums3) && nums1 >= nums3) { // num1 >= nums3
			nums1 = sub(nums1, nums3)
			result = addStrings(count, result) // 减一次就加一次count
		}
	}

	return result
}

func sub(nums1, nums2 string) string {
	result := ""

	borrow := 0
	l1 := len(nums1) - 1
	l2 := len(nums2) - 1
	for l1 >= 0 || l2 >= 0 {
		var val1, val2 int
		if l1 >= 0 {
			val1 = int(nums1[l1] - '0')
			l1--
		}
		if l2 >= 0 {
			val2 = int(nums2[l2] - '0')
			l2--
		}

		temp := val1 - val2 - borrow
		if temp < 0 {
			borrow = 1
			temp += 10
		} else {
			borrow = 0
		}

		result = strconv.Itoa(temp) + result
	}

	idx := 0
	for i, v := range result {
		if v != '0' {
			idx = i
			break
		}
	}
	result = result[idx:]

	return result
}

func main() {
	//nums1 := "9"
	//nums2 := "3"
	nums1 := "12345678"
	nums2 := "234"

	s := div(nums1, nums2)
	fmt.Println(s)
}
