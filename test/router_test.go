package test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"go_learning/router"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var Users []User = []User{
	User{Id: 1, Name: "Admin"},
	User{Id: 2, Name: "Moderator"},
	User{Id: 3, Name: "Viewer"},
}

func TestGetRoles(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/role", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	message, _ := json.Marshal(&Users)
	assert.Equal(t, string(message), w.Body.String())
}

func TestGetRole(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()

	ind := rand.Intn(len(Users)) + 1
	uri := fmt.Sprintf("/api/v1/role/%d", Users[ind-1].Id)
	fmt.Println(uri)
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	decodedUser, _ := json.Marshal(&Users[ind-1])
	assert.Equal(t, string(decodedUser), w.Body.String())

}

func TestGetRoleIdInvalid(t *testing.T) {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/role/aa", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Id invalid.", w.Body.String())

}

func TestGetRoleNoId(t *testing.T) {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/role/0", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "No user with this id.", w.Body.String())

}
