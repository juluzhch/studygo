package consistent_hashing

import (
	"fmt"
)

type SortArrayHashRing struct {
	keys         map[uint32]string //反向索引，hashcode->key
	hashCodeRing SortArrayRing     //排序队列实现数据环
	hashCodeFunc func(key string) uint32
}

func NewSortArrayHashRing() HashRing {
	ring := new(SortArrayHashRing)
	ring.keys = make(map[uint32]string)
	ring.hashCodeRing = NewSortArrayRing()
	ring.hashCodeFunc = defaultHashCodeFun
	return ring
}

func (hashRing *SortArrayHashRing) PutNode(key string) {
	hashCode := hashRing.getHashCode(key)
	hashRing.keys[hashCode] = key
	hashRing.hashCodeRing.Put(hashCode)
}

func (hashRing *SortArrayHashRing) DeleteNode(key string) {
	code := hashRing.getHashCode(key)
	if existKey, ok := hashRing.keys[code]; ok {
		if existKey == key { //find ，完全匹配，删除
			delete(hashRing.keys, code)        //map中移除
			hashRing.hashCodeRing.Delete(code) //数据环移除
		}
	}
}

func (hashRing *SortArrayHashRing) FindNearNode(key string) string {
	code := hashRing.getHashCode(key)
	findValue, find := hashRing.hashCodeRing.FindNear(code)
	if find {
		return hashRing.keys[findValue]
	} else {
		return ""
	}
}

func (hashRing *SortArrayHashRing) getHashCode(key string) uint32 {
	return hashRing.hashCodeFunc(key)
}
func (hashRing *SortArrayHashRing) PrintRing() {
	for i, code := range hashRing.hashCodeRing.GetData() {
		fmt.Println(i, ":", code, "->", hashRing.keys[code])
	}
}
