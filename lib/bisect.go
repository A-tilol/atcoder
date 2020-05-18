package lib

// keyword: binary search, bisect, 2分探索

func bisect(arr []int, v, l, r int) int {
	var m int
	for l < r {
		m = (l + r) / 2
		if arr[m] < v {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

// ソート済配列arrに含まれる数のうちvを超えない最大の数の配列のインデックスを返す
func bisectLeft(arr []int, v, l, r int) int {
	if v < arr[l] {
		return -1
	}

	var m int
	for l < r {
		m = (l+r)/2 + 1
		if arr[m] <= v {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

// ソート済配列arrに含まれる数のうちvを超える最小の数の配列のインデックスを返す
func bisectRight(arr []int, v, l, r int) int {
	if v >= arr[r] {
		return -1
	}

	var m int
	for l < r {
		m = (l + r) / 2
		if arr[m] <= v {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
