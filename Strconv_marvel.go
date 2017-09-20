package main

import (
	"strconv"
	"fmt"
)

func main() {

	b := make([]byte, 0)
	c := strconv.AppendBool(b, true)
	fmt.Println(string(c))
}
