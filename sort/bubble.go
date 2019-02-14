package sort

func bubbleSort(list []int) {
	count := len(list)

	for i := 0; i < count; i++ {
		for j := i; j < count; j++ {
			if list[i] > list[j] {
				//int类型的变量用异或换值 效率高 省空间
				swapValue(list, i, j)
			}
		}
	}
}
