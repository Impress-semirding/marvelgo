package main

import (
	"bytes"
	"bufio"
	"fmt"
)

// WriteRune 向 b 中写入 r 的 UTF8 编码
// 返回 r 的编码长度
func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteByte('H')
	bw.WriteByte('e')
	bw.WriteByte('l')
	bw.WriteByte('l')
	bw.WriteByte('o')
	bw.WriteByte(' ')
	bw.WriteRune('世')
	bw.WriteRune('界')
	bw.WriteRune('！')
	bw.Flush()
	fmt.Println(b)
}
