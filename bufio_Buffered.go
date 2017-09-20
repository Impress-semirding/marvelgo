package main

import (
	"strings"
	"bufio"
	"fmt"
)

// Buffered 返回缓存中数据的长度
func main() {
	s := strings.NewReader("你好，世界！")
	br := bufio.NewReader(s)

	fmt.Println(s.Size())
	fmt.Println(br.Buffered())

	br.Peek(1)
	fmt.Println(br.Buffered())
}
