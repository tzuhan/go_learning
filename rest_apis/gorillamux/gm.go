package gm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tzuhan/go_learning/rest_apis/db"
)

func GmRest() {
	db.Initialize()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/regist/{id}", RegistUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome\n")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	/*
		fmt.Fprintf(w, "%s", json.Marshal(UserList))
	*/
	json.NewEncoder(w).Encode(db.UserList)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, user := range db.UserList {
		if user.ID == id {
			//fmt.Fprintf(w, "%#v", json.Marshal(user))
			w.WriteHeader(http.StatusOK) //200
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	//user not found
	w.WriteHeader(http.StatusNoContent) //204: id not found
	fmt.Fprintf(w, "User not found\n")
}

func RegistUser(w http.ResponseWriter, r *http.Request) {
	var newUser db.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		fmt.Fprintf(w, "User data format wrong\n")
		return
	}
	//check if ID already exists

	json.Unmarshal(reqBody, &newUser)
	id := mux.Vars(r)["id"]
	for _, user := range db.UserList {
		if user.ID == id {
			w.WriteHeader(http.StatusBadRequest) //400
			fmt.Fprintf(w, "User ID already exists\n")
			return
		}
		if user.Account == newUser.Account {
			w.WriteHeader(http.StatusBadRequest) //400
			fmt.Fprintf(w, "User Account already exists\n")
			return
		}
	}

	db.UserList = append(db.UserList, newUser)
	w.WriteHeader(http.StatusCreated) //201
	json.NewEncoder(w).Encode(newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//id and account name can't be modified after registered
	var updateUser struct {
		FirstName   string
		LastName    string
		Birthday    string
		Description string
		Gender      int
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		fmt.Fprintf(w, "User data format wrong\n")
		return
	}
	//check if ID already exists
	json.Unmarshal(reqBody, &updateUser)

	updateId := mux.Vars(r)["id"]
	for id, user := range db.UserList {
		if user.ID == updateId {
			//update user
			db.UserList[id].FirstName = updateUser.FirstName
			db.UserList[id].LastName = updateUser.LastName
			db.UserList[id].Birthday = updateUser.Birthday
			db.UserList[id].Description = updateUser.Description
			db.UserList[id].Gender = updateUser.Gender
			w.WriteHeader(http.StatusOK) //200
			json.NewEncoder(w).Encode(db.UserList[id])
			return
		}
	}
	//User not found
	w.WriteHeader(http.StatusBadRequest) //400
	fmt.Fprintf(w, "User not found\n")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	deleteId := mux.Vars(r)["id"]

	for id, user := range db.UserList {
		if user.ID == deleteId {
			//delete user
			db.UserList = append(db.UserList[:id], db.UserList[id+1:]...)
			w.WriteHeader(http.StatusOK) //200
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	//User not found
	w.WriteHeader(http.StatusBadRequest) //400
	fmt.Fprintf(w, "User not found\n")
}
