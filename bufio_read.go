package main

import (
	"strings"
	"bufio"
	"fmt"
)

func main() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	b := make([]byte, 20)

	n, err := br.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)
	// ABCDEFGHIJKLMNOPQRST 20 <nil>

	n, err = br.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)
	// UVWXYZ1234567890     16 <nil>

	// UnreadByte 撤消最后一次读出的字节
	// 只有最后读出的字节可以被撤消
	// 无论任何操作，只要有内容被读出，就可以用 UnreadByte 撤消一个字节
	// 注意：只能调用一次
	br.UnreadByte()
	//br.UnreadByte()
	//br.UnreadByte()

	n, err = br.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)
}
