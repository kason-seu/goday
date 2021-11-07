package chapter4

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Addr string
}


func to_json(p *Person) []byte {


	jsonValue, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error")
	}

	return jsonValue
}

func to_object(b []byte) *Person {

	var p Person
	json.Unmarshal(b, &p)
	return &p
}