package main

import "fmt"

type B interface {
	foo()
}

/*
	结构体A 里面放了一个匿名字段，是接口B
 */
type A struct {
	B
}
/*
	定义一个结构体，实现 接口B 里面的方法
 */
type C struct {}

func (c C) foo()  {
	fmt.Println("我是结构体C，我实现了 B 接口里面的方法")
}

func main() {
	// 给结构体赋值
	a := A{C{}}
	// 实际上调用的结构体 C 里面的方法
	a.foo()
}
