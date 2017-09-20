package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()	//解析参数
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))

	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到 w 的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	if "GET" == r.Method {
		t, _ := template.ParseFiles("web/login.tpl")
		if t != nil {
			t.Execute(w, nil)
		} else {
			fmt.Fprintf(w, "出错了")
		}
	} else {
		/*
		   r.Form["username"]也可写成 r.FormValue("username")。调用 r.FormValue 时会自动调用
		   r.ParseForm，所以不必提前调用。r.FormValue 只会返回同名参数中的第一个，若参数不存
		   在则返回空字符串。
		 */
		// 获取表单数据的时候一定要有这句话，不然获取不到参数
		//r.ParseForm()
		//fmt.Println("username:", r.Form["username"]) //这样获取的参数都是带 []
		//fmt.Println("password:", r.Form["password"])
		//fmt.Println("selects:", r.Form["selects"])

		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))
		fmt.Println("selects:", r.FormValue("selects"))
		// 这样会原封不动的输出到页面中，应该过滤之后再输出到页面中去
		// fmt.Fprintf(w, "<script>alert('你好!')</script>")
		// 下面这样就可以过滤掉 html 标签
		template.HTMLEscape(w, []byte("<script>alert('你好!')</script>"))
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) //设置访问的路由
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
