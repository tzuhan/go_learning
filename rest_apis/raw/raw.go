package raw

import (
	"log"
	"net/http"

	"github.com/tzuhan/go_learning/rest_apis/db"
)

func RawRest() {
	db.Initialize()
	mux := http.NewServeMux()
	mux.Handle("/", http.HandleFunc(db.Index))
	mux.Handle("/users", http.HandleFunc(db.GetAllUsers))
	log.Fatal(http.ListenAndServe(":8082", mux))
}
