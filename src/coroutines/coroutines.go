package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @author xiongjl
 * @since 2023/11/29  14:47
 */

//func fib(n int) int {
//	if n <= 1 {
//		return n
//	}
//	a, b := 0, 1
//	for i := 2; i <= n; i++ {
//		a, b = b, a+b
//	}
//	return b
//}

func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func testCoroutines(n int, wait *sync.WaitGroup) {
	fib(n)
	wait.Done()
}

func main() {
	var n, nums = 30, 10
	startTime := time.Now()
	for i := 0; i < nums; i++ {
		fib(n)
	}
	endTime := time.Since(startTime)
	fmt.Println(endTime)
	startTime = time.Now()
	// 使用协程
	var wait sync.WaitGroup
	wait.Add(nums)
	for i := 0; i < nums; i++ {
		// 前面要加一个go
		go testCoroutines(n, &wait)
	}
	wait.Wait()
	endTime = time.Since(startTime)
	fmt.Println(endTime)
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}
}
