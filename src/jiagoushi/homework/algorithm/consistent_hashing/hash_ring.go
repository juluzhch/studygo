package consistent_hashing

//一致性哈希实现-算法-哈希环
//算法实现
//1:key求哈希后,放入排序数组，并用map存放 哈希值与key映射
//2:查找时通过2分查找法查找最近的大于等于 指定hash 值 下标，取出哈希，并从map中获取key
//3:环：小于等于第一个及大于数组最后一个的情况归返回第一个元素
//接口方法名调整
type HashRing interface {
	PutNode(key string)
	DeleteNode(key string)
	FindNearNode(dataKey string) string
	PrintRing() //输出节点在环上的分步-测试用
}
type HashRingVirtual interface {
	HashRing
	PutNodeExtend(nodeKey string, virtualCount int)
}

//哈希工厂
//算法，哈希函数，
func NewHashRing() HashRing {
	return NewSortArrayHashRing()
}

//
//type HashRingImpl struct {
//	hashCodes []uint32  //排序队列实现一致性哈希环
//    keys map[uint32]string //反向索引，hashcode->key
//}
//func NewHashRing() HashRing{
//	ring:=new(HashRingImpl)
//	ring.keys=make(map[uint32]string)
//	return ring
//}
//func(hashRing *HashRingImpl)Put(key string){
//	hashCode:=hashRing.getHashCode(key)
//	hashRing.keys[hashCode]=key
//	hashCodes:=hashRing.hashCodes
//	//按排序插入
//	if len(hashCodes)==0{
//		hashCodes= append(hashCodes, hashCode)
//		hashRing.hashCodes=hashCodes
//		return
//	}
//	findIndex:=hashRing.findNearIndex(key)
//	if findIndex==0{
//		if hashCode<=hashCodes[0]{//添加到头部
//			hashCodes=insertBefor(hashCodes,hashCode,0)
//		}else{//添加到尾部
//			hashCodes=insertBefor(hashCodes,hashCode,len(hashCodes))
//		}
//	}else{//中间添加
//		hashCodes=insertBefor(hashCodes,hashCode,findIndex)
//	}
//	hashRing.hashCodes=hashCodes
//}
//
//func (hashRing *HashRingImpl)Delete(key string){
//   //find ，完全匹配，删除
//	findIndex:=hashRing.findNearIndex(key)
//	if findIndex!=-1{
//		code:=hashRing.hashCodes[findIndex]
//		existKey:=hashRing.keys[code]
//		if existKey==key{
//			delete(hashRing.keys, code)
//			hashRing.hashCodes = deleteAt(hashRing.hashCodes,findIndex)
//		}
//	}
//}
//
//func (hashRing *HashRingImpl)FindNear(key string) string{
//	findIndex:=hashRing.findNearIndex(key)
//	if findIndex==-1{
//		return ""
//	}else{
//		code:=hashRing.hashCodes[findIndex]
//		return hashRing.keys[code]
//	}
//}
//
////查找最近的大于等于本key,哈希值的下标索引
////没有节点返回-1,非-1 表示查到节点的下标
//func (hashRing *HashRingImpl)findNearIndex(key string) int{
//	code:=hashRing.getHashCode(key)
//	keys:=hashRing.hashCodes
//	if len(keys)==0{
//		return -1
//	}
//	if keys[0]>=code||keys[len(keys)-1]<code{
//		return 0
//	}
//	if keys[len(keys)-1]==code{
//		return len(keys)-1
//	}
//	return hashRing.findNearInMiddle(0,len(keys)-1,code)
//}
//
////数组中 beginIndex 位置小于code,endIndex位置 大于code,从中间找
//func (hashRing *HashRingImpl)findNearInMiddle(beginIndex,endIndex int ,code uint32)int{
//	if endIndex-beginIndex<=1{
//		return endIndex
//	}
//	middle:=int((endIndex+beginIndex)/2)
//	if hashRing.hashCodes[middle]>code{
//		return hashRing.findNearInMiddle(beginIndex,middle,code)
//	}
//	if  hashRing.hashCodes[middle]==code{
//		return middle
//	}else{
//		return hashRing.findNearInMiddle(middle,endIndex,code)
//	}
//}
//func  (hashRing *HashRingImpl)getHashCode(key string )uint32{
//	var newKey=key+Md5([]byte(key))
//	return getHashCode(newKey)
//}
//func  (hashRing *HashRingImpl)PrintRing(){
//	//for code := range hashRing.keys {
//	//	fmt.Println( hashRing.keys[code], "->", code)
//	//}
//    for i,code:=range hashRing.hashCodes{
//    	fmt.Println(i,":",code,"->",hashRing.keys[code])
//	}
//	//fmt.Println(hashRing)
//}
