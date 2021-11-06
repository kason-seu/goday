package main

import (
	"fmt"

	"github.com/kason-seu/leetcode/module-main/entity"
)
func mm[T any](s []T) {
	for _, v := range s {
	 fmt.Print(v)
	}
}
   
func main() {
	fmt.Println("hello world")
	var n entity.ListNode = entity.ListNode{
		Val: 0,
	}
	fmt.Println(n.Val)

	mm([]string{"你好, ", "脑子进了煎鱼\n"})
	mm([]int64{1, 2, 3})
}
