package sort

func insertionSort(list []int) {
	length := len(list)
	for i := 1; i < length; i++ {
		curIndex := i
		cur := list[i]
		for {
			if curIndex > 0 && cur < list[curIndex-1] {
				list[curIndex] = list[curIndex-1]
				curIndex--
			} else {
				list[curIndex] = cur
				break
			}
		}
		printList(list)
	}
}
