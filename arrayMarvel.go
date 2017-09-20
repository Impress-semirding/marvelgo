package main

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := arr[2:5]
	s2 := arr[2:6:8]
	s3 := arr[3:6]

	for _, v := range s1 {
		print(v)
		print(" ")
	}

	println()
	for _, v := range s2 {
		print(v)
		print(" ")
	}
	println(cap(s2))
	for _, v := range s3 {
		print(v)
		print(" ")
	}
}
