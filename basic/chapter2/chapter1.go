package chapter2

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)


func P11() {
	fmt.Println(os.Args[0])
}

func P12() {
	for index, arg := range os.Args{

		fmt.Println(" index = ",index, " arg = ", arg)

	}
}

func P14(files []string) {
	lineStatic := make(map[string]int)
	for _, fileName := range files {
		if f , err := os.Open(fileName); err == nil {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				lineStatic[scanner.Text()]++
				if lineStatic[scanner.Text()] > 1 {
					fmt.Printf("FIND ONE FILE %s \n", fileName)
					break
				}
			}
		}
	}


}

/**
method 1 time diff  49084
method 2 time diff  500
 */
func P13TimeDiff() {
	t1 := time.Now()
	s,seq := "",""

	for _, arg := range os.Args{

		s += seq + arg
		seq = " "

	}
	var d time.Duration =  time.Since(t1)

	timediff := d.Nanoseconds()
	fmt.Println("method 1 time diff ", timediff)


	t1 = time.Now()
	s = ""

	s = strings.Join(os.Args, " ")
	timediff = time.Since(t1).Nanoseconds()
	fmt.Println("method 2 time diff ", timediff)
}


func OsArgs()  {

	s,seq := "",""

	for _, arg := range os.Args{

		s += seq + arg
		seq = " "

	}
	fmt.Println("osArgs: \n", s)
}
func OsArgs2()  {

	s := ""

	s = strings.Join(os.Args, " ")
	fmt.Println("osArgs: \n", s)
}

func L13()  {

	result := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		result[scanner.Text()]++
	}


	for k, v := range result {
		fmt.Println("over 1 write times line ", k, " times ", v)
	}

}
/**
流读文件
 */
func L13_2(files []string)  {
	lineStatic := make(map[string]int)
	for _, fileName := range files {
		if f , err := os.Open(fileName); err == nil {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				lineStatic[scanner.Text()]++
			}
		}
	}

	for word, times := range lineStatic {
		if times > 1 {
			fmt.Printf("the word is %s, the times is %d \n", word, times)
		}
	}
}

func L13_3(files []string)  {

	lineStatic := make(map[string]int)
	for _, fileName := range files {
		if data, err := ioutil.ReadFile(fileName); err == nil {
			line := string(data)
			for _, line := range strings.Split(line, "\n") {
				lineStatic[line]++
			}
		} else {
			fmt.Fprintf(os.Stderr, "err %v", err)
			continue
		}

	}
	for word, times := range lineStatic {
		if times > 1 {
			fmt.Printf("the word is %s, the times is %d \n", word, times)
		}
	}
}

func L15(urls []string)  {

	for _, url := range urls {
		if resp, err := http.Get(url); err == nil {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERR %v\n", err)
					os.Exit(1)
				}
			}(resp.Body)
			if data, err := ioutil.ReadAll(resp.Body); err == nil {
				fmt.Printf("data %s \n", data)
			}

		} else {
			fmt.Fprintf(os.Stderr, "ERR %v\n", err)
		}

	}
}
// 前面章节，后面问题
func P15_17(urls []string)  {
	start := time.Now()
	for _, url := range urls {
		if resp, err := http.Get(url); err == nil {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERR %v\n", err)
					os.Exit(1)
				}
			}(resp.Body)
			if copyData, err := io.Copy(os.Stdout, resp.Body); err == nil {
				fmt.Println("=============")
				fmt.Printf("copy data %v \n", copyData)
			}
		} else {
			fmt.Fprintf(os.Stderr, "ERR %v\n", err)
		}

	}
	cost := time.Since(start).Seconds()
	fmt.Println(" cost ", cost)
}

