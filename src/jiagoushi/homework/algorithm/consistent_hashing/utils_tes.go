package consistent_hashing

import (
	"fmt"
)

//test slince

func SlinceTest() {
	var s []uint32
	s = insertBefor(s, 1, 0)
	s = insertBefor(s, 3, 1)
	s = deleteAt(s, 0)
	printSlince(s)
}
func printSlince(s []uint32) {
	fmt.Println(fmt.Sprintf("s =%+v", s))
}

func StandDeviationTest() {
	data := []int{1, 2, 3, 4}
	fmt.Println(StandDeviation(data))
	//append(data, )
}

func ArrayRingTest() {
	arrayRinger := NewSortArrayRing()
	var i uint32
	for i = 1; i <= 3; i++ {
		arrayRinger.Put(i)
		arrayRinger.Put(100 - i)
	}
	for i = 1; i <= 5; i++ {
		j, _ := arrayRinger.FindNear(i)
		k, _ := arrayRinger.FindNear(100 - i)
		s := "找到%d 的匹配数据%d"
		fmt.Println(fmt.Sprintf(s, i, j))
		fmt.Println(fmt.Sprintf(s, 100-i, k))
	}

	for i = 1; i <= 2; i++ {
		arrayRinger.Delete(i)
		arrayRinger.Delete(100 - i)
	}
	for i = 1; i <= 5; i++ {
		j, _ := arrayRinger.FindNear(i)
		k, _ := arrayRinger.FindNear(100 - i)
		s := "找到%d 的匹配数据%d"
		fmt.Println(fmt.Sprintf(s, i, j))
		fmt.Println(fmt.Sprintf(s, 100-i, k))
	}
}
