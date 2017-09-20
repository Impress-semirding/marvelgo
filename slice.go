package main

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	for i, v := range s {
		if i == 0 {
			s = s[:3]
			//s[2] = 18

			/*for j, value := range s {
				println(j, "**", value)
			}*/
		}

		println(i, "**", v)
		println(len(s))
	}

}

