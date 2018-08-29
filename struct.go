package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

type notifier interface {
	notify()
}

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	user
	level string
}

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	zhujiajun := user{
		name:       "zhujiajun",
		email:      "zhujiajunup@163.com",
		ext:        123,
		privileged: true,
	}
	fred := admin{
		user:  zhujiajun,
		level: "super",
	}
	zhujiajun.notify()
	sendNotification(&zhujiajun)
	zhujiajun.changeEmail("zhujiajun@163.com")
	sendNotification(&zhujiajun)
	fred.user.notify()
	fred.changeEmail("zhujiajun")
	fred.notify()
}
