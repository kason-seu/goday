package skiplist

import "github.com/kason-seu/leetcode/module-main/entity"

type SkipList struct {
	Header *entity.HeaderNode
}

func (t *SkipList) addNode(key interface{}, data interface{}) *SkipList {

	if t.Header == nil {
		d := entity.DataNode{Key: key, Value: data}
		h := entity.HeaderNode{}
		h.Next = &d
		t.Header = &h
	}
	return t
}