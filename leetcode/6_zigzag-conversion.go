package main

// https://leetcode.cn/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := make([][]byte, numRows)
	for i := range res {
		res[i] = make([]byte, 0)
	}

	row := 0      // 行号
	reverse := -1 // 判断是否翻转遍历 0->n , n->0
	for _, v := range s {
		if row == numRows-1 || row == 0 {
			reverse *= -1 // 翻转
		}

		res[row] = append(res[row], byte(v))
		row += reverse
	}

	str := ""
	for _, row := range res {
		for _, b := range row {
			str += string(b)
		}
	}

	return str
}
