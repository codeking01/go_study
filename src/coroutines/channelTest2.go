package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @author xiongjl
 * @since 2023/11/30  14:47
 */

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始买东西\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 结束\n", name)
	moneyChan <- money
	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	startTime := time.Now()
	wait.Add(6)
	go pay("张三", 3, &wait)
	go pay("李四", 4, &wait)
	go pay("王五", 5, &wait)
	go pay("王五", 5, &wait)
	go pay("王五", 5, &wait)
	go pay("王五", 5, &wait)
	// 需要结束通道
	go func() {
		defer close(moneyChan)
		wait.Wait()
	}()
	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		break
	//	}
	//}
	var moneyList []int
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}
	endTime := time.Since(startTime)
	fmt.Println(endTime)
	fmt.Println(moneyList)
}
