package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function

func newRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")

	return r
}

func main() {

	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page!")

}
