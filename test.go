package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {


/*	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)*/
	Iterator()

}
func incCount() {
	defer wg.Done()
	for i := 0; i < 200000; i++ {
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		value++
		atomic.StoreInt32(&count, value)
	}
}

func Iterator() {
	slice := []string{"a", "b", "c"}

	for _, value := range slice{
		/*
			这里就相当于调用一个匿名函数
			var function = func() {fmt.Println(value)}
			go function()
		 */
		go func() {
			fmt.Println(value)
		}()
	}

}
