package problems

import (
	"fmt"
	"testing"
)

func Test_p4_3(t *testing.T) {

	a := [5]int{1,2,3,4,5}

	ptr := Q43Reverse(&a)

	fmt.Println(ptr)
}

func Test_Q45(t *testing.T)  {
	re := Q45RemoveSameLetters([]string{"abc", "abc", "dfg", "rrr", "rrr", "eee"})
	fmt.Println(re)

}

func Test_Q46RemoveAdjWhiteSpace (t *testing.T) {

	re:=Q46RemoveAdjWhiteSpace([]byte("abc   dd  ff g"))

	fmt.Println(re)

	fmt.Println(string(re[:]))

}

func TestQ44Rotate(t *testing.T) {
	arr := []int{1,2,3,4,5}

	Q44Rotate(arr, 2)

	fmt.Println(arr)

	result := Q44Rotate2([]int{1,2,3,4,5}, 2)
	fmt.Println(result)
}

func TestQ47Reverse(t *testing.T) {
	s := []byte("abcderfg")

	Q47Reverse(s)

	fmt.Println(string(s))
}