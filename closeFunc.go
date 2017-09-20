package main

import "fmt"

func main() {

	i, j := 10, 20
	func(i, j int) {
		fmt.Println(i)
		fmt.Println(j)
	}(i, j)


/*	a := func()(func()) {
		var i int = 10
		return func(j int) {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}(1)*/

}
