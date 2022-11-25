package main

import (
	"html/template"
	"os"
)

type Todo struct {
	Name        string
	Description string
}

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("Hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "MU",
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
	data.Name = "HSN"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
