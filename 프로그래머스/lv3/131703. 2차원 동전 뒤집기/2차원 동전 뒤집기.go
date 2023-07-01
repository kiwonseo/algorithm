func solution(beginning [][]int, target [][]int) int {
	var s []int
	if i := flipRowThenColumn(getDiff(beginning, target), true); i >= 0 {
		s = append(s, i)
	}
	if i := flipRowThenColumn(getDiff(beginning, target), false); i >= 0 {
		s = append(s, i)
	}
	if i := flipColumnThenRow(getDiff(beginning, target), true); i >= 0 {
		s = append(s, i)
	}
	if i := flipColumnThenRow(getDiff(beginning, target), false); i >= 0 {
		s = append(s, i)
	}

	if len(s) == 0 {
		return -1
	}
	min := s[0]
	for i := 1; i < len(s); i++ {
		if min > s[i] {
			min = s[i]
		}
	}

	return min
}

func flipRowThenColumn(diff [][]bool, b bool) int {
	row, cntRow := flipRow(diff, b)
	column, cntCol := flipColumn(row, true)

	for i := range column {
		for j := range column[i] {
			if column[i][j] {
				return -1
			}
		}
	}
	return cntRow + cntCol
}

func flipColumnThenRow(diff [][]bool, b bool) int {
	column, cntCol := flipColumn(diff, b)
	row, cntRow := flipRow(column, true)

	for i := range row {
		for j := range row[i] {
			if row[i][j] {
				return -1
			}
		}
	}
	return cntRow + cntCol
}

func flipRow(diff [][]bool, b bool) ([][]bool, int) {
	result := diff
	count := 0
	for i := range result {
		if result[i][0] == b {
			for j := range result[i] {
				result[i][j] = !result[i][j]
			}
			count++
		}
	}
	return result, count
}

func flipColumn(diff [][]bool, b bool) ([][]bool, int) {
	result := diff
	count := 0
	for i := range result[0] {
		if result[0][i] == b {
			for j := range result {
				result[j][i] = !result[j][i]
			}
			count++
		}
	}
	return result, count
}

func getDiff(beginning [][]int, target [][]int) [][]bool {
	diff := make([][]bool, len(beginning))
	for i := range diff {
		diff[i] = make([]bool, len(beginning[i]))
	}

	for i := range diff {
		for j := range diff[i] {
			diff[i][j] = beginning[i][j] != target[i][j]
		}
	}

	return diff
}