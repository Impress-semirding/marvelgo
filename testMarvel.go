package main

func length(s string) int {

	println("marvel")
	return len(s)
}

func main() {

	s := "abcd"
	for i, n := 0, length(s); i < n; i++ { // 使YíǛˁȜ length qƛ͉
		println(i, s[i])
	}
}
