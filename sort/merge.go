package sort

func MergeSort(data []int) []int {
	temp := make([]int, len(data))
	mergeSort(data, temp)
	return data
}

func mergeSort(data, temp []int) {
	if len(data) <= 1 {
		return
	}

	m := len(data) / 2

	mergeSort(data[:m], temp[:m])
	mergeSort(data[m:], temp[m:])

	merge(data[:m], data[m:], data, temp)
}

func merge(left, right, data, temp []int) {
	iLeft, iRight := 0, 0
	iTemp := 0
	for iLeft < len(left) || iRight < len(right) {
		if iLeft >= len(left) {
			temp[iTemp] = right[iRight]
			iTemp++
			iRight++
			continue
		}

		if iRight >= len(right) {
			temp[iTemp] = left[iLeft]
			iTemp++
			iLeft++
			continue
		}

		if left[iLeft] <= right[iRight] {
			temp[iTemp] = left[iLeft]
			iLeft++
		} else {
			temp[iTemp] = right[iRight]
			iRight++
		}
		iTemp++
	}

	copy(data, temp)
}
