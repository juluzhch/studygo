package consistent_hashing

import (
	"fmt"
)

//验证虚拟节点数变更的对方差影响
func NodeManagerStaticsTest() {
	//var vNodeSacle=1
	var maxVnodeSacle = 2000  //000
	var maxDataSize = 1000000 //10000//0000
	for vNodeSacle := 2; vNodeSacle <= maxVnodeSacle; vNodeSacle = vNodeSacle * 10 {
		//for dataSize:=100;dataSize<=maxDataSize;dataSize=dataSize*10{
		NodeManagerTest(10, vNodeSacle, maxDataSize)
		//}
		//NodeManagerTest(10,vNodeSacle,maxDataSize)
	}

}
func NodeManagerTest(nodeSize, vNodeSacle, dataSize int) {
	//t1:=GetCurrentTimestamp()

	nodeManager := GetVirtualHashRing(nodeSize, vNodeSacle)

	//fmt.Println("开始生成节点耗时",GetCurrentTimestamp()-t1)
	data := GetMockDataKey(dataSize)
	var result = make(map[string]int)
	for _, key := range data {
		node := nodeManager.FindNearNode(key)
		result[node]++
	}
	var resultData []int
	for _, v := range result {
		resultData = append(resultData, v)
	}
	sd := StandDeviation(resultData)

	fmt.Println(fmt.Sprintf("节点数%d,虚拟节点数%d, 测试数据量%d,标准差%.3f,标准差/数据量= %0.3f", nodeSize, vNodeSacle, dataSize, sd, sd/float64(dataSize)))

}
