package consistent_hashing

import (
	"fmt"
)

func RingStaticsTest() {
	for i := 100; i <= 10000; i = i * 10 {
		RingTest(10, i)
	}

}

func RingStaticsTestWithPrefix() {
	//nodePrefix="abcde"
	//RingStaticsTest()
	//nodePrefix="12345678"
	//RingStaticsTest()
	//nodePrefix="192.168.110."
	//RingStaticsTest()
	nodePrefix = "192.16fdjkdsfjkdfdsfjdskfdsklfjkdsfjklsdfkjl"
	RingStaticsTest()
	nodePrefix = "a"
	RingStaticsTest()
}
func RingTest(ringSize, dataSize int) {

	ring := GetSortArrayHashRing(ringSize)
	data := GetMockDataKey(dataSize)
	var result = make(map[string]int)
	for _, key := range data {
		node := ring.FindNearNode(key)
		result[node]++
	}
	var resultData []int
	for _, v := range result {
		resultData = append(resultData, v)
	}
	sd := StandDeviation(resultData)

	fmt.Println(fmt.Sprintf("节点数%d,测试数据量%d,标准差%.3f,标准差/数据量= %0.3f", ringSize, dataSize, sd, sd/float64(dataSize)))

}
func TestSortArrayRingManager() {
	hashRing := NewHashRing()
	s := GetMockNodeKey(10)
	for _, node := range s {
		hashRing.PutNode(node)
	}
	hashRing.PrintRing()
	for _, node := range s {
		fmt.Println("remove:" + node)
		hashRing.DeleteNode(node)
		hashRing.PrintRing()
	}
}

func PrintSortArrayHashRingTest() {
	hashRing := GetSortArrayHashRing(10)
	hashRing.PrintRing()
}

func SortArrayHashRingCorrectTest() {
	var ringSize = 10
	var dataSize = 2
	ring := GetSortArrayHashRing(ringSize)
	data := GetMockDataKey(dataSize)
	ring.PrintRing()
	for _, key := range data {
		node := ring.FindNearNode(key)
		s := fmt.Sprintf("%s hash =%d  in %s ,node hash =%d", key, defaultHashCodeFun(key), node, defaultHashCodeFun(node))
		fmt.Println(s)
	}
}
