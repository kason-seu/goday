package chapter4

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_to_json(t *testing.T) {
	type args struct {
		p *Person
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := to_json(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("to_json() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_PrintJson(t *testing.T) {

	val := to_json(&Person{"lucy", "nanjing"})

	fmt.Println(string(val))

}


func Test_ParseJson(t *testing.T) {
	val := to_json(&Person{"lucy", "nanjing"})


	v := to_object(val)

	var v2 *Person = to_object(val)

	var v3 *Person
	v3 = to_object(val)
	fmt.Println(*v)

	fmt.Println(v2.Addr)

	fmt.Println(v3.Name)
}