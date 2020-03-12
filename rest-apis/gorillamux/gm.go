package gm

import (
	"log"
	"net/http"
	"preset"

	"github.com/gorilla/mux"
)

func GmRest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", preset.Index)
	log.Fatal(http.ListenAndServe(":8081", router))
}
