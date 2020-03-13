package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tzuhan/go_learning/handler"
	"github.com/tzuhan/go_learning/router"
)

func TestIHelloGetRouter(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/role", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	message, _ := json.Marshal(&handler.Users)
	assert.Equal(t, message, w.Body.String())
}
