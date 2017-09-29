package main

import "fmt"

type B struct {
	nameB string
}

type C struct {
	nameC string
}

type A struct {
	b B
	c C
}

func Abc(jk *A)  {
	fmt.Println(jk)
}

func main() {
	b := B{"marvelB"}
	c := C{"marvelC"}

	a := new(A)
	a.b = b
	a.c = c

	Abc(a)
}
