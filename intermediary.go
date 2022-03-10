package algorithm

import (
	bubble "github.com/ahlemarg/BMI/src/algorithm/bubbleSort"
	quick "github.com/ahlemarg/BMI/src/algorithm/quickSort"
)

func Intermediary(rank []float64, sortName string) {
	switch sortName {
	case "quick":
		quick.QuickSort(rank, 0, len(rank))
	case "bubble":
		bubble.BubbleSort(rank)
	default:
		quick.QuickSort(rank, 0, len(rank))
	}
}
