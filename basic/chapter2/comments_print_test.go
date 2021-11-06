package chapter2

import (
	"fmt"
	"testing"
)

func Test_print(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			print()
		})
	}
}

type P struct {
    x string
y int
}
func Test_comment(t *testing.T) {

	var num int8 = 12
	fmt.Printf("value = %d \n", num)

	const (
		MONDEY int  = iota
		TUESDAY
		WEDSDAY
	)

	fmt.Println(MONDEY)

	
   b := P{"x", -1}
  
   c := new(P)
   fmt.Print(c)

   d := &P{"y", 10}


  fmt.Println(b.x)
  fmt.Printf("%p \n ", &b)

  b.y = 10

  fmt.Println(b.y)

  fmt.Println("d.value ", d.y)

  d.y = 111
  fmt.Println("d.value ", d.y)

  *&d.y = 2000

  fmt.Println(*d, *&d.x)
  fmt.Println("d.value ", d.y, *&d.y)
}
