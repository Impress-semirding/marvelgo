package main

var s int
var b int
var c = make(chan int, 10)

func t()  {
	s = 10
	b = 20
	c <- 20
}

func g()  {

	println(s)
	println(b)
	<-c
	println("a")
}

func main() {

	go t()
	g()
}
