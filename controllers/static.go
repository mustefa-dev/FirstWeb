package controllers

import (
	"awesomeMu0.2/views"
)

type Static struct {
	HomeView     *views.View
	ContactsView *views.View
}

func NewStatic() *Static {
	return &Static{
		HomeView:     views.NewView("bootstrap", "views/static/home.gohtml"),
		ContactsView: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}
