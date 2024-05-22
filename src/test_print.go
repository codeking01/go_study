package main

import "fmt"

/*
*
测试输入输出
*/
func main() {
	fmt.Print("测试")
	fmt.Printf("%d\n\n", 123)
	fmt.Printf("%#v %#v", "", "")
	fmt.Printf("%T %#T", "", "")
	var name string
	n, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	fmt.Println(n, err)
}
