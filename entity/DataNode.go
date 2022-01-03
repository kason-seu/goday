package entity

type DataNode struct {
	Key interface{}
	Value interface{}
	Next *DataNode
	Down *DataNode
}
