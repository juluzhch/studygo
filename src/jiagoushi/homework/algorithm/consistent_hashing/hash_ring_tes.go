package consistent_hashing

import (
	"fmt"
)

func TestAllRing() {
	var nodeSize = 10
	var virtualNodeSize = 200
	var dataSize = 1000000
	fmt.Println("预分槽模式-槽位数16384")
	var hashRing = GetSegmentHashRing(nodeSize, 16384)
	TestRingPerformanceAndDeviation(hashRing, dataSize, nodeSize)
	hashRing = GetSortArrayHashRing(nodeSize)
	fmt.Println("数组排序环")
	TestRingPerformanceAndDeviation(hashRing, dataSize, nodeSize)
	fmt.Println("数组排序环-每个节点200个虚拟节点")
	hashRing = GetVirtualHashRing(nodeSize, virtualNodeSize)
	TestRingPerformanceAndDeviation(hashRing, dataSize, nodeSize)

}

func TestRingPerformanceAndDeviation(ring HashRing, dataSize int, ringSize int) {
	data := GetMockDataKey(dataSize)
	var result = make(map[string]int)
	var begin = GetCurrentTimestamp()
	for _, key := range data {
		node := ring.FindNearNode(key)
		result[node]++
	}
	timeSpan := GetCurrentTimestamp() - begin
	fmt.Println(fmt.Sprintf("耗时%d ms", timeSpan))
	var resultData []int
	for _, v := range result {
		resultData = append(resultData, v)
	}
	sd := StandDeviation(resultData)
	fmt.Println(fmt.Sprintf("节点数%d,测试数据量%d,标准差%.3f,标准差/数据量= %0.3f", ringSize, dataSize, sd, sd/float64(dataSize)))

}

func RingTestHashCode() {
	defaultHashCodeFun = getCrcHash
	RingTest(10, 1000000)
	defaultHashCodeFun = getHashCode
	RingTest(10, 1000000)
	defaultHashCodeFun = getHashCode4KeyWithMd5
	RingTest(10, 1000000)
}
