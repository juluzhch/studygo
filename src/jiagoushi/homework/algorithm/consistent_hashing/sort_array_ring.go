package consistent_hashing

type SortArrayRing interface {
	Put(value uint32)
	Delete(value uint32)
	FindNear(value uint32) (findValue uint32, find bool)
	DeleteAt(index int)
	IsEmpty() bool
	GetData() []uint32
}

type SortArrayRingImpl struct {
	data []uint32
}

func NewSortArrayRing() SortArrayRing {
	return new(SortArrayRingImpl)
}

func (this *SortArrayRingImpl) Put(value uint32) {
	if len(this.data) == 0 {
		this.data = append(this.data, value)
		return
	}
	findIndex := this.findNearIndex(value)
	if findIndex == 0 {
		if value <= this.data[0] { //添加到头部
			this.insertBefore(value, 0)
		} else { //添加到尾部
			this.insertBefore(value, len(this.data))
		}
	} else { //中间添加
		this.insertBefore(value, findIndex)
	}
}
func (this *SortArrayRingImpl) Delete(value uint32) {
	if len(this.data) > 0 {
		index := this.findNearIndex(value)
		if this.data[index] == value {
			this.data = deleteAt(this.data, index)
		}
	}
}
func (this *SortArrayRingImpl) DeleteAt(index int) {
	this.data = deleteAt(this.data, index)
}
func (this *SortArrayRingImpl) GetData() []uint32 {
	return this.data
}

func (this *SortArrayRingImpl) FindNear(value uint32) (findValue uint32, find bool) {
	index := this.findNearIndex(value)
	if index >= 0 {
		return this.data[index], true
	}
	return 0, false
}
func (this *SortArrayRingImpl) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *SortArrayRingImpl) findNearIndex(code uint32) int {
	keys := this.data
	if len(keys) == 0 {
		return -1
	}
	if keys[0] >= code || keys[len(keys)-1] < code {
		return 0
	}
	if keys[len(keys)-1] == code {
		return len(keys) - 1
	}
	var beginIndex = 0
	var endIndex = len(keys) - 1
	for endIndex-beginIndex > 1 { //二分查找
		middle := int((endIndex + beginIndex) / 2)
		if keys[middle] > code {
			endIndex = middle
			continue
		}
		if keys[middle] == code {
			return middle
		} else {
			beginIndex = middle
		}
	}
	if keys[beginIndex] == code {
		return beginIndex
	}
	return endIndex
}

func (this *SortArrayRingImpl) insertBefore(newValue uint32, insertIndex int) {
	this.data = insertBefor(this.data, newValue, insertIndex)
}
