package sort

func SelectionSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		maxIdx := 0
		maxValue := data[0]

		for j := 1; j < len(data)-i; j++ {
			if data[j] >= maxValue {
				maxIdx = j
				maxValue = data[j]
			}
		}

		data[len(data)-i-1], data[maxIdx] = data[maxIdx], data[len(data)-i-1]
	}

	return data
}
