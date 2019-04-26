package main

import "fmt"

type notifier interface {
	notify()
}

type user1 struct {
	name  string
	email string
}

func (u *user1) notify() {
	fmt.Println("notify", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Println("admin", a.name, a.email)
}

func main() {
	u := user1{name: "wmh", email: "1032662429"}
	sendNotification(&u)

	a := admin{name: "admin", email: "admin@qq.com"}
	sendNotification(&a)
}
