package controllers

import (
	"awesomeMu0.2/views"
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

func NewUser() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/SignUp.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// New thi is used to render the form to the user
//a new user account
//Get /SignUp
func (U *Users) New(w http.ResponseWriter, r *http.Request) {
	U.NewView.Render(w, nil)
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (U *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
