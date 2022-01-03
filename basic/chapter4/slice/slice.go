package slice

import "fmt"

func AppSliceInit()  {


	//
	days := [7]string{"Mon", "Tue", "Wes", "Tue", "Fri", "Sat","Sun"}

	slice1 := days[1:4]
	slice2 := days[2:6]

	fmt.Printf("slice1 len %d, cap %d, value %v \n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2 len %d, cap %d, value %v \n", len(slice2), cap(slice2), slice2)

	// 底层数组还可以支撑slice扩展
	broadSlice := slice2[:5]
	fmt.Printf("broadSlice len %d, cap %d, value %v \n", len(broadSlice), cap(broadSlice), broadSlice)

	// 超过切片的容量则会抛错误
	/*broadSlice2 := slice2[:6]
	fmt.Printf("broadSlice2 len %d, cap %d, value %v \n", len(broadSlice2), cap(broadSlice2), broadSlice2)*/
	s := []int{1,2,3}

	fmt.Println(len(s), cap(s))


	slice4 := []int(nil)
	fmt.Printf("len %d, is null %t\n", len(slice4), slice4 == nil)


}