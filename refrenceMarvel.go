package main

type User struct {
	userName string `用户名`
	password string `密码`
}

func main() {

	user1 := User{"2452599288@qq.com", "123456"}
	user2 := User{"2452599288@qq.com", "123456"}

	println(&user1)
	println(&user2)

}
