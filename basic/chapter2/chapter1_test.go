package chapter2

import (
	"fmt"
	"testing"
)

func TestOsArgs(t *testing.T)  {

	OsArgs()

	fmt.Println("another write method")
	OsArgs2()

}

func TestP11(t *testing.T) {
	P11()
}

func TestP12(t *testing.T) {
	P12()
}

func TestP13TimeDiff(t *testing.T) {
	P13TimeDiff()
}

func TestL13(t *testing.T)  {

	L13()


}

func TestL13_2(t *testing.T) {
	L13_2([]string{"/Users/kason/gowork/leetcode/basic/files/hello.txt"})
}

func TestL13_3(t *testing.T) {
	L13_3([]string{"/Users/kason/gowork/leetcode/basic/files/hello.txt"})
}

func TestP14(t *testing.T) {
	P14([]string{"/Users/kason/gowork/leetcode/basic/files/hello.txt"})
}

func TestL15(t *testing.T) {
	L15([]string{"https://www.hufuvideo.cn/aliplay/hM94-1-7.html"})
}

func TestP15_17(t *testing.T) {
	P15_17([]string{"https://www.hufuvideo.cn/aliplay/hM94-1-7.html", "https://www.hufuvideo.cn/alitype/1.html"})
}

func TestP15_18(t *testing.T) {
	P15_18([]string{"www.hufuvideo.cn/aliplay/hM94-1-7.html"})
}

func TestP15_19(t *testing.T) {
	P15_19([]string{"https://www.hufuvideo.cn/aliplay/hM94-1-7.html", "https://www.hufuvideo.cn/alitype/1.html"})
}

func TestL16(t *testing.T) {
	L16([]string{"https://www.hufuvideo.cn/aliplay/hM94-1-7.html", "https://www.hufuvideo.cn/alitype/1.html"})
}

func TestP16_110(t *testing.T) {
	P16_110([]string{"https://www.hufuvideo.cn/alitype/1.html"})
}

func TestL17(t *testing.T) {
	L17()
}