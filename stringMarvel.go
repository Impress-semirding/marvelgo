package main

func main() {

	s := "str我爱你"
	a := []rune(s)
	/*for i, j := range s{
		println(i, j)
	}*/

	for i, j := 0, len(s) - 1; i < j ; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	println(string(a))
}
