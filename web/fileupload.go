package main

import (
	"net/http"
	"fmt"
	"text/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
	"strings"
)

func fileupload(w http.ResponseWriter, r *http.Request)  {

	method := r.Method
	fmt.Println(method)

	if "GET" == method {
		createTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(createTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("web/fileupload.html")
		t.Execute(w, token)
	} else {
		// 这里是设置文件大小
		r.ParseMultipartForm(32 << 20)
		// 根据 input 中的 name 属性获得上传的文件
		file, handler, err := r.FormFile("file")
		if err != nil {
			println("上传的文件有错误!")
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		filename := handler.Filename
		fileType := strings.Split(filename, ".")[1]
		// 文件格式的限制
		if fileType != "txt" {
			fmt.Fprintf(w, "%s", "文件格式有误，只能上传txt文件")
			return
		}
		// 给文件重命名
		filename = "a.txt"

		/*
			管道命令：

			可以用管道(pipe)操作符" | " 来连接进程。
			命令语法：
			command1 | command2 | command3 | ... | commandN
			功能：将命令 command1 的标准输出连接到 command2的标准输入 ...
		 */

		// 打开文件，该文件在web 目录下面，
		f, err := os.OpenFile("web/" + filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("打开文件有错误")
			return
		}
		defer f.Close()
		// 复制文件
		io.Copy(f, file)
	}

}

func main() {

	http.HandleFunc("/fileupload", fileupload)
	http.ListenAndServe(":9090", nil)

}
