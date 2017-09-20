package main

import (
	"os"
	"net"
	"fmt"
)

func main() {

	if len(os.Args) != 1 {
		fmt.Fprintf(os.Stderr,"Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[0]
	addr, err := net.LookupAddr(name)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	for _, value := range addr {
		fmt.Println(value)
	}

	os.Exit(0)
}
