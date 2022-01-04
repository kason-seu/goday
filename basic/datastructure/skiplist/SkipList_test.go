package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T)  {

	skipList := &SkipList{}

	node := skipList.AddNode(1, "foo")
	node = skipList.AddNode(5, "hello")
	node = skipList.AddNode(3, "world")
	fmt.Println("layer ", node.Layer)
	/*if node.Header != nil {
		if (node.Header.Next != nil) {
			fmt.Println("next value : ", node.Header.Next.Value)
		}
		if (node.Header.Down != nil) {
			fmt.Println("down nexr valye ", node.Header.Down.Next.Value)
		}
	}*/

	re := node.Search(5)

	fmt.Println("result ", re)


	re2 := node.Search(3)

	fmt.Println("result ", re2)

}