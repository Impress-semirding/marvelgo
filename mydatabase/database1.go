package main

import "fmt"

func main() {

	marvel := make(map[string]interface{}, 0)

	marvel = map[string]interface{} {
		"marvel1" : "12345",
		"marvel2" : "12345",
		"marvel3" : "12345",
		"marvel4" : "12345",
		"marvel5" : "12345",
	}

	// 第一个返回值是 value，第二个是 bool
	a, b := marvel["marvel1"]

	fmt.Println("a:", a, "b:", b)

	for k, v := range marvel {
		fmt.Println("k:", k, "v:", v)
	}
}
