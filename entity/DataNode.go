package entity

type DataNode struct {
	Key interface{}
	Value interface{}
	Next *DataNode
	Down *DataNode
	H int // 1 代表头， 0 代表数据节点
}
