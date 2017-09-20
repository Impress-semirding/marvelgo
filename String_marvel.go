package main

import "strings"

func main() {

	s := "We,      .are$ *frIen你好 "
	a := strings.Title(s)
	b := strings.ToTitle(s)
	c := strings.ToUpper(s)
	d := strings.Trim(s, "W")
	e := strings.Fields(s)
	f := strings.Split("v,v,v,v,v,v,a", ",")

	println(a)
	println(b)
	println(c)
	println(d)
	println(len(e))
	println(len(f))
	println(s)

	println("---------------------")
	g := make([]string, 7)
	g[6] = "a"
	for _, v := range g{
		println(v)
	}

	switch {

	case 1 > 2:
		println(1)
		println(1)
		println(1)
	case 1 < 2:
		println(2)
		println(2)
		println(2)
	}
}
