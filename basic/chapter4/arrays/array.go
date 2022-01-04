package arrays

import "fmt"

func ArrayCopy()  {


	var array1 [5]*string

	var array2 = [5]*string{new(string), new(string), new(string)}


	*array2[0] = "Red"
	*array2[1] = "blue"
	*array2[2] = "bLACK"

	array1 = array2

	for s := range array1 {
		fmt.Println(s)
	}
}