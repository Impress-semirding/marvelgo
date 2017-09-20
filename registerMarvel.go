package main

/**
	用户注册
	邮箱：email
	用户名：userName
	密码：password
	重复密码：repass

 */

/**
	用户 Model
 */
type user struct {
	email string `邮箱账号`
	userName string `用户名`
	password string `密码`
	repass string `重复密码`
}

func main() {


}

func controller ()  {
	// 获取页面输入的用户名或邮箱账号

	// 获取用户输出的密码

	// 获取用户输入的重复密码

	// 判断密码和重复密码是否相同

	// 如果相同则继续下面操作

	// 判断用户输入的是邮箱账号还是用户名

	// 1. 邮箱账号

	// 1.1 在缓存中查找该用户是否存在

	// 1.1.1 如果存在则提示该账号已经被注册

	// 1.1.2 不存在的话则发送邮件给用户并将邮箱和密码存到 redis 缓存中去

	// 2. 用户名

	// 2.1 在缓存中查找该用户是否存在

	// 2.2.1 如果存在则提示该账号已经被注册

	// 2.2.2 不存在的话则注册成功并将用户名和密码存到 redis 缓存中去

}