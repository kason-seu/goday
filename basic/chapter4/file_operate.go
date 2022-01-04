package chapter4

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(fileName string)  {

	fmt.Println("start to read file")
	fileContent, err := ioutil.ReadFile(fileName)

	if nil == err {
		fmt.Println("result ", string(fileContent))
	} else {
		fmt.Println("read file error ")
	}


}