package hr

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	//"html"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s\n", ps.ByName("name"))
}

func GetId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//get user info from id
	//query db

}
func HrIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome\n")
}

func HrRest() {
	router := httprouter.New()
	router.GET("/", HrIndex)
	router.GET("/hello/:name", Hello)
	/*
	       http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	           fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	   	    fmt.Println(r.URL.Path)
	       })
	*/

	log.Fatal(http.ListenAndServe(":8080", router))
}
