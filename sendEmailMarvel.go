package main

import (
	"strings"
	"net/smtp"
	"fmt"
)

/**
	用来发送邮件
 */

func SendToMail(user, password, host, to, subject, body, mailtype string) {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)

	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

func main() {
	user := "18946650769@163.com"
	password := "sunyixuan0819"
	/*user := "cotactwithanother@163.com"
	password := "qwertyuiop1"*/
	host := "smtp.163.com:25"
	to := "2452599288@qq.com"

	subject := "主题：送给你"

	body := `
		<html>
		<body>
		<h3>
		"你还好吗"
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")
	for i := 0; i < 100; i++ {
		go SendToMail(user, password, host, to, subject, body, "html")

	}
	fmt.Println("send emails")

}
