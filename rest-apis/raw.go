package restapi

import (
	"log"
	"net/http"
)

func RawRest() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
