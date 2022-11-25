package main

import (
	"awesomeMu0.2/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	staticC := controllers.NewStatic()
	userC := controllers.NewUser()
	r := mux.NewRouter()
	r.Handle("/", staticC.HomeView).Methods("GET")
	r.Handle("/C", staticC.ContactsView).Methods("GET")
	r.HandleFunc("/SignUp/", userC.New).Methods("GET")
	r.HandleFunc("/SignUp/", userC.Create).Methods("POST")

	err := http.ListenAndServe(":2000", r)
	if err != nil {
		return
	}

}
func must(err error) {
	if err != nil {
		panic(err)
	}
}

// 44 lines, before using static.go
