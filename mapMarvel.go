package main

func main() {
	var marvelMaps = map[string]int {
		"marvel1" : 1,
		"marvel2" : 2,
		"marvel3" : 3,
		"marvel4" : 4,
		"marvel5" : 5,
		"marvel6" : 6,
	}

	marvelMaps["marvel7"] = 7
	marvelMaps["marvel8"] = 8
	for k, _ := range marvelMaps{
		if "marvel2" == k {
			delete(marvelMaps, "marvel2")
			delete(marvelMaps, "marvel3")
			delete(marvelMaps, "marvel4")
			delete(marvelMaps, "marvel5")
			delete(marvelMaps, "marvel6")
		}
	}

	for k, v := range marvelMaps{
		println(k, v)
	}
}



/*	a := 0
	for j := 0; j < len(marvelMaps); j++ {
		a ++
		if j == 0 {
			delete(marvelMaps, "marvel2")
		}
	}
	println(a)*/