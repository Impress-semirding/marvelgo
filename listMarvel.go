package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()// 返回的是一个引用类型
	addElement(l, "a")

	removeElement(l)
}

/*
	添加一个元素
 */
func addElement(l *list.List, str string)  {
	l.PushBack(str)
	l.PushBack("b")
	l.PushBack("c")
}

/*
	删除所有节点
	循环删除节点的时候有一个坑，迭代条件不能是 e = e.Next() 因为 e 节点已经被删除了。
 */
func removeElement(l *list.List)  {

	var next *list.Element

	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		l.Remove(e)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

/*
	查找所有节点
 */
func findElement() (slice []string) {

	return
}

/*
	根据下标更新节点
 */
func updateElement(index int)  {

}
