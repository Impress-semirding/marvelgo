package main

import (
	"fmt"
	"reflect"
)

type boy struct {
	Name string
	Age int
}

type human interface {
	SeyName()
	SeyAge()
}


func (b *boy) SeyName()  {
	fmt.Println("名字: " + b.Name)
}

func (b *boy) SeyAge() {
	fmt.Println("年纪:",b.Age)
}



func main() {
	var q human
	marvel := &boy{"marvelous", 21}

	q = marvel

	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	fmt.Println(t)
	fmt.Println(v)
}
