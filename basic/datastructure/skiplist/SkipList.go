package skiplist

import (
	"github.com/kason-seu/leetcode/module-main/entity"
	"math/rand"
	"time"
)

type SkipList struct {
	Header *entity.DataNode
	Layer int
}

func (t *SkipList) Search(key int) interface{} {
	current := t.Header
	for  {
		/*if current.Down == nil {
			stop = true
			continue
		}*/
		if current == nil {
			break
		}
		if current.Next == nil {
			current = current.Down
		} else {
			if current.Next.Key.(int) < key {
				current = current.Next
			} else if current.Next.Key.(int) > key{
				current = current.Down

			} else {
				return current.Next.Value
			}
		}


	}
	return nil
}

func (t *SkipList) AddNode(key int, data interface{}) *SkipList {

	var top *entity.DataNode
	// 当前是一个空的跳表，尚不存在任何数据
	if t.Header == nil {
		d := entity.DataNode{Key: key, Value: data, H:0}
		h := entity.DataNode{H:1}
		h.Next = &d
		t.Header = &h
		top = &d
		t.Layer++
		// 抛硬币 升塔
		for flip() == 1 {
			t.Layer++
			dd := entity.DataNode{Key: key, Value: data, H:0}
			dd.Down = top
			hh := entity.DataNode{H:1}
			hh.Next = &dd
			hh.Down = t.Header
			t.Header = &hh
			top = &dd
		}
		return t
	} else {
		// 不空，找到插入的位置. 并且将需要处理的地方加入到站里面，后面还需要从层低不断向上处理
		current := t.Header
		var stack []*entity.DataNode
		stop := false
		for !stop {

			// 没有下一个节点，那么则向下走
			if current.Next == nil {
				// 记录下需要处理的地方
				stack = append(stack, current)
				current = current.Down
				if current == nil {
					stop = true
				}
			} else {
				if (current.Next.Key).(int) < key {
					current = current.Next
				} else {
					// 记录下需要处理的地方
					stack = append(stack, current)
					current = current.Down
					if current == nil {
						stop = true
					}
				}
			}

		}

		lowerLayerNode := stack[len(stack) - 1]
		d := entity.DataNode{Key: key, Value: data, H:0}
		d.Next = lowerLayerNode.Next
		lowerLayerNode.Next = &d
		top = &d
		stack = stack[:len(stack) - 1]
		// 针对站里面的其余数据判断
		for flip() == 1 {
			if len(stack) == 0 {
				dd := entity.DataNode{Key: key, Value: data, H:0}
				dd.Down = top
				hh := entity.DataNode{H:1}
				hh.Next = &dd
				hh.Down = t.Header
				t.Header = &hh
				top = &dd
			} else {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				dd := entity.DataNode{Key: key, Value: data, H: 0}
				dd.Next = pop.Next
				dd.Down = top
				pop.Next = &dd
				top = &dd
			}
		}
	}
	return t
}

func flip() int{
	//rand.Seed(time.Now().Unix())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(2)
}