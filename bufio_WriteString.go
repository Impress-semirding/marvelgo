package main

import (
	"bytes"
	"bufio"
	"fmt"
)

func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())

	bw.WriteString("ABCDEFGH")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

	// Flush 将缓存中的数据提交到底层的 io.Writer 中
	bw.Flush()
	// Available 返回缓存中的可以空间
	fmt.Println(bw.Available())
	// Buffered 返回缓存中未提交的数据长度
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)
}