func P15_18(urls []string)  {
	for _, url := range urls {
		if !strings.HasPrefix(url, "https://") {
			fmt.Printf("url do not has http prefix, origin %s\n", url)
			url = "https://" + url
		}
		if resp, err := http.Get(url); err == nil {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERR %v\n", err)
					os.Exit(1)
				}
			}(resp.Body)
			if copyData, err := io.Copy(os.Stdout, resp.Body); err == nil {
				fmt.Println("=============")
				fmt.Printf("copy data %v \n", copyData)
			}
		} else {
			fmt.Fprintf(os.Stderr, "ERR %v\n", err)
		}

	}
}


func P15_19(urls []string)  {
	start := time.Now()
	for _, url := range urls {
		if !strings.HasPrefix(url, "https://") {
			fmt.Printf("url do not has http prefix, origin %s\n", url)
			url = "https://" + url
		}
		if resp, err := http.Get(url); err == nil {
			fmt.Printf("response http code %d, %s\n", resp.StatusCode, resp.Status)
		} else {
			fmt.Fprintf(os.Stderr, "ERR %v\n", err)
		}

	}

	cost := time.Since(start).Seconds()
	fmt.Printf("time cost %.2f", cost)
}

type Url struct {
	URL string
	Content string
	ErrMsg string
	Len int64
}
func L16(urls []string) {
	start := time.Now()
	ch := make(chan Url)

	for _, url := range urls {

		// do get url content
		go doGetUrlContent(url, ch)
	}

	for i := 0; i < len(urls); i++ {

		fmt.Println("receive ", <- ch)

	}
	cost := time.Since(start).Seconds()

	fmt.Printf("cost %.2f \n", cost)
}

func doGetUrlContent(url string, ch chan Url) {

	if resp, err := http.Get(url); err == nil {

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				ch <- Url{URL: url, ErrMsg: fmt.Sprintf("err %v", err)}
				return
			}
		}(resp.Body)

		if data, err := ioutil.ReadAll(resp.Body); err == nil {
			ch <- Url{URL: url, Content: string(data)}
		} else {
			ch <- Url{URL: url, ErrMsg: fmt.Sprintf("err %v", err)}
		}
	} else {
		ch <- Url{URL: url, ErrMsg: fmt.Sprintf("err %v", err)}
	}
}

func P16_110(urls []string) {
	start := time.Now()
	ch := make(chan Url)

	for _, url := range urls {

		// do get url content
		go doGetUrlContent2(url, ch)
	}

	for i := 0; i < len(urls); i++ {

		fmt.Println("receive ", <- ch)

	}
	cost := time.Since(start).Seconds()

	fmt.Printf("cost %.2f \n", cost)
}
func doGetUrlContent2(url string, ch chan Url) {

	if resp, err := http.Get(url); err == nil {

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				ch <- Url{URL: url, ErrMsg: fmt.Sprintf("err %v", err)}
				return
			}
		}(resp.Body)

		open, err := os.Create("result")
		defer func(open *os.File) {
			err := open.Close()
			if err != nil {
				ch <- Url{URL: url, ErrMsg: fmt.Sprintf("err %v", err)}
			}
		}(open)
		if err != nil {
			ch <- Url{URL: url, ErrMsg: fmt.Sprintf("err %v", err)}
			return
		}

		if writtenLen, err := io.Copy(open, resp.Body); err == nil {
			ch <- Url{URL: url, Len: writtenLen}
		} else {
			ch <- Url{URL: url, ErrMsg: fmt.Sprintf("err %v", err)}
		}
	} else {
		ch <- Url{URL: url, ErrMsg: fmt.Sprintf("err %v", err)}
	}
}

var mu sync.Mutex
var count int64 = 0
func L17()  {


	http.HandleFunc("/", handler)
	http.HandleFunc("/count", precount)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatal("http server listen failed {}", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	fmt.Println("handle execute once")
	_, err := fmt.Fprintf(w, "URI %q", r.Host)
	if err != nil {
		return
	}
	count++
	mu.Unlock()

}

func precount(w http.ResponseWriter, r *http.Request)  {

	mu.Lock()
	fmt.Println("precount execute once")
	_, err := fmt.Fprintf(w, "click count %d", count)
	if err != nil {
		return
	}
	mu.Unlock()
}