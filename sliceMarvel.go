package main

import "fmt"

var marvelMap = map[string]int {
	"marvel1" : 1,
	"marvel2" : 2,
	"marvel3" : 3,
	"marvel4" : 4,
}


func main() {
	// 使用内建函数 delete 删除指定的键
	delete(marvelMap, "marvel1")

	for a, b := range marvelMap {
		fmt.Println(a,b)
	}

	a := []byte{1, 2, 3}
	b := []byte{4, 5, 6}

	c := Append(a, b)

	fmt.Println("len:", len(c))
	fmt.Println("cap:", cap(c))
	for i, value := range c  {
		fmt.Println(i, value)
	}

}

func Append(slice, data[] byte) []byte {

	l := len(slice)

	fmt.Println("l:", l)

	// 如果切片的长度 + 要加入数据的长度比 切片最大长度大的话，则扩容
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, l+len(data)*2)

		copy(newSlice, slice)
		slice = newSlice
	}

	fmt.Println("l+len(data)*2:", l+len(data)*2)

	for i, c := range data  {
		slice[l+i] = c
	}

	return slice
}

func printMarvel()  {
	str := "我爱你"

	fmt.Println(str, len(str))
	fmt.Println(str[0:3])
}
