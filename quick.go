package quick

func partitionDouble(arr []float64, startIndex int, endIndex int) int {
	left, right := startIndex, endIndex
	pivot := arr[startIndex]
	for left != right {
		for left < right && arr[right] > pivot {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[startIndex] = arr[left]
	arr[left] = pivot
	return left
}

func QuickSort(arr []float64, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partitionDouble(arr, startIndex, endIndex)
	QuickSort(arr, startIndex, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, endIndex)
}
