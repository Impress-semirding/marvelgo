package main

func main() {
	switch1()
	switch2()
}

func switch1()  {
	integer := 6

	switch integer {
	case 1:
		println("integer == 1")
	case 2:
		println("integer == 2")
	case 6:
		println("integer == 6")
	case 3:
		println("integer == 3")
	}
}

func switch2()  {
	integer := 6

	switch integer {
	case 4:
		println("integer == 4")
		fallthrough
	case 5:
		println("integer == 5")
		fallthrough
	case 6:
		println("integer == 6")
		fallthrough
	case 7:
		println("integer == 7")
	}
}