package main


var c = make(chan int, 10)
func f()  {
	println("hello world!")
	a := <-c
	println(a)

}

func hello()  {
	go f()

}

func main() {

	hello()
	c<-10
	println("你好")
}
