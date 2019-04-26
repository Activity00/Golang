package main

import "fmt"

type user2 struct {
	name  string
	email string
}

type notifierplus interface {
	notify()
}

func (u *user2) notify() {
	fmt.Println(u.name, u.email)
}

type adminPlus struct {
	user2
	level int
}

func sendNotify(n notifierplus) {
	n.notify()
}

func main() {
	adp := adminPlus{user2: user2{name: "wmh", email: "hhh"}, level: 1}
	adp.user2.notify()
	adp.notify()

	sendNotify(&adp)
}
