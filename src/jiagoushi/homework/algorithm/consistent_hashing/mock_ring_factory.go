package consistent_hashing

func GetSortArrayHashRing(size int) HashRing {
	hashRing := NewHashRing()
	s := GetMockNodeKey(size)
	for _, node := range s {
		hashRing.PutNode(node)
	}
	return hashRing
}

//func GetSortArrayHashRingWithHashFun(size int) HashRing {
//	hashRing := NewHashRing()
//	s := GetMockNodeKey(size)
//	for _, node := range s {
//		hashRing.PutNode(node)
//	}
//	return hashRing
//}
func GetSegmentHashRing(nodeSize int, slotSize uint32) HashRing {
	hashRing := NewSegmentHashRing(slotSize)
	s := GetMockNodeKey(nodeSize)
	for _, node := range s {
		hashRing.PutNode(node)
	}
	return hashRing
}
func GetVirtualHashRing(nodeSize, vNodeScale int) HashRing {
	nodeManager := NewVirtualHashRing()
	s := GetMockNodeKey(nodeSize)
	for _, node := range s {
		nodeManager.PutNodeExtend(node, vNodeScale)
	}
	return nodeManager
}
