package main

import "fmt"

func main() {
	rangeMarvel()
}

func rangeMarvel() {

	a := []string{"a", "b", "c"}

	for _, value := range a {
		fmt.Println(value)
	}
}
