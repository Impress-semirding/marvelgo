package main

import (
	"fmt"
	"os"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of " +
		"two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {
	args := os.Args

	if args == nil, len(args) < 2 {
		Usage()
		return
	}

	switch args[0] {
	case "add" :
		if len(args) != 3 {
			fmt.Println("USAGE: calc command [arguments] ...")
		}
		value1, err1 := strconv.Atoi(args[1])
		value2, err2 := strconv.Atoi(args[2])

		if err1 != nil || err2 != nil {
			return
		}

		sum := value1 + value2

		fmt.Println(sum)
	default:
		Usage()
	}
}
