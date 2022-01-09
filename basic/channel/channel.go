package channel

import (
	"fmt"
	"time"
)

func test(num int, ch chan int)  {

	for i := 0; i < num * 10; i++ {

		fmt.Println(num, "task hehe")
		time.Sleep(time.Second)

	}
	ch <- num
}

func ChannelSenderWait(c chan int)  {


	fmt.Println("chan send start")
	c <- 100

	fmt.Println("chan send finish")
}

func ChannelReceiverWait(c chan int) {

	time.Sleep(time.Second * 10)
	fmt.Println("sender gorountine prepare")
	c <- 100
	fmt.Println("sender gorountine send data to chan finish")
}
