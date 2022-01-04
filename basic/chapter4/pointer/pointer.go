package pointer

import "fmt"

func Pointer() {



	var ptr **int

	i := 5

	v := *&i

	fmt.Println(v)

	p1 := &i

	ptr = &p1

	fmt.Println(*ptr)
	fmt.Println(**ptr)
	/***ptr[1] = 50

	**ptr[2] = 500*/

	/*for i, v := range ptr{

		fmt.Println(i, v)

	}*/

	var pparr [3]**int

	one := &i
	pparr[0] = &one

	fmt.Println(**pparr[0])



	//arrptr := new([3]int)  // var arr [3]int arrptr := &arr
	var arr [3]int
	arrptr := &arr
	fmt.Printf("%T\n", arrptr)

	(*arrptr)[0] = 3

	for _,v := range *arrptr {
		fmt.Println(v)
	}


	// init array
	var arr0 [3]int
	arr0[0] = 100
	arr1 := [3]int{1,2,3}
	arr2 := [...]int{99:-1}

	fmt.Println(arr0)
	fmt.Println(arr1)
	fmt.Println(arr2)
	arr1 = [3]int{1,2,3}
	fmt.Printf("%p\n", &arr1)
	clearArr(arr1)
	//clearArr2(&arr1)
	clearArr3(&arr1)
}

func clearArr(arr [3]int)  {
	fmt.Printf("%p\n",&arr)
}

func clearArr2(arr *[3]int)  {
	fmt.Printf("%p\n",arr)
	fmt.Println("before", *arr)
	*arr = [3]int{}
	fmt.Println("after", *arr)
}

func clearArr3(arr *[3]int)  {
	fmt.Printf("%p\n",arr)
	fmt.Println("before", *arr)
	for i := range arr {
		(*arr)[i] = 0
	}
	fmt.Println("after", *arr)
}