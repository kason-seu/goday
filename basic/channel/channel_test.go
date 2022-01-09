package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_Chan(t *testing.T) {

	ch := make(chan int)
	for num := 0; num < 5; num++ {
		go test(num, ch)
	}
	for num := 0; num < 5; num++ {
		value := <-ch
		fmt.Println("receiver", value)
		time.Sleep(time.Second * 10)
		fmt.Println("main sleep finish")
	}


}


func Test_ChannelSenderWait(t *testing.T) {
	c := make(chan int)
	go ChannelSenderWait(c)
	time.Sleep(time.Second * 10)
	fmt.Println("after channel send, send cahnnel wait")
	fmt.Println("main rountine sleep over,styart to receive chan, release sender rountine")
	<- c
	fmt.Println("main gorountine finish!")
	time.Sleep(time.Second * 2)

}

func TestChannelReceiverWait(t *testing.T) {

	c := make(chan int)

	go ChannelReceiverWait(c)
	fmt.Println("main gorountine start to recive chan data")
	<- c
	fmt.Println("main gorountine receive the data, release main gorountine")


}

