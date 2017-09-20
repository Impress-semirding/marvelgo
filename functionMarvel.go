package main

var maps = map[int]func() int {
	1 : func() int {
		return 1
	},
	2 : func() int {
		return 2
	},
	3 : func() int {
		return 3
	},
}

func main() {
}
