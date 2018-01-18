package utils

import "gopkg.in/gomail.v2"

var username = "XXX@163.com"
var password = "XXX"
var D = gomail.NewDialer("smtp.163.com", 25, username, password)

func RegisterEmail(to string, name string, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", to)                                              //收件人
	m.SetHeader("Subject", "Welcome")                                  //标题
	m.SetBody("text/html", "<h1>Hello,"+name+"!</h1>Welcome to Iris!") //内容

	// Send the email to Bob, Cora and Dan.
	if err := D.DialAndSend(m); err != nil {
		panic(err)
	}
}
