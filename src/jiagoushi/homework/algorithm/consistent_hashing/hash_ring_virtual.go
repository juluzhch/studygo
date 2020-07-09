package consistent_hashing

import (
	"strconv"
	"strings"
)

type Node struct {
	NodeIdentity string
	VirtualCount int
}

type NodeManagerImpl struct {
	hashRing HashRing
	nodes    map[string]Node
}

func NewVirtualHashRing() HashRingVirtual {
	nodeManager := new(NodeManagerImpl)
	nodeManager.nodes = make(map[string]Node)
	nodeManager.hashRing = NewHashRing()
	return nodeManager
}
func (manager *NodeManagerImpl) PutNode(key string) {
	manager.PutNodeExtend(key, 1)
}
func (manager *NodeManagerImpl) PutNodeExtend(nodeKey string, virtualCount int) {
	var node Node
	node.NodeIdentity = nodeKey
	node.VirtualCount = virtualCount
	manager.nodes[nodeKey] = node
	//manager.hashRing.Put(nodeKey)
	vNods := manager.getVirtualNodes(nodeKey, virtualCount)
	for _, vNode := range vNods {
		manager.hashRing.PutNode(vNode)
	}
}
func (manager *NodeManagerImpl) DeleteNode(nodeKey string) {
	if node, ok := manager.nodes[nodeKey]; ok {
		delete(manager.nodes, nodeKey)
		vNods := manager.getVirtualNodes(nodeKey, node.VirtualCount)
		for _, vNode := range vNods {
			manager.hashRing.DeleteNode(vNode)
		}
	}
}
func (manager *NodeManagerImpl) getVirtualNodes(nodeKey string, virtualCount int) []string {
	var keys []string
	for i := 1; i <= virtualCount; i++ {
		var key string = nodeKey + "_" + strconv.Itoa(i)
		keys = append(keys, key)
	}
	return keys
}

func (manager *NodeManagerImpl) FindNearNode(dataKey string) string {
	var nodeIdentityWithPrefix = manager.hashRing.FindNearNode(dataKey)
	if nodeIdentityWithPrefix != "" {
		s := strings.Split(nodeIdentityWithPrefix, "_")
		return s[0]
	} else {
		return nodeIdentityWithPrefix //""
	}
}

func (manager *NodeManagerImpl) PrintRing() {
	manager.hashRing.PrintRing()
}
