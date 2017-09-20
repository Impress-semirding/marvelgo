package main

import (
	"os"
	"fmt"
	"net"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip- addr\n", os.Args[0])
		os.Exit(1)
	}

	// 由所学只是可知 os.Args[0] 是协议，os.Args[1] 是ip/域名 或 ip:post
	name := os.Args[1]
	// 将 域名 转换为 ip 地址
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The Address is", addr.String())
	}

	os.Exit(0)
}
