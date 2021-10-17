package sort

import "math/rand"

func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	m := median(data)

	iLeft, iRight := 0, len(data)-1
	for {
		for iLeft < len(data) && data[iLeft] < m {
			iLeft++
		}
		for iRight >= 0 && data[iRight] > m {
			iRight--
		}

		if iLeft > iRight {
			break
		}

		data[iLeft], data[iRight] = data[iRight], data[iLeft]
		iLeft++
		iRight--
	}

	QuickSort(data[:iRight+1])
	QuickSort(data[iLeft:])

	return data
}

func median(data []int) int {
	return data[rand.Intn(len(data))]
}
