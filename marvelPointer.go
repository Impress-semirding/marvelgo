package main

type S map[string][]string

func Summary(param string)(s*S) {
	s = &S{
		"name": []string{param},
		"profession": []string{"Javaprogrammer", "ProjectManager"},
		"interest(lang)": []string{"Clojure", "Python", "Go"},
		"focus(project)": []string{"UE", "AgileMethodology", "SoftwareEngineering"},
		"hobby(life)": []string{"Basketball", "Movies", "Travel"},
	}
	return s
}
func main(){
	s := Summary("Harry")

	var ss S
	ss = *s
	var value []string
	value = ss["name"]
	println(value[0])

	//fmt.Printf("Summary(address):%v\r\n",s)
	//fmt.Printf("Summary(content):%v\r\n",*s)
}

