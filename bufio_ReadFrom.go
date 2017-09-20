package main

import (
	"bytes"
	"strings"
	"bufio"
	"fmt"
)

func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	s := strings.NewReader("Hello 世界!")
	bw := bufio.NewWriter(b)
	bw.ReadFrom(s)
	//bw.Flush() // ReadFrom 无需使用 Flush, 其自己已经写入
	fmt.Println(b)
}
