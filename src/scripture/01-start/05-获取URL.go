package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * @author xiongjl
 * @since 2024/5/22  16:24
 */

func main() {
	start := time.Now()
	ch := make(chan string)
	var urlList []string
	for i := 0; i < 100; i++ {
		urlList = append(urlList, "http://www.baidu.com")
	}
	for _, url := range urlList {
		go fetch(url, ch) // start a goroutine
	}
	for range urlList {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
