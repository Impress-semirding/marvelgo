package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

/*
解析 XML 到 struct 的时候遵循如下的规则：
	• 如果 struct 的一个字段是 string 或者[]byte 类型且它的 tag 含
	有",innerxml"，Unmarshal 将会将此字段所对应的元素内所有内嵌的原始 xml 累加到此字
	段上，如上面例子 Description 定义。最后的输出是
	Shanghai_VPN127.0.0.1Beijing_VPN127.0.0.2

	• 如果 struct 中有一个叫做 XMLName，且类型为 xml.Name 字段，那么在解析的时
	候就会保存这个 element 的名字到该字段,如上面例子中的 servers。

	• 如果某个 struct 字段的 tag 定义中含有 XML 结构中 element 的名称，那么解析的时
	候就会把相应的 element 值赋值给该字段，如上 servername 和 serverip 定义。

	• 如果某个 struct 字段的 tag 定义了中含有",attr"，那么解析的时候就会将该结构所
	对应的 element 的与字段同名的属性的值赋值给该字段，如上 version 定义。

	• 如果某个 struct 字段的 tag 定义 型如"a>b>c",则解析的时候，会将 xml 结构 a 下面
	的 b 下面的 c 元素的值赋值给该字段。

	• 如果某个 struct 字段的 tag 定义了"-",那么不会为该字段解析匹配任何 xml 数据。

	• 如果 struct 字段后面的 tag 定义了",any"，如果他的子元素在不满足其他的规则的
	时候就会匹配到这个字段。

	• 如果某个 XML 元素包含一条或者多条注释，那么这些注释将被累加到第一个 tag
	含有",comments"的字段上，这个字段的类型可能是[]byte 或 string,如果没有这样的字段存
	在，那么注释将会被抛弃。

	上面详细讲述了如何定义 struct 的 tag。只要设置对了 tag，那么 XML 解析就如上面示例般
	简单，tag 和 XML 的 element 是一一对应的关系，如上所示，我们还可以通过 slice 来表示
	多个同级元素。

	注意： 为了正确解析，go 语言的 xml 包要求 struct 定义中的所有字段必须是可导出的（即
	首字母大写）
 */
type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
	Descriptoin string `xml:"innerxml"`
}

type server struct {
	XMLName xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}
func main() {

	// 打开文件
	file, err := os.Open("server.xml")
	if err != nil {
		fmt.Println(err)
	}

	// 将文件里面的内容读到 data []byte 里面
	data, err := ioutil.ReadAll(file)

	file.Close()

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := Recurlyservers{}

	// data 接收的是 XML 数据流，v 是需要输出的结构，定义为 interface，也就是可以把 XML 转换为任意的格式。
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}
