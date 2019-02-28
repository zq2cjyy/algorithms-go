package sort

func selectionSort(list []int) {
	lenth := len(list)
	for i := 0; i < lenth; i++ {
		current := i
		for j := i + 1; j < lenth; j++ {
			if list[j] < list[current] {
				current = j
			}
		}
		if current != i {
			swapValue(list, i, current)
		}
	}
}
