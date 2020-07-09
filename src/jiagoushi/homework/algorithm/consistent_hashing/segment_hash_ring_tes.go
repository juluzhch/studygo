package consistent_hashing

import (
	"fmt"
)

//验证排序
func TestSegmentAll() {
	var slotSize uint32 = 512 * 2 * 2 * 8
	//fmt.Println(slotSize)
	var dataSize = 1000000
	//var nodeSize=0
	//TestSegment(slotSize,10,100)
	for i := 10000; i <= dataSize; i = i * 10 {
		TestSegment(slotSize, 10, i)
	}

}
func TestSegmentVariance() {
	var slotSize uint32 = 512 * 2 * 2 * 8
	//fmt.Println(slotSize)
	var dataSize = 1000000
	//var nodeSize=0
	//TestSegment(slotSize,10,100)
	for i := 10000; i <= dataSize; i = i * 10 {
		TestSegment(slotSize, 10, i)
	}

}

func TestSegment(slotSize uint32, nodeSize, dataSize int) {

	ring := GetSegmentHashRing(nodeSize, slotSize)
	//ring.PrintRing()
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
	fmt.Println(fmt.Sprintf("节点数%d,测试数据量%d,标准差%.3f,标准差/数据量= %0.3f", nodeSize, dataSize, sd, sd/float64(dataSize)))
	////ring.PutNode()
}
