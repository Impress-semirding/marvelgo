package main

import (
	"strings"
	"bufio"
	"fmt"
)

// reset丢弃任何的缓存数据，丛植所有状态并且将缓存读切换到r
func main() {
	s := strings.NewReader("ABCEFG")
	str := strings.NewReader("123456")
	br := bufio.NewReader(s)
	b, _ := br.ReadString('\n')
	fmt.Println(b) // ABCEFG

	br.Reset(str)
	b, _ = br.ReadString('\n')
	fmt.Println(b) // 123456

}
