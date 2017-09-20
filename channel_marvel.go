package main

import "fmt"

/*
在这个例子中，我们定义了一个包含10个channel的数组（名为chs），并把数组中的每个
channel分配给10个不同的goroutine。在每个goroutine的Add()函数完成后，我们通过ch <- 1语
句向对应的channel中写入一个数据。在这个channel被读取前，这个操作是阻塞的。在所有的
goroutine启动完成后，我们通过<-ch语句从10个channel中依次读取数据。在对应的channel写入
数据前，这个操作也是阻塞的。这样，我们就用channel实现了类似锁的功能，进而保证了所有
goroutine完成后主函数才返回。是不是比共享内存的方式更简单、优雅呢?
 */

func Count(ch chan int)  {
	// 将数据写入 channel
	ch <- 1
	println("nihao")
}

func main() {

	chs := make([] chan int, 10)
	for i := 0; i < 10 ; i++ {
		chs[i] = make(chan int, 1)
		go Count(chs[i])
	}

	// 在没有将数据写入 channel 之前，该操作是阻塞的
	for _, ch := range chs {
		// 将数据从 channel 中读出来
		value := <- ch
		fmt.Println(value)
	}
}
