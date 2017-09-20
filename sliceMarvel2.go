package main

import "fmt"

func main() {
	arr := make([]int, 6, 100)
	arr = []int{1, 2, 3, 4}
	//s := arr[1:]

	fmt.Println("s:", cap(arr))
	fmt.Println("s:", len(arr))
	fmt.Println("s:", cap(arr)-len(arr))

	inserts(arr)
	//fmt.Println("s:", s)
	fmt.Println("arr:", arr)
}

func inserts(slice []int)  {
	slice = append(slice, 6)
}
