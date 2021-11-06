package questions

import (
	"fmt"
	"testing"
)

func Test_reorderedPowerOf2(t *testing.T) {
	n := 24

	result := reorderedPowerOf2(n)
	fmt.Println(result)

	nn := 218

	result2 := reorderedPowerOf2(nn)
	fmt.Println(result2)
}


func Test_reorderedPowerOf22(t *testing.T) {
	n := 10

	result := reorderedPowerOf2(n)
	fmt.Println(result)

	nn := 24

	result2 := reorderedPowerOf2(nn)
	fmt.Println(result2)
}
