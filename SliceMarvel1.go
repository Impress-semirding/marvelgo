package main

import "fmt"

/**
	TODO：Slice 和 append 连用的坑
	append 函数会改变 slice 所引用的数组内容，从而影响到引用同一数组的其它 slice，但当 slice 中没有剩余空间（即 （cap - len）== 0）
	时，此时将动态分配新的数组空间。返回的 slice 数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的 slice 则不受
	影响
 */
func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:2]
	fmt.Println("s cap:", cap(s))
	fmt.Println("s len:", len(s))
	fmt.Println("cap(s)-len(s):", cap(s)-len(s))

	insert(s)

	fmt.Println("s len:", len(s))
	fmt.Println("arr:", arr)
}

func insert(slices []int)  {
	slices = append(slices, 6)
	fmt.Println()
	fmt.Println("========insert start ============")
	fmt.Println("len(slices):", len(slices))
	fmt.Println("slices: ", slices)
	fmt.Println("slices[0]：", slices[0])
	fmt.Println("slices[1]：", slices[1])
	fmt.Println("========insert end ============")
	fmt.Println()
}
