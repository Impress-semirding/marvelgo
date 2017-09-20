package main

import (
	"fmt"
	"reflect"
)

type Bird struct{
	Name string
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying....")
}

func main() {

	sparrow := &Bird{"sparrow", 3}
	re := reflect.ValueOf(sparrow).Elem()
	typeOf := re.Type()

	for i := 0; i < re.NumField() ;i++ {
		f := re.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOf.Field(i).Name,
			f.Type(),
			f.Interface())
	}

}

