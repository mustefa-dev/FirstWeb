package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1>hello</h1>")
}
func contact(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "text/html")
	fmt.Fprint(w, "mailto:support@mu.com")

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact/", contact)
	http.ListenAndServe(":2000", r)
}
