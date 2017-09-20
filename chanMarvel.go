package main

func main() {

	data := make(chan int, 3)
	exit := make(chan bool)

	data <- 1
	data <- 2
	data <- 3

	go func() {
		for d := range data{
			println(d)
		}

		exit <- true
	} ()

	data <- 4
	data <- 5
	close(data)

	<- exit
}
