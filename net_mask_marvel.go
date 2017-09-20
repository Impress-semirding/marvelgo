package main

import (
	"os"
	"fmt"
	"net"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted- ip -addr\n", os.Args[0])
		os.Exit(1)
	}

	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)

	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	ipMask := addr.DefaultMask()
	network := addr.Mask(ipMask)

	ones, bits := ipMask.Size()
	fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits, bits,
		"Leading ones count is ", ones, ones,
		"Mask is (hex) ", ipMask.String(),
		" Network is ", network.String())
	os.Exit(0)
}
