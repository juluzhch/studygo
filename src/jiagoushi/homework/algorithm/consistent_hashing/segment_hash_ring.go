package consistent_hashing

import (
	"fmt"
	"sort"
)

type SegmentHashRing struct {
	slotSize     uint32
	slotNodeMap  map[uint32]string //slot->nodekey
	nodeList     NodeList          //节点列表（按槽位多少升序排列）
	hashCodeFunc func(key string) uint32
}

func NewSegmentHashRing(slotSize uint32) HashRing {
	ring := new(SegmentHashRing)
	ring.slotNodeMap = make(map[uint32]string)
	ring.slotSize = slotSize
	ring.hashCodeFunc = defaultHashCodeFun
	return ring
}

func (this *SegmentHashRing) PutNode(key string) {
	//创建节点，	给节点分配槽位，更新槽位map，节点添加到列表，节点列表排序
	var node = new(segmentNode)
	node.key = key
	offerData := this.offerDataForNextNode()
	node.sendBack(offerData)
	for i := 0; i < len(offerData); i++ {
		this.slotNodeMap[offerData[i]] = key
	}
	this.nodeList = append(this.nodeList, node)
	this.sortNode()
}

func (this *SegmentHashRing) DeleteNode(key string) {
	//获取节点并从列表删除节点，节点槽位归还并更新槽位map，节点列表排序
	node := this.deleteNodeFromList(key)
	if node == nil {
		return
	}
	this.sendBackData(node.slots)
	this.sortNode()
}

func (this *SegmentHashRing) FindNearNode(key string) string {
	code := this.getHashCode(key)
	slot := code % this.slotSize
	nodeName, find := this.slotNodeMap[slot]
	if find {
		return nodeName
	} else {
		return ""
	}
}

func (this *SegmentHashRing) offerDataForNextNode() []uint32 {
	//新节点添加时从现有节点列表中取槽位
	var offerData []uint32
	if len(this.nodeList) == 0 { //初始化全部槽位
		offerData = make([]uint32, 0, this.slotSize)
		var i uint32
		for i = 0; i < this.slotSize; i++ {
			offerData = append(offerData, i)
		}
		return offerData
	}
	//计算新节点应该获取多少：新节点个数=总个数/（当前节点数+1） ，平均数等于新节点个数（有余数时部分节点会多）。
	newNodeSlotSize := this.slotSize / uint32(len(this.nodeList)+1)
	averageSlotSize := newNodeSlotSize
	var offerTotal uint32 = 0
	for i := len(this.nodeList) - 1; i >= 0; i-- { //出借从底部开始（底部节点个slot多）
		//逐个节点取-//每个节点给多少？  取小（当前节点个数-平均数   新节需要个数-已经分配个数）
		node := this.nodeList[i]
		surplus := node.size() - averageSlotSize //当前节点个数-平均数=本节点多余
		needs := newNodeSlotSize - offerTotal    //新节点还差个数
		if needs <= surplus {                    //差的个数小于多余的 给差的个数就行-够了。
			offerData = append(offerData, node.offer(needs)...)
			return offerData

		} else {
			offerData = append(offerData, node.offer(surplus)...)
			offerTotal += surplus
		}
	}
	return offerData
}

func (this *SegmentHashRing) sendBackData(backData []uint32) {
	//节点删除，将数据归还给剩下的节点。
	//如果没有节点，退出，一个节点，全部归还
	nodeSize := len(this.nodeList)
	if nodeSize == 0 {
		return
	}
	if nodeSize == 1 {
		this.sendBackToNode(this.nodeList[0], backData)
		return
	}
	totalBackCount := len(backData)
	averageBackCound := totalBackCount / nodeSize
	remainder := totalBackCount % nodeSize
	currentIndex := 0
	haveBackCount := 0
	for i := 0; i < len(this.nodeList); i++ { //归还从头部还，头部节点slot少
		if remainder > 0 {
			//归还 平均数+1
			backCount := averageBackCound + 1
			nextDataIndex := currentIndex + backCount
			remainder = remainder - 1
			haveBackCount = haveBackCount + backCount
			this.sendBackToNode(this.nodeList[i], backData[currentIndex:nextDataIndex])
		} else {
			backCount := averageBackCound
			nextDataIndex := currentIndex + backCount
			haveBackCount = haveBackCount + backCount
			this.sendBackToNode(this.nodeList[i], backData[currentIndex:nextDataIndex])
		}
		if haveBackCount == len(backData) { //还完了
			return
		}
	}
}
func (this *SegmentHashRing) sendBackToNode(node *segmentNode, backData []uint32) {
	node.sendBack(backData)
	for _, v := range backData {
		this.slotNodeMap[v] = node.key
	}
}
func (this *SegmentHashRing) deleteNodeFromList(key string) *segmentNode {
	for i := 0; i < len(this.nodeList); i++ {
		node := this.nodeList[i]
		if node.key == key {
			this.nodeList = append(this.nodeList[:i], this.nodeList[i+1:]...)
			return node
		}
	}
	return nil
}
func (this *SegmentHashRing) sortNode() {
	sort.Sort(this.nodeList)
}

func (this *SegmentHashRing) getHashCode(key string) uint32 {
	return this.hashCodeFunc(key)
}
func (this *SegmentHashRing) PrintRing() {
	totalSlotSize := 0
	for _, node := range this.nodeList {
		totalSlotSize += int(node.size())
		fmt.Println(fmt.Sprintf("节点:%s 槽位数:%d", node.key, node.size()))
	}
	s := fmt.Sprintf("总槽位数:%d,map大小:%d", totalSlotSize, len(this.slotNodeMap))
	fmt.Println(s)

}

//================
type NodeList []*segmentNode

func (list NodeList) Len() int {
	return len(list)
}
func (list NodeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
func (list NodeList) Less(i, j int) bool {
	return list[i].size() < list[j].size()
}

type segmentNode struct {
	key   string
	slots []uint32
}

func (this *segmentNode) size() uint32 {
	return uint32(len(this.slots))
}

//借出
func (this *segmentNode) offer(count uint32) []uint32 {
	var offerData = make([]uint32, count)
	copy(offerData, this.slots[:count])
	var newData []uint32
	newData = append(newData, this.slots[count:]...)
	this.slots = newData
	return offerData
}

//归还
func (this *segmentNode) sendBack(backData []uint32) {
	this.slots = append(this.slots, backData...)
}
