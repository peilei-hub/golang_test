package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := make([][]byte, numRows)
	for i := 0; i < len(res); i++ {
		res[i] = make([]byte, 0)
	}

	row := 0
	reverse := -1
	for _, v := range s {
		if row == numRows-1 || row == 0 {
			reverse *= -1
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
