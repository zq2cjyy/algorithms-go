package sort

import "fmt"

var methodList map[int]MethodModel

func init() {
	methodList = make(map[int]MethodModel)
	methodList[1] = MethodModel{Name: "冒泡", Method: bubbleSort}
	methodList[2] = MethodModel{Name: "快速", Method: quickSort}
	methodList[3] = MethodModel{Name: "选择", Method: selectionSort}
	methodList[4] = MethodModel{Name: "插入", Method: insertionSort}
}

type MethodModel struct {
	Name   string
	Method func([]int)
}

func printList(list []int) {
	for _, v := range list {
		fmt.Print(v)
		fmt.Print(",")
	}
	fmt.Print("\n")
}

func swapValue(list []int, i, j int) {
	if i == j {
		return
	}
	if list[i] == list[j] {
		return
	}
	list[i] ^= list[j]
	list[j] ^= list[i]
	list[i] ^= list[j]
	printList(list)
}

func Run() {
	var cnt int
	fmt.Println("输入参数个数:")
	fmt.Scanf("%d", &cnt)

	list := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		fmt.Printf("输入第%d个参数:", i+1)
		fmt.Scanf("%d", &list[i])
	}

	fmt.Println("选择排序算法:")
	for k, v := range methodList {
		fmt.Printf("%d.%s\n", k, v.Name)
	}

	var m int
	fmt.Scanf("%d", &m)
	method, ok := methodList[m]
	if !ok {
		fmt.Printf("无此方法")
		return
	}

	printList(list)
	method.Method(list)
}
