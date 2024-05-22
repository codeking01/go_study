package main

import "fmt"

/**
 * @author xiongjl
 * @since 2023/11/29  15:35
 */
func main() {
	//var a = make([]int, 1, 2) // 这个是有初始值的
	// 定义一个传递整型的通道
	var intChannel chan int
	intChannel = make(chan int, 1)
	intChannel <- 1
	fmt.Println(<-intChannel)
	intChannel <- 2
	n, ok := <-intChannel
	fmt.Println(n, ok)
	close(intChannel)
	//intChannel <- 3 // 关闭之后就不能再写或读了  send on closed channel
	//fmt.Println(intChannel)
}
