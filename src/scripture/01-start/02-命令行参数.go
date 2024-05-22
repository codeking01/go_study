package main

import (
	"fmt"
	"os"
)

/**
 * @author xiongjl
 * @since 2024/5/22  15:07
 */

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println(index)
	}
	fmt.Println(s)
}
