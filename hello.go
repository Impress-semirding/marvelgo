package main

import "fmt"

func sum(values []int, resultChan chan int){

	sum := 0
	for _, value := range values  {
		sum += value
	}

	resultChan <- sum // 将计算结果发送到 channel 中
}

func main() {

	values := [] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultChan := make(chan int, 2)
	go sum(values[: len(values)/2], resultChan)
	go sum(values[len(values)/2 : ], resultChan)

	sum1, sum2 := <- resultChan, <-resultChan // 接收结果
	fmt.Println("Result:", sum1, sum2, sum1 + sum2)
}
