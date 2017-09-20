package main

import "fmt"

func results(marvel, marvelous int) (n int, v int) {
	n = marvel
	v = marvelous
	return
}

func main() {

	n, v := results(100, 200)

	fmt.Println(n, v)

}
