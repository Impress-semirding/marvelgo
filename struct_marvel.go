package main

type Person struct {
	Name string
}

type Person1 struct {
	Name string
}

type Person2 struct {
	Person
	Person1
}

func main() {

	per := Person2{Person{"marvel"},Person1{"hiaho"}}
	println(per.Person.Name)
	println(per.Person1.Name)
	fla := nil != nil
	println(fla)

}
