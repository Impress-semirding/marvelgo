package main

import (
	"strings"
	"bufio"
	"bytes"
	"fmt"
)

// WriteTo 实现了 io.WriterTo 接口
func main() {

	s := strings.NewReader("ABCEFG")
	br := bufio.NewReader(s)
	b := bytes.NewBuffer(make([]byte, 0))

	br.WriteTo(b)
	fmt.Printf("%s\n", b) // ABCEFG
}
