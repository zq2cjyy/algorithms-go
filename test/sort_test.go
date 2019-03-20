package test

import (
	"testing"
	"math/rand"
	"time"
	"algorithms/sort"
	"fmt"
)

func TestInsert(t *testing.T) {
	doTest(t, 1)
}

func TestQuick(t *testing.T) {
	doTest(t, 2)
}

func TestSelection(t *testing.T) {
	doTest(t, 3)
}

func doTest(t *testing.T, i int) {
	list := generalData(100)
	method := sort.GetSortMethod(i)
	method(list)
	checkData(t, list)
}

func generalData(cnt int) []int {
	fmt.Println("初始化数据")

	ret := make([]int, cnt)
	rand.Seed(time.Now().Unix())
	for i := 0; i < cnt; i++ {
		ret[i] = rand.Intn(100)
	}
	fmt.Println(ret)

	return ret
}

func checkData(t *testing.T, list []int) {
	length := len(list)
	for i := 0; i < length-1; i++ {
		if list[i] > list[i+1] {
			t.Log(fmt.Sprintf("%d > %d", list[i], list[i+1]))
			t.Fail()
		}
	}
}
