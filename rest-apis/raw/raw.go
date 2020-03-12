package restapi

import (
	"log"
	"net/http"
)

func RawRest() {
	mux := http.NewServeMux()
	mux.Handle("/", preset.Index)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
