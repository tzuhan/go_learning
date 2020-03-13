package hr

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tzuhan/go_learning/rest_apis/db"
	//"html"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s\n", ps.ByName("name"))
}

func GetUserById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//get user by id
}

func GetAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//get all users
}

func HrIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome\n")
}

func HrRest() {
	db.Initialize()
	router := httprouter.New()
	router.GET("/", HrIndex)
	router.GET("/users", GetAllUsers)
	router.GET("/users/:id", Hello)
	/*
	       http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	           fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	   	    fmt.Println(r.URL.Path)
	       })
	*/

	log.Fatal(http.ListenAndServe(":8080", router))
}
