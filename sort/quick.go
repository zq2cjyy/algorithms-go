package sort

func quickSort(list []int) {
	quickSortMethod(list, 0, len(list)-1)
}

func quickSortMethod(list []int, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	std := list[i]

	for right > left {

		//从右往左扫描
		for list[right] >= std && right > left {
			right--
		}

		//从左往右扫描
		for list[left] <= std && left < right {
			left++
		}

		//交换
		if right > left {
			swapValue(list,left,right)
		}
	}

	//循环结束后 表明 left==right
	if list[right] < std {
		swapValue(list,i,right)
	} else if left > i && list[left] > std { //表示左标识移动过 且当前值比参考值要大
		right--
		swapValue(list,i,right)
	}

	quickSortMethod(list, i, right-1)
	quickSortMethod(list, right+1, j)

}
