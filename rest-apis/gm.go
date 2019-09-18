package restapi

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GmRest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8081", router))
}
