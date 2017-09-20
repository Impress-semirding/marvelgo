package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Book struct {
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float32
	Times int64
	A time.Time
}

func main() {

	Authors := []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"}

	fmt.Println(time.Now().Unix())

	gobook := Book {
		"Marvel",
		Authors,
		"ituring.com.cn",
		true,
		9.99,
		time.Now().Unix(),
		time.Unix(0,0),
	}

	// 将一个结构体转换为 json 对象
	b, _ := json.Marshal(gobook)

	fmt.Println(string(b))

/*	var book Book

	err := json.Unmarshal(b, &book)

	if err != nil {
		fmt.Println(book)
	}*/
}
