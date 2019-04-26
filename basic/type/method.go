package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) sendEmail() {
	fmt.Println("send to email", u.name, u.email)
}

func (u *user) changeName() {
	u.name = "xxx"
}

func main() {
	u := user{name: "wmh", email: "1032662429@qq.com"}
	u.sendEmail()

	u.changeName()
	fmt.Println(u.name)
}
