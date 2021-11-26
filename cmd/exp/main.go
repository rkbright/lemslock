package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Bio  string
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Susan Smith",
		Age:  111,
		Bio:  "Haha, you have been hacked",
		Meta: UserMeta{
			Visits: 4,
		},
	}

	fmt.Println(user.Meta.Visits)

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
