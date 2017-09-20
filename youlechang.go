package main

import "time"

func init()  {
	for i := 0; i < 100000; i++ {
		go println("i:", i)
	}
}

func main() {
	// 让主线程等待一会
	time.Sleep(time.Second * 10000)
}
